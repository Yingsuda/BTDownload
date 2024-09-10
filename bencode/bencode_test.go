package bencode

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncodeString(t *testing.T) {

	buf := new(bytes.Buffer)
	i := EncodeString(buf, "abc")
	fmt.Println("i:", i)

	fmt.Println(buf.String())
	// fmt.Printf("Data:% 02X\n", data)
	// fmt.Println(string(data[:i]))

	str, err := DecodeString(buf)
	if err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Println("STR:", str)
}
