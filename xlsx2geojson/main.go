package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	geojson "github.com/paulmach/go.geojson"
	"github.com/tealeg/xlsx"
)

var xlsxPath = flag.String("f", "", "Path to an XLSX file")
var sheetIndex = flag.Int("i", 0, "Index of sheet to convert, zero based")
var delimiter = flag.String("d", ";", "Delimiter to use between fields")

type gruppe struct {
	Stadtverband    string `xlsx:"0"`
	Gruppenname     string `xlsx:"1"`
	Strasse         string `xlsx:"2"`
	Plz             string `xlsx:"3"`
	Ort             string `xlsx:"4"`
	Alter           string `xlsx:"5"`
	Treffpunkt      string `xlsx:"6"`
	Zeit            string `xlsx:"7"`
	Webseite        string `xlsx:"8"`
	Ansprechpartner string `xlsx:"9"`
	Email           string `xlsx:"10"`
	Telefon         string `xlsx:"11"`
}

func generateCSVFromXLSXFile(excelFileName string, sheetIndex int) error {
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
		return error
	}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return errors.New("This XLSX file contains no sheets")
	case sheetIndex >= sheetLen:
		return fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d", sheetIndex, sheetLen-1)
	}
	sheet := xlFile.Sheets[sheetIndex]
	g := gruppe{}
	var gruppen []gruppe
	for i, row := range sheet.Rows {
		if i != 0 {
			if row != nil {
				row.ReadStruct(&g)
				gruppen = append(gruppen, g)
				//fmt.Printf("%+v\n", g)
			}
		}
	}
	featureCollection := geojson.NewFeatureCollection()
	for _, t := range gruppen {
		//feature := geojson.NewPointFeature([]float64{t.lat, t.lon})
		feature := geojson.NewPointFeature([]float64{0, 0})
		e := reflect.ValueOf(&t).Elem()
		for i := 0; i < e.NumField(); i++ {
			key := e.Type().Field(i).Name
			value := e.Field(i).Interface()
			if key != string("lat") && key != string("lon") {
				feature.SetProperty(strings.ToLower(key), value)
			}
		}
		featureCollection.AddFeature(feature)
	}
	s, _ := featureCollection.MarshalJSON()
	fmt.Println(string(s))
	return nil
}

func main() {
	flag.Parse()
	if len(os.Args) < 3 {
		flag.PrintDefaults()
		return
	}
	flag.Parse()
	if err := generateCSVFromXLSXFile(*xlsxPath, *sheetIndex); err != nil {
		fmt.Println(err)
	}
}
