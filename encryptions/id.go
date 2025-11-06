package encryptions

import "accounts.workflecks.com/constants"

func EncryptID(text string) (string, error) {
	return EncryptAES(text, constants.EncryptionAESIDKey)
}

func DecryptID(text string) (string, error) {
	return DecryptAES(text, constants.EncryptionAESIDKey)
}

func Encrypt(text string) (string, error) {
	return EncryptAES(text, constants.EncryptionAESKey)
}

func Decrypt(text string) (string, error) {
	return DecryptAES(text, constants.EncryptionAESKey)
}

func EncryptWithRounds(text string) (string, error) {
	return EncryptAESWithRounds(text, constants.EncryptionAESKey, constants.EncryptionAESRounds)
}

func DecryptWithRounds(text string) (string, error) {
	return DecryptAESWithRounds(text, constants.EncryptionAESKey, constants.EncryptionAESRounds)
}
