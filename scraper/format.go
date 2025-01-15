package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func format() {
	count := 1
	file, err := os.OpenFile("info.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	formattedfile, err := os.OpenFile("formatted.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer formattedfile.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []VtuberStruct
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, obj := range data {
		if obj.Status != "Active" {
			fmt.Println("Removed " + obj.Name + " from list")
			continue
		}

		validOrgs := []string{"Phase-Connect", "NIJISANJI Project", "hololive production", "VShojo", "Independent", "Idol", "Mythic Talent", "V&U", "Opera GX"}
		isValidOrg := false
		for _, org := range validOrgs {
			if obj.Org == org {
				isValidOrg = true
				break
			}
		}
		if !isValidOrg {
			fmt.Println("Removed " + obj.Name + " from list")
			continue
		}

		obj.ID = count
		count++

		jsonData, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}
		_, err = formattedfile.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}
