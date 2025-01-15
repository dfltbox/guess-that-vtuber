package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/chromedp/chromedp"
)

func getImages() {
	var count = 0
	var imageUrl string
	//https://hololist.net/wp-content/themes/primary/assets/images/no-image.png

	file, err := os.OpenFile("formatted.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	imgfilelink, err := os.OpenFile("imagelinks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var data []VtuberStruct
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	for _, obj := range data {
		line := obj.Url
		fmt.Println("Visiting:", obj.Name)
		count++

		var fetchImage func(string) error
		fetchImage = func(url string) error {
			err := chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.WaitReady(`img.rounded-circle`, chromedp.ByQuery),
				chromedp.AttributeValue(`img.rounded-circle`, "src", &imageUrl, nil),
			)
			if err != nil {
				return err
			}
			if imageUrl == "https://hololist.net/wp-content/themes/primary/assets/images/no-image.png" {
				fmt.Println("Retrying")
				return fetchImage(url)
			}
			return nil
		}

		err := fetchImage(line)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Image URL:", imageUrl)
			imgfilelink.WriteString(imageUrl + "\n")
		}

		// Replace 722 with the actual amount of vtubers
		fmt.Println(strconv.Itoa(count) + "/" + "351")
	}

}
