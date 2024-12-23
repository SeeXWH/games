package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func crocodilePlay() {
	fmt.Println("Добро пожаловать в игру крокодил! Побеждает тот, кто наберет 10 очков")
	fmt.Println("Чтобы начать играть введите количество игроков, а затем напишите их имена")
	var input int
	fmt.Scan(&input)
	var a []player
	for i := 0; i < input; i++ {
		var name string
		fmt.Scan(&name)
		a = append(a, createForCrocodile(name))
	}
	CallClear()
	fmt.Println("Итоговый список игроков:")
	for i, name := range a {
		fmt.Printf("%v. %v\n", i+1, name.name)
	}
	fmt.Println("Начнем!")
	time.Sleep(3 * time.Second)
	CallClear()
	rang := 0
	var i int
	for {
		if rang > len(a)-1 {
			rang = 0
		}
		fmt.Printf("Ход достается игроку: %v", a[rang])
		fmt.Println()
		fmt.Println("Когда будете готовы введите уровень сложности 1/2/3")
		fmt.Scan(&i)
		CallClear()
		fmt.Println("Ход игрока " + a[rang].name)
		done := make(chan bool)
		go displayTimer(5, done)
		displayQuestions(&a[rang], i)

		timerRunning = false
		<-done
		timerRunning = true
		rang++
		var flag = false
		var p int
		for j := 0; j < len(a); j++ {
			if a[j].stats >= 10 {
				flag = true
				p = j
			}
		}
		if flag {
			fmt.Printf("Наш победитель набравший %v балло: %v", a[p].stats, a[p].name)
			fmt.Println()
			time.Sleep(3 * time.Second)
			return
		}
		fmt.Println("Продолжаем? y/n")
		var choise string
		if choise == "y" {
			continue
		} else if choise == "n" {
			return
		}
	}
}

var timerRunning bool = true

func displayTimer(minutes int, done chan<- bool) {
	duration := time.Duration(minutes) * time.Minute
	startTime := time.Now()
	endTime := startTime.Add(duration)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if !timerRunning {
			done <- true
			return
		}
		remaining := endTime.Sub(time.Now())
		if remaining <= 0 {
			fmt.Println("\r00:00")
			fmt.Println("Время вышло!")
			timerRunning = false
			done <- true
			return
		}
		m := int(remaining.Minutes())
		s := int(remaining.Seconds()) % 60
		fmt.Printf("\rТаймер: %02d:%02d", m, s)
	}
}

func displayQuestions(p *player, various int) {

	var filePath string
	switch various {
	case 1:
		filePath = "crocodileEasy.json"
	case 2:
		filePath = "crocodileMedium.json"
	case 3:
		filePath = "crocodileHard.json"
	default:
		fmt.Println("Неккоректный вывод")
		return
	}
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

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 154 && timerRunning; i++ {
		fmt.Println("\n")
		fmt.Println("--------------------------------------")
		fmt.Println(data[strconv.Itoa(rand.Intn(154))])
		fmt.Println("--------------------------------------")
		fmt.Println("Кто то угадал? y/n: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "стоп" {
			timerRunning = false
			return
		} else if input == "y" {
			p.stats = p.stats + various
		}
		CallClear()
	}
	if timerRunning {
		fmt.Println("\nВопросы закончились")
		timerRunning = false
	}
}
