package part4

import "fmt"

func performExport(exporter DataExporter) {
	data := string(exporter.Export())
	fmt.Println(data)
}

func Prob7(){
	j2 := JSONExporter{}
	c2 := CSVExporter{}

	performExport(j2)
	performExport(c2)
}