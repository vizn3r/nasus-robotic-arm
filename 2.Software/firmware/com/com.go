package com

import "fmt"

type COM struct {
	isEnabled bool
	isRunning bool
}

func (com *COM) IsRunning() bool {
	return com.isRunning
}

func (com *COM) IsEnabled() bool {
	return com.isEnabled
}

func (com *COM) Enable() {
	com.isEnabled = true
}

func (com *COM) Disable() {
	com.isEnabled = false
}

func (com *COM) Start(s string) bool {
	if !com.isEnabled {
		fmt.Println("'" + s + "' is not enabled.")
		return false
	} else if com.isRunning{
		fmt.Println("'" + s + "' is already running")
		return false
	}
	fmt.Println("Starting '" + s + "'")
	return true
}