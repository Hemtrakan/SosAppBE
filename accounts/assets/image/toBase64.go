package image

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
)

func ImageToBase64() (Base64 string, Error error) {
	// Read the entire file into a byte slice
	f, _ := os.Open("./assets/image/profile.png")

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	//var base64Encoding string

	// Determine the content type of the image file
	//mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	//switch mimeType {
	//case "image/jpeg":
	//	base64Encoding += "data:image/jpeg;base64,"
	//case "image/png":
	//	base64Encoding += "data:image/png;base64,"
	//}

	// Append the base64 encoded output
	//base64Encoding += toBase64(bytes)
	Base64 = encoded
	// Print the full base64 representation of the image
	return
}
