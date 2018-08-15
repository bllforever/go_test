package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int
	Price int
}

func main() {
	movie := Movie{"天使爱美丽", 2008, 30}
	json_str, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Marshal err", err)
	}
	//fmt.Println(json_str)
	fmt.Printf("%s\n", json_str)
	mymovie1 := Movie{}
	json.Unmarshal(json_str, &mymovie1)
	fmt.Println("=================")
	fmt.Printf("%+v\n", mymovie1)
	fmt.Println(mymovie1.Title)

	mymovie := map[string]interface{}{}
	json.Unmarshal(json_str, &mymovie)
	fmt.Println("=================")
	fmt.Printf("%+v\n", mymovie)
	fmt.Println(mymovie["Title"])

}
