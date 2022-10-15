//Menu for the main app

package main

import (
	"fmt"

	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//New App
	a := app.New()

	//New Windows
	w := a.NewWindow("Main Software")

	//Resize Windows
	w.Resize(fyne.NewSize(1000, 600))

	// Add button (contentb) /* i need to add this to fit into the left hand side of the app */
	//New Button
	btn1 := widget.NewButton("TOP BUTTON", func() {
	})
	btn2 := widget.NewButton("SIDE BUTTON", func() {
	})

	label1 := widget.NewLabel("Hello Fyne!")
	// NewHBox
	box1 := container.NewVBox( // Horizontal
		btn1,
		label1,
	)
	box2 := container.NewHBox( // Vertical
		btn2,
		label1,
	)

	//Menu Items - File | Edit | View | Help
	//File

	menuItem1 := fyne.NewMenuItem("Open", func() {
		fmt.Println("Main Menu---> Open")
	})

	//Edit
	menuItem2 := fyne.NewMenuItem("New", func() {
		fmt.Println("Main Menu---> Open")
	})
	//View
	menuItem3 := fyne.NewMenuItem("View", func() {
		fmt.Println("Main Menu---> View")
	})
	//Help
	menuItem4 := fyne.NewMenuItem("Devoloper", func() {
		openurl, _ := url.Parse("http://nafishaider.com.au")
		_ = a.OpenURL(openurl)
	})

	//New Menu
	newMenu := fyne.NewMenu("File", menuItem1)
	newMenu2 := fyne.NewMenu("Edit", menuItem1, menuItem2)
	newMenu3 := fyne.NewMenu("View", menuItem1, menuItem3)
	newMenu4 := fyne.NewMenu("Help", menuItem4)

	// Creating new main menu

	menu := fyne.NewMainMenu(newMenu, newMenu2, newMenu3, newMenu4)
	box := container.NewVBox(box1, box2)

	//Setting New Menu
	w.SetMainMenu(menu)
	w.SetContent(box)
	w.ShowAndRun()

}
