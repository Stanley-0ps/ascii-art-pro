# 🎨 ASCII Art CLI Program
A powerful and interactive command-line ASCII Art generator built with **Go**.
Transform text into beautifully styled ASCII art with multiple fonts, colors, and smooth terminal animations.

---

## ✨ Features

* 🖋️ Generate ASCII art from text input
* 🎭 Multiple fonts:

  * Standard
  * Shadow
  * Thinkertoy
* 🌈 ANSI color customization
* ✨ Text styling (Bold, Dim, Italic, Stripes)
* 🎬 Terminal animations:

  * Loading animation
  * Progress bar
  * Typing effect
  * Blink effect
  * Rainbow animation
* 🔒 Input validation
* 👤 Interactive CLI onboarding
* ⚡ Fast and lightweight

---

## 🛠️ Project Structure

```
Ascii-art-PRO/
│
├── animation/   # Terminal animations
├── ascii/       # ASCII rendering logic
├── style/       # ANSI styling utilities
├── main.go      # Entry point
```

---

## ✳ Font Samples

### Standard

```
      _           _                     _____                  
     | |         | |                   |  __ \                 
     | |   ___   | |__    _ __         | |  | |   ___     ___  
 _   | |  / _ \  |  _ \  | '_ \        | |  | |  / _ \   / _ \ 
| |__| | | (_) | | | | | | | | |       | |__| | | (_) | |  __/ 
 \____/   \___/  |_| |_| |_| |_|       |_____/   \___/   \___| 
```

### Shadow

```
      _|          _|                      _|_|_|                     
      _|   _|_|   _|_|_|   _|_|_|         _|    _|   _|_|     _|_|   
      _| _|    _| _|    _| _|    _|       _|    _| _|    _| _|_|_|_| 
_|    _| _|    _| _|    _| _|    _|       _|    _| _|    _| _|       
  _|_|     _|_|   _|    _| _|    _|       _|_|_|     _|_|     _|_|_| 
```

### Thinkertoy

```
    o     o               o-o           
    |     |               |  \          
    | o-o O--o o-o        |   O o-o o-o 
\   o | | |  | |  |       |  /  | | |-' 
 o-o  o-o o  o o  o       o-o   o-o o-o 
```

---

## 🎨 Customization

### Colors

`RED, YELLOW, GREEN, BLUE, CYAN, MAGENTA, PURPLE, DARK RED, DARK YELLOW, WHITE, GREY`

### Styles

`BOLD, DIM, ITALIC, STRIPES`

---

## 🚀 Usage

### Run the program

```bash
go run main.go
```

### Example

```
Fullname: John Doe
TEXT: Hello World
FONT: Shadow
COLOUR: Green
STYLE: Bold
```

### Output

* Styled ASCII art in terminal
* Animated transitions

---

## 📦 Requirements

* Go 1.18+
* Terminal with ANSI escape code support

---

## 🎯 Purpose

This project helps you practice:

* Go project structuring
* CLI application design
* File handling
* String manipulation
* ASCII rendering logic
* Terminal UI with ANSI codes

---

## 🧠 Future Improvements

* [ ] Custom font uploads
* [ ] Export output to file
* [ ] GUI version (web/desktop)
* [ ] Improve animation with goroutines
* [ ] More styles and themes

---

## 👨‍💻 Author

Built as a learning project to explore CLI design and ASCII rendering in Go.

---

## 🧾 License

MIT License

---

## 🙏 Thank You 

Thanks for checking out this project!
Feel free to ⭐ star the repo and contribute 🚀
