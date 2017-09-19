package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize"
)

//Constants
const fileName = "/home/hwe/testExcel.xlsx"
const myCol = 4 //Number of column

func main() {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")

	for _, row := range rows {

		for nrCol, colCell := range row {
			if nrCol == myCol-1 {

				fmt.Println(colCell, "\t", readFile(colCell))

			}

		}

	}

	os.Exit(1)

}

func readFile(file string) bool {

	obj, err := os.Open(file)
	if err != nil {
		b := false
		return b
	}
	defer obj.Close()

	b := true
	return b
}
