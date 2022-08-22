package xurl

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// "title":{...}1 item
// "excerpt":{...}1 item
// "date":{...}1 item
// "author":{...}1 item
// "language":{...}1 item
// "url":{...}1 item
// "effective_url":{...}1 item
// "content":{...}1 item
type result struct {
	Title              string      `json:"title"`
	Excerpt            string      `json:"excerpt"`
	Date               interface{} `json:"date"`
	Author             interface{} `json:"author"`
	Language           string      `json:"language"`
	URL                string      `json:"url"`
	EffectiveURL       string      `json:"effective_url"`
	Domain             string      `json:"domain"`
	WordCount          int         `json:"word_count"`
	OgURL              interface{} `json:"og_url"`
	OgTitle            interface{} `json:"og_title"`
	OgDescription      interface{} `json:"og_description"`
	OgImage            interface{} `json:"og_image"`
	OgType             interface{} `json:"og_type"`
	TwitterCard        interface{} `json:"twitter_card"`
	TwitterSite        interface{} `json:"twitter_site"`
	TwitterCreator     interface{} `json:"twitter_creator"`
	TwitterImage       interface{} `json:"twitter_image"`
	TwitterTitle       interface{} `json:"twitter_title"`
	TwitterDescription interface{} `json:"twitter_description"`
	Content            string      `json:"content"`
}

var dryRun = true
func URLToText() error{
  url:="www.medialens.org/index.php/alerts/alert-archive/2018/868-douma-part-1.html"

  api, present := os.LookupEnv("RAPIDAPIKEY")
  if !present{
    return useFree(url)
  }
  return usePaid(url, api)
}

func useFree(u string) error{
  base:="https://txtify.it"
  apiURL,_:=url.JoinPath(base,u)
  if dryRun {
    fmt.Printf("curl %s",apiURL)
    return nil
  }

  tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
  client := &http.Client{Transport: tr}
  resp, err := client.Get(apiURL)
  if err!=nil{
    return err
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err!=nil{
    return err
  }
  fmt.Println(string(body))
  return nil
}

func usePaid(u, key string) error{
	apiURL := "https://full-text-rss.p.rapidapi.com/extract.php"
	payload := strings.NewReader("url="+u)
	req, _ := http.NewRequest("POST", apiURL, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "full-text-rss.p.rapidapi.com")
	res, err := http.DefaultClient.Do(req)
  if err!=nil{
    return err
  }
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
  return nil
}

