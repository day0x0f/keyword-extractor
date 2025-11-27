package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"text-stats/handlers"
)

// FileExists verifica se um arquivo existe no sistema de arquivos.
// Retorna true caso o arquivo exista, ou false se não existir.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// GenerateOutputFilename gera automaticamente o nome do arquivo de saída.
// O nome é baseado no arquivo de entrada + um timestamp para evitar sobrescrições.
func GenerateOutputFilename(inputPath string) string {
	base := filepath.Base(inputPath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	timestamp := time.Now().Format("20060102-150405")
	return fmt.Sprintf("%s_stats_%s.txt", name, timestamp)
}

// SaveStatsToFile recebe as estatísticas de texto e salva tudo em um arquivo.
// Cria o arquivo de relatório com informações formatadas, incluindo dados sobre
// tamanho, contagem de palavras, linhas, caracteres e detalhes de Markdown.
func SaveStatsToFile(stats *handlers.TextStats, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf(`RELATÓRIO DE ESTATÍSTICAS DE TEXTO
===============================

Arquivo: %s
Data: %s

ESTATÍSTICAS:
------------
Tamanho do arquivo: %d bytes
Número de linhas: %d
Número de palavras: %d
Número de caracteres: %d
Palavra mais longa: %d caracteres

`,
		stats.Filename,
		time.Now().Format("02/01/2006 15:04:05"),
		stats.FileSize,
		stats.LineCount,
		stats.WordCount,
		stats.CharCount,
		stats.LongestWord,
	)

	if stats.IsMarkdown {
		content += fmt.Sprintf("Arquivo Markdown: Sim\nElementos Markdown: %d\n", stats.MarkdownElements)
	}

	content += "\n---\nGerado por Text Stats\n"

	_, err = file.WriteString(content)
	return err
}
