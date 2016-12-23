package main

import "log"
import (
	"structs"
)

func main() {
	log.Println("All Working fine")
	player := structs.Player{Name:"Ashish", Email:"ashish@ashish.com", Age:21}
	log.Println("name is: " + player.Name)
	var playerOps structs.PlayerOperation
	playerOps = player
	playerOps.Spawn()
	playerOps.Jump()
	playerOps.MoveForward()
	playerOps.MoveBackward()
	playerOps.Duck()
	playerOps.Die()
}
