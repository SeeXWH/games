package main

import (
	"fmt"
	"math/rand"
	"time"
)

func botlePlay() {
	fmt.Println("Добро пожаловать в игру Бутылочка")
	fmt.Println("Чтобы начать играть введите количество игроков, а затем напишите их имена")
	var input int
	fmt.Scan(&input)
	var a []string
	for i := 0; i < input; i++ {
		var name string
		fmt.Scan(&name)
		a = append(a, name)
	}
	CallClear()
	fmt.Println("Итоговый список игроков:")
	for i, name := range a {
		fmt.Printf("%v. %v\n", i+1, name)
	}
	fmt.Println("Начнем!")
	time.Sleep(3 * time.Second)
	CallClear()
	for {
		rn := rand.Intn(len(a))
		fmt.Printf("И сейчас крутит бутылочку: %v", a[rn])
		fmt.Println()
		time.Sleep(2 * time.Second)
		bottleSpin(rn, a)

		fmt.Println("Продолжим? 1/0")
		fmt.Scan(&input)
		if input == 1 {
			CallClear()
			continue
		} else {
			CallClear()
			break
		}

	}
}

func bottleSpin(spinnerIndex int, players []string) {
	playersWithoutSpinner := append([]string{}, players[:spinnerIndex]...)
	playersWithoutSpinner = append(playersWithoutSpinner, players[spinnerIndex+1:]...)
	index := 0
	spins := rand.Intn(len(playersWithoutSpinner)*2) + len(playersWithoutSpinner)
	delay := 50
	for i := 0; i < spins; i++ {
		CallClear() // Очистка экрана
		fmt.Println("Игра в бутылочку!")
		fmt.Println("Бутылочка крутится...")
		for j := 0; j < len(playersWithoutSpinner); j++ {
			if j == index {
				fmt.Printf("> %s <\n", playersWithoutSpinner[j])
			} else {
				fmt.Printf("  %s\n", playersWithoutSpinner[j])
			}
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)
		index = (index + 1) % len(playersWithoutSpinner)
		if spins-i < len(playersWithoutSpinner) {
			delay += 25
		}
	}
	if index-1 < 0 {
		index = len(playersWithoutSpinner) - 1
	} else {
		index--
	}
	fmt.Printf("\nБутылочка указала на: %s\n", playersWithoutSpinner[index])
}
