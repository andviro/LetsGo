package main

import "fmt"

type Conference struct {
	Name   string
	income int // HL
}

func main() {
	confs := []Conference{
		{"Стачка", 0},
		{"RailsClub", 100500},
		{"RailsClub'Ulyanovsk", 300},
	}
	for _, conf := range confs {
		if conf.income >= 10000 {
			conf.income = conf.income * 0.8 // HL
		} else {
			conf.income = conf.income * 0.9 // HL
		}
		fmt.Printf("%#v\n", conf)
	}
}
