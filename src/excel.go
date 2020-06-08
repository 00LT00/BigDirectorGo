package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"strconv"
	"unicode"
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

		if _, ok := ProcessTypeMap[row[0]]; !ok {
			return s.makeErrJSON(500, 50001, "A"+strconv.Itoa(k+1)+" 填写错误，请正确填写")
		}
		processType, err := strconv.Atoi(row[0])
		process.ProcessType = int64(processType)
		if err != nil {
			process.ProcessType = ProcessTypeMap[row[0]]
		}
		if process.ProcessType < 0 || process.ProcessType > 5 {
			return s.makeErrJSON(500, 50002, "A"+strconv.Itoa(k+1)+" 填写错误，请正确填写")
		}

		name := row[1]
		chineseCount := 0
		for k, v := range name {
			if unicode.Is(unicode.Han, v) {
				chineseCount++
				fmt.Println(string(name[k : k+3]))
				fmt.Println(k)
			}
		}

		fmt.Println(len(name) - chineseCount*3 + chineseCount)
		if len(name)-chineseCount*3+chineseCount <= 15 {
			process.ProcessName = name
		} else {
			return s.makeErrJSON(500, 50006, "B"+strconv.Itoa(k+1)+"过长，请勿超过15个字符(中英等长)")
		}

		if row[2] == "" {
			process.MicHand = 0
		} else {
			micHand, err := strconv.Atoi(row[2])
			if err != nil {
				return s.makeErrJSON(500, 50003, "C"+strconv.Itoa(k+1)+" 不是数字，请正确填写")
			}
			process.MicHand = int64(micHand)
		}

		if row[3] == "" {
			process.MicEar = 0
		} else {
			micEar, err := strconv.Atoi(row[3])
			if err != nil {
				return s.makeErrJSON(500, 50004, "D"+strconv.Itoa(k+1)+" 不是数字，请正确填写")
			}
			process.MicEar = int64(micEar)
		}

		if len(row[4]) > 100 {
			return s.makeErrJSON(500, 50005, "E"+strconv.Itoa(k+1)+" 备注过长，请勿超过100字")
		}
		process.Remark = row[4]

		processes = append(processes, process)
	}

	return s.makeSuccessJSON(processes)
}
