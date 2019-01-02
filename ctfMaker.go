package main

import(
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"strings"
	"flag"
	"./folderManage"
	"io/ioutil"
	"log"
)

type Challenge struct{
        Name    string  `json:"name"`
        Desc    string  `json:"desc"`
        Value   int     `json:"value"`
        Flag    string  `json:"flag"`
        Type    string  `json:"type"`
}

type CTF struct{
	Name	string		`json:"name"`
	Challs	[]Challenge	`json:"challs"`
	Tags	[]string	`json:"tags"`
}

func jsonify(name string,challs []Challenge,tags []string){
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

// Function that lists all folders (and files) at CTFs folder
func listsCTF(){
    files, err := ioutil.ReadDir("CTFs")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            fmt.Println(" ->", f.Name())
    }
}

// Function to handle CTF creation
func createCTF(){
	// Receives user input and processes it
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type the name of the CTF")
	scanner.Scan()
	name := scanner.Text()
	path := "CTFs/" + name

	// Check if the CTF exists
	if _, err := os.Stat(path); !os.IsNotExist(err){
		fmt.Println("CTF exists!")
		os.Exit(1)
	}

	// Create folder of the CTF
	folderManage.CreateDir(path)

	// Create tags for the CTF
	fmt.Println("Type the tags separed by a space")
	scanner.Scan()
	tags := strings.Split(scanner.Text(), " ")

	var challs []Challenge
	jsonify(name, challs, tags)
}

func main(){
	// Dealing with the flags at the cli
	createPtr := flag.Bool("create", false, "create CTF")
	//viewPtr := flag.String("view", "", "view CTF with provided name")
	//editPtr := flag.String("edit", "", "edit CTF with provided name")
	listPtr := flag.Bool("list", false, "list existing CTFs")
	flag.Parse()

	// Creates directory CTFs if it doesn't exists
	folderManage.CreateDir("CTFs")
	fmt.Println("Executing ...")
	if(*listPtr){
		listsCTF()
	}
	if(*createPtr){
		createCTF()
	}
}


