package main

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"strconv"
)

type excelProcess struct {
	Order       int64  `gorm:"not null" json:"order" binding:"required"`
	ProcessID   string `gorm:"not null;type:varchar(40)" json:"process_id" binding:"required"` //流程id 自己生成
	ProcessName string `gorm:"not null" json:"process_name" binding:"required"`
	ProcessType int64  `gorm:"not null" json:"process_type" binding:"required"`
	MicHand     int64  `json:"mic_hand" binding:"-"` //可选
	MicEar      int64  `json:"mic_ear" binding:"-"`  //可选
	Remark      string `json:"remark" binding:"-"`   //可选
}

func (s *Service) GetExcel(c *gin.Context) (int, interface{}) {
	excelFileHeader, err := c.FormFile("excel")
	if err != nil {
		return s.makeErrJSON(404, 40400, err.Error())
	}
	if excelFileHeader == nil {
		return s.makeErrJSON(404, 40401, "none file")
	}
	excelFile, err := excelFileHeader.Open()
	excel, err := excelize.OpenReader(excelFile)
	if err != nil {
		return s.makeErrJSON(500, 50000, err.Error())
	}
	var processes []*excelProcess
	sheetName := excel.GetSheetName(1)
	rows := excel.GetRows(sheetName)
	for k, row := range rows {
		//第一行不做处理
		if k == 0 {
			continue
		}
		process := new(excelProcess)
		process.Order = int64(k)

		processType, err := strconv.Atoi(row[0])
		process.ProcessType = int64(processType)
		if err != nil {
			process.ProcessType = ProcessTypeMap[row[0]]
		}
		if process.ProcessType < 0 || process.ProcessType > 5 {
			return s.makeErrJSON(500, 50001, errors.New("A"+string('1'+k)+" error"))
		}

		process.ProcessName = row[1]

		micHand, err := strconv.Atoi(row[2])
		if err != nil {
			return s.makeErrJSON(500, 50002, errors.New("C"+string('1'+k)+" error"))
		}
		process.MicHand = int64(micHand)

		micEar, err := strconv.Atoi(row[3])
		if err != nil {
			return s.makeErrJSON(500, 50003, errors.New("D"+string('1'+k)+" error"))
		}
		process.MicEar = int64(micEar)

		process.Remark = row[4]

		processes = append(processes, process)
	}

	return s.makeSuccessJSON(processes)
}
