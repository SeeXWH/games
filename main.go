package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		var input int
		fmt.Println("Добро пожаловать в игры для компании!")
		fmt.Println("Выберите игру в которую хотите поиграть")
		fmt.Println("1. Правда или действие")
		fmt.Println("2. Крокодил")
		fmt.Println("3. Я никогда не")
		fmt.Println("4. Бутылочка")
		fmt.Println("Ваш выбор: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Неверный ввод")
			break
		} else {
			fmt.Println("Отличный выбор :)")
		}
		time.Sleep(1 * time.Second)
		CallClear()
		AnimateLoadingBar()
		CallClear()

		switch input {
		case 1:
		case 2:
		case 3:
			iveNeverPlay()
		case 4:
			botlePlay()
		default:
			fmt.Println("Такого выбора нет")
			break
		}
	}
}

func AnimateLoadingBar() {
	totalDivisions := 20
	delay := 80 * time.Millisecond

	os.Stdout.WriteString("[")
	for completedDivisions := 0; completedDivisions <= totalDivisions; completedDivisions++ {

		os.Stdout.WriteString("\r[")
		for i := 0; i < completedDivisions; i++ {
			os.Stdout.WriteString("#")
		}
		for i := 0; i < totalDivisions-completedDivisions; i++ {
			os.Stdout.WriteString("-")
		}
		os.Stdout.WriteString("]")

		time.Sleep(delay)
	}
	fmt.Println("\nЗавершено!")
	time.Sleep(1 * time.Second)
}
