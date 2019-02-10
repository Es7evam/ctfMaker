package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func uploadChall(w http.ResponseWriter, r *http.Request) {
	challFile := ""

	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
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
			f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

			// In case any saving error occurs
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			defer file.Close()

			// Copies the received data into the file
			io.Copy(f, file)
			challFile = handler.Filename
		}

		if len(r.Form["name"][0]) == 0 {
			fmt.Fprintf(w, "Please type a challenge name")
			//template.HTMLEscape(w, "The challenge name is invalid")) // responded to clients
		}

		// TODO -> Insert at Challenge struct, validate category input
		challName := template.HTMLEscapeString(r.Form.Get("name"))
		challDesc := template.HTMLEscapeString(r.Form.Get("desc"))
		challFlag := template.HTMLEscapeString(r.Form.Get("flag"))
		challType := template.HTMLEscapeString(r.Form.Get("category"))

		var challValue int
		fmt.Sscan(r.Form.Get("value"), &challValue)
		//challFile := template.HTMLEscapeString(handler.Filename)
		fmt.Printf("Name: %s\n Desc: %s\n Flag: %s\n Type: %s\n Value: %d\n", challName, challDesc, challFlag, challType, challValue)
		fmt.Printf("Filename: %s\n", challFile)
	}

}

func main() {
	// setting router rule
	http.HandleFunc("/", uploadChall)
	http.HandleFunc("/login", login)

	// Setting up listener
	err := http.ListenAndServe(":9090", nil)
	// Error setting up listener
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
