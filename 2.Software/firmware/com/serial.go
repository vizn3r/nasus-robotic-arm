package com

import "sync"

type SERIAL struct {
	COM
}

func StartSERIAL(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
}