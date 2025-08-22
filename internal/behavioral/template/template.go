package template

import "fmt"

type DataProcessor interface {
	ReadData()
	ProcessData()
	SaveData()
}

type BaseProcessor struct{}

func (b *BaseProcessor) Execute(p DataProcessor) {
	p.ReadData()
	p.ProcessData()
	p.SaveData()
}

type CSVProcessor struct {
	BaseProcessor
}

func (c *CSVProcessor) ReadData()    { fmt.Println("Reading CSV file") }
func (c *CSVProcessor) ProcessData() { fmt.Println("Processing CSV data") }
func (c *CSVProcessor) SaveData()    { fmt.Println("Saving CSV result") }

type JSONProcessor struct {
	BaseProcessor
}

func (j *JSONProcessor) ReadData()    { fmt.Println("Reading JSON file") }
func (j *JSONProcessor) ProcessData() { fmt.Println("Processing JSON data") }
func (j *JSONProcessor) SaveData()    { fmt.Println("Saving JSON result") }
