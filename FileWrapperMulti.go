package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/cavaliercoder/grab"
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
	if turbo == true {
		speed = 0 //very fast
	} else {
		speed = 3 //rather slow
	}

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
	// 0 for the batch size means everythin in parallel!
	respch, err := grab.GetBatch(speed, savePath, urls...) //respch is a channel!
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// start a ticker to update progress every 200ms
	t := time.NewTicker(20 * time.Millisecond)

	// monitor downloads
	completed := 0
	inProgress := 0

	responses := make([]*grab.Response, 0)
	for completed < len(urls) {
		select {
		case resp := <-respch:
			// a new response has been received and has started downloading
			// (nil is received once, when the channel is closed by grab)
			if resp != nil {
				responses = append(responses, resp)
			}

		case <-t.C:
			// clear lines
			if inProgress > 0 {
				fmt.Printf("%d\n", inProgress)

			}

			// update completed downloads
			for i, resp := range responses {
				if resp != nil && resp.IsComplete() {
					// print final result
					if resp.Err() != nil {
						fmt.Fprintf(os.Stderr, "Error downloading %s\n", resp.Request.URL())
					}

					// mark completed
					responses[i] = nil
					completed++
					// Set Value in Progress bar!
					proBar.SetValue(float64(completed) / float64(len(urls)))
				}
			}

			// update downloads in progress
			inProgress = 0
			for _, resp := range responses {
				if resp != nil {
					inProgress++
					//fmt.Printf("Downloading %s ", resp.Filename)
				}
			}

		}
	}
	proBar.SetValue(1.0)
	t.Stop()
	fmt.Println(completed)
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
