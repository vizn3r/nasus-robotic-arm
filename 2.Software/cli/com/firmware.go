package com

import (
	"fmt"
	"io"
	"net/http"
)

func Send(m string, p string) {
	res, err := http.Get("http://localhost" + p + "/" + m)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return
		}
		fmt.Println(string(b))
	}
}