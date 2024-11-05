package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// func add(a int, b int) int {
// 	return a + b

// }
// var sum int










func main(){
	// fmt.Println("add",add(3,4))
	// sum = add(5,4)
	// fmt.Printf("sum: %v \n", sum)
	// fmt.Printf("la multiplication de 3 et 4: %v \n", multip(3,4))
	// resultas, err := division(3,0)
	//  if err != nil {
	// 	log.Fatalf("Aletre la division par zero : %v", err)
	//  }
	//  log.Printf("la divisionde 3 et 4: %v \n", resultas)

	//var task string
	var reponse string
	var taskinput string
	log.Printf("Ajouter une tache : (y/n)")
	fmt.Scanln(&reponse)
	switch reponse {
	case "n", "no","non":
		log.Printf("reponse non")
	case "y", "o", "yes","oui":
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanLines)
		log.Println("reponse yes")
		log.Println("saisir la description de la tache :")
        scan.Scan()
		input := scan.Text()
		taskinput = input
		log.Printf("valeur saisi: %v", input)

		
	default:
		log.Printf("reponse inconnue : %v", reponse)

	}
	log.Printf("la tache saisi est: %v", taskinput)

	var tasks []string
	tasks = append(tasks, "hello","world")

	log.Println(tasks)
	




}

	
