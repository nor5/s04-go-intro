Aujourd'hui, nous allons approfondir notre compréhension de Go en implémentant des structures de données plus complexes et en utilisant des concepts avancés pour améliorer notre bot Discord.

## Objectifs pédagogiques

- **Appliquer** les concepts de structures, d'interfaces et de goroutines en Go
- **Créer** des handlers plus avancés pour notre bot Discord

## Étapes

1. Définissez des structures Go pour représenter les tâches.
go
```package main

import (
    "fmt"
)

type Task struct {
   //  Ajoutez ici les champs pour représenter une tâche
}

var tasks []Task
var nextID int

func main() {
    var response string

    for {
        fmt.Print("Voulez-vous ajouter une tâche ? (y/n): ")
        fmt.Scanln(&response)

        if response == "y" {
            var description string
            for {
                fmt.Print("Entrez la description de la tâche: ")
                fmt.Scanln(&description)
                if description != "" {
                    task := CreateTask(description)
                    fmt.Println("Tâche créée:", task)
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
    for i, task := range tasks {
        fmt.Printf("%d. %s (Complétée: %v)\n", task.ID, task.Description, task.Completed)
    }
}
```
 2. Implémentez un CRUD sur ces structures.
go
```
func createTask(task Task) Task {
    // Ajoutez ici le code pour créer une tâche
}

func readTask(id int) Task {
    // Ajoutez ici le code pour lire une tâche
}

func updateTask(task Task) Task {
    // Ajoutez ici le code pour mettre à jour une tâche
}

func deleteTask(id int) {
    // Ajoutez ici le code pour supprimer une tâche
}
```