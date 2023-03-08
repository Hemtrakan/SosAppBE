package image

import (
	"encoding/base64"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ImageToBase64() (Base64 string, Error error) {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile("profile.png")
	if err != nil {
		Error = err
		return
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)
	Base64 = base64Encoding
	// Print the full base64 representation of the image
	return
}
