package main

import (
	"log"
	"os"
	"testing"
)

var proxyInfoTest = ProxyInfo{"crolek", "sweetPassword", "chuckrolek.com", "80", "http://crolek:sweetPassword@chuckrolek.com:80", "http://crolek:sweetPassword@chuckrolek.com:80"}
var testHTTP = "PP_TEST_HTTP"
var testHTTP_Value = "testhttp"
var testHTTPS = "PP_TEST_HTTPS"
var testHTTPS_Value = "testhttps"

func TestCreateNewFile(t *testing.T) {
	var err error
	var isTestFileCreated bool
	var testConfiguration = NPM_Config
	testConfiguration.proxyInfo = proxyInfoTest
	testConfiguration.configFilePath, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	testConfiguration.configFilePath = testConfiguration.configFilePath + "/test_files/test_create_file.txt"
	createNewFile(testConfiguration)
	if _, err := os.Stat(testConfiguration.configFilePath); err != nil {
		if os.IsNotExist(err) {
			isTestFileCreated = false
		}
	} else {
		isTestFileCreated = true
	}

	IsTrueOrFalse(t, isTestFileCreated, true, "test file was created")
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

func TestDoesProxyFileExist(t *testing.T) {
	//currently performing a check for a false file until I have a better integraiton test
	//should return false
	IsTrueOrFalse(t, doesProxyFilesExist(".fileThatDoesNotExist"), false, "correctly detected the lack of a file")
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
