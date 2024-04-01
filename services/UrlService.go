package services

import (
	"bytes"
	"github.com/go-chi/render"
	"math/rand"
	"net/http"
)

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// Base62 потому что это Base64 без двух символов "=/"
func encodeBase62(num uint32) string {
	length := len(alphabet)
	var result bytes.Buffer
	for ; num > 0; num /= uint32(length) {
		result.WriteByte(alphabet[(num % uint32(length))])
	}
	return result.String()
}

func CreateShortUrl(writer http.ResponseWriter, request *http.Request) {
	// Получаем url из request's body
	buf := new(bytes.Buffer)
	_, err2 := buf.ReadFrom(request.Body)
	if err2 != nil {
		return
	}

	url := buf.String()
	randomInt := rand.Uint32()
	shortUrl := encodeBase62(randomInt)

	// На случай коллизии генерируем повторно
	for ItemExists(shortUrl) {
		randomInt = rand.Uint32()
		shortUrl = encodeBase62(randomInt)
	}

	// Когда убедились, что коллизии нет - сохраняем
	SetToStorage(shortUrl, url)

	_, err := writer.Write([]byte(shortUrl))
	if err != nil {
		return
	}
}

func GetFullUrl(writer http.ResponseWriter, request *http.Request) {
	url := request.Context().Value("url").(string)

	result := GetFromStorage(url)
	// Если такой сокращённый url существует - возвращаем его, иначе 404
	if len(result) > 0 {
		writer.Write([]byte(result))
	} else {
		render.Status(request, 404)
		render.PlainText(writer, request, "")
	}
}
