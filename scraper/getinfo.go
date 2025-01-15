package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type VtuberStruct struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Org         string   `json:"org"`
	Gender      string   `json:"gender"`
	Language    string   `json:"language"`
	Datedebuted string   `json:"datedebuted"`
	Status      string   `json:"status"`
	Nicknames   []string `json:"nicknames"`
	Url         string   `json:"url"`
}

func formatInfo(data string) string {
	output := ""
	return output
}

func getInfo() {
	file, err := os.OpenFile("names.txt", os.O_RDONLY, 0644)
	var count = 0
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	infofile, err := os.OpenFile("info.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer infofile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Visiting:", line)
		var name, country, org, gender, language, datedebuted, status string
		var nicknames []string
		url := line
		c := colly.NewCollector()

		c.OnHTML("h1.fs-4.m-0", func(e *colly.HTMLElement) {
			name = e.Text
		})
		c.OnHTML("div.d-flex.gap-2", func(e *colly.HTMLElement) {
			country = e.Text[1:]
		})
		//since it keeps detecting multiple debut dates we need to make it like.. not do that
		processed := false
		c.OnHTML("div#debut span:first-of-type", func(e *colly.HTMLElement) {
			if !processed {
				fullText := e.Text
				parts := strings.Split(fullText, "(")
				datedebuted = strings.TrimSpace(parts[0])
				processed = true
			}
		})
		c.OnHTML("div#affiliation a", func(e *colly.HTMLElement) {
			org = e.Text
		})
		c.OnHTML("div#gender a", func(e *colly.HTMLElement) {
			gender = e.Text
		})
		c.OnHTML("div#status", func(e *colly.HTMLElement) {
			status = strings.TrimSpace(e.DOM.Contents().Last().Text())
		})
		c.OnHTML("div#language", func(e *colly.HTMLElement) {
			language = strings.Replace(e.Text, "\nLanguage\n", "", -1)
		})
		c.OnHTML("div#nickname div", func(e *colly.HTMLElement) {
			if e.Text != "" {
				nicknames = strings.Split(e.Text, "\n")
			}
		})

		c.Visit(line)

		vtuber := VtuberStruct{
			Name:        name,
			Country:     country,
			Org:         org,
			Gender:      gender,
			Language:    language,
			Datedebuted: datedebuted,
			Nicknames:   nicknames,
			Status:      status,
			Url:         url,
		}
		jsonData, err := json.MarshalIndent(vtuber, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}
		count++
		//replace 722 with the actual amount of vtubers i was just lazy lol
		fmt.Println(strconv.Itoa(count) + "/" + "461")
		//fmt.Println(string(jsonData))
		_, err = infofile.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Rememeber to manually format the json file otherwise itll be invalid")

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
