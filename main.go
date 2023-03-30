package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type user struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Proffesion string `json:"prof"`
}

type config struct {
	MeName string `json:"me"`
	Users  []user `json:"users"`
}

func main() {
	_, err := os.Stat("config.json")
	if err != nil {
		fmt.Printf("You started as new user. You need to write your name (nick), and emails and manes (micks) of all peoples, with whom you want to work.")
		var resultStruct config
		var isToWrite bool
		for !isToWrite {
			var result string
			var isReaded bool
			for i := 0; i < 2; i++ {
				for !isReaded {
					fmt.Printf("Please, write your name (nick): ")
					fmt.Scanln(&result)
					switch result {
					case "":
						continue
					case "*":
						continue
					case "****":
						continue
					case "fuck":
						continue
					case "Fuck":
						continue
					}
					resultStruct.MeName = result
					isReaded = true
				}
				for !isReaded {
					var result2 string
					var isReaded2 bool
					for !isReaded2 {
						fmt.Printf("How many users do you want to add? ")
						fmt.Scanln(&result2)
						if result2 == "" {
							continue
						}
						reaultInt, err := strconv.Atoi(result2)
						if err == nil {
							continue
						}
						if reaultInt < 0 {
							continue
						}
						resultStruct.Users = make([]user, reaultInt)
						for i := 0; i < reaultInt; i++ {
							var result3 string
							var result4 string
							var isReaded3 bool
							var isReaded4 bool
							for !isReaded3 {
								fmt.Printf("Please, write name (nick) of user " + strconv.Itoa(i+1) + ": ")
								fmt.Scanln(&result3)
								if result3 == "" {
									continue
								}
								isReaded3 = true
							}
							for !isReaded4 {
								fmt.Printf("Please, write email of user " + strconv.Itoa(i+1) + ": ")
								fmt.Scanln(&result4)
								if result4 == "" {
									continue
								}
								isReaded4 = true
							}
							resultStruct.Users[i].Name = result3
							resultStruct.Users[i].Email = result4
						}
						isReaded2 = true
					}
				}
			}
		}

		jsonData, err := json.Marshal(resultStruct)
		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}

		file, err := os.Create("config.json")
		if err != nil {
			panic(err)
		}
		file.Write(jsonData)
	}
}
