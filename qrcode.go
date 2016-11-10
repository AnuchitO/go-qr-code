package main

import qrcode "github.com/skip2/go-qrcode"

func CreateQRCode(data string) ([]byte, error) {
	png, err := qrcode.Encode(data, qrcode.Medium, 256)
	return png, err
}
