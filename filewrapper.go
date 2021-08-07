// GopherPEDS catch filewrapper tool (USPTO)
// Copyright (C) 2021 Vinz Frauchiger
//
// This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either version 3 of the License, or any later version.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/cavaliercoder/grab"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Document struct {
	ApplId    string `json:"applicationNumberText"`
	MRDate    string `json:"mailRoomDate"`
	DocCode   string `json:"documentCode"`
	DocDesc   string `json:"documentDescription"`
	DocCate   string `json:"documentCategory"`
	AccessLev string `json:"accessLevelCategory"`
	DocIdent  string `json:"documentIdentifier"`
	PagCount  int    `json:"pageCount"`
	PdfURL    string `json:"pdfUrl"`
}

func GetFileWrapper(applId string) error {

	// cms link
	url_down := "https://ped.uspto.gov/api/queries/cms/"
	// base url for file list
	url_list_files := "https://ped.uspto.gov/api/queries/cms/public/"
	comb_url_list := url_list_files + applId
	fmt.Println(comb_url_list)

	res, err := http.Get(comb_url_list)
	if err != nil {

		return err
	}
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {

		return err
	}

	var documents []Document
	json.Unmarshal(body, &documents)

	//iterate through every document and download it into the present folder
	for i := 0; i < len(documents); i++ {
		fmt.Println("Application Id: " + documents[i].ApplId)
		fmt.Println("Document Type: " + documents[i].DocCate)
		fmt.Println("Mail Room Date: " + documents[i].MRDate)
		t, _ := time.Parse("01-02-2006", documents[i].MRDate)
		fmt.Println(t)
		fmt.Println("Document Identifier: " + documents[i].DocIdent)
		fmt.Println("pdfUrl: " + documents[i].PdfURL)
		doc_url := url_down + documents[i].PdfURL
		if documents[i].PdfURL == "" {
			fmt.Println("THERE IS NO URL FOR THE FILE")
			continue
		}
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		resp, err := grab.Get(dirname, doc_url)
		if err != nil {
			return err
		}
		// define, refine, and rename the file name of the downloaded files, less cryptic
		filename := documents[i].ApplId + "_" + documents[i].MRDate + "_" + documents[i].DocDesc + ".pdf"
		filename = strings.ReplaceAll(filename, "/", "_")
		filename = strings.ReplaceAll(filename, ":", "_")
		filename = dirname + "/" + filename
		os.Rename(resp.Filename, filename)

		fmt.Println("Download saved to", resp.Filename, filename)
	}

	return nil
}
