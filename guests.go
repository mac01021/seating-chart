package main

import "fmt"
import "strings"


type Guest struct {
	Surname string
	GivenName string
}

func (g Guest) Name() string {
	return fmt.Sprintf("%v, %v", g.Surname, g.GivenName)
}

func unpack(name string) Guest {
	var sur string
	var given string
	if strings.Contains(name, ",") {
		parts := strings.SplitN(name, ",", 2)
		sur = strings.TrimSpace(parts[0])
		given = strings.TrimSpace(parts[1])
	} else {
		parts := strings.Fields(name)
		last := len(parts) - 1
		sur = parts[last]
		gParts := parts[:last]
		given = strings.Join(gParts, " ")
	}
	return Guest{sur, given}
}


type Table string

type Chart map[Guest]Table

func jsonable(chart Chart) (it map[string]string) {
	it = make(map[string]string)
	for k, v := range chart {
		it[k.Name()] = string(v)
	}
	return
}

