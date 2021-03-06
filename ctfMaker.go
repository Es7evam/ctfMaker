package main

import (
	"bufio"
	"encoding/json"

	"fmt"
	"github.com/es7evam/ctfmaker/libctfmaker"
	"github.com/es7evam/ctfmaker/models"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// The CTF struct stores the CTF structure
// It has Name, a list of challenges (Challs) and Tags.
type CTF struct {
	Name   string             `json:"name"`
	Challs []models.Challenge `json:"challs"`
	Tags   []string           `json:"tags"`
}

// Gets the name, challenges and tags of a CTF and saves it into a .json file
// The file will be at CTFs/ctfname/config.json
func jsonify(name string, challs []models.Challenge, tags []string) {
	ctf := CTF{name, challs, tags}

	// Create json file
	bs, err := json.Marshal(ctf)
	// Error check when creating json file
	if err != nil {
		panic(err)
	}

	path := "CTFs/" + name + "/"
	jsonName := path + "config.json"

	// Writes json to file CTFs/name/config.json
	ioutil.WriteFile(jsonName, bs, 0644)
}

// ListCTF Function that lists all folders (and files) at CTFs folder
func ListCTF() {
	files, err := ioutil.ReadDir("CTFs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(" ->", f.Name())
	}
}

// CreateCTF Function to handle CTF creation
func CreateCTF() {
	// Receives user input and processes it
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type the name of the CTF")
	scanner.Scan()
	name := scanner.Text()
	path := "CTFs/" + name

	// Check if the CTF exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		fmt.Println("CTF exists!")
		os.Exit(1)
	}

	// Create folder of the CTF
	libctfmaker.CreateDir(path)

	// Create tags for the CTF
	fmt.Println("Type the tags separed by a space")
	scanner.Scan()
	tags := strings.Split(scanner.Text(), " ")

	var challs []models.Challenge
	jsonify(name, challs, tags)
}

// ViewCTF Function to view the given ctf attributes.
// It will get the content from the config.json file.
func ViewCTF(name string) (ctf CTF) {
	// Casts into json file
	jsonName := "CTFs/" + name + "/config.json"

	jsonFile, err := os.Open(jsonName)
	if err != nil {
		fmt.Println("Non existing file")
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &ctf)

	fmt.Println("\nName: ", ctf.Name)

	/*
			List CTF models.Challenges -> TODO
		if(ctf.Challs != nil){
			fmt.Println("models.Challenges: ", strings.Join(ctf.Challs.Name, ", "))
		}
	*/

	if ctf.Tags != nil {
		fmt.Println("Tags: ", strings.Join(ctf.Tags, ", "))
	}
	return ctf
}

// Main function.
// It deals with the arguments passed using flags.
// In case it is the first execution, it creates the necessary folders.
/*
func main() {
	// Dealing with the flags at the cli
	createPtr := flag.Bool("create", false, "create CTF")
	viewCtfPtr := flag.String("viewctf", "", "view provided CTF informations")
	//editPtr := flag.String("edit", "", "edit CTF with provided name")
	listPtr := flag.Bool("list", false, "list existing CTFs")
	flag.Parse()

	// No arguments provided
	if len(os.Args) < 2 {
		fmt.Println("No arguments provided.")
		fmt.Println("Usage of ", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Creates directory CTFs if it doesn't exists
	libctfMaker.CreateDir("CTFs")
	fmt.Println("Executing ...")
	if *listPtr {
		listCTF()
	}
	if *createPtr {
		CreateCTF()
	}
	if *viewCtfPtr != "" {
		ViewCTF(*viewCtfPtr)
	}
}
*/
