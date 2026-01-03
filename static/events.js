// HTMX Loading State for Contact Form
document.addEventListener('htmx:beforeRequest', function(evt) {
    var form = evt.detail.elt;
    if (form.id === 'contact-form') {
        var btn = document.getElementById('submit-btn');
        if (btn) {
            var text = btn.querySelector('.btn-text');
            var icon = btn.querySelector('.btn-icon');
            var loader = btn.querySelector('.btn-loader');
            if (text) text.style.visibility = 'hidden';
            if (icon) icon.style.visibility = 'hidden';
            if (loader) loader.style.display = 'flex';
        }
    }
});

document.addEventListener('htmx:afterRequest', function(evt) {
    var form = evt.detail.elt;
    if (form.id === 'contact-form') {
        var btn = document.getElementById('submit-btn');
        if (btn) {
            var text = btn.querySelector('.btn-text');
            var icon = btn.querySelector('.btn-icon');
            var loader = btn.querySelector('.btn-loader');
            if (text) text.style.visibility = 'visible';
            if (icon) icon.style.visibility = 'visible';
            if (loader) loader.style.display = 'none';
        }
    }
});

// Email Copy Functionality
function copyEmail() {
    navigator.clipboard.writeText("meet.bhesaniya.prof@gmail.com").then(() => {
        const popup = document.createElement("div");
        popup.textContent = "Email Copied!";
        popup.style.position = "fixed";
        popup.style.bottom = "20px";
        popup.style.right = "20px";
        popup.style.backgroundColor = "#bbf7d0";
        popup.style.border = "3px solid #0f0f0f";
        popup.style.padding = "10px 20px";
        popup.style.fontFamily = "'VT323', monospace";
        popup.style.fontSize = "20px";
        popup.style.boxShadow = "4px 4px 0px 0px #0f0f0f";
        popup.style.zIndex = "1000";
        document.body.appendChild(popup);
        setTimeout(() => {
            popup.remove();
        }, 2000);
    });
}