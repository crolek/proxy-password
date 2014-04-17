package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var proxyInfoTest = ProxyInfo{"crolek", "sweetPassword", "chuckrolek.com", "80", "", ""}
var testHTTP = "PP_TEST_HTTP"
var testHTTP_Value = "testhttp"
var testHTTPS = "PP_TEST_HTTPS"
var testHTTPS_Value = "testhttps"
var testHTTP_ProxyString = "http://crolek:sweetPassword@chuckrolek.com:80"
var testHTTPS_ProxyString = testHTTP_ProxyString //yes, it's the same in this case.

func TestBuildConfig(t *testing.T) {
	//that lovely integration test
}

func TestBuildProxyInfo(t *testing.T) {
	var testingConfig = NPM_Config

	testingConfig.proxyInfo = proxyInfoTest
	httpResult, httpsResult := buildProxyString(testingConfig)

	EqualString(t, httpResult, testHTTP_ProxyString, "properly built the http string from proxyInfo")
	EqualString(t, httpsResult, testHTTPS_ProxyString, "properly built the https string from proxyInfo")

}

func TestCreateNewFile(t *testing.T) {
	var err error
	var testConfiguration = NPM_Config
	testConfiguration.proxyInfo = proxyInfoTest
	testConfiguration.configFilePath, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	testConfiguration.configFilePath = testConfiguration.configFilePath + "/test_files/test_create_file.txt"
	//remove the file if its there already
	_ = os.Remove(testConfiguration.configFilePath)

	createNewFile(testConfiguration)
	isTestFileCreated := doesFileExist(testConfiguration.configFilePath)

	IsTrueOrFalse(t, isTestFileCreated, true, "test file was created")

	//remove the test file(s) to keep a clean testing area.
	_ = os.Remove(testConfiguration.configFilePath)
}

func TestSetWindowsVariables(t *testing.T) {
	resetTestSystemVariables()
	setWindowsVariables(testHTTP, testHTTP_Value)
	EqualString(t, os.Getenv(testHTTP), testHTTP_Value, testHTTP+" was properly set")
	setWindowsVariables(testHTTPS, testHTTPS_Value)
	EqualString(t, os.Getenv(testHTTPS), testHTTPS_Value, testHTTPS+" was properly set")
}

func TestUpdateUsernamePassword(t *testing.T) {
	EqualString(t, updateUsernamePassword(PROXY_REPLACE_STRING, proxyInfoTest), "http://crolek:sweetPassword@url:port", "updating clean username/password")
}

func TestDoesFileExist(t *testing.T) {
	var testFileLocation = "test_files/test-file.txt"

	IsTrueOrFalse(t, doesFileExist(".fileThatDoesNotExist"), false, "correctly detected the lack of a file")

	err := ioutil.WriteFile(testFileLocation, []byte("sweet data"), 0644)
	if err != nil {
		log.Println(err)
		t.Fail() //the Fail() might be redudant, but i guess thats okay
	}
	IsTrueOrFalse(t, doesFileExist(testFileLocation), true, "correctly detected the there was a test file")

	_ = os.Remove(testFileLocation)
}

func resetTestSystemVariables() {
	err := os.Setenv(testHTTP, "")
	if err != nil {
		log.Println(err)
	}
	err = os.Setenv(testHTTPS, "")
	if err != nil {
		log.Println(err)
	}
}

func IsTrueOrFalse(t *testing.T, actual bool, expected bool, message string) {
	var result string
	if actual == expected {
		t.Log("Passed - " + message)
	} else {
		//quick bool to string
		if actual {
			result = "true"
		} else {
			result = "false"
		}
		t.Log("Failed - " + message)
		t.Log("Actual: " + result)
		t.Fail()
	}
}

func EqualString(t *testing.T, actual string, expected string, message string) {
	if actual == expected {
		t.Log("Passed - " + message)
	} else {
		t.Log("Failed - " + message)
		t.Log("Actual:   " + actual)
		t.Log("Expected: " + expected)
		t.Fail()
	}
}
