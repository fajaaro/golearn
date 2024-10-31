package config

type ConstantStruct struct {
	UploadExcelStartFromIndex int
}

var Constant = ConstantStruct{
	UploadExcelStartFromIndex: 3,
}