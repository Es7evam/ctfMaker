package main

import (
	"crypto/md5"
	"fmt"
	"github.com/es7evam/ctfmaker/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type category struct {
	Name  string
	Value string
}

type toSubmit struct {
	Types []category
	Token string
}

func uploadChall(w http.ResponseWriter, r *http.Request) {
	challFile := ""
	categories := []category{
		{Name: "Cryptography", Value: "crypto"},
		{Name: "Reverse Engineering", Value: "reveng"},
		{Name: "Web", Value: "web"},
		{Name: "Linux", Value: "linux"},
		{Name: "Programming", Value: "prog"},
		{Name: "Networking", Value: "network"},
		{Name: "Pwning", Value: "pwn"},
	}
	var send toSubmit

	send.Types = categories

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		send.Token = token

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, send)
	} else if r.Method == "POST" {
		// Get uploaded file
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		// If any parsing error occurs
		if err != nil {
			fmt.Println("No file has been sent")
			fmt.Println(err)
		} else {
			// Writes the file into ./uploads/filename
			fmt.Fprintf(w, "%v", handler.Header)
			reg, err := regexp.Compile("[^A-Za-z0-9]+")
			challFile = reg.ReplaceAllString(handler.Filename, "")

			f, err := os.OpenFile("./uploads/"+challFile, os.O_WRONLY|os.O_CREATE, 0666)

			// In case any saving error occurs
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			defer file.Close()

			// Copies the received data into the file
			io.Copy(f, file)
		}

		if len(r.Form["name"][0]) == 0 {
			fmt.Fprintf(w, "Please type a challenge name")
			//template.HTMLEscape(w, "The challenge name is invalid")) // responded to clients
		}

		// TODO -> Insert at Challenge struct, validate category input
		// 		-> Use files
		var challValue int

		// Regex to remove path traversal vulnerability
		reg, err := regexp.Compile("[^A-Za-z0-9]+")

		challName := reg.ReplaceAllString(template.HTMLEscapeString(r.Form.Get("name")), "")
		challDesc := template.HTMLEscapeString(r.Form.Get("desc"))
		challFlag := template.HTMLEscapeString(r.Form.Get("flag"))
		challType := template.HTMLEscapeString(r.Form.Get("category"))
		fmt.Sscan(r.Form.Get("value"), &challValue)

		fmt.Printf("Name: %s\n Desc: %s\n Flag: %s\n Type: %s\n Value: %d\n", challName, challDesc, challFlag, challType, challValue)
		fmt.Printf("Filename: %s\n", challFile)

		chall := models.Challenge{challName, challDesc, challValue, challFlag, challType}
		JsonifyChall(chall, "webServer")
	}
}

func main() {
	// setting router rule
	http.HandleFunc("/", uploadChall)

	// Setting up listener
	err := http.ListenAndServe(":9090", nil)
	// Error setting up listener
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
