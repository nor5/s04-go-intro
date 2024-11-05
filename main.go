package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
   //  Ajoutez ici les champs pour représenter une tâche
   ID int
   Description string
   Completed bool

}
var tasks []Task

func (task *Task) setDescription(description string){
	task.Description = description
}

func (task *Task) setID(){
	task.ID = len(tasks)
}

func (task *Task) setStatus(status bool){
	task.Completed = status
}

func  createTask(){
	
	task := Task{}
	for {
		fmt.Print("Entrez la description de la tâche: ")
		scan := bufio.NewScanner(os.Stdin)
		scan.Split(bufio.ScanLines)
		scan.Scan()
		description := scan.Text()

		if description != "" {
			task.setID()
			task.setDescription(description)
			fmt.Println("Tâche créée:", task)
			tasks = append(tasks, task)
			break
		} else {
			fmt.Println("La tâche ne peut pas être vide. Veuillez réessayer.")
		}
	}
}

func  readTask(ID int)interface{}{
	if ID < len(tasks){
		fmt.Printf("%d. %s (Complétée: %v)\n", tasks[ID].ID, tasks[ID].Description, tasks[ID].Completed)
		return tasks[ID]

	}else{
		err := errors.New("index out of range")
		return err
	}


}

func  updateTask(ID int)interface{} {
	if ID < len(tasks){
		task := tasks[ID]
		var response string
		outerLoop:
		for{
			fmt.Print("Voulez-vous modifier la [D]escription ou le [S]tatus [E]xit ? : ")
        	fmt.Scanln(&response)
			switch (response){
				case "D", "d", "description":
					for {
						fmt.Print("Modifier la description de la tâche: ")
						scan := bufio.NewScanner(os.Stdin)
						scan.Split(bufio.ScanLines)
						scan.Scan()
						description := scan.Text()
				
						if description != "" {
							
							task.setDescription(description)
							fmt.Println("Tâche modifié:", task)
							
							break
						} else {
							fmt.Println("La tâche ne peut pas être vide. Veuillez réessayer.")
						}
					}
				case "S", "s", "status":
					for {
						fmt.Print("Modifier le status de la tâche: true or false : ")
						scan := bufio.NewScanner(os.Stdin)
						scan.Split(bufio.ScanLines)
						scan.Scan()
						status := scan.Text()
						boolStatus, err := strconv.ParseBool(status)
						if err != nil {
							fmt.Println("Invalid input please enter true or false")
						}else{
							task.setStatus(boolStatus)
							fmt.Println("Tâche modifié:", task)
							
							break
						}

					}
				case "E", "e", "exit":
					break outerLoop
						
				
						
					



			}
		}
			
				



		

		fmt.Printf("%d. %s (Complétée: %v)\n", task.ID, task.Description, task.Completed)
		return task

	}else{
		err := errors.New("index out of range")
		return err
	}


}



func  deleteTask(ID int) {
	
	if ID < len(tasks){
		task := tasks[ID]
		tasks = append(tasks[:ID], tasks[ID+1:]... )
		 fmt.Printf("task %v succusfuly deleted \n",task.ID)
		

	}else{
		err := errors.New("index out of range")
		fmt.Print(err) 
	}


}



func main() {
    var response string

    for {
        fmt.Print("Voulez-vous ajouter une tâche ? (y/n): ")
        fmt.Scanln(&response)

        if response == "y" {
            
           createTask()
        } else if response == "n" {
            break
        } else {
            fmt.Println("Réponse invalide. Veuillez entrer 'y' ou 'n'.")
        }
    }
	deleteTask(0)

    fmt.Println("Liste des tâches enregistrées:")
    for _, task := range tasks {
        fmt.Printf("%d. %s (Complétée: %v)\n", task.ID, task.Description, task.Completed)
    }
	//fmt.Println(readTask(0))
	//fmt.Println(updateTask(0))

}