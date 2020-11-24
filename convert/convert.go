package convert

import (
	"KindleHighlightsReader/reader"
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"log"
	"reflect"
	"strings"
)

//Todo* Return error or bool with the bytes?

//TODO 1) Remove BOM and /r characters
func ToJSON(highlights []reader.Highlight) []byte {
	b, err := json.Marshal(highlights)
	if err != nil {
		log.Println(err)
	}
	return b
}

func ToText(highlights []reader.Highlight) []byte {
	layout := "%s\n\n%s, %s\n________________________________\n\n"
	sb := strings.Builder{}
	for _, v := range highlights {
		highlight := fmt.Sprintf(layout, v.Text, v.Author, v.Title)
		sb.WriteString(highlight)
	}
	b := []byte(sb.String())
	return b
}

func ToCSV(highlights []reader.Highlight) []byte {
	var headers []string
	h := highlights[0]
	v := reflect.ValueOf(&h).Elem()
	typeOf := v.Type()
	for i := 0; i < v.NumField(); i++ {
		header := typeOf.Field(i).Name
		headers = append(headers, header)
	}

	var b bytes.Buffer
	writer := csv.NewWriter(&b)
	if err := writer.Write(headers); err != nil {
		log.Println("Failed to write csv headers", err)
	}

	for _, v := range highlights {
		row := []string{v.Text, v.Author, v.Title}
		if err := writer.Write(row); err != nil {
			log.Println("Failed to write csv values:", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("Failed to flush csv writer,", err)
		return nil
	}
	return b.Bytes()
}

func ToPDF(highlights []reader.Highlight) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 14)
	layout := "%s, %s\n\n%s\n\n\n"
	for i, v := range highlights {
		if i%3 == 0 {
			pdf.AddPage()
		}

		highlight := fmt.Sprintf(layout, v.Author, v.Title, v.Text)
		highlight = strings.TrimRight(highlight, "\r\n")
		pdf.MultiCell(0, 10, highlight, "0", "0", false)
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if err := pdf.Output(w); err != nil {
		log.Println("Failed to write PDF")
	}
	return b.Bytes()
}

func removeBOM(text string) string {
	return ""
}

func removeNewLines(text string) string {
	return ""
}

func removeCarriageReturns(text string) string {
	return ""
}
