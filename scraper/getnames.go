package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	pushtodb()
}

func groupPage() {
	maxpages := 11
	currentpage := 1
	file, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for currentpage < maxpages {
		c := colly.NewCollector()

		c.OnHTML("div.mb-4 div.ratio a[href]", func(e *colly.HTMLElement) {
			fmt.Println(e.Attr("href"))
			_, err := file.WriteString(e.Attr("href") + "\n")
			if err != nil {
				fmt.Println(err)
			}
		})
		c.Visit("https://hololist.net/group/nijisanji-project/page/" + strconv.Itoa(currentpage) + "/")
		currentpage++
	}
}

func topPage() {
	maxpages := 4
	currentpage := 1
	file, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for currentpage < maxpages {
		c := colly.NewCollector()

		c.OnHTML("div.col-12.col-sm-6.col-lg-4 div.d-flex.mb-4.rounded a[href]", func(e *colly.HTMLElement) {
			fmt.Println(e.Attr("href"))
			_, err := file.WriteString(e.Attr("href") + "\n")
			if err != nil {
				fmt.Println(err)
			}
		})
		c.Visit("https://hololist.net/top/page/" + strconv.Itoa(currentpage) + "/?platform=twitch")
		currentpage++
	}
}
