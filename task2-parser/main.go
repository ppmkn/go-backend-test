package main

import (
	"encoding/csv"
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Открываем CSV файл для записи
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создаем писатель CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовки CSV
	headers := []string{"Rank", "Username", "Name", "Category", "Followers", "Country", "Eng.(Auth.)", "Eng.(Avg.)"}
	if err := writer.Write(headers); err != nil {
		log.Fatal(err)
	}

	// URL страницы для парсинга
	url := "https://hypeauditor.com/top-instagram-all-russia/"

	// Создаем HTTP-клиент
	client := &http.Client{}

	// Отправляем GET-запрос
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Парсим каждый блок данных
	doc.Find(".row__top").Each(func(i int, s *goquery.Selection) {
		// Инициализируем пустой слайс для хранения данных строки
		rowData := make([]string, 0)

		// Добавляем Rank
		rank := i + 1
		rowData = append(rowData, strconv.Itoa(rank))

		// Добавляем Username
		username := s.Find(".row-cell.contributor .contributor__name-content").Text()
		rowData = append(rowData, username)

		// Добавляем Name
		name := strings.TrimSpace(s.Find(".row-cell.contributor .contributor__title").Text())
		rowData = append(rowData, name)

		// Добавляем Category
		category := s.Find(".row-cell.category .tag__content").Text()
		rowData = append(rowData, category)

		// Добавляем Followers
		subscribers := s.Find(".row-cell.subscribers").Text()
		rowData = append(rowData, subscribers)

		// Добавляем Country
		country := s.Find(".row-cell.audience").Text()
		rowData = append(rowData, country)

		// Добавляем Eng.(Auth.)
		authentic := s.Find(".row-cell.authentic").Text()
		rowData = append(rowData, authentic)

		// Добавляем Eng.(Avg.)
		avgEngagement := s.Find(".row-cell.engagement").Text()
		rowData = append(rowData, avgEngagement)

		// Записываем данные в CSV файл
		if err := writer.Write(rowData); err != nil {
			log.Fatal(err)
		}
	})

	fmt.Println("CSV file created successfully!")
}
