package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string
	City  string
	State string
	// in order for code outside of the models package to use the struct, the fields must be exportable, so capital letters
}

type ZipSlice []*Zip

type ZipIndex map[string]ZipSlice

func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil { //something went wrong
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read() //_ ignores the value returned in golang w/o compile errors
	if err != nil {
		return nil, fmt.Errorf("error reading header row %v", err)
	}

	zips := make(ZipSlice, 0, 43000) //pre-allocating array to be 43000 long, much more efficient
	for {
		fields, err := reader.Read() //reads one line of csv file at a time, splitting row into separate string fields
		if err == io.EOF {           //signal that reader is done with the input string, so we know we're done
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)

	}
}
