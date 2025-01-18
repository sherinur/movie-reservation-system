package utilits

import (
	"encoding/base64"
	"errors"
	"github.com/skip2/go-qrcode"
	"net/smtp"
	"os"
)

func GenerateQR(QRdata string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(QRdata, qrcode.Medium, 256)
	if err != nil {
		return "", errors.New("Error generating QR Code: " + err.Error())
	}

	return base64.StdEncoding.EncodeToString(png), nil
}

func SendMail(to, subject, content string) error {
	from := "dbeysembaev@gmail.com"
	password := os.Getenv("PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + content)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
