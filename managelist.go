package main

import "fmt"

func ManageList(publnolist []string) string {
	filename := "noname"

	for _, pblno := range publnolist {
		fmt.Println("Publikation" + pblno)
	}

	return filename
}
