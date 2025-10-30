package app

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func (c *ChestNotes) Load() error {
	data, err := os.ReadFile(note_path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
		}
		c.notes = make([]Note, 0)
		return nil
	}

	var notes []Note
	if err := json.Unmarshal(data, &notes); err != nil {
		fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
		c.notes = make([]Note, 0)
		return err
	}

	c.notes = notes
	return nil
}

func EnterValue(msg string) string {
	fmt.Print(msg)
	var enterString string
	fmt.Scanln(&enterString)
	return enterString
}

func (c *ChestNotes) Create() error {
	var note Note
	maxID := 0
	for _, n := range c.notes {
		if n.ID > maxID {
			maxID = n.ID
		}
	}
	note.ID = maxID + 1
	note.Title = EnterValue("Введите название заметки:")
	note.Body = EnterValue("Введите текст заметки:")
	fmt.Printf("\033[35mВведите текст заметки:\n\033[0m")
	c.notes = append(c.notes, note)
	return nil
}

func (c *ChestNotes) Save(filePath string) error {
	data, err := json.MarshalIndent(c.notes, "", "  ")
	if err != nil {
		fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func Crop(s string) string {
	newString := []rune(s)
	if len(newString) <= 20 {
		return s
	}
	return string(newString[:17]) + "..."
}

func (c *ChestNotes) List() error {
	var selected *Note
	for _, note := range c.notes {
		fmt.Printf("%v - %-20s - %s\n", note.ID, Crop(note.Title), note.CreateDate.Format("02-01-2006"))
	}
	fmt.Printf("\n\033[35m0 - Назад\n\033[0m")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		value, err := strconv.Atoi(input)
		if err != nil || value < 0 || value > len(c.notes) {
			return errors.New("введите номер заметки")
		}

		if value == 0 {
			return nil
		}

		for i := range c.notes {
			if c.notes[i].ID == value {
				selected = &c.notes[i]
				break
			}
		}
		if selected == nil {
			fmt.Println("Заметка не найдена.")
			continue
		}

		fmt.Printf("%v - %-20s - %s\n", selected.ID, selected.Title, selected.CreateDate.Format("02-01-2006"))
		for _, line := range LineWrap(selected.Body) {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("\n\033[35m0 - Назад\n\033[0m")
	}
	return nil
}

func LineWrap(text string) []string {
	var res []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{""}
	}
	currentLine := words[0]
	for _, word := range words[1:] {

		if len([]rune(currentLine+" "+word)) <= 100 {
			currentLine += " " + word
		} else {
			res = append(res, currentLine)
			currentLine = word
		}
	}
	res = append(res, currentLine)
	return res
}

func (c *ChestNotes) Edit() error {

	for _, note := range c.notes {
		fmt.Printf("%v - %-20s - %s\n", note.ID, Crop(note.Title), note.CreateDate.Format("02-01-2006"))
	}

	fmt.Printf("\033[35mВведите номер заметки для редактирования:\n\033[0m")
	var input string
	fmt.Scan(&input)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	id, err := strconv.Atoi(input)
	if err != nil || id <= 0 {
		return errors.New("введите НОМЕР заметки.")
	}

	found := false
	for i := range c.notes {
		if c.notes[i].ID == id {
			found = true
			title := EnterValue("Введите новый заголовок (оставьте пустым, чтобы не менять): ")
			if title != "" {
				c.notes[i].Title = title
			}

			body := EnterValue("Введите новый текст (оставьте пустым, чтобы не менять): ")
			if body != "" {
				c.notes[i].Body = body
			}

			dateStr := EnterValue("Введите новую дату (формат: 02-01-2006): ")
			if dateStr != "" {
				if parsed, err := time.Parse("02-01-2006", dateStr); err == nil {
					c.notes[i].CreateDate = parsed
				}
			}
			break
		}
	}

	if !found {
		return errors.New("заметка не найдена")
	}

	return nil
}

func (c *ChestNotes) Stop() {
	c.Save(note_path)
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
}
