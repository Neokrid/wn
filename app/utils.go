package app

import "fmt"

type menuFunction struct {
	Title    string
	Function func() error
}

func generateMenu(commands map[string]menuFunction) string {
	menu := "Меню:\n"
	for k, v := range commands {
		menu += fmt.Sprintf("%s - %s\n", k, v.Title)
	}
	menu += "Введите номер команды: "
	return menu
}
