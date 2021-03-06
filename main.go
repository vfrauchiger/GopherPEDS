// GopherPEDS main file with GUI (USPTO)
// Copyright (C) 2021 Vinz Frauchiger
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// v0.13.0 List Processor!

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	canvas2 "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Versioning!
var ReleaseVersion string = "0.13.0 Process!"

func modifyText(rawText string) string {
	// Function removes Country Code and Kind Code from Patent Number
	if strings.ToUpper(rawText[:2]) == "US" {
		rawText = rawText[2:]
	}
	l := len(rawText)
	if strings.ToUpper(string(rawText[l-2])) == "A" || strings.ToUpper(string(rawText[l-2])) == "B" {
		rawText = rawText[:l-2]
	} else if strings.ToUpper(string(rawText[l-1])) == "A" {
		rawText = rawText[:l-1]
	}
	return rawText
}

func removeChars(rawText string) string {
	// Function removes unwanted chars from a string
	rawText = strings.ReplaceAll(rawText, "/", "")
	rawText = strings.ReplaceAll(rawText, "-", "")
	rawText = strings.ReplaceAll(rawText, " ", "")
	return rawText
}

func treatEarlAppNumb(pubnum string) string {
	// this function checks for the correct length and
	// the correct kind code
	l := len(pubnum)
	if l == 14 {
		pubnum = pubnum[:6] + "0" + pubnum[6:]
	}
	if pubnum[l-2:] == "AA" {
		pubnum = pubnum[:l-2] + "A1"
	} else if pubnum[l-2:] == "AB" {
		pubnum = pubnum[:l-2] + "A2"
	}
	return pubnum
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

func chooseFile(w fyne.Window, fname *widget.Label) {
	dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
		if file == nil {
			return
		}
		fileP := file.URI().Path()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fileP)
		fname.SetText(fileP)
		file.Close()
	}, w)
}

//Variables
var save_dir string = "$HOME"
var theApplId2 string = ""

//
// MAIN
//

