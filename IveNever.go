package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func iveNeverPlay() {
	filePath := "never.json"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}
	var data map[string]string
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("Ошибка десериализации JSON: %v", err)
	}
	fmt.Println("Добро пожаловать в игру Я никогда не!")
	fmt.Println("Давайте начинать!")
	time.Sleep(2 * time.Second)
	CallClear()
	for {
		varibal := rand.Intn(155)
		fmt.Println("Вопрос:")
		fmt.Println("Я никогда не " + data[strconv.Itoa(varibal)])
		var input string
		fmt.Println("Продолжаем? y/n")
		fmt.Scan(&input)
		if input == "y" {
			CallClear()
			continue
		} else {
			CallClear()
			return
		}
	}

}
