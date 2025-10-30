package main

import "cli/app"

func main() {
	chest := app.NewChestNotes()
	chest.Load()
	chest.RunApp()
}
