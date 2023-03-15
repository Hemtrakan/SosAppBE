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
	
	encoded := base64.StdEncoding.EncodeToString(content)

	Base64 = encoded
	return
}
