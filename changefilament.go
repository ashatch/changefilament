package main

import "bufio"
import "fmt"
import "os"
import "flag"
import "strings"

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Must supply a .gcode file and a comma-separated list of layer numbers")
		os.Exit(1)
	}

	gcodeFile := args[0]

	layerNumbers := strings.Split(args[1], ",")
	layerMatches := make(map[string]struct{}, len(layerNumbers))

	for i := 0; i < len(layerNumbers); i++ {
		key := ";LAYER:" + layerNumbers[i]
		fmt.Println(key)
		layerMatches[key] = struct{}{}
	}

	f, _ := os.Open(gcodeFile)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()		
		fmt.Println(line)

		_, isPresent := layerMatches[line]

		if isPresent {
			fmt.Println("M600")
		}
	}
}
