package des

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

type Encrypter interface {
	Encrypt(key string, message string) string
}

type Decrypter interface {
	Decrypt(key string, message string) string
}

type desCypher struct {
}

func getDecNumber(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal("Error while converting string to int")
	}
	return out
}

func (d *desCypher) Encrypt(key string, message string) string {

	message = permute(message, IPTable)
	keys := generateKeys(key)
	left, right := message[0:32], message[32:]

	for i := 0; i < 16; i++ {

		rightExp := permute(right, extensionTable)

		xored := xorStrings(rightExp, keys[i])

		var buffer bytes.Buffer
		for j := 0; j < 48; j += 6 {
			tmp := xored[j : j+6]
			num := int(j / 6)

			row, _ := strconv.ParseInt(string(tmp[0])+string(tmp[5]), 2, 8)
			col, _ := strconv.ParseInt(string(tmp[1:5]), 2, 8)

			sBoxValue := SBoxes[num][row][col]
			s := fmt.Sprintf("%.4b", sBoxValue)
			buffer.WriteString(s)
		}

		sBoxOutputString := buffer.String()
		sBoxOutputString = permute(sBoxOutputString, PTable)

		result := xorStrings(left, sBoxOutputString)

		left = result

		left, right = right, left

		//fmt.Println(i+1, " ", left, " ", right)
	}

	combined := right + left
	encipheredText := permute(combined, finalPermutation)

	return encipheredText
}

func (d *desCypher) Decrypt(key string, message string) string {

	message = permute(message, IPTable)
	keys := generateKeys(key)
	left, right := message[0:32], message[32:]

	for i := 0; i < 16; i++ {

		rightExp := permute(right, extensionTable)

		xored := xorStrings(rightExp, keys[len(keys)-i-1])

		var buffer bytes.Buffer
		for j := 0; j < 48; j += 6 {
			tmp := xored[j : j+6]
			num := int(j / 6)

			row, _ := strconv.ParseInt(string(tmp[0])+string(tmp[5]), 2, 8)
			col, _ := strconv.ParseInt(string(tmp[1:5]), 2, 8)

			sBoxValue := SBoxes[num][row][col]
			s := fmt.Sprintf("%.4b", sBoxValue)
			buffer.WriteString(s)
		}

		sBoxOutputString := buffer.String()
		sBoxOutputString = permute(sBoxOutputString, PTable)

		result := xorStrings(left, sBoxOutputString)

		left = result

		left, right = right, left

		//fmt.Println(i+1, " ", left, " ", right)
	}

	combined := right + left
	encipheredText := permute(combined, finalPermutation)

	return encipheredText
}

func NewDesEncrypter() Encrypter {
	return &desCypher{}
}

func NewDesDecrypter() Decrypter {
	return &desCypher{}
}

func permute(in string, permutationArray []int) string {

	var buffer bytes.Buffer

	for i := 0; i < len(permutationArray); i++ {
		buffer.WriteByte(in[permutationArray[i]-1])
	}
	return buffer.String()

}

func shiftLeft(in string, shifts int) string {

	var buffer bytes.Buffer
	for i := 0; i < shifts; i++ {
		for j := 1; j < 28; j++ {
			buffer.WriteByte(in[j])
		}
		buffer.WriteByte(in[0])
		in = buffer.String()
		buffer.Reset()
	}
	return in
}

func xorStrings(a string, b string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			buffer.WriteString("0")
		} else {
			buffer.WriteString("1")
		}
	}
	return buffer.String()
}

func generateKeys(key string) []string {
	var keys [16]string

	key = permute(key, PC1Table)
	for i := 0; i < 16; i++ {
		key = shiftLeft(key[0:28], ShiftTable[i]) + shiftLeft(key[28:], ShiftTable[i])

		keys[i] = permute(key, PC2Table)
	}
	return keys[:]
}
