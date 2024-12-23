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

func trueOrActionPlay() {
	fmt.Println("Добро пожаловать в игру правда или действие")
	fmt.Println("Чтобы начать играть введите количество игроков, а затем напишите их имена")
	var input int
	fmt.Scan(&input)
	var a []player
	for i := 0; i < input; i++ {
		var name string
		fmt.Scan(&name)
		a = append(a, createForActionOrTrue(name))
	}
	CallClear()
	fmt.Println("Итоговый список игроков:")
	for i, name := range a {
		fmt.Printf("%v. %v\n", i+1, name.name)
	}
	fmt.Println("Начнем!")
	time.Sleep(3 * time.Second)
	CallClear()
	turn := 0
	for {
		if turn > len(a)-1 || turn < 0 {
			turn = 0
		}
		var choise string
		fmt.Println("Ход игрока: ", a[turn].name)
		fmt.Println("Правда или действе t/a")
		fmt.Scan(&choise)
		switch choise {
		case "t":
			if a[turn].truth > 2 {
				fmt.Println("Вы не можете выбрать Правда")
				time.Sleep(1 * time.Second)
				turn--
				CallClear()
				continue
			} else {
				a[turn].truth++
				filePath := "truth.json"
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

				question := data[strconv.Itoa(rand.Intn(52))]

				fmt.Println("Вам достался вопрос: " + question)
				fmt.Println("Продолжаем? y/n")
				turn++
				var contin string
				fmt.Scan(&contin)
				switch contin {
				case "y":
					CallClear()
					continue
				case "n":
					CallClear()
					return

				}
			}
		case "a":
			a[turn].truth = 0
			filePath := "action.json"
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

			question := data[strconv.Itoa(rand.Intn(62))]

			fmt.Println("Вам досталось действие: " + question)
			fmt.Println("Продолжаем? y/n")
			turn++
			var contin string
			fmt.Scan(&contin)
			switch contin {
			case "y":
				CallClear()
				continue
			case "n":
				CallClear()
				return
			}
		default:
			fmt.Println("Некорректнйы выбор")
			time.Sleep(1 * time.Second)
			CallClear()
		}
	}
}
