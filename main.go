package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/totoafreeca/goencrypter/des"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}
func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

func byteArrayToString(bytes []byte) string {
	out := ""
	for _, n := range bytes {
		out += fmt.Sprintf("%08b", n)
	}
	return out
}

var ErrRange = errors.New("Value out of range")

func bitStringToBytes(s string) ([]byte, error) {
	b := make([]byte, (len(s)+(8-1))/8)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '1' {
			return nil, ErrRange
		}
		b[i>>3] |= (c - '0') << uint(7-i&7)
	}
	return b, nil
}

func main() {

	flags := flag.String("type", "encrypt")

	//message := "0100100101000101010011110100011001001001010101000010001100110001"
	//String is the same as Message

	stringKey := "IEOFIT#1"
	binaryKey := stringToBin(stringKey)

	//message := "TESTIT#1"
	//messageBin := stringToBin(message)
	//fmt.Printf("KEY: %s is %s in HEX\n", stringKey, parseBinToHex(stringToBin(stringKey)))
	//fmt.Printf("MSG: %s is %s in HEX\n\n", message, parseBinToHex(stringToBin(message)))

	//binaryKey := stringToBin(stringKey)

	// encrypter := des.NewDesEncrypter()
	// encryptedMsg := encrypter.Encrypt(binaryKey, messageBin)
	// fmt.Println("Encrypted msg: " + encryptedMsg)
	// fmt.Println("Encrypted msg HEX: " + parseBinToHex(encryptedMsg))

	// decrypter := des.NewDesDecrypter()
	// decryptedMsg := decrypter.Decrypt(binaryKey, encryptedMsg)
	// fmt.Println("Decrypted msg: " + decryptedMsg)
	// fmt.Println("Decrypted msg HEX: " + parseBinToHex(decryptedMsg))

	//Opening and reading

	encrypter := des.NewDesEncrypter()
	decrypter := des.NewDesDecrypter()

	file, err := os.Open("2.bin")
	defer file.Close()
	check(err)

	stats, statsErr := file.Stat()
	check(statsErr)

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	//Writing encrypted (bytes are read data)

	//opening
	f, err := os.Create("encrypted2.bin")
	check(err)

	for i := int64(0); i < size; i += 8 {
		fileMSG := bytes[i : i+8]

		val := encrypter.Encrypt(binaryKey, byteArrayToString(fileMSG))

		encryptedBytes, err := bitStringToBytes(val)
		check(err)

		f.Write(encryptedBytes)
		f.Sync()
	}
	f.Close()

	file, err = os.Open("encrypted2.bin")
	defer file.Close()
	check(err)

	stats, statsErr = file.Stat()
	check(statsErr)

	size = stats.Size()
	bytes = make([]byte, size)

	bufr = bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	//opening
	f, err = os.Create("decrypted2.bin")
	check(err)

	for i := int64(0); i < size; i += 8 {
		fileMSG := bytes[i : i+8]

		val := decrypter.Decrypt(binaryKey, byteArrayToString(fileMSG))

		decryptedBytes, err := bitStringToBytes(val)
		check(err)

		f.Write(decryptedBytes)
		f.Sync()
	}
	f.Close()

}
