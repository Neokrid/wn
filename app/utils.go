package app

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type menuFunction struct {
	Title    string
	Function func() error
}

type Note struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	CreateDate time.Time `json:"create_date"`
}

func generateMenu(commands map[string]menuFunction) string {
	keys := make([]int, 0)
	menu := "Меню:\n"
	for k, _ := range commands {
		key, _ := strconv.Atoi(k)
		keys = append(keys, key)

	}
	sort.Ints(keys)
	for _, v := range keys {

		menu += fmt.Sprintf("%s - %s\n", strconv.Itoa(v), commands[strconv.Itoa(v)].Title)
	}
	menu += "Введите номер команды: "
	return menu
}
