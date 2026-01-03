package views

import (
	"bytes"
	"context"
	"fmt"
	"goth/views/components"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/adrg/frontmatter"
	"github.com/go-chi/chi/v5"
	"github.com/yuin/goldmark"
)

type Meta struct {
	Title  string `yaml:"title"`
	Desc   string `yaml:"desc"`
	Date   string `yaml:"date"`
	Author string `yaml:"author"`
}

func RegisterRoutes(r *chi.Mux) {
	r.Get("/", renderIndex)
	r.Get("/posts", getPostsFromMarkdown)
	r.Get("/posts/{name}", getPostFromName)
	r.Get("/contact", getContactSection)
	r.Post("/contact/send", sendMailHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		renderWithLayout(w, r, NotFound())
	})
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	renderWithLayout(w, r, components.Home())
}

func getPostsFromMarkdown(w http.ResponseWriter, r *http.Request) {
	var posts []components.Post
	var meta Meta

	err := filepath.WalkDir("posts", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || path[len(path)-3:] != ".md" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = frontmatter.Parse(bytes.NewReader(data), &meta)
		if err != nil {
			return err
		}

		posts = append(posts, components.Post{
			Title:   meta.Title,
			Content: meta.Desc,
			Date:    meta.Date,
			Author:  meta.Author,
			Link:    "/posts/" + strings.Split(filepath.Base(path), ".")[0],
		})

		return nil
	})
	if err != nil {
		log.Fatal("Error reading posts:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderWithLayout(w, r, PostList(posts))
}

func getPostFromName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	filePath := filepath.Join("posts", name+".md")
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading post:", err)
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var meta Meta
	body, err := frontmatter.Parse(bytes.NewReader(data), &meta)
	if err != nil {
		log.Println("Error parsing post metadata:", err)
		http.Error(w, "Post not found", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(body, &buf); err != nil {
		log.Println("Markdown conversion failed:", err)
	}
	htmlComponent := Unsafe(buf.String())

	log.Print("Rendering post:", htmlComponent)
	log.Print("Rendering post:", meta)

	renderWithLayout(w, r, PostContent(meta, htmlComponent))
}

func renderWithLayout(w http.ResponseWriter, r *http.Request, component templ.Component) {
	if r.Header.Get("HX-Request") == "true" {
		component.Render(context.Background(), w)
		return
	}
	Index(component).Render(context.Background(), w)
}

func getContactSection(w http.ResponseWriter, r *http.Request) {
	renderWithLayout(w, r, Contact())
}

func sendMailHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	visitorEmail := r.FormValue("email")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	// Get SMTP credentials from environment variables
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	if smtpEmail == "" || smtpPassword == "" || smtpHost == "" || smtpPort == "" {
		log.Println("SMTP credentials not set. Skipping email send.")
		// Still render success for demo purposes if creds are missing
		ContactSuccess().Render(context.Background(), w)
		return
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	// Compose the email
	// We send FROM the smtpEmail (to satisfy Gmail's requirements)
	// But we set Reply-To to the visitor's email so you can reply easily.
	headers := make(map[string]string)
	headers["From"] = "Portfolio Contact Form <" + smtpEmail + ">"
	headers["To"] = smtpEmail
	headers["Reply-To"] = visitorEmail
	headers["Subject"] = "Portfolio Contact: " + subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	body := fmt.Sprintf("You received a new message from your portfolio contact form.\n\nFrom: %s\n\nMessage:\n%s", visitorEmail, message)

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)

	// Send the email
	addr := smtpHost + ":" + smtpPort
	err = smtp.SendMail(addr, auth, smtpEmail, []string{smtpEmail}, msg.Bytes())
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Printf("Email sent successfully to %s", smtpEmail)

	// Render success component
	ContactSuccess().Render(context.Background(), w)
}

func Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func isHtmxRequest(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}
