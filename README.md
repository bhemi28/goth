# GOTH - Go, Templ, HTMX Portfolio

A retro-futuristic, Y2K-inspired portfolio and blog built with the **GOTH** stack (Go, Templ, HTMX) and Tailwind CSS.

## 🌐 Live Demo

Check out the live version hosted on Railway:
**[https://goth-production.up.railway.app/](https://goth-production.up.railway.app/)**

## ✨ Features

*   **Retro Y2K Design**: Custom Tailwind configuration for a nostalgic desktop interface look.
*   **Server-Side Rendering**: Fast and SEO-friendly pages using [Templ](https://templ.guide/).
*   **SPA-like Experience**: Smooth transitions and partial page updates powered by [HTMX](https://htmx.org/).
*   **Markdown Blog**: Renders blog posts from markdown files with frontmatter support.
*   **Responsive Layout**: Fully responsive design that works on mobile and desktop.
*   **Contact Form**: Functional contact form integrated with email delivery (SMTP).
*   **Dockerized**: Ready for deployment with a multi-stage Dockerfile.

## 🛠️ Tech Stack

*   **Language**: [Go (Golang)](https://go.dev/)
*   **Templating**: [Templ](https://templ.guide/)
*   **Interactivity**: [HTMX](https://htmx.org/)
*   **Styling**: [Tailwind CSS](https://tailwindcss.com/)
*   **Routing**: [Chi Router](https://github.com/go-chi/chi)
*   **Markdown**: [Goldmark](https://github.com/yuin/goldmark)

## 🚀 Getting Started

### Prerequisites

*   Go 1.21+
*   Make (optional, for using Makefile)
*   [Templ CLI](https://templ.guide/quick-start/installation) (`go install github.com/a-h/templ/cmd/templ@latest`)

### Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/yourusername/goth.git
    cd goth
    ```

2.  **Install dependencies**
    ```bash
    go mod download
    ```

3.  **Environment Setup**
    Copy the example environment file and configure your secrets (if needed for email).
    ```bash
    cp .env.example .env
    ```

### Running Locally

Use the Makefile for convenience:

*   **Generate Templates & Run**:
    ```bash
    make dev
    ```
    This will compile the `.templ` files and start the server on `http://localhost:8080`.

*   **Just Generate Templates**:
    ```bash
    make generate
    ```

*   **Build Binary**:
    ```bash
    make build
    ```

### Docker

Build and run the container:

```bash
docker build -t goth-portfolio .
docker run -p 8080:8080 goth-portfolio
```

## 📂 Project Structure

```
.
├── main.go           # Entry point
├── go.mod            # Go dependencies
├── Makefile          # Build commands
├── Dockerfile        # Docker configuration
├── views/            # Templ templates & Go view logic
│   ├── components/   # Reusable UI components (Header, Sidebar, etc.)
│   ├── index.templ   # Main layout
│   └── ...           # Page templates
├── static/           # Static assets (CSS, JS, Images)
└── posts/            # Markdown blog posts
```

## 📝 License

This project is open source and available under the [MIT License](LICENSE).
