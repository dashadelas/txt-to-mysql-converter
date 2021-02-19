package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//// os settings e putarias --------------
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//// stdin --------------
	fmt.Print("PATH -> ")
	fp := bufio.NewScanner(os.Stdin)
	fp.Scan()
	dbPath := fp.Text()
	fmt.Println("\n[ + ] FILEPATH DEFINIDO [ + ] \n   -> " + dbPath)

	//// mysql connection --------------
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/parsedb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//// fopen --------------
	file, err := os.Open(dbPath)
	//path manual, comente todo o *stdin
	//file, err := os.Open("C:/Users/ederm/Desktop/db.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//// Parser & Query --------------
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		explode := strings.Split(scanner.Text(), "|")
		//seta as vars
		var CPF, NOME, DATA string = explode[0], explode[1], explode[2]

		//query
		insert, err := db.Query("INSERT INTO teste (nome, cpf, data) VALUES (?, ?, ? )", NOME, CPF, DATA)
		if err != nil {
			panic(err.Error())
			//progress log
		} else {
			fmt.Println(
				"\nNOME: "+NOME,
				"\nCPF: "+CPF,
				"\nNASCIMENTO: "+DATA,
			)
		}
		defer insert.Close()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
