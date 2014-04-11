package main

import (
	//"bytes"
	//"log"
	//"os"
	//"os/exec"
	"testing"
)

var proxyInfoTest = ProxyInfo{"crolek", "sweetPassword", "chuckrolek.com", "80", "http://crolek:sweetPassword@chuckrolek.com:80", "http://crolek:sweetPassword@chuckrolek.com:80"}

func TestUpdateUsernamePassword(t *testing.T) {

	EqualString(t, updateUsernamePassword(PROXY_REPLACE_STRING, proxyInfoTest), "http://crolek:sweetPassword@url:port", "updating clean username/password")
}

func TestDoesProxyFileExist(t *testing.T) {
	//currently performing a check for a false file until I have a better integraiton test
	//should return false
	IsTrueOrFalse(t, doesProxyFilesExist(".fileThatDoesNotExist"), false, "correctly detected the lack of a file")
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
