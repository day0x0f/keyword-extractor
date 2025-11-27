package handlers

import (
	"os"
	"strings"
	"unicode"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/text/unicode/norm"
)



type TextStats struct {
	Filename         string
	FileSize         int64
	LineCount        int
	WordCount        int
	CharCount        int
	LongestWord      int
	IsMarkdown       bool
	MarkdownElements int
}

func AnalyzeTextFile(filePath string) (*TextStats, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	stats := &TextStats{
		Filename: filePath,
		FileSize: int64(len(content)),
	}

	text := string(content)

	// Normaliza texto
	text = norm.NFC.String(text)

	// Analisa estatísticas básicas
	analyzeBasicStats(stats, text)

	// Verifica se é Markdown e analisa
	if isMarkdownFile(filePath) {
		analyzeMarkdown(stats, text)
	}

	return stats, nil
}

func analyzeBasicStats(stats *TextStats, text string) {
	lines := strings.Split(text, "\n")
	stats.LineCount = len(lines)

	words := strings.Fields(text)
	stats.WordCount = len(words)
	stats.CharCount = len(text)

	// Encontra palavra mais longa
	for _, word := range words {
		cleanWord := strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		if len(cleanWord) > stats.LongestWord {
			stats.LongestWord = len(cleanWord)
		}
	}
}

func analyzeMarkdown(stats *TextStats, text string) {
	stats.IsMarkdown = true

	// Converte Markdown para HTML
	html := blackfriday.Run([]byte(text))

	// Sanitiza HTML e conta elementos
	sanitized := bluemonday.UGCPolicy().SanitizeBytes(html)
	stats.MarkdownElements = strings.Count(string(sanitized), "<") / 2
}

func isMarkdownFile(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".md") ||
		strings.HasSuffix(strings.ToLower(filename), ".markdown")
}
