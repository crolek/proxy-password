package proxypassword

import (
	//"log"
	"fmt"
	"io/ioutil"
	"strings"
)

type ProxyInfo struct {
	username string
	password string
	proxyUrl string
	port     string
}

var (
	proxyInfo = new(ProxyInfo)
)

func MockMain() {
	testPath := "C:\\dev\\go_dev\\src\\main\\.npmrc"
	bindData()
	createNewFile(testPath)
}

const (
	FILE_HTTP_START      = "proxy = "
	FILE_HTTPS_START     = "https-proxy = "
	NPMRC_LOCATION_START = "c:/Users/"
	PROXY_REPLACE_STRING = "http://username:password@url:port"
)

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

func createNewFile(path string) {
	var data string

	replacer := strings.NewReplacer("username", proxyInfo.username,
		"password", proxyInfo.password,
		"url", proxyInfo.proxyUrl,
		"port", proxyInfo.port)

	proxyString := replacer.Replace(PROXY_REPLACE_STRING)

	data += FILE_HTTP_START + proxyString + "\n"
	data += FILE_HTTPS_START + proxyString

	fmt.Println(data)
	err := ioutil.WriteFile(path, []byte(data), 0644)

	if err != nil {
		fmt.Println(err)
	}
}
