package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
)

type Document struct {
	ApplId    string    `json:"applicationNumberText"`
	MRDate    time.Time `json:"mailRoomDate"`
	DocCode   string    `json:"documentCode"`
	DocDesc   string    `json:"documentDescription"`
	DocCate   string    `json:"documentCategory"`
	AccessLev string    `json:"accessLevelCategory"`
	DocIdent  string    `json:"documentIdentifier"`
	PagCount  int       `json:"pageCount"`
	PdfURL    string    `json:"pdfUrl"`
}

func GetFileWrapperMulti(applId string, save_dir string, proBar *widget.ProgressBar, turbo bool) error {
	//set the download speed
	var speed int
	if turbo {
		speed = 0 //very fast
	} else {
		speed = 3 //rather slow
	}
	fmt.Println(speed)
	url_list_files := "https://ped.uspto.gov/api/queries/cms/public/"
	comb_url_list := url_list_files + applId

	// Get Document list from USPTO
	res, err := http.Get(comb_url_list)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//Read the JSON Response (body) into the program
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	//documents is a slice of Document structs
	var documents []Document
	err = json.Unmarshal(
		body,
		&documents,
	)
	if err != nil {
		return err
	}
	//... end of reading the JSON into the program

	//Create a list of File URLS for Download
	var urls []string

	for i := 0; i < len(documents); i++ {
		fmt.Println(documents[i].DocIdent)
		if documents[i].PdfURL == "" {
			continue
		}
		claims_url := "https://ped.uspto.gov/api/queries/cms/" + documents[i].PdfURL
		urls = append(urls, claims_url)

	}
	fmt.Println(len(urls))
	// define the location to save the files to
	var savePath string
	if save_dir == "$HOME" {
		savePath, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		savePath = save_dir
	}

	/*
		The following routine starts the batch download
	*/

	for i, url := range urls {
		var reader io.Reader

		client := &http.Client{}

		fmt.Println(url)

		req, err := http.NewRequest(
			"GET",
			url,
			reader,
		)
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		saveFilePath := savePath + "/" + documents[i].DocIdent

		file, err := os.Create(saveFilePath)
		if err != nil {
			fmt.Println(err)
		}

		writtenBytes, err := io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(writtenBytes)
		file.Close()

	}

	fmt.Printf("%d files successfully downloaded.\n", len(urls))

	// renaming of all downloaded files
	for i := 0; i < len(documents); i++ {
		newFilename := documents[i].ApplId + "_" + documents[i].MRDate.Format("20060101") +
			"_" + documents[i].DocIdent + "_" + documents[i].DocDesc + ".pdf"
		newFilename = strings.ReplaceAll(newFilename, " ", "_")
		newFilename = strings.ReplaceAll(newFilename, "/", "_")
		newFilename = strings.ReplaceAll(newFilename, ":", "_")
		newFilename = savePath + "/" + newFilename
		oldFilename := savePath + "/" + documents[i].DocIdent
		os.Rename(oldFilename, newFilename)
	}
	return nil
}
