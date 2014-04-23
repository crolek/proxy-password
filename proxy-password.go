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
	configInfo.proxyInfo.proxyHTTP_String, configInfo.proxyInfo.proxyHTTPS_String = getProxyString(configInfo)
	createNewFile(configInfo.configFilePath, getProxyFileContent(configInfo))

	setVariablesError := setProxyConfigVariables(configInfo)

	if setVariablesError != nil {
		log.Println("Error updating system variables: ")
		log.Println(setVariablesError)
	}

	updateCreateError := updateOrCreateProxyFile(configInfo)

	if updateCreateError != nil {
		log.Println("Error updating/creating proxy file(s): ")
		log.Println(updateCreateError)
	} else {
		log.Println("Updated/Created proxy file(s)")
	}
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

/*
------------------------------Get------------------------------
*/

func getProxyFileContent(configInfo ConfigInfo) (content string) {
	data := configInfo.FILE_HTTP_START + configInfo.proxyInfo.proxyHTTP_String + "\n"
	data += configInfo.FILE_HTTPS_START + configInfo.proxyInfo.proxyHTTPS_String

	return data
}

func getProxyString(configInfo ConfigInfo) (http string, https string) {
	replacer := strings.NewReplacer(
		"username", configInfo.proxyInfo.username,
		"password", configInfo.proxyInfo.password,
		"url", configInfo.proxyInfo.proxyUrl,
		"port", configInfo.proxyInfo.port)

	h := replacer.Replace(PROXY_REPLACE_STRING)
	hs := h //a little bit of cheating

	return h, hs
}

//currently only checking for .npmrc
func getUserHomeDirectory() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	return currentUser.HomeDir
}

/*
------------------------------Set------------------------------
*/

func setProxyConfigVariables(configInfo ConfigInfo) (err error) {
	httpError := setWindowsVariables(configInfo.systemVariableHTTP_key, configInfo.proxyInfo.proxyHTTP_String)
	if httpError != nil {
		return httpError
	}

	httpsError := setWindowsVariables(configInfo.systemVariableHTTPS_key, configInfo.proxyInfo.proxyHTTPS_String)

	if httpsError != nil {
		return httpsError
	}

	return nil
}

func setWindowsVariables(key string, value string) (err error) {
	setEnvError := os.Setenv(key, value)
	return setEnvError
}

/*
------------------------------Update------------------------------
*/

func updateOrCreateProxyFile(configInfo ConfigInfo) (err error) {
	//kick off the http and https set commands.

	return nil
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
