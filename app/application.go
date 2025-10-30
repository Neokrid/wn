package app

import (
	"bufio"
	"fmt"
	"os"
)

type ChestNotes struct {
	notes []Note
}

func NewChestNotes() *ChestNotes {
	return &ChestNotes{
		notes: make([]Note, 0),
	}
}

func (c *ChestNotes) RunApp() {
	funcTable := map[string]menuFunction{
		"0": {"Выход", func() error { c.Stop(); return nil }},
		"1": {"Создать заметку", c.Create},
		"2": {"Просмотр всех заметок", c.List},
		"3": {"Редактирование заметок", c.Edit},
	}
	fmt.Println("\033[33mДобро пожаловать \"Записки Ластоногих\"\033[0m")
	var command string
	for {
		fmt.Println("----------------------------")
		fmt.Print(generateMenu(funcTable))
		fmt.Scan(&command)
		fmt.Println("----------------------------")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		targetF, ok := funcTable[command]
		if !ok {
			fmt.Println("\033[31mКоманда не найдена\033[0m")
			continue
		}
		err := targetF.Function()
		if err != nil {
			fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())

			continue
		}

	}
}
