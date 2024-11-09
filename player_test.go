package main
import (
	
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
