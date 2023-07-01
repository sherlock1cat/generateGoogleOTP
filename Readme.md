# generateGoogleOTP

This is a simple Go program that generates a QR code for a TOTP secret key.

## Usage

To use this program, run the following command:

```
go run main.go <totp_secret>
```

Replace `<totp_secret>` with the TOTP secret key that you want to generate a QR code for.

The program will output a data URL for a PNG image of the QR code to the console. Copy the `data:image/png;base64...` part of the URL and paste it into your browser to view the QR code.

The TOTP secret key will also be printed to the console. Make sure to keep this key secret and store it in a safe place.


## Dependencies

This program uses the following third-party packages:

- [dgoogauth](https://github.com/dgryski/dgoogauth) for generating TOTP keys
- [rsc/qr](https://github.com/rsc/qr) for generating QR codes
- [boombuler/barcode](https://github.com/boombuler/barcode) for scaling QR codes

Make sure to install these packages before running the program.