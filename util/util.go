package util

import (
	"fmt"
	"os"
	"strings"

	"template/service/custom_error"
	"template/service/logger"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func EncryptedText(text string) string {
	return text
}

func DecryptedText(text string) string {
	return text
}

func IsBusinessError(err error) bool {
	isBusinessError := false

	if err == nil {
		return isBusinessError
	}

	if _, ok := err.(custom_error.BusinessError); ok {
		isBusinessError = true
	}

	return isBusinessError
}

func ParseValidtionErrorToString(err error) (bool, string) {
	isValidationError := false
	message := ""

	if err == nil {
		return isValidationError, message
	}

	switch castType := err.(type) {
	case validator.ValidationErrors:
		isValidationError = true
		for _, validateError := range castType {
			message += fmt.Sprintf("field: %s validate error on tag: %s\n", validateError.Field(), validateError.Tag())
		}
	}

	return isValidationError, strings.TrimSpace(message)
}

func InitUtil() {
	// Init Cipher
	key := viper.GetString("Encryption.Key")
	if key == "" {
		logger.Logger.Error("Failed to get encryption key")
		os.Exit(1)
	}
}
