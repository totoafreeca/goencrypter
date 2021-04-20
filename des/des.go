package des

import "bytes"

type Encrypter interface {
	Encrypt(key string, message string)
}

type Decrypter interface {
	Decrypt()
}

type desCypher struct {
}

func (d *desCypher) Encrypt(key string, message string) {

	message = permute(message, IPTable)
	keys := GenerateKeys(key)
	left, right := message[0:32], message[32:]

	for i := 0; i < 16; i++ {

		rightExp := permute(right, extensionTable)

		xored := xorStrings(rightExp)
	}
}

func (d *desCypher) Decrypt() {

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
