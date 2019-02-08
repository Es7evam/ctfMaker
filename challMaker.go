package main

import (
	"./libctfMaker"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// The Challenge struct stores a challenge options.
// It has Name, Description (Desc), Value, Flag and Category (Type)
type Challenge struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Value int    `json:"value"`
	Flag  string `json:"flag"`
	Type  string `json:"type"`
}

// Receives user input and creates a challenge
func create(challCTF string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type the name of the challenge")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Type the description of the challenge and end with line being a dot alone \".\"")
	var desc []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "." {
			break
		}
		desc = append(desc, line)
	}

	var valor int
	fmt.Println("Type the point value of the challenge")
	fmt.Scan(&valor)

	fmt.Println("Type the flag of the challenge")
	scanner.Scan()
	flag := scanner.Text()

	fmt.Println("Type the category of the challenge")
	scanner.Scan()
	category := scanner.Text()

	description := strings.Join(desc, "\n")
	fmt.Println("Please check the challenge attributes")
	fmt.Println("Name: ", name)
	//fmt.Println("Desc: ", strings.Join(desc, "\n"))
	fmt.Println("Desc: ", description)
	fmt.Println("Value: ", valor)
	fmt.Println("Flag: ", flag)
	fmt.Println("Category: ", category)

	jsonify(name, description, valor, flag, category, challCTF)
}

// Auxiliary function to get the path of a challenge.
func getpath(name string, challCTF string) string {
	if challCTF == "" {
		challCTF = "CTFs/standalone"
		// Creates standalone directory if it doesn't exist
		libctfMaker.CreateDir(challCTF)
	} else {
		challCTF = "CTFs/" + challCTF
		// If the CTF exists
		exists, _ := libctfMaker.FileExists(challCTF)
		if !exists {
			fmt.Println("CTF does not exist")
			os.Exit(1)
		}
	}
	path := challCTF + "/" + name + ".json"
	return path
}

// Turns the parameters into json and writes them into a file named "name.json"
func jsonify(name string, description string, valor int, flag string, category string, challCTF string) {
	chall := Challenge{name, description, valor, flag, category}
	bs, err := json.Marshal(chall)

	if err != nil {
		panic(err)
	}

	// if the challenge belongs to a ctf
	// then the path will be CTFs/ctfname/challenge.json
	jsonName := getpath(name, challCTF)

	// Writes json to file with permission only to the user
	ioutil.WriteFile(jsonName, bs, 0600)
}

// Function to visualize challenge with given name.
func view(name string, challCTF string) (chall Challenge) {
	jsonName := getpath(name, challCTF)
	jsonFile, err := os.Open(jsonName)
	if err != nil {
		fmt.Println("Non existing file")
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &chall)
	fmt.Println("\nName: " + chall.Name)
	fmt.Println("Description: " + chall.Desc)
	fmt.Println("Value: ", chall.Value)
	fmt.Println("Flag: ", chall.Flag)
	fmt.Println("Category: ", chall.Type)
	return chall
}

// Function to edit the challenge with given name.
func edit(name string, challCTF string) {
	chall := view(name, challCTF)
	fmt.Println("What do you want to edit?")
	fmt.Println("1 - Description")
	fmt.Println("2 - Value")

	var val, option int
	fmt.Scan(&option)

	scanner := bufio.NewScanner(os.Stdin)
	switch option {
	case 1:
		fmt.Println("Type the new description")
		var desc []string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "." {
				break
			}
			desc = append(desc, line)
		}
		chall.Desc = strings.Join(desc, "\n")
	case 2:
		fmt.Println("Type the new value")
		fmt.Scan(&val)
		chall.Value = val
	}
	jsonify(chall.Name, chall.Desc, chall.Value, chall.Flag, chall.Type, challCTF)
}

var ctfPtr string
var createPtr, viewPtr, editPtr *bool

// Init function
// Treats the argument flags.
func init() {
	createPtr = flag.Bool("create", false, "create challenge")
	viewPtr = flag.Bool("view", false, "view challenge")
	editPtr = flag.Bool("edit", false, "edit challenge")
	flag.StringVar(&ctfPtr, "ctf", "", "select ctf to associate the challenge of the string")
}

// Main function.
// It receives user input and calls the wanted functions.
func main() {
	flag.Parse()

	// No arguments provided
	fmt.Println("Ctfptr ", ctfPtr)

	if len(os.Args) < 2 {
		fmt.Println("No arguments provided.")
		fmt.Println("Usage of ", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *createPtr {
		fmt.Println(ctfPtr)
		create(ctfPtr)
	} else if *viewPtr {
		fmt.Println("\nType the name of the challenge")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		view(name, ctfPtr)
	} else if *editPtr {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		edit(name, ctfPtr)
	}
}
