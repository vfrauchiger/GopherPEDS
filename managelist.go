package main

import "fmt"

func ManageList(publnolist []string) string {
	filename := "noname"
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

	for key, value := range pbmap {
		fmt.Println("publno: " + key)
		fmt.Println("Kind: " + value)
	}
	return filename
}
