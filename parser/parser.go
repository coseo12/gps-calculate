package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Article struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetJson(articles []Article) {
	doc, _ := json.Marshal(articles) // data를 JSON 문서로 변환

	err := ioutil.WriteFile("./articles.json", doc, os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		log.Panic(err)
	}
}

func GetJson() []Article {
	b, err := ioutil.ReadFile("./articles.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		log.Panic(err)
	}

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장

	return data
}
