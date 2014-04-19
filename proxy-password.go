package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"

	//"github.com/howeyc/gopass" //can't get it to function with username, re-evaluate later
)

var (
	proxyInfo ProxyInfo
	npmrcPath string
)

const (
	NPMRC_LOCATION_START = "c:/Users/"
	PROXY_REPLACE_STRING = "http://username:password@url:port"
)

func main() {
	fmt.Println("Starting the update process")
	//gearing up for the next feature :D
	//buildConfig(configInfo)
	fmt.Println("Your proxy info has been updated. :)")
}

func buildConfig(configInfo ConfigInfo) {
	//build info
	//i should probably have every method throw erros up to here and fail the build if
	//anyone errors
	configInfo.proxyInfo.proxyHTTP_String, configInfo.proxyInfo.proxyHTTPS_String = buildProxyString(configInfo)
	createNewFile(configInfo.configFilePath, getProxyFileContent(configInfo))

	setVariablesError := setProxyConfigVariables(configInfo)

	if setVariablesError != nil {
		log.Println("Error updating system variables: ")
		log.Println(setVariablesError)
	}

	status, updateCreateError := updateOrCreateProxyFile(configInfo)

	if updateCreateError != nil {
		log.Println("Error updating/creating proxy file(s): ")
		log.Println(updateCreateError)
	} else {
		log.Println(status + " proxy file(s)")
	}
}

func buildProxyString(configInfo ConfigInfo) (http string, https string) {
	replacer := strings.NewReplacer(
		"username", configInfo.proxyInfo.username,
		"password", configInfo.proxyInfo.password,
		"url", configInfo.proxyInfo.proxyUrl,
		"port", configInfo.proxyInfo.port)

	h := replacer.Replace(PROXY_REPLACE_STRING)
	hs := h //a little bit of cheating

	return h, hs
}

func setProxyConfigVariables(configInfo ConfigInfo) (err error) {
	httpError := setWindowsVariables("HTTP_PROXY", configInfo.proxyInfo.proxyHTTP_String)
	if httpError != nil {
		return httpError
	}

	httpsError := setWindowsVariables("HTTPS_PROXY", configInfo.proxyInfo.proxyHTTPS_String)

	if httpsError != nil {
		return httpsError
	}

	return nil
}

func setWindowsVariables(key string, value string) (err error) {
	setEnvError := os.Setenv(key, value)
	return setEnvError
}

func updateOrCreateProxyFile(configInfo ConfigInfo) (status string, err error) {
	//This should go into a func or something for all of them.
	//configInfo.configFilePath = getUserHomeDirectory() + "\\" + configInfo.configFileName

	if doesFileExist(configInfo.configFilePath) {
		err := updateProxyFiles(configInfo)
		if err != nil {
			log.Println(err)
			return "", err
		}

		return "Updated", nil
	} else {
		createNewFile(configInfo.configFilePath, getProxyFileContent(configInfo))

		return "Created", nil
	}

	//return a new error saying update/create file failed
}

//currently only checking for .npmrc
func getUserHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	return currentUser.HomeDir
}

func doesFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createNewFile(filepath string, content string) {
	err := ioutil.WriteFile(filepath, []byte(content), 0644)

	if err != nil {
		log.Println(err)
	}
}

func getProxyFileContent(configInfo ConfigInfo) (content string) {
	data := configInfo.FILE_HTTP_START + configInfo.proxyInfo.proxyHTTP_String + "\n"
	data += configInfo.FILE_HTTPS_START + configInfo.proxyInfo.proxyHTTPS_String

	return data
}

func updateProxyFiles(configInfo ConfigInfo) (err error) {
	var contents []byte
	contents, e := ioutil.ReadFile(configInfo.configFilePath)
	fileContents := string(contents)
	if e != nil {
		fmt.Println(e)
	}

	e = ioutil.WriteFile(configInfo.configFilePath, []byte(updateUsernamePassword(fileContents, configInfo.proxyInfo)), 0644)

	return e
}

func updateUsernamePassword(proxyString string, info ProxyInfo) string {
	regex := regexp.MustCompile("(https?://)(.*?):(.*?)(@.*)")
	results := regex.ReplaceAllString(proxyString, "${1}"+info.username+":"+info.password+"${4}")
	return results
}

//@todo this needs done
func updateUrlProxy(proxyString string, info ProxyInfo) string {

	return ""
}
