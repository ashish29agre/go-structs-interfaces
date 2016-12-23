package structs

import "log"

type Player struct {
	Name  string
	Age   int
	Email string
}


type PlayerOperation interface {
	Jump()
	MoveForward()
	MoveBackward()
	Duck()

	Spawn()
	Die()
}


func (player Player) Spawn() {
	log.Println("Player " + player.Name + " is spawned")
}

func (player Player) Jump() {
	log.Println("Player " + player.Name + " is jumping now")
}


func (player Player) MoveForward() {
	log.Println("Player " + player.Name + " is moving forward now")
}


func (player Player) MoveBackward() {
	log.Println("Player " + player.Name + " is moving backward now")
}

func (player Player) Die() {
	log.Println("Player " + player.Name + " Died")
}

func (player Player) Duck() {
	log.Println("Player " + player.Name + " Ducked")
}
