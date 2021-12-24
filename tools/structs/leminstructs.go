package structs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RoomsAndAnts struct {
	StartRoom string
	EndRoom   string
	Rooms     []string
	Edges     []string
	Ants      int
}

//FileRead - reads file and returns struct with rooms and count of ants
func FileRead(filepath string) *RoomsAndAnts {
	roomsAndAnts := &RoomsAndAnts{}
	file, err := os.Open("examples/" + filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		if scanner.Text() == "" {
			continue
		}
		if ants, err := strconv.Atoi(scanner.Text()); err == nil {
			roomsAndAnts.Ants = ants
			continue
		}
		if scanner.Text() == "##start" {
			fmt.Println(scanner.Text())
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			roomsAndAnts.StartRoom = temp[0]
			fmt.Println(scanner.Text())
			continue
		}
		if scanner.Text() == "##end" {
			fmt.Println(scanner.Text())
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			roomsAndAnts.EndRoom = temp[0]
			fmt.Println(scanner.Text())
			continue
		}
		if scanner.Text()[:1] == "#" {
			continue
		}
		if strings.Contains(scanner.Text(), "-") {
			roomsAndAnts.Edges = append(roomsAndAnts.Edges, scanner.Text())
			fmt.Println(scanner.Text())
			continue
		}
		temp := strings.Split(scanner.Text(), " ")
		roomsAndAnts.Rooms = append(roomsAndAnts.Rooms, temp[0])
		fmt.Println(scanner.Text())
	}
	return roomsAndAnts
}
