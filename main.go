package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	CallClear()
	for {
		var input int
		fmt.Println("Добро пожаловать в игры для компании!")
		fmt.Println("Выберите игру в которую хотите поиграть")
		fmt.Println("1. Правда или действие")
		fmt.Println("2. Крокодил")
		fmt.Println("3. Я никогда не")
		fmt.Println("4. Бутылочка")
		fmt.Println("5. Узнать об играх")
		fmt.Println("0. Выход")
		fmt.Println("Ваш выбор: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Неверный ввод")
			break
		} else if input != 5 && input != 0 {
			fmt.Println("Отличный выбор :)")
			time.Sleep(1 * time.Second)
		}
		CallClear()
		switch input {
		case 1:
			AnimateLoadingBar()
			CallClear()
			trueOrActionPlay()
		case 2:
			AnimateLoadingBar()
			CallClear()
			crocodilePlay()
		case 3:
			AnimateLoadingBar()
			CallClear()
			iveNeverPlay()
		case 4:
			AnimateLoadingBar()
			CallClear()
			botlePlay()
		case 5:
			about()
		case 0:
			CallClear()
			return
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

func about() {
	fmt.Println("О какой игре хотите узнать подробнее?")
	choise := 0
	ad := ""
	fmt.Println("1. Правда или действие")
	fmt.Println("2. Крокодил")
	fmt.Println("3. Я никогда не")
	fmt.Println("4. Бутылочка")
	fmt.Println("0. Выход")
	fmt.Scan(&choise)
	switch choise {
	case 1:
		CallClear()
		fmt.Println("Игра правда или действие")
		fmt.Println("Правда или Действие – это игра, созданная для развлечения и непринужденного общения.\nЦель игры - узнать своих друзей или знакомых с новой стороны, раскрепоститься и повеселиться вместе. \nБлагодаря выбору между Правдой и Действием игра дает возможность либо открыться и рассказать о себе, либо проявить себя в смешном и необычном амплуа. \nГлавное - помнить, что игра должна приносить удовольствие всем участникам.")
		fmt.Println("---------------------------")
		fmt.Println("Чтобы вернуться в меню выбора игры нажмите y, иначе a")
		fmt.Scan(&ad)
		if ad == "y" {
			CallClear()
			return
		} else {
			CallClear()
			about()
		}
	case 2:
		CallClear()
		fmt.Println("Игра Крокодил")
		fmt.Println("Крокодил - это популярная командная игра, основанная на пантомиме. \nОдин игрок получает слово или фразу и должен изобразить их без использования слов, звуков или предметов. \nОстальные игроки должны угадать, что именно показывается. Игра требует находчивости, актерских способностей и умения понимать друг друга без слов. \nКрокодил отлично подходит для развлечения на вечеринках, в семейном кругу или в качестве тимбилдинга.")
		fmt.Println("---------------------------")
		fmt.Println("Чтобы вернуться в меню выбора игры нажмите y, иначе a")
		fmt.Scan(&ad)
		if ad == "y" {
			CallClear()
			return
		} else {
			CallClear()
			about()
		}
	case 3:
		CallClear()
		fmt.Println("Игра Я никогда не")
		fmt.Println("---------------------------")
		fmt.Println("Я никогда не… – это популярная игра, которая помогает сблизиться и лучше узнать своих друзей или знакомых. \nИгроки по очереди начинают фразу со слов \"Я никогда не…\" и рассказывают о том, чего они никогда не делали. \nЕсли кто-то из других игроков совершал то, о чём говорит ведущий, он должен что-то сделать (например, загнуть палец, выпить глоток напитка и т.д.). \nИгра отлично подходит для вечеринок, посиделок с друзьями и для создания непринужденной атмосферы.")
		fmt.Println("Чтобы вернуться в меню выбора игры нажмите y, иначе a")
		fmt.Scan(&ad)
		if ad == "y" {
			CallClear()
			return
		} else {
			CallClear()
			about()
		}
	case 4:
		CallClear()
		fmt.Println("Игра Бутылочка")
		fmt.Println("---------------------------")
		fmt.Println("Бутылочка – это классическая игра, в которой участники сидят в кругу, а в центре располагается бутылка. \nОдин из игроков вращает бутылку, и когда она останавливается, горлышко или донышко бутылки указывает на двух игроков. \nЭти двое должны выполнить заранее оговоренное действие, чаще всего – поцеловаться. Игра часто используется для создания романтической или кокетливой атмосферы \nи может включать в себя разные задания.")
		fmt.Println("Чтобы вернуться в меню выбора игры нажмите y, иначе a")
		fmt.Scan(&ad)
		if ad == "y" {
			CallClear()
			return
		} else {
			CallClear()
			about()
		}
	case 0:
		CallClear()
		return
	default:
		CallClear()
		fmt.Println("Такой игры нет")
		time.Sleep(1 * time.Second)
		CallClear()
		return
	}
}
