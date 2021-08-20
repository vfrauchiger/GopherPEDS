package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"fyne.io/fyne/v2/widget"
	"github.com/cavaliercoder/grab"
)

// Document struct for json document list
// from USPTO Peds

func get_ApplId(publNo, queryKind string) string {
	searchText := queryKind + ":(" + publNo + ")"
	query := Query{
		SearchText: searchText,
		Fl:         "*",
		Mm:         "20",
		Qf: `appEarlyPubNumber applId appLocation appType appStatus_txt appConfrNumber appCustNumber appGrpArtNumber 
		appCls appSubCls appEntityStatus_txt patentNumber patentTitle primaryInventor firstNamedApplicant appExamName 
		appExamPrefrdName appAttrDockNumber appPCTNumber appIntlPubNumber wipoEarlyPubNumber pctAppType firstInventorFile 
		appClsSubCls rankAndInventorsList`,
		Facet: "false",
		Sort:  "applId asc",
		Start: "0",
	}

	json_data, err := json.Marshal(query)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://ped.uspto.gov/api/queries", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(err)
	}

	var res AutoGenerated

	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	applId := res.QueryResults.SearchResponse.Response.Docs[0].ApplID
	fmt.Println(applId)
	return applId
}

func getLatestClaims(applIdExt, savePath string, proBar *widget.ProgressBar) {
	applId := applIdExt
	proBar.SetValue(0.2)
	url_list_files := "https://ped.uspto.gov/api/queries/cms/public/"
	comb_url_list := url_list_files + applId

	res, err := http.Get(comb_url_list)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	proBar.SetValue(0.4)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var documents []Document
	err = json.Unmarshal(
		body,
		&documents,
	)
	if err != nil {
		log.Fatal(err)
	}

	var maxDate time.Time
	var maxI int
	for i := 0; i < len(documents); i++ {
		if documents[i].DocDesc == "Claims" {
			fmt.Println(documents[i])
			if documents[i].MRDate.After(maxDate) {
				fmt.Println("Plopp")
				maxDate = documents[i].MRDate
				maxI = i
			} else if documents[i].MRDate.Equal(maxDate) {
				fmt.Println("Second of claims at identical date found!!")
			}
		}
	}
	fmt.Println(maxDate)
	fmt.Println(maxI)
	fmt.Println(documents[maxI])
	proBar.SetValue(0.6)
	claims_url := "https://ped.uspto.gov/api/queries/cms/" + documents[maxI].PdfURL
	if savePath == "$HOME" {
		savePath, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
	}

	resp, err := grab.Get(savePath, claims_url)
	if err != nil {
		log.Fatal(err)
	}
	proBar.SetValue(0.8)
	filename := savePath + "/" + documents[maxI].ApplId + "_" + documents[maxI].MRDate.Format("20060101") + "_latest_claims.pdf"
	fmt.Println(filename)
	os.Rename(resp.Filename, filename)
	proBar.SetValue(1.0)
}

func discNumber(docNo, queryKind, savePath string, proBar *widget.ProgressBar) {
	proBar.SetValue(0)
	if queryKind == "applId" {
		getLatestClaims(docNo, savePath, proBar)
	} else {
		applId := get_ApplId(docNo, queryKind)
		getLatestClaims(applId, savePath, proBar)
	}
}
