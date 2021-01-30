package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	strTest := "I love this beautiful world!"
	strEncrypted := "98b4fc4538115c4980a8b859ff3d27e1"
	fmt.Println(Check(strTest, strEncrypted))
}

func Check(content, encrypted string) bool {
	return strings.EqualFold(Encode(content), encrypted)
}
func Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
