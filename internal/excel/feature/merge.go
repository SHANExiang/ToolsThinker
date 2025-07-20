package feature

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"tools-thinker/support/file"
)

// Handle
//
//	@Description:
//	@param inputDir
func Handle(inputDir string) {
	sourceDir := "data"          // 数据源目录
	outputExcel := "merged.xlsx" // 输出文件

	excelFiles, err := file.GetFiles(sourceDir, file.XLSX)
	if err != nil {
		panic(err)
	}

	if len(excelFiles) == 0 {
		fmt.Println("没有找到 Excel 文件")
		return
	}

	mergedFile := excelize.NewFile()
	sheet := "Sheet1"
	mergedFile.SetSheetName("Sheet1", sheet)
	rowIndex := 1

	for i, filePath := range excelFiles {
		fmt.Println("读取文件：", filePath)
		f, err := excelize.OpenFile(filePath)
		if err != nil {
			fmt.Printf("无法打开文件 %s: %v\n", filePath, err)
			continue
		}

		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Printf("读取数据失败 %s: %v\n", filePath, err)
			continue
		}

		// 遍历每一行
		for j, row := range rows {
			// 除了第一个文件，其他文件跳过标题行
			if i != 0 && j == 0 {
				continue
			}
			// 写入合并文件
			cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
			if err := mergedFile.SetSheetRow(sheet, cell, &row); err != nil {
				fmt.Printf("写入失败: %v\n", err)
			}
			rowIndex++
		}
		f.Close()
	}

	err = mergedFile.SaveAs(outputExcel)
	if err != nil {
		fmt.Println("保存合并文件失败：", err)
	} else {
		fmt.Println("成功保存文件：", outputExcel)
	}
}
