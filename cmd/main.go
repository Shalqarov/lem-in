package main

import path "git.01.alem.school/MangoMango/lem-in/paths"

type Lem struct {
	StartRoom string
	EndRoom   string
	Rooms     []string
	Edges     []string
	Ants      int
}

func main() {
	// lem := Lem{}
	// args := os.Args[1:]
	// file, err := os.Open("test/" + args[0])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()
	// for scanner := bufio.NewScanner(file); scanner.Scan(); {
	// 	if scanner.Text() == "##start" {
	// 		fmt.Println(scanner.Text())
	// 		scanner.Scan()
	// 		temp := strings.Split(scanner.Text(), " ")
	// 		lem.StartRoom = temp[0]
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if scanner.Text() == "##end" {
	// 		fmt.Println(scanner.Text())
	// 		scanner.Scan()
	// 		temp := strings.Split(scanner.Text(), " ")
	// 		lem.EndRoom = temp[0]
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if strings.Contains(scanner.Text(), "-") {
	// 		lem.Edges = append(lem.Edges, scanner.Text())
	// 		fmt.Println(scanner.Text())
	// 		continue
	// 	}
	// 	if len(scanner.Text()) == 1 {
	// 		temp, _ := strconv.Atoi(scanner.Text())
	// 		lem.Ants = temp
	// 		continue
	// 	}
	// 	temp := strings.Split(scanner.Text(), " ")
	// 	lem.Rooms = append(lem.Rooms, temp[0])
	// 	fmt.Println(scanner.Text())
	// }
	// g := Graph{}
	// g.AddVertex(lem.StartRoom)
	// g.AddVertex(lem.EndRoom)
	// for _, v := range lem.Rooms {
	// 	g.AddVertex(v)
	// }
	// for _, v := range lem.Edges {
	// 	temp := strings.Split(v, "-")
	// 	g.AddEdge(temp[0], temp[1])
	// }

	g := &path.Graph{}
	for i := '1'; i <= '7'; i++ {
		g.AddVertex(string(i))
	}

	r := &path.Graph{}
	for i := '1'; i <= '7'; i++ {
		r.AddVertex(string(i))
	}

	g.AddEdge("1", "2")
	g.AddEdge("1", "3")
	g.AddEdge("4", "3")
	g.AddEdge("4", "2")
	g.AddEdge("4", "5")
	g.AddEdge("7", "5")
	g.AddEdge("4", "6")
	g.AddEdge("7", "6")

	r.AddEdge("1", "2")
	r.AddEdge("1", "3")
	r.AddEdge("4", "3")
	r.AddEdge("4", "2")
	r.AddEdge("4", "5")
	r.AddEdge("7", "5")
	r.AddEdge("4", "6")
	r.AddEdge("7", "6")

	g.PrintGraph()

	g.AllPathsBFS(r, "1", "7")
}

// g := &path.Graph{}
// 	for i := 'a'; i <= 'n'; i++ {
// 		g.AddVertex(string(i))
// 	}
// 	g.AddEdge("b", "a")
// 	g.AddEdge("d", "a")
// 	g.AddEdge("d", "i")
// 	g.AddEdge("j", "i")
// 	g.AddEdge("b", "e")
// 	g.AddEdge("b", "c")
// 	g.AddEdge("f", "c")
// 	g.AddEdge("e", "h")
// 	g.AddEdge("e", "g")
// 	g.AddEdge("h", "l")
// 	g.AddEdge("n", "l")
// 	g.AddEdge("n", "m")
// 	g.AddEdge("g", "f")
// 	g.AddEdge("g", "j")
// 	g.AddEdge("g", "k")
// 	g.AddEdge("m", "k")
// 	g.AddEdge("j", "m")

// 	r := &path.Graph{}
// 	for i := 'a'; i <= 'n'; i++ {
// 		r.AddVertex(string(i))
// 	}
// 	r.AddEdge("b", "a")
// 	r.AddEdge("d", "a")
// 	r.AddEdge("d", "i")
// 	r.AddEdge("j", "i")
// 	r.AddEdge("b", "e")
// 	r.AddEdge("b", "c")
// 	r.AddEdge("f", "c")
// 	r.AddEdge("e", "h")
// 	r.AddEdge("e", "g")
// 	r.AddEdge("h", "l")
// 	r.AddEdge("n", "l")
// 	r.AddEdge("n", "m")
// 	r.AddEdge("g", "f")
// 	r.AddEdge("g", "j")
// 	r.AddEdge("g", "k")
// 	r.AddEdge("m", "k")
// 	r.AddEdge("j", "m")
