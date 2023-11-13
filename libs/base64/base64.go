package base64

import (
	b64 "encoding/base64"
	"fmt"
)

func Encode(str string) string {
	return b64.URLEncoding.EncodeToString(([]byte)(str))
}

func Decode(str string, target *string) error {
	byteArray, err := b64.URLEncoding.DecodeString(str)
	if err != nil {
		return err
	}
	final := string(byteArray)
	fmt.Println("final", final)
	*target = final
	return nil
}
