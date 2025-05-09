<p align="center">
  <img src="/docs/logo.png" alt="stringalign logo" width="525"/>
</p>

`stringalign` is a lightweight Go package for wrapping and aligning multi‑line text to a fixed width. It supports left, right, center, and full‑justify modes, preserves ANSI escape sequences (via `ansiwalker`), and handles full Unicode width correctly using `go-runewidth`.

## ✨ **Features**

* **Left, Right, Center & Justify**: wrap text then align each line within your width limit.
* **ANSI‑aware**: skips over CSI, OSC, DCS, SOS, PM, APC, and C1 codes so colors and styles don’t break wrapping.
* **Unicode safe**: computes visual width with full rune support (including emojis and wide CJK characters).
* **Simple API**: four top‑level functions—`LeftAlign`, `RightAlign`, `CenterAlign`, `Justify`—each returning `string, error`.

## 🚀 **Getting Started**

```bash
go get github.com/galactixx/stringalign@latest
```

## 📚 **Usage**

```go
import (
    "fmt"
    "github.com/galactixx/stringalign"
)

text := "Hello, 🌍!\nThis is a long line that will be wrapped and aligned."
width := 30
```

### ***Left Align***

```go
left, _ := stringalign.LeftAlign(text, width)
fmt.Println(left)
```

### Output:

```text
Hello, 🌍!
This is a long line that will
be wrapped and aligned.
```

---

### ***Right Align***

```go
right, _ := stringalign.RightAlign(text, width)
fmt.Println(right)
```

### Output:

```text
                    Hello, 🌍!
 This is a long line that will
       be wrapped and aligned.
```

---

### ***Center Align***

```go
center, _ := stringalign.CenterAlign(text, width)
fmt.Println(center)
```

### Output:

```text
          Hello, 🌍!
This is a long line that will
   be wrapped and aligned.
```

---

### ***Justify***

```go
justified, _ := stringalign.Justify(text, width)
fmt.Println(justified)
```

### Output:

```text
Hello, 🌍!
This  is a long line that will
be wrapped and aligned.
```

## 🔍 **API**

```go
func LeftAlign(str string, limit int, tabSize int) (string, error)
func RightAlign(str string, limit int, tabSize int) (string, error)
func CenterAlign(str string, limit int, tabSize int) (string, error)
func Justify(str string, limit int, tabSize int) (string, error)
```

* **str**: input multi‑line text (may include ANSI escapes).
* **limit**: maximum visible width per line.
* **tabSize**: number of spaces to replace tabs when aligning.
* **returns**: aligned text with `\n`‑separated lines, or an error on wrap failure.

## 🤝 **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 📞 **Contact**

For questions or support, please open an issue on the [GitHub repository](https://github.com/galactixx/stringalign/issues).
