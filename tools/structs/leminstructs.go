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
func FileRead(filepath string) (*RoomsAndAnts, error) {
	roomsAndAnts := &RoomsAndAnts{}
	file, err := os.Open("examples/" + filepath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for scanner, step := bufio.NewScanner(file), 0; scanner.Scan(); {
		if step == 0 {
			step++
			if ants, err := strconv.Atoi(scanner.Text()); err == nil {
				roomsAndAnts.Ants = ants
				continue
			} else if ants <= 0 || err != nil {
				return nil, fmt.Errorf("invalid number of ants")
			}
		}
		if scanner.Text() == "" {
			continue
		}
		if scanner.Text() == "##start" {
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			roomsAndAnts.StartRoom = temp[0]
			continue
		}
		if scanner.Text() == "##end" {
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			roomsAndAnts.EndRoom = temp[0]
			continue
		}
		if scanner.Text()[:1] == "#" {
			continue
		}
		if strings.Contains(scanner.Text(), "-") {
			roomsAndAnts.Edges = append(roomsAndAnts.Edges, scanner.Text())
			continue
		}
		temp := strings.Split(scanner.Text(), " ")
		roomsAndAnts.Rooms = append(roomsAndAnts.Rooms, temp[0])
	}
	return roomsAndAnts, nil
}
