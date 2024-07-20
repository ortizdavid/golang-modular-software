package helpers

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tealeg/xlsx"
)


type ExcelGenerator struct {
	File  		*xlsx.File
	Sheet 		*xlsx.Sheet
	HeaderColor	string
}


func NewExcelGenerator() *ExcelGenerator {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	return &ExcelGenerator{
		File:        file,
		Sheet:       sheet,
		HeaderColor: "CC9F26",
	}
}


func (eg *ExcelGenerator) AddTitle(title string) {
	row := eg.Sheet.AddRow()
	cell := row.AddCell()
	cell.Value = title
	titleStyle := xlsx.NewStyle()
	titleStyle.Font = *xlsx.NewFont(14, "Arial")
	titleStyle.Font.Bold = true
	cell.SetStyle(titleStyle)
}


func (eg *ExcelGenerator) AddHeaderRow(columns ...string) {
	row := eg.Sheet.AddRow()
	for _, col := range columns {
		cell := row.AddCell()
		cell.Value = col
		cell.SetStyle(eg.getHeaderBorderStyle())
	}
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", eg.HeaderColor, eg.HeaderColor) 
	style.Fill = fill
	for _, cell := range row.Cells {
		cell.SetStyle(style)
	}
}


func (eg *ExcelGenerator) AddDataRow(data ...interface{}) {
	row := eg.Sheet.AddRow()
	for idx, d := range data {
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%v", d)
		cell.SetStyle(eg.getCellBorderStyle())

		textLength := len(fmt.Sprintf("%v", d))
		eg.adjustColumnWidth(idx, textLength)
	}
}


// getHeaderBorderStyle returns a style with header cell borders.
func (eg *ExcelGenerator) getHeaderBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}


// getCellBorderStyle returns a style with data cell borders.
func (eg *ExcelGenerator) getCellBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}

// getTitleBorderStyle returns a style with title cell borders.
func (eg *ExcelGenerator) getTitleBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}


// adjustColumnWidth adjusts the column width based on the text length.
func (eg *ExcelGenerator) adjustColumnWidth(colIndex int, textLength int) {
	// Calculate column width based on text length
	colWidth := float64(textLength) * 1.2 // Adjust this multiplier based on your requirements

	// Set the column width
	eg.Sheet.Col(colIndex).Width = colWidth
}


// getCellBackgroundColorStyle returns a style with a specific background color.
func (eg *ExcelGenerator) getCellBackgroundColorStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	style.Fill = *xlsx.NewFill("solid", "FF0000", "FF0000")
	return style
}

func (eg *ExcelGenerator) SaveToFile(ctx *fiber.Ctx, fileName string) error {
	var buf bytes.Buffer

	// Write the Excel data to the buffer
	err := eg.File.Write(&buf)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error writing Excel data")
	}

	// Set response headers for Excel download
	ctx.Response().Header.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Response().Header.Set("Content-Disposition", "attachment; filename="+fileName+"")

	// Send the Excel data to the client
	_, err = ctx.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}