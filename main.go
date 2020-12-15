package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	fmt.Println("JSON")
	//Usado para converter um Map ou Struct para JSON
	//json.Marshal()
	//Faz o processo reverso
	//json.Unmarshal()

	c := cachorro{"Lulu", "Basset", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "tobby",
		"raca": "vira-lata",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("------------------------------")
	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
