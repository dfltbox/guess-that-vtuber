package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func downloadImages() {
	var count = 1
	file, err := os.OpenFile("imagelinks.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if count < 303 {
			count++
			continue
		}
		imageurl := scanner.Text()

		resp, err := http.Get(imageurl)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		if err := os.MkdirAll("images/"+strconv.Itoa(count), os.ModePerm); err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		file, err := os.Create("images/" + strconv.Itoa(count) + "/main.jpg")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		count++
		// Replace 352 with the actual amount of vtubers

		fmt.Println("Downloaded " + strconv.Itoa(count) + "/" + "352")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