func main() {

	a := app.New()
	w := a.NewWindow("USPTO PEDS Tool Go!")
	hello := widget.NewLabel("Hello Dude!")
	hello.TextStyle = fyne.TextStyle{Bold: true}
	progress := widget.NewProgressBar()
	progress.SetValue(0)

	labApplId := widget.NewLabel("Application ID")
	inpApplId := widget.NewEntry()
	inpApplId.SetPlaceHolder("12321123")

	labPatentNum := widget.NewLabel("Patent Number")
	inpPatentNum := widget.NewEntry()
	inpPatentNum.SetPlaceHolder("7123456 or 11321123")

	labEarlPubNum := widget.NewLabel("Early Publication Number")
	inpEarlPubNum := widget.NewEntry()
	inpEarlPubNum.SetPlaceHolder("Us20080123456A1")

	// Check "Turbo Mode"

	checkTurbo := widget.NewCheck("Turbo", func(value bool) {
		fmt.Println(value)
	})

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
		"?? Vinz Frauchiger, 2021",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)
	labExclPatents := widget.NewLabelWithStyle(
		" Proudly for Ypsomed Patents!",
		fyne.TextAlignCenter,
		fyne.TextStyle{Italic: true},
	)
	// Button Application ID
	butTermApplID := widget.NewButton("Get Term Ext.", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		termdays, discl, _, err := GetTermDisc("applId", modifiedText)
		if err != nil {
			log.Fatal(err)
		}
		if discl == "" {
			discl = "No terminal disclaimer found!"
		}
		if termdays == "" {
			termdays = "0"
		}
		hello.SetText(termdays + " days and " + discl)
	})
	butWrapApplId := widget.NewButton("Get FileWrapper", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		fmt.Println(checkTurbo.Checked)
		_, _, theApplId, err := GetTermDisc("applId", modifiedText)
		if err != nil {
			hello.SetText("wrong number format!")
		} else if theApplId == "number!" {
			hello.SetText("wrong number!(kind code?)")
		} else {
			theApplId2 = theApplId
			hello.SetText("Getting FileWrapper for " + theApplId2)
			err = GetFileWrapperMulti(theApplId2, save_dir, progress, checkTurbo.Checked)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	butApplLatClaims := widget.NewButton("Get Latest Claims", func() {
		modifiedText := modifyText(inpApplId.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		discNumber(modifiedText, "applId", save_dir, progress)
	})

	// Buttons Patents
	butPatNumTerm := widget.NewButton("Get Term Ext.", func() {
		var termMonths float64
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
		if termdays == "" {
			termdays = "0"
		}
		// convert term extension from days to months.
		termMonths, err = strconv.ParseFloat(termdays, 64)
		termMonths = termMonths / 365.25 * 12.0
		if err != nil {
			log.Fatal(err)
		}
		hello.SetText(termdays + " days (" + fmt.Sprintf("%.1f", termMonths) + "months) and " + discl + " / ApplID " + theApplId)
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
			err = GetFileWrapperMulti(theApplId2, save_dir, progress, checkTurbo.Checked)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	butPatLatClaims := widget.NewButton("Get Latest Claims", func() {
		modifiedText := modifyText(inpPatentNum.Text)
		modifiedText = removeChars(modifiedText)
		hello.SetText(modifiedText)
		discNumber(modifiedText, "patentNumber", save_dir, progress)
	})

	// Buttons Early Publication
	butEarlPubNumTerm := widget.NewButton("Get Term Ext.", func() {
		modifiedText := removeChars(inpEarlPubNum.Text)
		modifiedText = strings.ToUpper(modifiedText)
		modifiedText = treatEarlAppNumb(modifiedText)
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
		modifiedText = treatEarlAppNumb(modifiedText)
		hello.SetText(modifiedText)
		_, _, theApplId, err := GetTermDisc("appEarlyPubNumber", modifiedText)
		if err != nil {
			hello.SetText("wrong number format!")
		} else if theApplId == "number!" {
			hello.SetText("wrong number!(kind code?)")
		} else {
			theApplId2 = theApplId
			hello.SetText("Getting FileWrapper for " + theApplId2)

			err = GetFileWrapperMulti(theApplId2, save_dir, progress, checkTurbo.Checked)
			if err != nil {
				hello.SetText(err.Error())
			}
		}
	})

	butEarlPubLatClaims := widget.NewButton("Get Latest Claims", func() {
		modifiedText := removeChars(inpEarlPubNum.Text)
		modifiedText = strings.ToUpper(modifiedText)
		modifiedText = treatEarlAppNumb(modifiedText)
		hello.SetText(modifiedText)
		discNumber(modifiedText, "appEarlyPubNumber", save_dir, progress)
	})

	fmt.Println(checkTurbo.Checked)
	//button for directory to save to
	labSavDir := widget.NewLabel("$HOME")
	butSaveDir := widget.NewButton("Get Save Directory!", func() {
		chooseDirectory(w, labSavDir) // Text of hello updated by return value
	})

	//
	// List processor
	labListProc := widget.NewLabel("no file chosen")
	butListProc := widget.NewButton("Get File", func() {
		chooseFile(w, labListProc)
	})
	butGoList := widget.NewButton("Go List Proc.", func() {
		publicationList := LoadExcel(labListProc.Text)
		fmt.Println(publicationList)
		text := labListProc.Text
		response := ManageList(w, labListProc, publicationList, text)
		fmt.Println(response)
	})

	// CONTENT
	content := container.NewVBox(
		container.NewHBox(
			image,
			container.NewVBox(
				labTitle,
				labCopyRight,
				labExclPatents,
				imageYps,
				widget.NewLabelWithStyle(
					"Version: "+ReleaseVersion,
					fyne.TextAlignCenter,
					fyne.TextStyle{Italic: true},
				),
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
			butPatLatClaims,
		),
		labEarlPubNum,
		inpEarlPubNum,
		container.NewHBox(
			butEarlPubNumTerm,
			butEarlPubNumWrap,
			butEarlPubLatClaims,
		),
		widget.NewSeparator(),
		widget.NewLabel("Please not that the list processor takes (early & patent) publication numbers only!"),
		container.NewHBox(
			widget.NewLabel("File List Proc."),
			labListProc,
			widget.NewSeparator(),
			butListProc,
			butGoList,
		),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewLabel("Directory to which the files are saved: "),
			widget.NewSeparator(),
			layout.NewSpacer(),
			widget.NewSeparator(),
			checkTurbo,
		),
		labSavDir,
		butSaveDir,
		progress,
		widget.NewSeparator(),
		widget.NewButton("Quit", w.Close))
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
	w.SetOnClosed(func() {
		os.Exit(1)
	})
}
