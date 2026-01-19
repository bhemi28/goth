# Static Portfolio

This is a static version of the portfolio website, designed to be deployed on GitHub Pages or any static site host. It uses HTML, CSS (Tailwind via CDN), and vanilla JavaScript with htmx for dynamic content loading.

## Structure

*   `index.html`: The main entry point. It sets up the layout and loads the initial content.
*   `css/`: Contains the stylesheets.
*   `js/`: Contains JavaScript files.
*   `images/`: Contains images assets.
*   `partials/`: Contains HTML fragments loaded by htmx.
    *   `home.html`: The main homepage content.
    *   `under-development.html`: Placeholder for pages not yet static (Blog, Contact).
    *   `projects/`: Individual project pages.

## How to Run Locally

Since this site uses HTMX to load partials, you cannot simply open `index.html` in your browser due to CORS (Cross-Origin Resource Sharing) restrictions. You must serve it over HTTP.

### Using Python

If you have Python installed, you can easily start a local server:

1.  Open a terminal in this directory (`portfolio/`).
2.  Run one of the following commands:
    *   Python 3: `python3 -m http.server`
    *   Python 2: `python -m SimpleHTTPServer`
3.  Open your browser to `http://localhost:8000`.

### Using Node.js (http-server)

1.  Install `http-server` globally if you haven't already: `npm install -g http-server`
2.  Run `http-server .` in this directory.
3.  Open the URL shown in the terminal.

### Using VS Code

1.  Install the "Live Server" extension.
2.  Right-click `index.html` and select "Open with Live Server".

## Deployment to GitHub Pages

1.  Push this folder to a GitHub repository.
    *   If this is your main repository, you might want to push the contents of this folder to the root of a `gh-pages` branch, or configure GitHub Pages to serve from the `/portfolio` folder on the `main` branch.
2.  Go to your repository settings on GitHub.
3.  Navigate to the "Pages" section.
4.  Select the source branch and folder (e.g., `main` branch, `/portfolio` folder, or `gh-pages` branch depending on how you set it up).
5.  Click Save. Your site will be live shortly.
