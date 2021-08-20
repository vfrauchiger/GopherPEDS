// GopherPEDS main file with GUI (USPTO)
// Copyright (C) 2021 Vinz Frauchiger
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either version 3 of the License, or any later version.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func modifyText(rawText string) string {
	if strings.ToUpper(rawText[:2]) == "US" {
		rawText = rawText[2:]
	}
	l := len(rawText)
	if strings.ToUpper(string(rawText[l-2])) == "A" || strings.ToUpper(string(rawText[l-2])) == "B" {
		rawText = rawText[:l-2]
	}
	return rawText
}

func removeChars(rawText string) string {
	rawText = strings.ReplaceAll(rawText, "/", "")
	rawText = strings.ReplaceAll(rawText, "-", "")
	rawText = strings.ReplaceAll(rawText, " ", "")
	return rawText
}

func chooseDirectory(w fyne.Window, h *widget.Label) {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		save_dir = "$HOME"
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if dir != nil {
			fmt.Println(dir.Path())
			save_dir = dir.Path() // here value of save_dir shall be updated!
		}
		fmt.Println(save_dir)
		h.SetText(save_dir)
	}, w)
}

var save_dir string = "$HOME"
var theApplId2 string = ""

func main() {

	a := app.New()
	w := a.NewWindow("USPTO PEDS Tool Go!")
	hello := widget.NewLabel("Hello Dude!")
	hello.TextStyle = fyne.TextStyle{Bold: true}

	labApplId := widget.NewLabel("Application ID")
	inpApplId := widget.NewEntry()
	inpApplId.SetPlaceHolder("12321123")

	labPatentNum := widget.NewLabel("Patent Number")
	inpPatentNum := widget.NewEntry()
	inpPatentNum.SetPlaceHolder("7123456 or 11321123")

	labEarlPubNum := widget.NewLabel("Early Publication Number")
	inpEarlPubNum := widget.NewEntry()
	inpEarlPubNum.SetPlaceHolder("Us20080123456A1")

	// images
	//image := canvas2.NewImageFromFile("gopherli.png")
	image := canvas2.NewImageFromResource(resourceGopherliPng)
	image.FillMode = canvas2.ImageFillOriginal
	//imageYps := canvas2.NewImageFromFile("Ypsomed_big.png")
	imageYps := canvas2.NewImageFromResource(resourceYpsomedbigPng)
	imageYps.FillMode = canvas2.ImageFillOriginal

	labTitle := widget.NewLabelWithStyle(
		"GopherPEDS",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	labCopyRight := widget.NewLabelWithStyle(
		"© Vinz Frauchiger, 2021",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	labExclPatents := widget.NewLabelWithStyle(
		" Proudly for Ypsomed Patents!",
		fyne.TextAlignCenter,
		fyne.TextStyle{Italic: true},
	)
	// Button Application ID
	butTermApplID := widget.NewButton("Go Appl Id", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		termdays, discl, _, err := GetTermDisc("applId", modifiedText)
		if err != nil {
			log.Fatal(err)
		}
		if discl == "" {
			discl = "No terminal disclaimer found!"
		}
		hello.SetText(termdays + " days and " + discl)
	})
	butWrapApplId := widget.NewButton("Get FileWrapper", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		_, _, theApplId, err := GetTermDisc("applId", modifiedText)
		if err != nil {
			hello.SetText("wrong number format!")
		} else if theApplId == "number!" {
			hello.SetText("wrong number!(kind code?)")
		} else {
			theApplId2 = theApplId
			hello.SetText("Getting FileWrapper for " + theApplId2)
			err = GetFileWrapper(theApplId2, save_dir)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	butApplLatClaims := widget.NewButton("Get Latest Claims", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		discNumber(modifiedText, "applId", save_dir)
	})

	// Buttons Patents
	butPatNumTerm := widget.NewButton("Go Pat Num", func() {
		modifiedText := modifyText(inpPatentNum.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		termdays, discl, theApplId, err := GetTermDisc("patentNumber", modifiedText)
		if err != nil {
			log.Fatal(err)
		}
		theApplId2 = theApplId
		if discl == "" {
			discl = "No terminal disclaimer found!"
		}
		hello.SetText(termdays + " days and " + discl + " /ApplID " + theApplId)
	})
	butPatNumWrap := widget.NewButton("Get FileWrapper for Patent", func() {
		modifiedText := modifyText(inpPatentNum.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		_, _, theApplId, err := GetTermDisc("patentNumber", modifiedText)
		if err != nil {
			hello.SetText("wrong number format!")
		} else if theApplId == "number!" {
			hello.SetText("wrong number!(kind code?)")
		} else {
			theApplId2 = theApplId
			hello.SetText("Getting FileWrapper for " + theApplId2)
			err = GetFileWrapper(theApplId2, save_dir)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	// Buttons Early Publication
	butEarlPubNumTerm := widget.NewButton("Go Publ Num", func() {
		modifiedText := removeChars(inpEarlPubNum.Text)
		modifiedText = strings.ToUpper(modifiedText)
		hello.SetText(modifiedText)
		termdays, discl, theApplId, err := GetTermDisc("appEarlyPubNumber", modifiedText)
		if err != nil {
			fmt.Println(termdays, discl, theApplId, err)
		}
		if discl == "" {
			discl = "No terminal disclaimer found!"
		}
		if termdays == "" {
			termdays = "not granted yet!"
		}
		hello.SetText(termdays + " days and " + discl + " / " + theApplId)
	})
	butEarlPubNumWrap := widget.NewButton("Get FileWrapper Publ", func() {
		modifiedText := removeChars(inpEarlPubNum.Text)
		modifiedText = strings.ToUpper(modifiedText)
		hello.SetText(modifiedText)
		_, _, theApplId, err := GetTermDisc("appEarlyPubNumber", modifiedText)
		if err != nil {
			hello.SetText("wrong number format!")
		} else if theApplId == "number!" {
			hello.SetText("wrong number!(kind code?)")
		} else {
			theApplId2 = theApplId
			hello.SetText("Getting FileWrapper for " + theApplId2)
			err = GetFileWrapper(theApplId2, save_dir)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	//button for directory to save to
	labSavDir := widget.NewLabel("$HOME")
	butSaveDir := widget.NewButton("Get Save Directory!", func() {
		chooseDirectory(w, labSavDir) // Text of hello updated by return value
	})

	// CONTENT
	content := container.NewVBox(
		container.NewHBox(
			image,
			container.NewVBox(
				labTitle,
				imageYps,
				labCopyRight,
				labExclPatents,
			),
		),
		widget.NewSeparator(),
		widget.NewLabel("Find Output below:"),
		widget.NewSeparator(),
		hello,
		widget.NewSeparator(),
		labApplId,
		inpApplId,
		container.NewHBox(
			butTermApplID,
			butWrapApplId,
			butApplLatClaims,
		),
		labPatentNum,
		inpPatentNum,
		container.NewHBox(
			butPatNumTerm,
			butPatNumWrap,
		),
		labEarlPubNum,
		inpEarlPubNum,
		container.NewHBox(
			butEarlPubNumTerm,
			butEarlPubNumWrap,
		),
		widget.NewSeparator(),
		widget.NewLabel("Directory to which the files are saved: "),
		labSavDir,
		butSaveDir,
		widget.NewSeparator(),
		widget.NewButton("Quit", w.Close))
	w.SetContent(content)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
	w.SetOnClosed(func() {
		os.Exit(1)
	})
}
