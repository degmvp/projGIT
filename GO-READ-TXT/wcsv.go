package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data/email.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records := [][]string{
		{"degmar barbosa", "degmvp@gmail.com"},
		{"Amanda Moreira", "amanda.moreira@teste.io"},
		{"João Santos", "joao.santos@email.me"},
		{"Valentina Silva", "valentina.silva@tutorial.com"},
		{"Lawrence Barbosa", "lawrencebarbosa@gmail.com"},
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
