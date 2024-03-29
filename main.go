package wolframgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

var AppID string

func SetAppId(id string) {
	AppID = id
}

func GetSimpleResult(question string) (string, error) {
	link := "http://api.wolframalpha.com/v2/result?appid=" + AppID + "&output=json&input=" + strings.ReplaceAll(question, " ", "+")
	
	res, err := http.Get(link)
	defer res.Body.Close()

	if err != nil {
		return "", err
	}
	
	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func GetComplexResult(question string) (*ComplexResult, error) {
	link := "http://api.wolframalpha.com/v2/query?appid=" + AppID + "&includepodid=Result&output=json&input=" + strings.ReplaceAll(question, " ", "+")

	res, err := http.Get(link)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	a := baseResult{}

	if err := json.Unmarshal(result, &a); err != nil {
		return nil, err
	}

	return &a.Result, nil
}
