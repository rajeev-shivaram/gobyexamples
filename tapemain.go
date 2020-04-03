package main

import "github.com/gobyexample/gadget"

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range(songs){
		device.Play(song)
	}
	device.Stop()
}

func main(){
	player := gadget.TapePlayer{}
	mixtape := []string{"8miles", "Scream and shout"}
	playList(player, mixtape)
}