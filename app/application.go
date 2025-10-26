package app

import (
	"bufio"
	"fmt"
	"os"
)

var funcTable = map[string]menuFunction{
	"0": menuFunction{Title: "Выход", Function: StopApp},
	"1": menuFunction{Title: "Создать заметку", Function: CreateNotes},
	"2": menuFunction{Title: "Просмотр всех  заметок", Function: ListNotes},
	"3": menuFunction{Title: "Редактирование заметок", Function: EditNotes},
}

func RunApp() {
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
