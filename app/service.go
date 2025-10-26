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

func LoadNotes() []Note {
	data, err := os.ReadFile("../data/notes.json")
	if err != nil {
		fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
	}
	var notes []Note
	err = json.Unmarshal(data, &notes)
	return notes
}

func CreateNotes() error {
	scanner := bufio.NewScanner(os.Stdin)
	var notes = LoadNotes()
	var note Note
	note.ID = len(notes) + 1
	note.CreateDate = time.Now()
	fmt.Printf("\033[35mВведите название заметки:\n\033[0m")
	if scanner.Scan() {
		note.Title = scanner.Text()
	}
	fmt.Printf("\033[35mВведите текст заметки:\n\033[0m")
	if scanner.Scan() {
		note.Body = scanner.Text()
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	SaveNote(note, notes)
	return nil
}

func SaveNote(note Note, notes []Note) {
	notes = append(notes, note)
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
	}
	os.WriteFile("../data/notes.json", data, 0644)
}

func Crop(s string) string {
	newString := []rune(s)
	if len(newString) <= 20 {
		return s
	}
	return string(newString[:17]) + "..."
}
func ListNotes() error {
	notes := LoadNotes()
	for i := range notes {
		if len(notes[i].Title) > 17 {
			fmt.Printf("%v - %s - %s\n", int(notes[i].ID), Crop(notes[i].Title), notes[i].CreateDate.Format("02-01-2006"))
		} else {
			fmt.Printf("%v - %-20s - %s\n", int(notes[i].ID), Crop(notes[i].Title), notes[i].CreateDate.Format("02-01-2006"))
		}

	}
	fmt.Printf("\n\033[35m0 - Назад\n\033[0m")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil || value > len(notes) {
			err = errors.New("введите номер заметки")
			return err
		} else if value == 0 {
			return nil
		}
		fmt.Printf("%v - %-20s - %v\n", notes[value-1].ID, notes[value-1].Title, notes[value-1].CreateDate.Format("02-01-2006"))
		for _, line := range LineWrap(notes[value-1].Body) {
			fmt.Printf("%s\n", line)
		}
	}
	err := ViewingNote(notes)
	return err
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

func ViewingNote(notes []Note) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil || value > len(notes) {
			err = errors.New("введите номер заметки")
			return err
		} else if value == 0 {
			return nil
		}
		fmt.Printf("%v - %-20s - %v\n", notes[value-1].ID, notes[value-1].Title, notes[value-1].CreateDate.Format("02-01-2006"))
	}
	return nil
}

func EditNotes() error {
	notes := LoadNotes()
	var nunber int
	for i := range notes {
		if len(notes[i].Title) > 17 {
			fmt.Printf("%v - %s - %s\n", int(notes[i].ID), Crop(notes[i].Title), notes[i].CreateDate.Format("02-01-2006"))
		} else {
			fmt.Printf("%v - %-20s - %s\n", int(notes[i].ID), Crop(notes[i].Title), notes[i].CreateDate.Format("02-01-2006"))
		}
	}
	fmt.Printf("\033[35mВведите номер заметки для редактирования:\n\033[0m")
	fmt.Scan(&nunber)
	if nunber == 0 || nunber > len(notes) {
		err := errors.New("введите номер заметки для редактирования")
		return err
	}
	scanner := bufio.NewScanner(os.Stdin)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//Доделать редактирование
	for i := range notes {
		if notes[i].ID == nunber {
			fmt.Printf("\033[35mВведите заголовок:\033[0m")
			if scanner.Scan() {
				if scanner.Text() != "" {
					notes[i].Title = scanner.Text()
				}
			}
			fmt.Printf("\033[35mВведите текст заметки:\033[0m")
			if scanner.Scan() {
				if scanner.Text() != "" {
					notes[i].Body = scanner.Text()
				}

				fmt.Printf("\033[35mВведите дату создания заметки:\033[0m")
				if scanner.Scan() {
					if scanner.Text() != "" {
						notes[i].CreateDate, _ = time.Parse("02-01-2006", scanner.Text())
					}

				}

			}
			data, err := json.MarshalIndent(notes, "", "  ")
			if err != nil {
				fmt.Printf("\033[31mОшибка: %s\n\033[0m", err.Error())
			}
			os.WriteFile("../data/notes.json", data, 0644)
		}
	}
	return nil
}
func StopApp() error {
	fmt.Println("\033[33mДо свидания!\033[0m")
	os.Exit(0)
	return nil
}
