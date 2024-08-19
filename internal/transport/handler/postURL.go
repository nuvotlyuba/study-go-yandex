package handler

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/study-go-yandex/internal/models"
)

func (h Handler) PostURL(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "text/plain" {
		http.Error(w, "Неверный тип данных", http.StatusBadRequest)
	}
	if r.Method == http.MethodPost {
		bytesData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
			return
		}
		strData := string(bytesData)

		shortURL, err := h.service.CreateURL(models.URL(strData).Point())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, string(*shortURL))
	}
}
