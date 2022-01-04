package structs

import (
	"bufio"
	"fmt"
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
		return nil, fmt.Errorf(err.Error())
	}
	defer file.Close()
	for scanner, step := bufio.NewScanner(file), 0; scanner.Scan(); {
		if step == 0 {
			step++
			ants, err := strconv.Atoi(scanner.Text())
			if ants <= 0 || err != nil {
				return nil, fmt.Errorf("invalid data format")
			}
			roomsAndAnts.Ants = ants
			continue
		}
		if scanner.Text() == "" {
			continue
		}
		if scanner.Text() == "##start" {
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			if len(temp) != 3 {
				return nil, fmt.Errorf("invalid data format")
			}
			roomsAndAnts.StartRoom = temp[0]
			continue
		}
		if scanner.Text() == "##end" {
			scanner.Scan()
			temp := strings.Split(scanner.Text(), " ")
			if len(temp) != 3 {
				return nil, fmt.Errorf("invalid data format")
			}
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
		if len(temp) != 3 {
			return nil, fmt.Errorf("invalid data format")
		}
		roomsAndAnts.Rooms = append(roomsAndAnts.Rooms, temp[0])
	}
	if roomsAndAnts.StartRoom == "" || roomsAndAnts.EndRoom == "" {
		return nil, fmt.Errorf("invalid data format")
	}
	return roomsAndAnts, nil
}
