package main

import (
	//"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Player struct{
	Name string
	Username string
	Pseudo string
	Age int
	Health int
	State string
} 

var (
	players = make(map[string]*Player)
)

func (player *Player) save(){
	fileName := player.Name + ".yml"
	yamlData, err := yaml.Marshal(&player)
	if err != nil {
		log.Fatalf("Error when marshling data: %v",err)
		
	}
	//fmt.Println(string(yamlData))

	yamlFile, err := os.Create(fileName)
	if err != nil {
		
		log.Fatalf("Error when creating file: %v", err)
	}
	defer yamlFile.Close()
	_, err = io.Writer.Write(yamlFile,yamlData)
	//os.WriteFile(fileName,yamlData,0666)
	
	if err != nil {
		
		log.Fatalf("error when writing file: %v",err)
	}
	

}