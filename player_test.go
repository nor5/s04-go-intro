package main

import (
	"fmt"
	"log"
	"testing"
)
func TestSave(t *testing.T){
	
		p := Player{
			Name: "noor",
			Username: "nor",
			Pseudo: "nor5",
			Age: 55,
			Health: 100,
			State: "active",
		}
		p.save()
		
	}

func TestDel(t *testing.T){
	p := Player{
		Name: "noor",
		Username: "nor",
		Pseudo: "nor5",
		Age: 55,
		Health: 100,
		State: "active",
	}
	p.del()
}

func TestDisplay(t *testing.T){
	p := Player{
		Name: "noor",
		Username: "nor",
		Pseudo: "nor5",
		Age: 55,
		Health: 100,
		State: "active",
	}
	p.display()
}

func TestLoadPlayer(t *testing.T){
	f,err := loadPlayer("amir")
	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Print(f)
	}

	
}