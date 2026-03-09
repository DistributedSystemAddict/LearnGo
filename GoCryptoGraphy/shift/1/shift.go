package shift

func Encipher(plaintext []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] + 1
	}
	return ciphertext
}
