package goutils


import (
  "bytes"
  //"fmt"

  "net/http"
  "io/ioutil"
)


func GetWithHeaders(url string, headers map[string]string) ([]byte, error) {
  return makeRequestWithHeaders(url, "GET", nil, headers)
}

func PostWithHeaders(url string, data []byte, headers map[string]string) ([]byte, error) {
  return makeRequestWithHeaders(url, "POST", data, headers)
}


func makeRequestWithHeaders(url string, method string, data []byte, headers map[string]string) ([]byte, error) {
  client := &http.Client{}
  req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
  if err != nil {
    return nil, err
  }

  for k, v := range headers {
    req.Header.Add(k, v)
  }


  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  resp_data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return resp_data, nil
}
