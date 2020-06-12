package main

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	readyName := strings.Replace("rmq::queue::[{queue}]::ready", "{queue}", "tasks", 1)
	println(readyName)
}
