package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

const version = "1.0.0"

func main() {
	checkVersionFlag()

	inputPath, outputPath := parseArgs()

	mdContent, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading input file %s: %v\n", inputPath, err)
		os.Exit(1)
	}

	pdf := initPDF()

	err = renderMarkdownToPDF(mdContent, pdf)
	if err != nil {
		fmt.Printf("Error processing Markdown content: %v\n", err)
		os.Exit(1)
	}

	err = pdf.OutputFileAndClose(outputPath)
	if err != nil {
		fmt.Printf("Error generating PDF file %s: %v\n", outputPath, err)
		os.Exit(1)
	}

	fmt.Printf("Success! PDF generated at: %s\n", outputPath)
}

// checkVersionFlag checks for -v/--version before any other argument
// validation, so it works regardless of how many other arguments are given.
func checkVersionFlag() {
	if len(os.Args) < 2 {
		return
	}
	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		fmt.Printf("butterfly version %s\n", version)
		os.Exit(0)
	}
}

// parseArgs validates CLI arguments and formats the output filename
func parseArgs() (string, string) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Insufficient arguments.")
		fmt.Println("Usage: butterfly <input_file.md> <output_file.pdf>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	if !strings.HasSuffix(strings.ToLower(outputPath), ".pdf") {
		outputPath += ".pdf"
	}

	return inputPath, outputPath
}

// initPDF configures the PDF document and sets up the watermark footer
func initPDF() *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.SetFont("Arial", "I", 8)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(0, 10, "butterfly", "", 0, "L", false, 0, "")
		pdf.SetTextColor(0, 0, 0)
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	return pdf
}

// renderMarkdownToPDF parses the MD syntax tree and writes components into the PDF
func renderMarkdownToPDF(mdContent []byte, pdf *fpdf.Fpdf) error {
	translator := pdf.UnicodeTranslatorFromDescriptor("")

	parser := goldmark.New()
	reader := text.NewReader(mdContent)
	doc := parser.Parser().Parse(reader)

	return ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch node.Kind() {
			case ast.KindHeading:
				heading := node.(*ast.Heading)
				switch heading.Level {
				case 1:
					pdf.SetFont("Arial", "B", 20)
					pdf.Ln(8)
				case 2:
					pdf.SetFont("Arial", "B", 16)
					pdf.Ln(6)
				default:
					pdf.SetFont("Arial", "B", 14)
					pdf.Ln(4)
				}
			case ast.KindParagraph:
				pdf.SetFont("Arial", "", 12)
				pdf.Ln(4)
			case ast.KindText:
				textNode := node.(*ast.Text)
				content := string(textNode.Value(mdContent))
				content = strings.ReplaceAll(content, "\n", " ")
				pdf.Write(6, translator(content))
			}
		} else {
			switch node.Kind() {
			case ast.KindHeading:
				pdf.Ln(6)
			case ast.KindParagraph:
				pdf.Ln(4)
			}
		}
		return ast.WalkContinue, nil
	})
}
