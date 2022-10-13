package main

import (

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

)

func main() {
	//New App
	a := app.New()
	
	//New Windows
	w := a.NewWindow("Main Software")
	
	//Resize Windows
	w.Resize(fyne.NewSize(640, 480))

	//Menu Items
	menuItem1 := fyne.NewMenuItem("Open", func() {
		fmt.Println("Main Menu---> Open")
	})

	//New Menu
	newMenu := fyne.NewMenu("File", menuItem1)

	// Creating new main menu

	menu := fyne.NewMainMenu(newMenu)

	//Setting New Menu
	w.SetMainMenu(menu)
	w.ShowAndRun()


}