package main

import "fmt"

func main() {
	var mymap map[string]string
	if mymap == nil {
		fmt.Println("mymap是个指向为nil的指针")
	}
	mymap = make(map[string]string)
	if mymap != nil {
		fmt.Println("mymap is not nil")
	}
	mymap["aaa"] = "1111"
	mymap["bbb"] = "2222"
	mymap["ccc"] = "3333"
	mymap["ddd"] = "4444"
	for key, value := range mymap {
		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("=====================")
	for key, _ := range mymap {
		fmt.Println(mymap[key])
	}

	Student := map[string]interface{}{
		"name": "马薇薇",
		"Age":  28,
	}
	for key, value := range Student {
		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("=====================")
	Student["Age"] = 30
	for key, _ := range Student {
		fmt.Println(Student[key])
	}
	fmt.Println("=====================")
	delete(Student, "name")
	for key, _ := range Student {
		fmt.Println(Student[key])
	}

}
