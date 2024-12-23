package main

import "fmt"

type player struct {
	name  string
	stats int
	truth int
}

func (p player) String() string {
	return fmt.Sprintf("Имя: %s, баллы: %d", p.name, p.stats)
}

func createForCrocodile(name string) player {
	return player{name, 0, 0}
}

func createForActionOrTrue(name string) player {
	return player{name, 0, 0}
}
