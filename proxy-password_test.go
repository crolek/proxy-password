package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var testTempFileLocation = "test_files/test-file.dont-track"
var testHTTP_key = "PP_TEST_HTTP"
var testHTTPS_key = "PP_TEST_HTTPS"
var testProxyInfo = ProxyInfo{"crolek", "sweetPassword", "chuckrolek.com", "80", "", ""}
var testConfigInfo = ConfigInfo{
	configFilePath:          testTempFileLocation,
	systemVariableHTTP_key:  testHTTP_key,
	systemVariableHTTPS_key: testHTTPS_key,
	proxyInfo:               testProxyInfo,
}
var testHTTP_Value = "testhttp"
var testHTTPS_Value = "testhttps"
var testHTTP_ProxyString = "http://crolek:sweetPassword@chuckrolek.com:80"
var testHTTPS_ProxyString = testHTTP_ProxyString //yes, it's the same in this case.

func TestBuildConfig(t *testing.T) {
	//that lovely integration test
	removeTempFileLocation()
	resetTestUpdateProxyFile()
	resetTestSystemVariables()
	buildConfig(testConfigInfo)
	IsTrueOrFalse(t, doesFileExist(testConfigInfo.configFilePath), true, "BuildConfig() created the test proxy file")

	/*these won't work work until I absractd out the variables from setProxyConfigVariables()*/
	EqualString(t, os.Getenv(testHTTP_key), testHTTP_ProxyString, "BuildConfig() set HTTP_PROXY")
	EqualString(t, os.Getenv(testHTTPS_key), testHTTPS_ProxyString, "BuildConfig() set HTTPS_PROXY")

	removeTempFileLocation()
	resetTestUpdateProxyFile()
	resetTestSystemVariables()
}

func TestDoesFileExist(t *testing.T) {
	IsTrueOrFalse(t, doesFileExist(".fileThatDoesNotExist"), false, "correctly detected the lack of a file")

	err := ioutil.WriteFile(testTempFileLocation, []byte("sweet data"), 0644)
	if err != nil {
		log.Println(err)
		t.Fail() //the Fail() might be redudant, but i guess thats okay
	}
	IsTrueOrFalse(t, doesFileExist(testTempFileLocation), true, "correctly detected the there was a test file")

	removeTempFileLocation()
}

func TestCreateNewFile(t *testing.T) {
	var err error
	var testConfiguration = NPM_Config

	testConfiguration.proxyInfo = testProxyInfo

	testConfiguration.configFilePath = testTempFileLocation
	//remove the file if its there already
	removeTempFileLocation()

	createNewFile(testConfiguration.configFilePath, "some test content")
	isTestFileCreated := doesFileExist(testConfiguration.configFilePath)

	contents, err := ioutil.ReadFile(testConfiguration.configFilePath)
	fileContents := string(contents)
	if err != nil {
		t.Fail()
	}

	EqualString(t, fileContents, "some test content", "The correct data was writtent to a file")
	IsTrueOrFalse(t, isTestFileCreated, true, "test file was created")

	//remove the test file(s) to keep a clean testing area.
	removeTempFileLocation()
}

func TestGetProxyString(t *testing.T) {
	var testingConfig = NPM_Config

	testingConfig.proxyInfo = testProxyInfo
	httpResult, httpsResult := getProxyString(testingConfig)

	EqualString(t, httpResult, testHTTP_ProxyString, "properly built the http string from proxyInfo")
	EqualString(t, httpsResult, testHTTPS_ProxyString, "properly built the https string from proxyInfo")

}

func TestSetWindowsVariables(t *testing.T) {
	resetTestSystemVariables()
	setWindowsVariables(testHTTP_key, testHTTP_Value)
	EqualString(t, os.Getenv(testHTTP_key), testHTTP_Value, testHTTP_key+" was properly set")
	setWindowsVariables(testHTTPS_key, testHTTPS_Value)
	EqualString(t, os.Getenv(testHTTPS_key), testHTTPS_Value, testHTTPS_key+" was properly set")
	resetTestSystemVariables()
}

func TestUpdateProxyFile(t *testing.T) {
	//setting up a test file location
	testConfigInfo.configFilePath = "test_files/TestUpdateProxyFile-actual.npmrc"

	resetTestUpdateProxyFile()

	updateProxyFiles(testConfigInfo)
	contents, err := ioutil.ReadFile(testConfigInfo.configFilePath)

	if err != nil {
		log.Println(err)
	}

	actual := string(contents)

	contents, err = ioutil.ReadFile("test_files/TestUpdateProxyFile-expected.npmrc")
	expected := string(contents)

	if err != nil {
		t.Fail()
	}

	EqualString(t, actual, expected, "proxy file was updated")

	resetTestUpdateProxyFile()
}

func TestUpdateUsernamePassword(t *testing.T) {
	EqualString(t, updateUsernamePassword(PROXY_REPLACE_STRING, testProxyInfo), "http://crolek:sweetPassword@url:port", "updating clean username/password")
}

/*
------------------------------Utils------------------------------
*/

func removeTempFileLocation() {
	_ = os.Remove(testTempFileLocation)
}

func resetTestUpdateProxyFile() {
	//resetting the file for the test
	contents, err := ioutil.ReadFile("test_files/TestUpdateProxyFile-reset.npmrc")

	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile("test_files/TestUpdateProxyFile-actual.npmrc", contents, 0644)

	if err != nil {
		log.Println(err)
	}
}

func resetTestSystemVariables() {
	err := os.Setenv(testHTTP_key, "")
	if err != nil {
		log.Println(err)
	}
	err = os.Setenv(testHTTPS_key, "")
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
