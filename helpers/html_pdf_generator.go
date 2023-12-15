package helpers

import (
	"bytes"
	"html/template"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gofiber/fiber/v2"
)

type HtmlPdfGenenerator struct {
}

func (gen *HtmlPdfGenenerator) GeneratePDF(htmlTemplate string, data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	// Load HTML template
	tmpl, err := gen.LoadHtmlTemplate(htmlTemplate)
	if err != nil {
		return nil, err
	}
	// Execute the template with the data
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	// Create new PDF generator
	pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}
	// Set global options
	pdfGen.Dpi.Set(100)
	pdfGen.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfGen.Grayscale.Set(true)
	// Create a new input page from HTML content
	page := wkhtmltopdf.NewPageReader(&buf)
	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)
	// Add to the document
	pdfGen.AddPage(page)
	// Create PDF document in the internal buffer
	err = pdfGen.Create()
	if err != nil {
		return nil, err
	}
	// Get the PDF bytes
	pdfBytes := pdfGen.Bytes()
	return pdfBytes, nil
}


func (gen *HtmlPdfGenenerator) LoadHtmlTemplate(filePath string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(filePath)
	return tmpl, err
}


func (gen *HtmlPdfGenenerator) SetOutput(ctx *fiber.Ctx, pdfBytes []byte, fileName string) error {
	ctx.Response().Header.SetContentType("application/pdf")
	ctx.Response().Header.Set("Content-Disposition", "attachment; fileName=" +fileName)
	_, err := ctx.Write(pdfBytes)
	return err
}
