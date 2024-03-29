package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

type ErrorDetail struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	isPasswordCorrect := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if isPasswordCorrect != nil && isPasswordCorrect == bcrypt.ErrMismatchedHashAndPassword {
		return isPasswordCorrect
	}

	if isPasswordCorrect != nil {
		return isPasswordCorrect
	}

	return nil
}

func FileTypeValidation(filetype string) bool {
	filetypes := []string{"image/jpeg", "image/png", "image/jpg"}
	lowerTheFiletype := strings.ToLower(filetype)
	for _, val := range filetypes {
		if lowerTheFiletype == val {
			return true
		}
	}
	return false
}

// AppResponse is for response config show to Frontend side
func AppResponse(w http.ResponseWriter, code int, data interface{}) {
	response := Response{
		Success: code <= 400,
		Data:    data,
		Error:   &ErrorDetail{},
	}
	if !response.Success {
		response.Data = nil
		response.Error = &ErrorDetail{
			Code:    code,
			Message: fmt.Sprintf("%v", data),
		}
	}

	jsonResponse, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(code)
	w.Write(jsonResponse)
}
