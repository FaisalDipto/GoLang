package part4

import "fmt"

type DataExporter interface{
	Export() []byte
}

type JSONExporter struct{}
type CSVExporter struct{}

func (j JSONExporter) Export() []byte{
	return []byte("Exporting to JSON...")
}

func (c CSVExporter) Export() []byte{
	return []byte("Exporting to CSV...")
}

func Prob6(){
	j1 := JSONExporter{}
	c1 := CSVExporter{}
	fmt.Println(j1.Export())
	fmt.Println(c1.Export())
}