package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ManageList(w fyne.Window, success *widget.Label, publnolist []string, file string) string {
	//filename := "noname"
	pbmap := make(map[string]string)
	for _, pblno := range publnolist {
		fmt.Println("Publikation: " + pblno)

		pblno = removeChars(pblno)
		l := len(pblno)
		fmt.Printf("korrigierte Länge: %d\n", l)
		//early publication number
		if l >= 14 {
			if l == 14 {
				pblno = pblno[:6] + "0" + pblno[6:]
				l += 1 // length has to be increased due to addition of 0
			}
			fmt.Println(pblno[l-2:])
			if pblno[l-2:] == "AA" {
				pblno = pblno[:l-2] + "A1"
			} else if pblno[l-2:] == "AB" {
				pblno = pblno[:l-2] + "A2"
			}
			pbmap[pblno] = "appEarlyPubNumber"
		} else if l < 14 { //patentNumbers
			pblno = modifyText(pblno)
			pbmap[pblno] = "patentNumber"
		}
	}
	fmt.Println(pbmap)
	var noTermDiscAppl [][]string
	noTermDiscAppl = append(noTermDiscAppl, []string{"Publ. No", "Term Ext. [days]", "Disclaimer/Date", "Appl ID"})
	for key, value := range pbmap {
		s := make([]string, 4)
		fmt.Println("publno: " + key)
		fmt.Println("Kind: " + value)
		termdays, discl, theapplId, err := GetTermDisc(value, key)
		if err != nil {
			fmt.Println(err)
		} else {
			s[0] = key
			if termdays == "0" || termdays == "" {
				s[1] = "no extension"
			} else {
				s[1] = termdays
			}
			if discl == "" {
				s[2] = "no disclaimer"
			} else {
				s[2] = discl
			}
			s[3] = theapplId
			noTermDiscAppl = append(noTermDiscAppl, s)
			fmt.Println(s)
		}
	}
	//fmt.Println(noTermDiscAppl)
	response := toExcel(noTermDiscAppl, file)
	success.SetText(response)
	return response
}
