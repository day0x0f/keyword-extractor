Keyword Extractor
https://img.shields.io/badge/version-1.0.0-blue
https://img.shields.io/badge/docker-ready-green
https://img.shields.io/badge/CI/CD-enabled-success

Um extractor de keywords automatizado que processa documentos de texto e identifica as palavras-chave mais relevantes usando técnicas de processamento de linguagem natural.

🚀 Funcionalidades
Extração automática de keywords de documentos de texto

Suporte a múltiplos formatos (markdown, txt, etc.)

Processamento em lote de múltiplos arquivos

Containerização Docker para fácil deploy

Pipeline CI/CD automatizado com GitHub Actions

📦 Instalação
Usando Docker (Recomendado)
bash
# Puxar a imagem mais recente
docker pull ghcr.io/day0x0f/keyword-extractor:latest

# Executar em um arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest arquivo.md
Instalação Manual
bash
# Clonar o repositório
git clone https://github.com/day0x0f/keyword-extractor.git
cd keyword-extractor

# Instalar dependências
pip install -r requirements.txt

# Executar
python main.py arquivo.md
🛠️ Uso
Comando Básico
bash
# Processar um único arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md

# Processar múltiplos arquivos
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest *.md

# Salvar output em um arquivo
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md > keywords.txt
Exemplo de Output
text
Arquivo: exemplo.md
Keywords extraídas:
1. processamento linguagem natural (score: 0.95)
2. extração keywords (score: 0.87)
3. análise textual (score: 0.82)
4. machine learning (score: 0.78)
5. python (score: 0.75)
🔧 Desenvolvimento
Estrutura do Projeto
text
keyword-extractor/
├── Dockerfile
├── .github/
│   └── workflows/
│       └── docker.yml
├── src/
│   ├── main.py
│   ├── extractor.py
│   └── utils/
├── requirements.txt
├── examples/
└── README.md
Build Local
bash
# Build da imagem Docker
docker build -t keyword-extractor .

# Executar localmente
docker run --rm -v "${PWD}:/data" keyword-extractor exemplo.md
Testes
bash
# Executar testes unitários
python -m pytest tests/

# Testar com arquivo de exemplo
docker run --rm -v "${PWD}:/data" keyword-extractor examples/sample.md
📋 CI/CD Pipeline
O projeto utiliza GitHub Actions para automatizar:

✅ Build automático da imagem Docker em pushes para main e dev

✅ Push para GitHub Container Registry com múltiplas tags

✅ Cache de builds para otimização de performance

✅ Testes automatizados em cada pull request

Status do Workflow
https://github.com/day0x0f/keyword-extractor/actions/workflows/docker.yml/badge.svg

Tags Disponíveis
latest - Última versão estável

sha-<commit-hash> - Builds específicos por commit

vX.X.X - Versões semânticas

🐛 Troubleshooting
Problemas Comuns
Imagem não encontrada:

bash
# Verificar tags disponíveis
curl -s https://ghcr.io/v2/day0x0f/keyword-extractor/tags/list

# Usar tag específica
docker run --rm -v "${PWD}:/data" ghcr.io/day0x0f/keyword-extractor:sha-243a861 arquivo.md
Permissão negada:

bash
# Dar permissão de execução no host
chmod +x arquivo.md
Arquivo não encontrado:

bash
# Verificar se o arquivo existe no diretório atual
ls -la

# Usar caminho absoluto
docker run --rm -v "/caminho/completo:/data" ghcr.io/day0x0f/keyword-extractor:latest documento.md
🤝 Contribuição
Fork o projeto

Crie uma branch para sua feature (git checkout -b feature/AmazingFeature)

Commit suas mudanças (git commit -m 'Add some AmazingFeature')

Push para a branch (git push origin feature/AmazingFeature)

Abra um Pull Request

Convenções de Commit
feat: Novas funcionalidades

fix: Correções de bugs

docs: Documentação

style: Formatação, missing semi colons, etc.

refactor: Refatoração de código

test: Adicionando testes

chore: Manutenção do build


👥 Autores
Dayvid Dias - day0x0f

Fernando Franca Filho - FernandofrancaFilho


<div align="center">
Feito com ❤️ e muito ☕

</div>
