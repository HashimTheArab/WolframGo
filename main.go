package wolframgo

import (
  "fmt"
  "strings"
  "io/ioutil"
  "net/http"
)

var appid string

func GetSimpleResult(question string) (string, error) {

  args := strings.Split(question, " ")
  question = strings.Join(args, "+")
  link := fmt.Sprintf("http://api.wolframalpha.com/v2/result?appid=%s&output=json&input=%s", appid, question)

  res, err := http.Get(link)

  if err != nil {
    return "", err
  }

  defer res.Body.Close()

  result, err := ioutil.ReadAll(res.Body)
    
  if err != nil {
    return "", err
  }

  return string(result), nil
}