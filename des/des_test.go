package des

import (
	"fmt"
	"testing"
)

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return
}

func TestPermuteIP(t *testing.T) {
	fmt.Println("TESTING IP PERMUTATION")
	message := "0100100101000101010011110100011001001001010101000010001100110001"
	permutedByIP := permute(message, IPTable)

	if permutedByIP != "0011111110100000001011101101011100000000110000000001010101001100" {
		t.Errorf("Message wrongly permuted")
	}

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Permuted: %s\n", permutedByIP)

}

func TestKeyGeneration(t *testing.T) {

	fmt.Println("TESTING KEY GENERATION")

	stringKey := "IEOFIT#1"
	binaryKey := stringToBin(stringKey)
	keys := generateKeys(binaryKey)

	if keys[0] != "111100001001001010100010100010011010010000010011" {
		t.Errorf("Wrong key number 0")
	}
	if keys[15] != "111000001000001000101010110000100011001010011000" {
		t.Errorf("Wrong key number 15")
	}
}
