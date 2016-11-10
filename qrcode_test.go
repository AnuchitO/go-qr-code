package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCreateQRCode(t *testing.T) {
	png, err := CreateQRCode("http://www.google.com")
	if err != nil {
		t.Error(err.Error())
	}

	b := bytes.NewBuffer(png)
	fmt.Printf("% #v\n", b.String())
}
