# 🎨 ASCII Art Web Stylize

## ✨ Description

Bring your text to life with ASCII Art Web—a playful Go-powered web app that transforms your words into eye-catching ASCII banners. Choose between three unique fonts:

* **standard**: Classic and clean.
* **shadow**: Bold with a dramatic drop-shadow.
* **thinkertoy**: Whimsical and fun.

## 🚀 Quick Start

````bash
# Clone your project
```bash
git clone https://learn.reboot01.com/git/ymuhamma/ascii-art-web-stylize.git
cd ascii-art-web-stylize

# Get dependencies & launch the server
go mod tidy
go run main.go
````

Then, open your browser and go to [http://localhost:8080](http://localhost:8080). Enter your text, pick a font, and click **Generate** to see the magic!

## 🛠️ Features

* **Responsive UI**: Built with Bootstrap for a sleek experience on any device.
* **Real-time Preview**: Instantly render ASCII art in your browser.
* **Error Handling**: Friendly messages guide you if something goes wrong.

## 📂 Project Structure

```plaintext
ascii-art-web-stylize/
├── assets/
│   ├── fonts/       ✏️  Banner definitions (standard, shadow, thinkertoy)
│   ├── templates/   📄  HTML templates styled with localized CSS
│   └── static/      🎨  Static assets like CSS, videos and images
├── helpers/         🔧  General-purpose logic used across the project
├── handlers/        🌐  HTTP request routing and handler functions
├── main.go          🖥️  Server entry point and handlers
└── README.md        📘  This guide

```

## ✍️ Authors

* **Yousif Muhammad**
* **Ahmed Alsafseef**
* **Lora Mohamed**

---

Happy creating!
