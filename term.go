package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Query struct {
	SearchText string `json:"searchText"`
	Fl         string `json:"fl"`
	Mm         string `json:"mm"`
	Qf         string `json:"qf"`
	Facet      string `json:"facet"`
	Sort       string `json:"sort"`
	Start      string `json:"start"`
}

type AutoGenerated struct {
	QueryResults struct {
		IndexLastUpdatedDate string `json:"indexLastUpdatedDate"`
		SearchResponse       struct {
			ResponseHeader struct {
				ZkConnected bool `json:"zkConnected"`
				Status      int  `json:"status"`
				QTime       int  `json:"QTime"`
			} `json:"responseHeader"`
			Response struct {
				NumFound      int  `json:"numFound"`
				Start         int  `json:"start"`
				NumFoundExact bool `json:"numFoundExact"`
				Docs          []struct {
					CorrAddrCountryName string `json:"corrAddrCountryName"`
					ApplID              string `json:"applId"`
					TotalPtoDays        string `json:"totalPtoDays"`
					Assignments         []struct {
						ReelNumber           string `json:"reelNumber"`
						FrameNumber          string `json:"frameNumber"`
						AddressNameText      string `json:"addressNameText"`
						AddressLineOneText   string `json:"addressLineOneText"`
						AddressLineTwoText   string `json:"addressLineTwoText"`
						AddressLineThreeText string `json:"addressLineThreeText"`
						AddressLineFourText  string `json:"addressLineFourText"`
						MailDate             string `json:"mailDate"`
						ReceivedDate         string `json:"receivedDate"`
						RecordedDate         string `json:"recordedDate"`
						PagesCount           string `json:"pagesCount"`
						ConveryanceName      string `json:"converyanceName"`
						SequenceNumber       string `json:"sequenceNumber"`
						Assignors            []struct {
							AssignorName string `json:"assignorName"`
							ExecDate     string `json:"execDate"`
						} `json:"assignors"`
						Assignee []struct {
							AssigneeName      string `json:"assigneeName"`
							StreetLineOneText string `json:"streetLineOneText"`
							StreetLineTwoText string `json:"streetLineTwoText"`
							CityName          string `json:"cityName"`
							CountryCode       string `json:"countryCode"`
							PostalCode        string `json:"postalCode"`
						} `json:"assignee"`
					} `json:"assignments"`
					AppFilingDate          time.Time `json:"appFilingDate"`
					AppExamName            string    `json:"appExamName"`
					AppExamNameFacet       string    `json:"appExamNameFacet"`
					PublicInd              string    `json:"publicInd"`
					AppInd                 string    `json:"APP_IND"`
					InventorName           string    `json:"inventorName"`
					InventorNameFacet      string    `json:"inventorNameFacet"`
					AppEarlyPubDate        time.Time `json:"appEarlyPubDate"`
					CorrAddrGeoRegionCode  string    `json:"corrAddrGeoRegionCode"`
					AppLocation            string    `json:"appLocation"`
					AppEarlyPubNumber      string    `json:"appEarlyPubNumber"`
					ID                     string    `json:"id"`
					AppGrpArtNumber        string    `json:"appGrpArtNumber"`
					AppGrpArtNumberFacet   string    `json:"appGrpArtNumberFacet"`
					ApplIDStr              string    `json:"applIdStr"`
					ApplIDTxt              string    `json:"appl_id_txt"`
					InventorsFullName      []string  `json:"inventorsFullName"`
					InventorsFullNameFacet string    `json:"inventorsFullNameFacet"`
					AppSubCls              string    `json:"appSubCls"`
					PatentNumber           string    `json:"patentNumber"`
					LastModTs              time.Time `json:"LAST_MOD_TS"`
					Transactions           []struct {
						RecordDate  string `json:"recordDate"`
						Code        string `json:"code"`
						Description string `json:"description"`
					} `json:"transactions"`
					LastInsertTime    time.Time `json:"LAST_INSERT_TIME"`
					AppCls            string    `json:"appCls"`
					AppStatus         string    `json:"appStatus"`
					AppStatusFacet    string    `json:"appStatusFacet"`
					AppStatusTxt      string    `json:"appStatus_txt"`
					PtaPteInd         string    `json:"ptaPteInd"`
					PtaPteTranHistory []struct {
						Number              string `json:"number"`
						PtaOrPteDate        string `json:"ptaOrPteDate"`
						ContentsDescription string `json:"contentsDescription"`
						PtoDays             string `json:"ptoDays"`
						ApplDays            string `json:"applDays"`
						Start               string `json:"start"`
					} `json:"ptaPteTranHistory"`
					PatentTitle       string    `json:"patentTitle"`
					ApplDelay         string    `json:"applDelay"`
					CDelay            string    `json:"cDelay"`
					AppStatusDate     time.Time `json:"appStatusDate"`
					AppAttrDockNumber string    `json:"appAttrDockNumber"`
					Inventors         []struct {
						NameLineOne string `json:"nameLineOne"`
						NameLineTwo string `json:"nameLineTwo"`
						Suffix      string `json:"suffix"`
						StreetOne   string `json:"streetOne"`
						StreetTwo   string `json:"streetTwo"`
						City        string `json:"city"`
						GeoCode     string `json:"geoCode"`
						Country     string `json:"country"`
						RankNo      string `json:"rankNo"`
					} `json:"inventors"`
					InventorsFacet        string    `json:"inventorsFacet"`
					PtoAdjustments        string    `json:"ptoAdjustments"`
					CorrAddrStreetLineOne string    `json:"corrAddrStreetLineOne"`
					FirstInventorFile     string    `json:"firstInventorFile"`
					OverlapDelay          string    `json:"overlapDelay"`
					ADelay                string    `json:"aDelay"`
					AppType               string    `json:"appType"`
					AppTypeFacet          string    `json:"appTypeFacet"`
					CorrAddrPostalCode    string    `json:"corrAddrPostalCode"`
					AppCustNumber         string    `json:"appCustNumber"`
					AppClsSubCls          string    `json:"appClsSubCls"`
					AppClsSubClsFacet     string    `json:"appClsSubClsFacet"`
					PatentIssueDate       time.Time `json:"patentIssueDate"`
					BDelay                string    `json:"bDelay"`
					CorrAddrNameLineOne   string    `json:"corrAddrNameLineOne"`
					AttrnyAddr            []struct {
						ApplicationID  interface{} `json:"applicationId"`
						RegistrationNo string      `json:"registrationNo"`
						FullName       string      `json:"fullName"`
						PhoneNum       string      `json:"phoneNum"`
						RegStatus      string      `json:"regStatus"`
					} `json:"attrnyAddr"`
					ParentContinuity []struct {
						ClaimApplicationNumberText   string `json:"claimApplicationNumberText"`
						ApplicationNumberText        string `json:"applicationNumberText"`
						FilingDate                   string `json:"filingDate"`
						AiaIndicator                 string `json:"aiaIndicator"`
						PatentNumberText             string `json:"patentNumberText"`
						ApplicationStatus            string `json:"applicationStatus"`
						ApplicationStatusDescription string `json:"applicationStatusDescription"`
					} `json:"parentContinuity"`
					CorrAddrCustNo         string    `json:"corrAddrCustNo"`
					CorrAddrCountryCd      string    `json:"corrAddrCountryCd"`
					PtoDelay               string    `json:"ptoDelay"`
					CorrAddrCity           string    `json:"corrAddrCity"`
					AppEntityStatus        string    `json:"appEntityStatus"`
					AppConfrNumber         string    `json:"appConfrNumber"`
					LastUpdatedTimestamp   time.Time `json:"lastUpdatedTimestamp"`
					AppAttrDockNumberFacet string    `json:"appAttrDockNumberFacet"`
					AppEntityStatusFacet   string    `json:"appEntityStatusFacet"`
					AppCustNumberFacet     string    `json:"appCustNumberFacet"`
					Version                int64     `json:"_version_"`
					FirstInventorFileFacet string    `json:"firstInventorFileFacet"`
					AppLocationFacet       string    `json:"appLocationFacet"`
					AppEarlyPubNumberFacet string    `json:"appEarlyPubNumberFacet"`
					PatentNumberFacet      string    `json:"patentNumberFacet"`
				} `json:"docs"`
			} `json:"response"`
		} `json:"searchResponse"`
		QueryID string `json:"queryId"`
	} `json:"queryResults"`
	JobStatus          interface{} `json:"jobStatus"`
	QueryID            string      `json:"queryId"`
	Page               int         `json:"page"`
	Count              int         `json:"count"`
	CreateQueryRequest struct {
		SearchText       string      `json:"searchText"`
		Facet            string      `json:"facet"`
		FacetField       interface{} `json:"facetField"`
		FacetLimit       interface{} `json:"facetLimit"`
		FacetMissing     interface{} `json:"facetMissing"`
		FacetDate        interface{} `json:"facetDate"`
		FacetDateGap     interface{} `json:"facetDateGap"`
		FacetDateStart   interface{} `json:"facetDateStart"`
		FacetDateEnd     interface{} `json:"facetDateEnd"`
		FacetDateOther   interface{} `json:"facetDateOther"`
		FacetDateInclude interface{} `json:"facetDateInclude"`
		Mm               string      `json:"mm"`
		Sort             string      `json:"sort"`
		Qf               string      `json:"qf"`
		Wt               interface{} `json:"wt"`
		Df               interface{} `json:"df"`
		Fl               string      `json:"fl"`
		Start            string      `json:"start"`
		Fq               interface{} `json:"fq"`
		Rows             interface{} `json:"rows"`
		Parameters       struct {
			Mm    string   `json:"mm"`
			Qf    string   `json:"qf"`
			Fl    []string `json:"fl"`
			Start string   `json:"start"`
			Sort  string   `json:"sort"`
			Rows  int      `json:"rows"`
			Facet string   `json:"facet"`
		} `json:"parameters"`
	} `json:"createQueryRequest"`
	Links []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

func GetTermDisc(queryKind string, queryNumber string) (string, string, string, error) {
	searchText := queryKind + ":(" + queryNumber + ")"
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
		return "", "", "", err

	}
	//fmt.Println(json_data)

	resp, err := http.Post("https://ped.uspto.gov/api/queries", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", "", "", err
	}

	//fmt.Println(resp.Body)
	var res AutoGenerated

	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", "", "", err
	}

	// Attention: themap is a map with submaps an different types to dig deeper: themap.(map[string]interface{})
	//fmt.Println(res)
	if len(res.QueryResults.SearchResponse.Response.Docs) == 0 {
		return "enter", "useful", "number!", nil
	}
	themap := res.QueryResults.SearchResponse.Response.Docs[0]

	// trans is a slice of maps! Iteraste through it with a counter
	trans := themap.Transactions
	theapplId := themap.ApplID
	termdays := themap.TotalPtoDays
	var discl string = ""
	for i := 0; i < len(trans); i++ {
		if trans[i].Code == "DIST" {
			discl = trans[i].Description + trans[i].RecordDate
		}
	}
	return termdays, discl, theapplId, nil
}
