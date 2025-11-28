Um extractor de keywords automatizado que processa documentos de texto e identifica as palavras-chave mais relevantes usando técnicas de processamento de linguagem natural.

🚀 Funcionalidades
Extração automática de keywords de documentos de texto
Suporte a múltiplos formatos (markdown, txt, etc.)
Processamento em lote de múltiplos arquivos
Containerização Docker para fácil deploy
Pipeline CI/CD automatizado com GitHub Actions

📦 Instalação
Usando Docker (Recomendado)

# Puxar a imagem mais recente
docker pull ghcr.io/day0x0f/keyword-extractor:latest

# Executar em um arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest arquivo.md

# Clonar o repositório
git clone https://github.com/day0x0f/keyword-extractor.git
cd keyword-extractor

🛠️ Uso
Comando Básico

# Processar um único arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md

# Processar múltiplos arquivos
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest *.md

# Salvar output em um arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md > keywords.txt

# Build da imagem Docker
docker build -t keyword-extractor .

# Executar localmente
docker run --rm -v "${PWD}:/data" keyword-extractor exemplo.md
Testes

# Testar com arquivo de exemplo
docker run --rm -v "${PWD}:/data" keyword-extractor examples/sample.md

📋 CI/CD Pipeline
O projeto utiliza GitHub Actions para automatizar:

✅ Build automático da imagem Docker em pushes para main e dev

✅ Push para GitHub Container Registry com múltiplas tags

✅ Cache de builds para otimização de performance

✅ Testes automatizados em cada pull request


Tags Disponíveis
v1.0.0 - Primeira versão

bash
# Verificar tags disponíveis
curl -s https://ghcr.io/v2/day0x0f/keyword-extractor/tags/list

bash
# Dar permissão de execução no host
chmod +x arquivo.md
Arquivo não encontrado:

bash
# Verificar se o arquivo existe no diretório atual
ls -la

# Usar caminho absoluto
docker run --rm -v "/caminho/completo:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md


👥 Autores
Dayvid Dias - day0x0f
Fernando Franca Filho - FernandofrancaFilho


<div align="center">
Feito com ❤️ e muito ☕

</div>
