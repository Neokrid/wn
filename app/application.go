package app

import (
	"errors"
	"fmt"
	"os"
)

var funcTable = map[string]menuFunction{
	"0": menuFunction{Title: "Выход", Function: StopApp},
	"1": menuFunction{Title: "Тестовая функция с ошибкой", Function: TestFuncWithError},
	"2": menuFunction{Title: "Тестовая успешная функция", Function: AccessFunc},
}

func RunApp() {
	fmt.Println("\033[33mДобро пожаловать \"Записки Ластоногих\"\033[0m")
	var command string
	for {
		fmt.Println("----------------------------")
		fmt.Print(generateMenu(funcTable))
		fmt.Scan(&command)
		fmt.Println("----------------------------")
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
		fmt.Println("\033[32mУспешно!\033[0m")
	}
}

func AccessFunc() error {
	return nil
}

func TestFuncWithError() error {
	return errors.New("тестовая ошибка")
}

func StopApp() error {
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
	return nil
}
