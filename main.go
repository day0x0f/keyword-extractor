package main

import (
	"fmt"
	"log"
	"os"

	"text-stats/handlers"
	"text-stats/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: text-stats <arquivo>")
		return
	}

	filePath := os.Args[1]

	if !utils.FileExists(filePath) {
		log.Fatalf("Arquivo não encontrado: %s", filePath)
	}

	stats, err := handlers.AnalyzeTextFile(filePath)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	// Exibe resultados
	fmt.Printf("\n📊 ESTATÍSTICAS DO TEXTO\n")
	fmt.Printf("=======================\n")
	fmt.Printf("📄 Arquivo: %s\n", stats.Filename)
	fmt.Printf("📦 Tamanho: %d bytes\n", stats.FileSize)
	fmt.Printf("📝 Linhas: %d\n", stats.LineCount)
	fmt.Printf("🔤 Palavras: %d\n", stats.WordCount)
	fmt.Printf("🔡 Caracteres: %d\n", stats.CharCount)
	fmt.Printf("📏 Palavra mais longa: %d caracteres\n", stats.LongestWord)

	if stats.IsMarkdown {
		fmt.Printf("🗃️ Markdown: Sim (%d elementos)\n", stats.MarkdownElements)
	}

	// Salva relatório
	outputFile := utils.GenerateOutputFilename(filePath)
	if err := utils.SaveStatsToFile(stats, outputFile); err != nil {
		log.Printf("Não foi salvar relatório: %v", err)
	} else {
		fmt.Printf("💾 Relatório salvo com sucesso: %s\n", outputFile)
	}
}
