package main

import (
	"fmt"
)

type song struct {
	name   string
	author string
	next   *song
}

// func printSongList(songList *song) {
// 	for s := songList; s != nil; s = s.next {
// 		fmt.Println(s)
// 	}
// }

// func (s *song) addSong(newSong *song) *song {
// 	// for s := songList; s != nil; s = s.next {
// 	// 	if s.next == nil {
// 	// 		s.next = newSong
// 	// 		return songList
// 	// 	}
// 	// }
// 	// return songList
// }

func main() {
	// soWhat := &song{"so what", "pink", nil}
	// justGiveMeAReason := &song{"just give me a reason", "pink", nil}
	// walkMeHome := &song{"walk me home", "pink", nil}

	songList := new(song)

	// printSongList(songList)
	fmt.Println(songList)
}
