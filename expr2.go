package main

import (
	"bufio"
	"fmt"
	"os"
)

func expr2() {
	// In their uninitialized forms they are (nil)
	var dict map[string]string
	//dict["x"] = "y" // can not assign to a nil map
	var list []int

	fmt.Println(dict == nil)
	fmt.Println(list == nil)

	// map literal
	newDict := map[string]string{
		"x": "y",
	}
	newDict["y"] = "z"

	// if a key is not found the default value is going to be returned
	val := newDict["not-found"]
	fmt.Println(val == "")

	// the comma ok syntax for maps
	if _, ok := newDict["not-found"]; !ok {
		fmt.Println("does not exist")
	}

	dict1 := map[int]int{
		1:  2,
		10: 5,
		2:  3,
	}
	dict2 := map[int]int{
		10: 5,
		2:  3,
		1:  2,
	}
	// can hack map comparison like below
	first := fmt.Sprintf("%#v", dict1)
	second := fmt.Sprintf("%#v", dict2)
	fmt.Println(first == second)

	// map declaration (uninitialized)
	var myMap2 map[string]int
	_ = myMap2

	// map literal (initialized)
	myDict := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}

	fmt.Printf("%#v\n", myDict)
	for k := range myDict {
		delete(myDict, k)
		//delete(myDict, k) // deleting a non-existing key is a noop
	}

	fmt.Printf("%#v\n", myDict)

	// make (initialized)
	myMap := make(map[string]int)
	myMap["x"] = 1
	myMap["y"] = 2
	myMap["z"] = 3

	fmt.Printf("%#v\n", myMap)

	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		input := in.Text()
		fmt.Printf("line no %d: %s\n", lines, input)
		lines++
	}

	if err := in.Err(); err != nil {
		fmt.Println(err)
	}

	type person struct {
		name, lastname string
		age            int
	}

	picasso := person{
		name:     "Pablo",
		lastname: "Picasso",
		age:      91,
	}

	var freud person
	freud.name = "Sigmund"
	freud.lastname = "Freud"
	freud.age = 83

	fmt.Printf("%+v\n%+v\n", picasso, freud)

	type song struct {
		title, artist string
	}

	type playlist struct {
		genre string
		songs []song
	}

	song1 := song{title: "wonderwall", artist: "oasis"}
	song2 := song{title: "super sonic", artist: "oasis"}

	songs := []song{song1, song2}
	myPlaylist := playlist{genre: "xd", songs: songs}
	fmt.Printf("%+v\n", myPlaylist)

	type text struct {
		title string
		words int
	}

	type book struct {
		text // embedding
		isbn string
	}

	moby := book{
		text: text{title: "moby dick", words: 206052}, // Initialize the embedded struct
		isbn: "102030",
	}
	fmt.Printf("%+v", moby)
	fmt.Println(moby.title) // this works because of embedding

	// no need to specify book{} for every item in the list
	myBooks := []book{
		{
			text: text{title: "xd", words: 69},
			isbn: "holy mother of god",
		},
		{
			text: text{title: "xdd", words: 70},
			isbn: "holy mother of jesus",
		},
	}

	fmt.Printf("%+v\n", myBooks)
}
