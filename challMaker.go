package main

import(
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"bufio"
	"strings"
)

type Challenge struct{
	Name	string	`json:"name"`
	Desc	string	`json:"desc"`
	Value	int	`json:"value"`
	Flag	string	`json:"flag"`
	Type	string	`json:"type"`
}

// Receives user input and creates a challenge
func create(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type the name of the challenge")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Type the description of the challenge and end with line being a dot alone \".\"")
	var desc []string
	for scanner.Scan(){
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

	description := strings.Join(desc,"\n")
	fmt.Println("Confirma os valores do chall?")
	fmt.Println("Name: ", name)
	//fmt.Println("Desc: ", strings.Join(desc, "\n"))
	fmt.Println("Desc: ", description)
	fmt.Println("Value: ", valor)
	fmt.Println("Flag: ", flag)
	fmt.Println("Category: ", category)

	jsonify(name, description, valor, flag, category)
}

// Turns the parameters into json and writes them into a file named "name.json"
func jsonify(name string, description string, valor int, flag string, category string){
	chall := Challenge{name, description, valor, flag, category}
	bs, _ := json.Marshal(chall)

	jsonName := name + ".json"

	// Writes json to file with permission only to the user
	ioutil.WriteFile(jsonName, bs, 0600)
}

func view(name string)(chall Challenge){
	jsonName := name + ".json"
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

func edit(name string){
	chall := view(name)
	fmt.Println("What do you want to edit?")
	fmt.Println("1 - Description")
	fmt.Println("2 - Value")

	var val, option int
	fmt.Scan(&option)

        scanner := bufio.NewScanner(os.Stdin)
	switch option{
	case 1:
		fmt.Println("Type the new description")
		var desc []string
		for scanner.Scan(){
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
	jsonify(chall.Name, chall.Desc, chall.Value, chall.Flag, chall.Type)
}

func main(){
	fmt.Println("Do you want to create, edit or view a challenge?")
	fmt.Println("1 	- Create")
	fmt.Println("2 	- View")
	fmt.Println("3 	- Edit")
	fmt.Println("99	- Exit")
	var option int
	fmt.Scan(&option)
	switch option{
	case 1:
		create()
	case 2:
		fmt.Println("\nType the name of the challenge")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		view(name)
	case 3:
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		edit(name)
	case 99:
		break;
	default:
		fmt.Println("\nInvalid option, try again")
		main()
	}
}

