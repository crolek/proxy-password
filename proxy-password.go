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
	configInfo.proxyInfo.proxyHTTP_String, configInfo.proxyInfo.proxyHTTPS_String = buildProxyString(configInfo)
	createNewFile(configInfo)
	setProxyConfigVariables(configInfo)

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

func setProxyConfigVariables(configInfo ConfigInfo) {
	setWindowsVariables("HTTP_PROXY", configInfo.proxyInfo.proxyHTTP_String)
	setWindowsVariables("HTTPS_PROXY", configInfo.proxyInfo.proxyHTTPS_String)
}

func setWindowsVariables(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		log.Println(err)
	}
}

func updateOrCreateProxyFile(configInfo ConfigInfo) (status string, err error) {
	if doesProxyFilesExist(configInfo.configFilePath) {
		updateProxyFiles()
	}

	//return a new error saying update/create file failed
}

//currently only checking for .npmrc
func doesProxyFilesExist(fileName string) bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	npmrcPath := currentUser.HomeDir + fileName

	if _, err := os.Stat(npmrcPath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createNewFile(configInfo ConfigInfo) {
	var data string

	data += configInfo.FILE_HTTP_START + configInfo.proxyInfo.proxyHTTP_String + "\n"
	data += configInfo.FILE_HTTPS_START + configInfo.proxyInfo.proxyHTTPS_String

	log.Println(data)
	err := ioutil.WriteFile(configInfo.configFilePath, []byte(data), 0644)

	if err != nil {
		log.Println(err)
	}
}

//again, only checking for .npmrc
func updateProxyFiles(configInfo ConfigInfo) {
	var contents []byte
	contents, err := ioutil.ReadFile(npmrcPath)
	fileContents := string(contents)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(npmrcPath, []byte(updateUsernamePassword(fileContents, configInfo.proxyInfo)), 0644)
}

func updateUsernamePassword(proxyString string, info ProxyInfo) string {
	regex := regexp.MustCompile("(https?://)(.*?):(.*?)(@.*)")
	results := regex.ReplaceAllString(proxyString, "${1}"+info.username+":"+info.password+"${4}")
	return results
}
