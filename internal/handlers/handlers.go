package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../index.html")
	}
}

func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			http.Error(w, "Ошибка обработки файла", http.StatusBadRequest)
			return
		}
		defer file.Close()

		buffer := new(bytes.Buffer)
		_, err = buffer.ReadFrom(file)
		if err != nil {
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		input := buffer.String()
		converted, err := service.AutoConvert(input)
		if err != nil {
			http.Error(w, "Ошибка преобразования", http.StatusInternalServerError)
			return
		}

		// Формируем уникальное имя файла
		filename := time.Now().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)

		// Создаем файл и записываем результат
		outFile, err := os.Create(filename)
		if err != nil {
			http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, bytes.NewReader([]byte(converted)))
		if err != nil {
			http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
			return
		}

		// Возвращаем результат пользователю
		w.Write([]byte("Результат: " + converted))
	}
}
