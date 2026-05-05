# 🎨 ASCII Art CLI Program

A powerful and interactive **command-line ASCII Art generator** built with Go.
This project allows users to transform input text into stylized ASCII art using different banner fonts, colors, and styles — enhanced with smooth terminal animations for a modern CLI experience.

---

## ✨ Features

* 🖋️ ASCII art generation from text input
* 🎭 Multiple banner fonts:

  * Standard
  * Shadow
  * Thinkertoy
* 🌈 Text styling with colors and ANSI effects
* 🎬 Smooth terminal animations:

  * Loading animations
  * Progress bars
  * Typing effects
  * Blink effects
  * Rainbow text animation
* 🔒 Input validation for safe and clean user input
* 👤 User-friendly CLI onboarding flow (name, options, configuration)
* ⚡ Fast and lightweight Go implementation

---

## 🛠️ Project Structure

```
Ascii-art-PRO/
│
├── animation/   # Terminal animations (loading, progress bar, effects)
├── ascii/       # ASCII art generation logic and banner parsing
├── style/       # ANSI color and text styling utilities
├── main.go      # CLI entry point and user interaction flow
```

---

## 🚀 How It Works

1. The user enters their name and configuration settings.
2. The program validates input and loads the selected font banner.
3. The input text is converted into ASCII art.
4. Output is styled with selected colors and text effects.
5. Animated CLI feedback enhances user experience throughout the process.

---

## ▶️ Usage

### Run the program:

```bash
go run main.go
```

### Example flow:

```
⚬ Fullname: John Doe
⚬ TEXT: Hello World
⚬ FONT: Shadow
⚬ COLOUR: Green
⚬ STYLE: Bold
```

Output:

* Styled ASCII art rendered in terminal
* Animated transitions and progress effects

---

## 📦 Requirements

* Go 1.18+
* Terminal with ANSI escape code support

---

## 🎯 Purpose

This project was built to strengthen understanding of:

* Go package structuring
* File I/O handling
* CLI design
* String manipulation
* ASCII rendering logic
* Terminal UI effects using ANSI codes

---

## 🧠 Future Improvements

* Add custom font uploads
* Export ASCII output to files
* Add GUI version (web or desktop)
* Improve animation concurrency (goroutines)
* Add more ASCII styles and themes

---

## 👨‍💻 Author

Built as a learning project to explore CLI design and ASCII rendering in Go.

---
