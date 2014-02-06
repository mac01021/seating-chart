package main

import "bufio"
import "fmt"
import "os"
import "strings"

func seatingChart(filename string) (chart map[Guest]Table){
	chart = make(map[Guest]Table)
	var seatingFile, err = os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer seatingFile.Close()
	scanner := bufio.NewScanner(seatingFile)
	var curTab = ""
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				last := len(line) - 1
				curTab = line[1:last]
			} else {
				if curTab == "" {
					panic("Name without table")
				}
				var guest = unpack(line)
				chart[guest] = Table(curTab)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading 'seating_chart':", err)
	}
	return
}

func getInitialArrivals(filename string) (priorArrivals []Guest){
	var arrivalsFile, err = os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Proceeding with no initial arrival history...")
		return
	}
	defer arrivalsFile.Close()
	scanner := bufio.NewScanner(arrivalsFile)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			g := unpack(line)
			priorArrivals = append(priorArrivals, g)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	return
}


