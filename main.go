package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
   //  Ajoutez ici les champs pour représenter une tâche
   ID int
   Description string
   Completed bool

}
var tasks []Task

func (task *Task) CreateTask(description string){
	task.Description = description
}

func (task *Task) setID(){
	task.ID = len(tasks)
}

func (task *Task) setStatus(status bool){
	task.Completed = status
}



func main() {
    var response string

    for {
        fmt.Print("Voulez-vous ajouter une tâche ? (y/n): ")
        fmt.Scanln(&response)

        if response == "y" {
            
			task := Task{}

            for {
                fmt.Print("Entrez la description de la tâche: ")
				scan := bufio.NewScanner(os.Stdin)
				scan.Split(bufio.ScanLines)
				scan.Scan()
				description := scan.Text()

				if description != "" {
					task.setID()
                    task.CreateTask(description)
                    fmt.Println("Tâche créée:", task)
					tasks = append(tasks, task)
                    break
                } else {
                    fmt.Println("La tâche ne peut pas être vide. Veuillez réessayer.")
                }
            }
        } else if response == "n" {
            break
        } else {
            fmt.Println("Réponse invalide. Veuillez entrer 'y' ou 'n'.")
        }
    }

    fmt.Println("Liste des tâches enregistrées:")
    for _, task := range tasks {
        fmt.Printf("%d. %s (Complétée: %v)\n", task.ID, task.Description, task.Completed)
    }
}