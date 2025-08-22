package main

import "github.com/barantoraman/DesignPatterns/internal/behavioral/template"

func main() {
	csv := &template.CSVProcessor{}
	json := &template.JSONProcessor{}

	csv.Execute(csv)
	json.Execute(json)
}
