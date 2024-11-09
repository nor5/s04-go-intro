package main

import (
	//"fmt"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Player struct{
	Name string // unique
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
func (player *Player) del () {
	delete(players, player.Name)
	fileName := player.Name + ".yml"
	fmt.Printf("this FileExist(fileName) funct %v : ", FileExist(fileName))
	if FileExist(fileName){
		err := os.Remove(fileName)
		if err != nil {
				log.Fatalf("Error while removing %v file",fileName)
			}

	}else{
		fmt.Printf("File %v does not exist !", fileName)
	}

}

func FileExist(fileName string) bool {
    var files []string
	dir, err := os.Getwd()    //get the current directory using the built-in function
	fileFound := false
   if err != nil {
	log.Fatalf("Error to get the current directory %V:",err)
      
   }
    err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
       if !d.IsDir() && filepath.Ext(path) == ".yml" {
          files = append(files, filepath.Base(path))
       }
       return nil
    })
    if err != nil {
       log.Fatal(err)
    }

	for _,file := range files {
		if file == fileName {
			fileFound = true
			break
		}
		
		
	}
	
	

    return fileFound
}