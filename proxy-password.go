package proxypassword

import (
	//"log"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
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
	proxyInfo = new(ProxyInfo)
	npmrcPath string
)

const (
	FILE_HTTP_START      = "proxy = "
	FILE_HTTPS_START     = "https-proxy = "
	NPMRC_LOCATION_START = "c:/Users/"
	PROXY_REPLACE_STRING = "http://username:password@url:port"
)

func MockMain() {
	parseCommandlineFlags()
	buildProxyString()
	if doesProxyFilesExist() {
		//update file
		updateProxyFiles()
	} else {
		//create new file
		createNewFile(npmrcPath)
	}
	setWindowsVariables()
}

//proxy = http://username:password@url:80
//https-proxy
//c:\Users\%USERNAME%\.npmrc
func parseCommandlineFlags() {
	flag.StringVar(&proxyInfo.username, "username", "username", "the username that the proxy account uses")
	flag.StringVar(&proxyInfo.password, "password", "password", "the password that the proxy account uses")
	flag.StringVar(&proxyInfo.proxyUrl, "url", "proxy.testing.com", "the url for the proxy you will be using")
	flag.StringVar(&proxyInfo.port, "port", "80", "the port number the proxy is using, usually its 80")

	flag.Parse()
}

func buildProxyString() {
	replacer := strings.NewReplacer("username", proxyInfo.username,
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

//This is a hack to get the variables set to the System and not just the instance
//of this program.
func setWindowsVariables() {
	var out bytes.Buffer
	cmd := exec.Command("setx", "HTTP_PROXY", proxyInfo.proxyHTTP_String, "/m")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("setx", "HTTPS_PROXY", proxyInfo.proxyHTTPS_String, "/m")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out.String())
}

//currently only checking for .npmrc
func doesProxyFilesExist() bool {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	npmrcPath = currentUser.HomeDir + "\\.npmrc"

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

	err = ioutil.WriteFile(npmrcPath, []byte(updatePassword(fileContents)), 0644)
}

//this needs symbol support for finding the password
func updatePassword(proxyString string) string {
	//i'm horrible at regex so this will do for now.
	regex := regexp.MustCompile("(:)((?:[a-z][a-z0-9_]*))(@)")
	results := regex.ReplaceAllLiteralString(proxyString, ":"+proxyInfo.password+"@")
	return results
}
