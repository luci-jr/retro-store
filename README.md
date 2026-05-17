<div align="center">

<img src="https://raw.githubusercontent.com/marcusolsson/gophers/master/pixel_gopher.png" width="100" alt="Pixel Gopher">

# 🕹️ R E T R O   S T O R E 🕹️
**[ WEB APP EDITION ]**

<br>
<p><i>>> P R E S S   S T A R T   T O   P L A Y <<</i></p>

---

</div>

```text
🎮 [ B O O T _ S E Q U E N C E ]
> INICIANDO_SISTEMA.................... [OK]
> CARREGANDO_MODULOS_GO............... [OK]
> VERIFICANDO_ROTAS_HTTP.............. [OK]
> CONECTANDO_POSTGRESQL............... [OK]
> MODO_RETRO_8BIT_ATIVADO............. [OK]
> SISTEMA_PRONTO.
```

<br>

## 👾 [ M I S S Ã O _ D O _ P R O J E T O ]

A **Retro Store** é uma aplicação web desenvolvida em **Go (Golang)** com foco em aprendizado prático de:
- Arquitetura MVC
- CRUD completo
- Manipulação de banco de dados
- Rotas HTTP
- Templates HTML
- Estruturação backend sem frameworks

Toda a interface foi criada com inspiração em:
- 🎮 Arcades dos anos 80
- 🕹️ Consoles retrô
- 💾 Terminais antigos
- 👾 Pixel Art + Neon UI

<br>

## 🧩 [ E N G I N E _ D O _ S I S T E M A ]

```text
╔══════════════════════════════════╗
║        CORE TECHNOLOGIES         ║
╚══════════════════════════════════╝
```

| ⚡ CAMADA | 🎯 TECNOLOGIA |
| :--- | :--- |
| 👨‍💻 Linguagem | Go (Golang) |
| 🌐 Frontend | HTML5 + CSS3 |
| 🎨 Interface | Retro Pixel UI |
| 🗄️ Banco | PostgreSQL |
| 📦 Container | Docker |
| 🔌 Driver SQL | lib/pq |
| 🚨 Alertas | SweetAlert2 |

<br>

## 🏛️ [ A R Q U I T E T U R A _ M V C ]

```text
  [ PLAYER ]
      ↓
[ CONTROLLERS ]
      ↓
  [ MODELS ]
      ↓
[ POSTGRESQL ]
```

### 📂 Estrutura do Projeto
```text
📦 retro-store
 ┣ 📂 controllers
 ┃ ┗ 📜 produtosController.go
 ┣ 📂 models
 ┃ ┗ 📜 produtosModel.go
 ┣ 📂 routes
 ┃ ┗ 📜 routes.go
 ┣ 📂 templates
 ┃ ┣ 📜 _head.html
 ┃ ┣ 📜 _menu.html
 ┃ ┣ 📜 index.html
 ┃ ┣ 📜 produto.html
 ┃ ┗ 📜 edit.html
 ┣ 📂 static
 ┃ ┗ 📂 css
 ┃   ┗ 📜 retro.css
 ┣ 📜 main.go
 ┗ 📜 go.mod
```

<br>

## 🛠️ [ M Ó D U L O S _ U T I L I Z A D O S ]

*   **⚙️ `net/http`**: Responsável por criar o servidor web, gerenciar rotas, escutar requisições HTTP e responder páginas ao navegador.
*   **📁 `http.FileServer`**: Usado para servir arquivos estáticos como `/static/` e `/images/`, permitindo carregar CSS, Imagens e Assets visuais.
*   **🧠 `html/template`**: Sistema de renderização SSR (Server Side Rendering). Renderiza páginas HTML dinamicamente pelo backend Go.
*   **🗄️ `database/sql`**: Camada de comunicação com o PostgreSQL. Funções utilizadas: `db.Query()`, `db.Exec()`, `db.Prepare()`.
*   **🔢 `strconv`**: Conversão de tipos recebidos via formulários HTML (exemplo: `strconv.Atoi(id)` e `strconv.ParseFloat(preco, 64)`).

<br>

## 🐘 [ B A N C O _ D E _ D A D O S ]

### 🐳 PostgreSQL + Docker
O banco foi executado em container Docker para:
- ✅ Ambiente isolado
- ✅ Fácil replicação
- ✅ Configuração rápida
- ✅ Padronização do ambiente

### 🔒 Segurança SQL
A aplicação utiliza a função `.Prepare()` para criar *Prepared Statements*. Isso evita de forma extremamente eficaz ataques de **SQL Injection**.

<br>

## 📦 [ I N S T A L A Ç Ã O ]

**🎯 1. Clone o projeto**
```bash
$ git clone <repo-url>
```

**🎯 2. Entre na pasta**
```bash
$ cd retro-store
```

**🎯 3. Instale dependências**
```bash
$ go mod tidy
```

**🎯 4. Execute o servidor**
```bash
$ go run main.go
```

**🎯 5. Abra no navegador**  
🔗 [http://localhost:8000](http://localhost:8000)

<br>

## 🕹️ [ G A M E P L A Y ]

### ✅ Funcionalidades (Features)
- [x] Criar produtos
- [x] Listar produtos
- [x] Editar produtos
- [x] Remover produtos
- [x] Interface retro 8-bit
- [x] Integração PostgreSQL
- [x] Arquitetura MVC

<br>

## 🚀 [ O B J E T I V O _ D O _ P R O J E T O ]

Este projeto foi criado para consolidar conhecimentos em Backend com Go, Estrutura MVC, CRUD completo, Banco relacional, Docker, Renderização server-side, Organização de código e HTTP na prática.
Tudo isso utilizando principalmente a biblioteca padrão do Go, entendendo como a web funciona **sem abstrações pesadas**.

<br>
<div align="center">

```text
█▀▀ ▄▀█ █▀▄▀█ █▀▀   █▀█ █░█ █▀▀ █▀█
█▄█ █▀█ █░▀░█ ██▄   █▄█ ▀▄▀ ██▄ █▀▄
```
<p><b>>> GAME OVER? NEGATIVO.</b></p>
<p><b>>> NEXT LEVEL UNLOCKED...</b></p>

<img src="https://media.giphy.com/media/13HgwGsXF0aiGY/giphy.gif" width="220" alt="Pixel Art GIF"/> 

</div>