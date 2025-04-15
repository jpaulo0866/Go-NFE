package utils

import "github.com/skip2/go-qrcode"

func GenerateQRCode(data string, filePath string) error {
	return qrcode.WriteFile(data, qrcode.Medium, 256, filePath)
}
