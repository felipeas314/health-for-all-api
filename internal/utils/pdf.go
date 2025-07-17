package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(file multipart.File, size int64) (string, error) {
	// Lê todo conteúdo em memória
	buf, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Cria um reader que implementa ReaderAt e Seeker
	reader := bytes.NewReader(buf)

	pdfReader, err := pdf.NewReader(reader, int64(len(buf)))
	if err != nil {
		return "", err
	}

	var text strings.Builder
	for i := 1; i <= pdfReader.NumPage(); i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		content, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}
		text.WriteString(content)
	}

	return text.String(), nil
}
