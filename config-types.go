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
	FILE_HTTP_COMMAND       string
	FILE_HTTPS_COMMAND      string
	configFilePath          string
	configFileName          string
	configFilePathIsUserDir bool
	systemVariableHTTP_key  string
	systemVariableHTTPS_key string
	proxyInfo               ProxyInfo
}

var NPM_Config = ConfigInfo{
	FILE_HTTP_COMMAND:       "npm config set proxy ",
	FILE_HTTPS_COMMAND:      "npm config set https-proxy ",
	configFileName:          ".npmrc",
	configFilePathIsUserDir: true,
}

func SetDefaultConfigurations() {

}
