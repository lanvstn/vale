package nlp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SegmentResult struct {
	Sents []string
}

func post(url string) ([]byte, error) {
	var body []byte

	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

func segment(text, lang, apiURL string) (SegmentResult, error) {
	var result SegmentResult

	data := url.Values{"lang": {lang}, "text": {text}}
	path := apiURL + "/segment?" + data.Encode()

	body, err := post(path)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
