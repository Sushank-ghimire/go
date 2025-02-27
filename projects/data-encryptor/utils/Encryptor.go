package Encryption

const (
	ORIGINAL_KEYS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~"
	SHIFT         = 5
)

func EncryptData(data string, shift int) string {
	encrypted := make([]byte, len(data))
	for i, char := range data {
		index := findIndex(ORIGINAL_KEYS, char)
		if index != -1 {
			encrypted[i] = ORIGINAL_KEYS[(index+SHIFT)%len(ORIGINAL_KEYS)]
		} else {
			encrypted[i] = byte(char)
		}
	}
	return string(encrypted)
}

func DecryptData(encryptedData string, shift int) string {
	decrypted := make([]byte, len(encryptedData))
	for i, char := range encryptedData {
		index := findIndex(ORIGINAL_KEYS, char)
		if index != -1 {
			decrypted[i] = ORIGINAL_KEYS[(index-SHIFT+len(ORIGINAL_KEYS))%len(ORIGINAL_KEYS)]
		} else {
			decrypted[i] = byte(char)
		}
	}
	return string(decrypted)
}

func findIndex(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1
}
