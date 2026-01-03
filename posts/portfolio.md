---
title: "Building My Portfolio: A Retro-Futuristic Journey with Go, Templ & HTMX"
desc: "How I built a Y2K-inspired portfolio website using Go, Templ, HTMX, and Tailwind CSS - and why I chose this stack."
date: 2026-01-03
author: Meet Bhesaniya
---

# Building My Portfolio: A Retro-Futuristic Journey

Welcome to the behind-the-scenes of this very website you're browsing! In this post, I'll walk you through the decisions, challenges, and learnings from building my personal portfolio with a unique retro-futuristic aesthetic.

## The Vision: Dreamscape OS

I wanted something different. Not another minimal portfolio, not another dark mode developer site. I wanted something that felt like **booting up a nostalgic operating system** - think Windows 95 meets Y2K aesthetics, but built with modern tech.

The result? A "Dreamscape OS" theme with:
- **Pastel color palette** (pinks, lavenders, mints, and yellows)
- **Hard shadows** that make elements look like physical stickers
- **Pixel fonts** (VT323) mixed with clean monospace (Space Mono)
- **Window chrome** that mimics old desktop applications

## The Tech Stack

### Why Go?

As a backend developer transitioning toward Golang-centric development, building my portfolio in Go was a no-brainer. It's:
- **Fast** - Compiles to a single binary
- **Simple** - No complex runtime or dependencies
- **Deployable anywhere** - Docker, bare metal, you name it

```go
func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)
    
    views.RegisterRoutes(router)
    
    log.Println("Starting server on :8080")
    http.ListenAndServe(":8080", router)
}
```

### Why Templ?

[Templ](https://templ.guide) is a game-changer for Go web development. It's a templating language that:
- Compiles to Go code (type-safe!)
- Has excellent IDE support
- Feels natural if you know JSX/React

```go
templ Home() {
    <div class="container">
        <h1>Hello World!</h1>
        <p>I'm a { "Backend Developer" }</p>
    </div>
}
```

### Why HTMX?

I wanted an SPA-like experience without shipping megabytes of JavaScript. [HTMX](https://htmx.org) lets you:
- Make AJAX requests with HTML attributes
- Swap content dynamically
- Keep the server in control

```html
<button 
    hx-get="/posts" 
    hx-target="#content" 
    hx-swap="innerHTML">
    Load Posts
</button>
```

The entire site feels like a single-page app, but there's **zero JavaScript framework**. Just 14KB of HTMX.

### Tailwind CSS (via CDN)

For rapid styling, Tailwind was perfect. The utility-first approach meant I could iterate on the design quickly without context-switching to CSS files.

## Challenges & Solutions

### Challenge 1: The Sidebar Duplication Bug

When clicking navigation links, the sidebar would duplicate! The issue? HTMX was replacing the entire layout instead of just the content area.

**Solution:** Detect `HX-Request` header and return partial HTML:

```go
func renderWithLayout(w http.ResponseWriter, r *http.Request, content templ.Component) {
    if r.Header.Get("HX-Request") == "true" {
        // HTMX request - return only the content
        content.Render(r.Context(), w)
    } else {
        // Full page load - return with layout
        Index(content).Render(r.Context(), w)
    }
}
```

### Challenge 2: Loading States

The contact form needed a loading spinner, but HTMX's built-in indicators weren't playing nice with Tailwind's CDN.

**Solution:** Custom CSS classes triggered by HTMX events:

```css
#submit-btn.htmx-request .btn-text,
#submit-btn.htmx-request .btn-icon {
    visibility: hidden !important;
}

#submit-btn.htmx-request .btn-loader {
    display: flex !important;
}
```

### Challenge 3: Code Block Styling

Goldmark renders markdown to HTML, but code blocks looked terrible with the default styles.

**Solution:** Custom CSS targeting the `.markdown-content` wrapper:

```css
.markdown-content pre {
    background-color: #1e1b4b; /* Dark Indigo */
    color: #e0e7ff; /* Lavender */
    border: 2px solid #000;
    box-shadow: 6px 6px 0px 0px #ffade3; /* Pink shadow */
}
```

## The Contact Form: Real Email Delivery

I implemented actual email delivery using Go's `net/smtp` package with Gmail's SMTP server:

```go
auth := smtp.PlainAuth("", smtpEmail, smtpPassword, "smtp.gmail.com")

msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
    from, to, subject, body)

smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
```

## Deployment: Docker + Railway

The site is containerized with a multi-stage Dockerfile:

```dockerfile
# Build stage
FROM golang:1.24-bookworm AS build
RUN go install github.com/a-h/templ/cmd/templ@latest
COPY . .
RUN templ generate
RUN CGO_ENABLED=0 go build -o server .

# Runtime stage  
FROM gcr.io/distroless/base-debian12
COPY --from=build /app/server /app/server
ENTRYPOINT ["/app/server"]
```

Deployed on Railway with environment variables for SMTP credentials. No secrets in the image!

## Lessons Learned

1. **HTMX is powerful** - You can build modern UX without complex JS frameworks
2. **Templ > html/template** - Type safety and IDE support matter
3. **Design systems help** - Having a consistent "Dreamscape OS" theme made decisions easier
4. **Ship it** - Perfect is the enemy of done

## What's Next?

- Add more blog posts (like this one!)
- Implement dark mode toggle
- Add project showcase section
- Maybe RSS feed support?

Thanks for reading! Feel free to [check out the source code](https://github.com/bhemi28/goth) or [reach out](/contact) if you have questions.

---

*Built with ❤️ and lots of coffee in Ahmedabad, Gujarat*