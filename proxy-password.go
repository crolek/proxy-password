package proxypassword

import (
	//"log"
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
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
)

const (
	FILE_HTTP_START      = "proxy = "
	FILE_HTTPS_START     = "https-proxy = "
	NPMRC_LOCATION_START = "c:/Users/"
	PROXY_REPLACE_STRING = "http://username:password@url:port"
)

func MockMain() {
	testPath := "C:\\dev\\go_dev\\src\\main\\.npmrc"
	bindData()
	buildProxyString()
	createNewFile(testPath)
	setWindowsVariables()
}

//proxy = http://username:password@url:80
//https-proxy
//c:\Users\%USERNAME%\.npmrc
func bindData() {

	//for testing
	proxyInfo.username = "Chuck"
	proxyInfo.password = "testing123"
	proxyInfo.proxyUrl = "proxy.testing.com"
	proxyInfo.port = "80"
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
