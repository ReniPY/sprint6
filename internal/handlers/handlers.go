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

// ServeIndex служит для отображения HTML-страницы index.html
func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "C:\\Users\\Данил\\Desktop\\Go\\sprint6\\index.html")
	}
}

// UploadHandler обрабатывает форму и загружаемый файл
func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20) // Ограничиваем размер формы

		// Читаем файл
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			logger.Printf("Ошибка обработки файла: %v\n", err)
			http.Error(w, "Ошибка обработки файла", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Чтение содержимого файла
		buffer := new(bytes.Buffer)
		_, err = buffer.ReadFrom(file)
		if err != nil {
			logger.Printf("Ошибка чтения файла: %v\n", err)
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}

		// Конвертируем содержимое
		input := buffer.String()
		converted, err := service.AutoConvert(input)
		if err != nil {
			logger.Printf("Ошибка преобразования: %v\n", err)
			http.Error(w, "Ошибка преобразования", http.StatusInternalServerError)
			return
		}
		// Формируем корректное имя файла
		filename := time.Now().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)

		// Создаем файл
		outFile, err := os.Create(filename)
		if err != nil {
			logger.Printf("Ошибка создания файла: %v\n", err)
			http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
			return
		}
		defer outFile.Close()

		// Записываем преобразованное содержимое в файл
		_, err = io.Copy(outFile, bytes.NewReader([]byte(converted)))
		if err != nil {
			logger.Printf("Ошибка записи в файл: %v\n", err)
			http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
			return
		}

		// Сообщаем пользователю об успешном завершении
		w.Write([]byte("Файл успешно обработан."))
	}
}
