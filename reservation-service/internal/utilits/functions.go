package utilits

import (
	"encoding/base64"
	"errors"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(QRdata string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(QRdata, qrcode.Medium, 256)
	if err != nil {
		return "", errors.New("Error generating QR Code: " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(png), nil
}
