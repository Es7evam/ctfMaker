package main

import(
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"strings"
	"flag"
	"./libctfMaker"
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
	libctfMaker.CreateDir(path)

	// Create tags for the CTF
	fmt.Println("Type the tags separed by a space")
	scanner.Scan()
	tags := strings.Split(scanner.Text(), " ")

	var challs []Challenge
	jsonify(name, challs, tags)
}

func viewCtf(name string)(ctf CTF){
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
		List CTF Challenges -> TODO
	if(ctf.Challs != nil){
		fmt.Println("Challenges: ", strings.Join(ctf.Challs.Name, ", "))
	}
	*/

	if(ctf.Tags != nil){
		fmt.Println("Tags: ", strings.Join(ctf.Tags, ", "))
	}
	return ctf
}

func main(){
	// Dealing with the flags at the cli
	createPtr := flag.Bool("create", false, "create CTF")
	viewCtfPtr := flag.String("viewctf", "", "view provided CTF informations")
	//editPtr := flag.String("edit", "", "edit CTF with provided name")
	listPtr := flag.Bool("list", false, "list existing CTFs")
	flag.Parse()

	// No arguments provided 
        if len(os.Args) < 2{
                fmt.Println("No arguments provided.")
                fmt.Println("Usage of ", os.Args[0])
                flag.PrintDefaults()
                os.Exit(0)
        }

	// Creates directory CTFs if it doesn't exists
	libctfMaker.CreateDir("CTFs")
	fmt.Println("Executing ...")
	if(*listPtr){
		listsCTF()
	}
	if(*createPtr){
		createCTF()
	}
	if(*viewCtfPtr != ""){
		viewCtf(*viewCtfPtr)
	}
}


