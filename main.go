package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/dgryski/dgoogauth"
	"image/png"
	"os"
)

func generateQRCode(secret string) string {
	// Generate TOTP key
	otp := dgoogauth.OTPConfig{
		Secret:      secret,
		HotpCounter: 0,
	}
	qrCodeURL := otp.ProvisionURIWithIssuer("Sherlock1cat", "generateGoogleOTP")

	qrCode, err := qr.Encode(qrCodeURL, qr.H, qr.Auto)
	if err != nil {
		panic(err)
	}
	// Generate PNG image bytes
	bCode, err := barcode.Scale(qrCode, 300, 300)
	if err != nil {
		panic(err)
	}
	// Create PNG image buffer
	var buf bytes.Buffer
	err = png.Encode(&buf, bCode)
	if err != nil {
		panic(err)
	}
	// Encode image as base64 string
	imgData := base64.StdEncoding.EncodeToString(buf.Bytes())
	// Return QR code data
	return fmt.Sprintf("data:image/png;base64,%s", imgData)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <totp_secret>")
		return
	}
	// Generate QR code for TOTP secret
	qrCode := generateQRCode(os.Args[1])
	fmt.Println(qrCode)
	fmt.Println("Please copy the data:image/png;base64... and open it in your browser. Please keep the Secret_key safe and secure.")
	fmt.Println("Secret_key:" + os.Args[1])
}
