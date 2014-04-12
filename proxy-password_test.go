package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var proxyInfoTest = ProxyInfo{"crolek", "sweetPassword", "chuckrolek.com", "80", "http://crolek:sweetPassword@chuckrolek.com:80", "http://crolek:sweetPassword@chuckrolek.com:80"}
var testHTTP = "PP_TEST_HTTP"
var testHTTP_Value = "testhttp"
var testHTTP_SetString = testHTTP + "=" + testHTTP_Value
var testHTTPS = "PP_TEST_HTTPS"
var testHTTPS_Value = "testhttps"
var testHTTPS_SetString = testHTTPS + "=" + testHTTPS_Value

func TestSetWindowsVariables(t *testing.T) {
	resetTestSystemVariables()

	//setting HTTP
	setWindowsVariables(testHTTP, testHTTP_Value)
	results, _ := WindowsCMD_Contains("set "+testHTTP_SetString, testHTTP_Value)
	IsTrueOrFalse(t, results, true, testHTTP+" was not set properly")
	//setting HTTPS
	setWindowsVariables(testHTTPS_Value, testHTTPS_Value)
	results, _ = WindowsCMD_Contains("set "+testHTTPS_SetString, testHTTPS_Value)
	IsTrueOrFalse(t, results, true, testHTTPS+" was not set properly")
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
	//resetting the test Variables
	results, err := WindowsCMD_Contains("set "+testHTTP+"=", "%"+testHTTP+"%")
	if err != nil {
		log.Println(err)
	}
	results, err = WindowsCMD_Contains("set "+testHTTPS+"=", "%"+testHTTPS+"%")
	if err != nil {
		log.Println(err)
	}
	results, err = WindowsCMD_Contains("echo %"+testHTTP+"%", "%"+testHTTP+"%")
	if results == false || err != nil {
		log.Println(err)
	}
	results, err = WindowsCMD_Contains("echo %"+testHTTPS+"%", "%"+testHTTPS+"%")
	if results == false || err != nil {
		log.Println(err)
	}
}

func WindowsCMD_Contains(command string, contains string) (IsInOutput bool, outputError error) {
	output, _, err := WindowsCMD(command) //maybe i shouldn't burry the second param?

	if err != nil {
		return false, err
	}

	if strings.Contains(output, contains) {
		return true, nil
	} else {
		return false, nil
	}

}

func WindowsCMD(command string) (consoleOutput string, consoleError string, cmdErr error) {
	var outputBuffer bytes.Buffer
	var errorBuffer bytes.Buffer

	//cmd := exec.Command("echo", "\\%HTTP_PROXY\\%")
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer
	err := cmd.Run()

	return outputBuffer.String(), errorBuffer.String(), err
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
