package main

import (
	"log"
	"testing"

	"github.com/crolek/proxy-password/proxy"
	"github.com/crolek/proxy-password/ui"
)

func TestMain(t *testing.T) {
	log.Println("we goood")
	ui.Tango()
	proxy.ProxyTango()
}
