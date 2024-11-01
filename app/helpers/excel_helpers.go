package helpers

import (
	"learn/config"
	"mime/multipart"
	"reflect"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadExcel(file *multipart.FileHeader) ([][]string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	xlsx, err := excelize.OpenReader(fileContent)
	if err != nil {
		return nil, err
	}

	rows := xlsx.GetRows("Sheet1")

	fileContent.Close()

	return rows, nil
}

func GetExcelRowsData(rows [][]string) [][]string {
	var rowsData [][]string
	if len(rows) > config.Constant.UploadExcelStartFromIndex {
		rowsData = rows[config.Constant.UploadExcelStartFromIndex:]
	} else {
		rowsData = [][]string{}
	} 

	return rowsData
}

func ExtractModelExcelColIndexes(structType interface{}) map[string]int {
	result := make(map[string]int)
	val := reflect.TypeOf(structType)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		excelTag := field.Tag.Get("excel-col-index")
		if excelTag != "" {
			if index, err := strconv.Atoi(excelTag); err == nil {
				result[field.Name] = index
			}
		}
	}
	return result
}