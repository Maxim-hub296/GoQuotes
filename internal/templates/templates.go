package templates

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var Tmpl *template.Template

func LoadTemplates() {
	Tmpl = template.New("")

	// Проходим по всем файлам в папке templates и вложенных
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Только файлы с расширением .html
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err = Tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal("Ошибка при загрузке шаблонов:", err)
	}

	log.Println("Шаблоны загружены успешно")
}
