package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const TIMER int8 = 24

func getEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Panic(".env файл не найден")
		return "", err
	}

	if api, exists := os.LookupEnv("API_KEY"); !exists {
		log.Panic("Записи об api ключе не существует")
		return "", errors.New("Записи об api ключе не существует")

	} else {
		return api, nil
	}

}

func main() {

	log.Println("Начало работы")
	log.Println("Получение данных из .env файла")

	api_key, err := getEnv()
	if err != nil {
		log.Fatal("Невозможно продолжить работу программы")
	}
	log.Println("api ключ (", api_key, ") получен.")

	// TODO: инициализация бота

	// TODO: вызов парсера
}
