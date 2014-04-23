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
	fileHTTPCommand         string
	fileHTTP_value          string
	fileHTTPSCommand        string
	fileHTTPS_value         string
	FILE_HTTP_START         string
	FILE_HTTPS_START        string
	configFilePath          string
	configFileName          string
	configFilePathIsUserDir bool
	systemVariableHTTP_key  string
	systemVariableHTTPS_key string
	proxyInfo               ProxyInfo
}

var NPM_Config = ConfigInfo{
	fileHTTPCommand:         "npm config set proxy ",
	fileHTTPSCommand:        "npm config set https-proxy ",
	configFileName:          ".npmrc",
	configFilePathIsUserDir: true,
}

func SetDefaultConfigurations() {

}
