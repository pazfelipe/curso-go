package main

import "fmt"

const (
	_ = 2022 + iota
	ano2020
	ano2021
	ano2022
	ano2023
)

func main() {
	fmt.Println(ano2020, ano2021, ano2022, ano2023)
}
