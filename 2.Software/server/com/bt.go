package com

import (
	"fmt"
	"server/util"
	"strconv"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func hasDevice(arr []bluetooth.ScanResult, device bluetooth.ScanResult) bool {
	for _, d := range arr {
		if d.Address.MAC.String() == device.Address.MAC.String() {
			return true
		}
	}
	return false
}

func ConnectBT() {
	devices := []bluetooth.ScanResult{}
	err := adapter.Enable()
	if err != nil {
		fmt.Println(err)
	}
	adapter.SetConnectHandler(func(device bluetooth.Address, connected bool) {
		if connected {
			fmt.Println("Connected to", device.String())
		} else {
			fmt.Println("Couldn't connect to", device.String())
		}
	})
	println("Scanning...")
	i := 0
	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if len(devices) == 0 {
			devices = append(devices, device)
		}
		if !hasDevice(devices, device) {
			devices = append(devices, device)
		}
		if i == 50 {
			adapter.StopScan()
		}
		i++
	})
	if err != nil {
		fmt.Println(err)
	}
	names := []string{}
	for _, v := range devices {
		names = append(names, v.LocalName() + " " + v.Address.String())
	}
	if len(devices) == 0 {
		fmt.Println("No devices found.")
	}
	addr := devices[util.UserIndexPrompt("What device do you want to connect to", names)].Address
	dev, err := adapter.Connect(addr, bluetooth.ConnectionParams{ConnectionTimeout: 1000,})
	if err != nil {
		fmt.Println(err)
	}
	srvcs, err := dev.DiscoverServices(nil)
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 255)
	for _, srvc := range srvcs {
		println("- service", srvc.UUID().String())

		chars, err := srvc.DiscoverCharacteristics(nil)
		if err != nil {
			println(err)
		}
		for _, char := range chars {
			println("-- characteristic", char.UUID().String())
			n, err := char.Read(buf)
			if err != nil {
				println("    ", err.Error())
			} else {
				println("    data bytes", strconv.Itoa(n))
				println("    value =", string(buf[:n]))
			}
		}
	}
	err = dev.Disconnect()
	if err != nil {
		println(err)
	}
}

func PrintBT() {

}