package main
import(
	
	"fmt"
	"strings"
	"os"
	"bufio"
	
)

func task() []string {
    tasks := []string{}
    var userInput string
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Voulez-vous ajouter une tache ? (y/n):")
        userInput, _ = reader.ReadString('\n')
        userInput = strings.TrimSpace(userInput)

        if userInput == "n" {
            fmt.Println("Vous avez choisi de ne pas ajouter de tâche. Bon repos!")
            break
        } else if userInput == "y" {
            fmt.Println("Écrivez une description pour votre tâche:")
            taskDescription, _ := reader.ReadString('\n')
            taskDescription = strings.TrimSpace(taskDescription)
            tasks = append(tasks, taskDescription)  // Fix: Assign the result of append
        } else {
            fmt.Println("Entrée invalide, veuillez entrer 'y' ou 'n'.")
        }
    }

    return tasks
}

