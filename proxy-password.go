package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"

	//"github.com/howeyc/gopass" //can't get it to function with username, re-evaluate later
)

type ProxyInfo struct {
	username          string
	password          string
	proxyUrl          string
	port              string
	proxyHTTP_String  string
	proxyHTTPS_String string
}

var (
	proxyInfo ProxyInfo
	npmrcPath string
)

const (
	FILE_HTTP_START      = "proxy = "
	FILE_HTTPS_START     = "https-proxy = "
	NPMRC_LOCATION_START = "c:/Users/"
	PROXY_REPLACE_STRING = "http://username:password@url:port"
)

func main() {
	caputerUsernamePassword()
	//parseCommandlineFlags()
	fmt.Println("Starting the update process")
	buildProxyString()
	if doesProxyFilesExist("\\.npmrc") {
		//update file
		updateProxyFiles()
	} else {
		//create new file
		createNewFile(npmrcPath)
	}
	setWindowsVariables("HTTP_PROXY", proxyInfo.proxyHTTP_String)
	setWindowsVariables("HTTPS_PROXY", proxyInfo.proxyHTTPS_String)
	fmt.Println("Your proxy info has been updated. :)")
}

func caputerUsernamePassword() {
	fmt.Print("Username: ")
	fmt.Scan(&proxyInfo.username)
	fmt.Print("Password: ")
	fmt.Scan(&proxyInfo.password)
	fmt.Println(proxyInfo.username)
	fmt.Println(proxyInfo.password)
}

func parseCommandlineFlags() {
	flag.StringVar(&proxyInfo.username, "username", "username", "the username that the proxy account uses")
	flag.StringVar(&proxyInfo.password, "password", "password", "the password that the proxy account uses")
	flag.StringVar(&proxyInfo.proxyUrl, "url", "proxy.testing.com", "the url for the proxy you will be using")
	flag.StringVar(&proxyInfo.port, "port", "80", "the port number the proxy is using, usually its 80")

	flag.Parse()
}

func buildProxyString() {
	replacer := strings.NewReplacer(
		"username", proxyInfo.username,
		"password", proxyInfo.password,
		"url", proxyInfo.proxyUrl,
		"port", proxyInfo.port)

	proxyInfo.proxyHTTP_String = replacer.Replace(PROXY_REPLACE_STRING)
	proxyInfo.proxyHTTPS_String = proxyInfo.proxyHTTP_String //a little bit of cheating
}

func createNewFile(path string) {
	var data string

	data += FILE_HTTP_START + proxyInfo.proxyHTTP_String + "\n"
	data += FILE_HTTPS_START + proxyInfo.proxyHTTPS_String

	fmt.Println(data)
	err := ioutil.WriteFile(path, []byte(data), 0644)

	if err != nil {
		fmt.Println(err)
	}
}

func setWindowsVariables(key string, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		log.Println(err)
	}
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

//again, only checking for .npmrc
func updateProxyFiles() {
	var contents []byte
	contents, err := ioutil.ReadFile(npmrcPath)
	fileContents := string(contents)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(npmrcPath, []byte(updateUsernamePassword(fileContents, proxyInfo)), 0644)
}

func updateUsernamePassword(proxyString string, info ProxyInfo) string {
	regex := regexp.MustCompile("(https?://)(.*?):(.*?)(@.*)")
	results := regex.ReplaceAllString(proxyString, "${1}"+info.username+":"+info.password+"${4}")
	return results
}
