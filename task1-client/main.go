package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	apiURL       = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	updatePeriod = 10 * time.Minute // Время обновления курса валюты == 10 минут
)

// Структура для хранения данных о крипте
type CoinData struct {
	ID     string  `json:"id"`
	Symbol string  `json:"symbol"`
	Name   string  `json:"name"`
	Price  float64 `json:"current_price"`
}

func main() {
	var wg sync.WaitGroup
	for {
		// Получаем мапу доступных криптовалют
		cryptoIDMap := getAvailableCryptos()
		// Выводим таблицу доступных криптовалют
		printCryptoTable(cryptoIDMap)

		// Получаем ввод пользователя
		selectedCrypto := getUserChoice(cryptoIDMap)
		if selectedCrypto != nil {
			wg.Add(1)
			go monitorCrypto(*selectedCrypto, &wg)
		}

		fmt.Println("\nWaiting for the next update...")
		fmt.Println("-----------------------------")
		time.Sleep(updatePeriod)
	}
}

// Получаем доступные криптовалюты из API
func getAvailableCryptos() map[string]string {
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var coins []CoinData
	if err := json.Unmarshal(body, &coins); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	// Создаем мапу для хранения названий криптовалют и ID (порядковый номер)
	cryptoIDMap := make(map[string]string)
	for _, coin := range coins {
		cryptoIDMap[coin.Name] = coin.ID
	}

	return cryptoIDMap
}

// Выводим таблицу доступных криптовалют
func printCryptoTable(cryptoIDMap map[string]string) {
	fmt.Printf("%-3s | %-15s\n", "ID", "Cryptocurrency")
	fmt.Println("-------------------------------")

	var keys []string
	for key := range cryptoIDMap {
		keys = append(keys, key)
	}

	// Сортируем названия криптовалют по алфавиту
	sort.Strings(keys)

	// Выводим отсортированные названия с их порядковыми номерами
	for i, key := range keys {
		fmt.Printf("%-3d | %-15s\n", i+1, key)
	}
}

// Получаем ввод пользователя
func getUserChoice(cryptoIDMap map[string]string) *CoinData {
	var choiceStr string
	fmt.Print("Enter the ID of the cryptocurrency you want to monitor: ")
	fmt.Scanln(&choiceStr)

	// Обработка ошибок при парсинге строки в число
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice <= 0 {
		fmt.Println("Invalid input. Please enter a valid number.")
		return getUserChoice(cryptoIDMap)
	}

	var keys []string
	for key := range cryptoIDMap {
		keys = append(keys, key)
	}

	// Сортируем названия криптовалют по алфавиту
	sort.Strings(keys)

	// Проверяем, что ввод пользователя валиден
	if choice >= 1 && choice <= len(keys) {
		selectedCryptoID := cryptoIDMap[keys[choice-1]]
		return &CoinData{ID: selectedCryptoID, Name: keys[choice-1], Symbol: ""}
	} else {
		fmt.Println("Invalid choice. Please enter a valid number.")
		return getUserChoice(cryptoIDMap)
	}
}

// Мониторим выбранную криптовалюту
func monitorCrypto(coin CoinData, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Monitoring cryptocurrency with %s (%s)..\n", coin.Name, coin.ID)
	fmt.Println("-----------------------------")

	for {
		price := getCryptoPrice(coin.ID)
		currentTime := time.Now().Format("15:04:05")
		fmt.Printf("[%s] %s: %.2f USD\n", currentTime, coin.Name, price)
		waitForUpdate()
	}
}

// Получаем текущую цену выбранной криптовалюты
func getCryptoPrice(cryptoID string) float64 {
	apiURL := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", cryptoID)
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return 0.0
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0.0
	}

	var priceData map[string]map[string]interface{}
	if err := json.Unmarshal(body, &priceData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0.0
	}

	if priceMap, ok := priceData[cryptoID]; ok {
		if price, ok := priceMap["usd"]; ok {
			switch v := price.(type) {
			case float64:
				return v
			case string:
				priceFloat, err := strconv.ParseFloat(v, 64)
				if err != nil {
					fmt.Println("Error converting price to float:", err)
					return 0.0
				}
				return priceFloat
			default:
				fmt.Println("Unexpected type for price:", reflect.TypeOf(price))
				fmt.Println("Full API response:", string(body))
			}
		} else {
			fmt.Println("USD price not found in the API response.")
			fmt.Println("Full API response:", string(body))
		}
	} else {
		fmt.Println("Crypto ID not found in the API response.")
		fmt.Println("Full API response:", string(body))
	}

	return 0.0
}

// Ожидание n-го времени для повторного вывода актуального курса
func waitForUpdate() {
	time.Sleep(updatePeriod)
}