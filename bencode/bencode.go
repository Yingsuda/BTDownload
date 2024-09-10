package bencode

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"time"
)

// type BType uint8

// const (
// 	BSTR  BType = 0x01
// 	BINT  BType = 0x02
// 	BLIST BType = 0x03
// 	BDIST BType = 0x04
// )

// type obj struct {
// 	val_ interface{}
// }

// type Btype interface {
// 	Bstr | Bint | BList | BDist
// }

// type Bstr struct {
// 	val_ string
// }

// type Bint struct {
// 	val_ int
// }

// type BList struct {
// 	val_ []*obj
// }

// type BDist struct {
// 	val_ map[string]*obj
// }

// type Bobject[bt Btype] struct {
// 	value_ bt
// }

// func (b *Bobject[bt]) GetValue() bt {
// 	return b.value_
// }

type BBT interface {
	string | int | []*Bobject | map[string]*Bobject
}

type Bobject struct {
	//type_  BType
	value_ interface{}
}

func GetObjValue[bbt BBT](obj *Bobject) bbt {
	return obj.value_.(bbt)
}

func SetObjValue[bbt BBT](obj *Bobject, val bbt) {
	obj.value_ = val
}

var ErrB = fmt.Errorf("BObject change Errors")

// func Value[bv BValue](obj *Bobject) bv {
// 	return obj.value_.(bv)
// }

// func (b *Bobject) Str() (string, error) {
// 	if b.type_ != BSTR {
// 		return "", ErrB
// 	}
// 	return b.value_.(string), nil
// }

// func (b *Bobject) Int() (int, error) {
// 	if b.type_ != BINT {
// 		return 0, ErrB
// 	}
// 	return b.value_.(int), nil
// }

// func (b *Bobject) List() ([]*Bobject, error) {
// 	if b.type_ != BLIST {
// 		return nil, ErrB
// 	}
// 	return b.value_.([]*Bobject), nil
// }

// func (b *Bobject) Dist() (map[string]*Bobject, error) {
// 	if b.type_ != BDIST {
// 		return nil, ErrB
// 	}
// 	return b.value_.(map[string]*Bobject), nil
// }

func writeDecimal(bw *bufio.Writer, sl int) int {
	buf := []byte(strconv.Itoa(sl))
	bw.Write(buf)
	return len(buf)
}

func EncodeString(w io.Writer, val string) int {
	strLen := len(val)
	bw := bufio.NewWriter(w)
	wlen := writeDecimal(bw, strLen)
	bw.WriteByte(':')
	wlen++
	bw.WriteString(val)
	wlen += strLen

	err := bw.Flush()
	if err != nil {
		return 0
	}
	return wlen
}

func DecodeString(r io.Reader) (string, error) {
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}

	strlen := 0
	num := ""
	for {
		b, err := br.Peek(1)
		if err != nil {
			return "", err
		}
		if b[0] >= '0' && b[0] <= '9' {
			bb, err := br.ReadByte()
			if err != nil {
				return "", err
			}
			num += string(bb)

		} else {
			strlen, err = strconv.Atoi(num)
			if err != nil {
				return "", err
			}
			break
		}
		time.Sleep(time.Second)
	}

	b1, err := br.ReadByte()
	if err != nil {
		return "", err
	}

	if b1 != ':' {
		return "", fmt.Errorf(" not string")
	}

	b2 := make([]byte, strlen)
	_, err = io.ReadAtLeast(br, b2, strlen)
	if err != nil {
		return "", err
	}
	return string(b2), nil
}
