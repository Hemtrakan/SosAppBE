package assets

import (
	"bufio"
	"encoding/base64"
	"io"
	"os"
)

func ImageToBase64(pathImageName string) (Base64 string) {
	// Read the entire file into a byte slice
	//f, _ := os.Open("./assets/image/profile.png")
	f, _ := os.Open(pathImageName)

	reader := bufio.NewReader(f)
	content, _ := io.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	Base64 = encoded
	return
}
