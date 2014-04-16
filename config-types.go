package main

type ProxyInfo struct {
	username          string
	password          string
	proxyUrl          string
	port              string
	proxyHTTP_String  string
	proxyHTTPS_String string
}

type ConfigInfo struct {
	FILE_HTTP_START  string
	FILE_HTTPS_START string
	configFilePath   string
	proxyInfo        ProxyInfo
}

var NPM_Config = ConfigInfo{
	FILE_HTTP_START:  "proxy = ",
	FILE_HTTPS_START: "https-proxy = ",
}

func SetDefaultConfigurations() {

}
