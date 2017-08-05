package sage

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"fmt"

	sr "github.com/novikk/splitbot/sage/sageResponses"
)

//SageClient is sageClient
type SageClient struct {
	RefreshToken    string
	AccessToken     string
	ResourceOwnerID string
	ExpirationDate  int
}

var baseURL = "https://hutomasage.azurewebsites.net/api/headers?code=MPWTXmQcWxclveAHs6ye6vOtgQRjqKnRSVLgaSL09SlNrdalRJUATA=="

type sageOauth struct {
	Credentials struct {
		RefreshToken    string `json:"refreshToken"`
		AccessToken     string `json:"accessToken"`
		ResourceOwnerId string `json:"resourceOwnerId"`
		ExpirationDate  int    `json:"expirationDate"`
	} `json:"credentials"`
	Url        string      `json:"url"`
	HttpMethod string      `json:"httpMethod"`
	Parameters interface{} `json:"parameters,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

func (sc *SageClient) getHeaders(apiURL, methodType string, params interface{}) (sr.HeadersResponse, error) {

	var response sr.HeadersResponse
	var headersBody sageOauth
	headersBody.Credentials.RefreshToken = sc.RefreshToken
	headersBody.Credentials.AccessToken = sc.AccessToken
	headersBody.Credentials.ResourceOwnerId = sc.ResourceOwnerID
	headersBody.Credentials.ExpirationDate = sc.ExpirationDate
	headersBody.Url = apiURL
	headersBody.HttpMethod = methodType

	if methodType == "POST" {
		headersBody.Parameters = params
	} else {
		headersBody.Body = params
	}

	jsonValue, err := json.Marshal(headersBody)
	//fmt.Println("JSONVALUE", string(jsonValue))
	if err != nil {
		return response, errors.New("Failed unmarshaling " + err.Error())
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return response, errors.New("Http error: " + err.Error())
	}

	//fmt.Println(req.Host)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return response, errors.New("Http error: " + err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	if err != nil {
		return response, errors.New("Error reading body")
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	sc.AccessToken = response.Credentials.AccessToken
	sc.RefreshToken = response.Credentials.RefreshToken
	sc.ResourceOwnerID = response.Credentials.ResourceOwnerID
	sc.ExpirationDate = response.Credentials.ExpirationDate

	fmt.Println("SC", sc)
	return response, nil
}

/*func (sc *SageClient) AddContact(name, contact_type_id string) error {
	params := //name, contact_type_id
	sc.getHeaders("https://api.columbus.sage.com/fr/sageone/accounts/v3/contacts", "POST",)
	return nil
}*/
func (sc *SageClient) ShowContacts() error {
	var headers sr.HeadersResponse
	headers, err := sc.getHeaders("https://api.columbus.sage.com/fr/sageone/accounts/v3/contacts", "GET", nil)
	if err != nil {
		return errors.New("error calling getHeaders: " + err.Error())
	}
	//fmt.Println(headers)

	contactsURL := "https://api.columbus.sage.com/fr/sageone/accounts/v3/contacts"
	req, err := http.NewRequest("GET", contactsURL, nil)
	if err != nil {
		return errors.New("Http error: " + err.Error())
	}

	for k, v := range headers.Headers {
		fmt.Println(k, v)
		req.Header.Set(k, v)
	}

	// hdstr, _ := json.Marshal(req.Header)
	// fmt.Println(string(hdstr))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return errors.New("Http error: " + err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("Error reading body")
	}
	fmt.Println(string(body))

	return nil
} /*
func (sc *SageClient) AddExpenditure() error {
	sc.getHeaders("https://api.columbus.sage.com/fr/sageone/accounts/v3/purchase_invoices", "POST",)

	return nil
}*/
