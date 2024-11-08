package main

type Player struct{
	Nom string
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

}