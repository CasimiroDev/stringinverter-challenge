# 🔤 Go String Reverser: Safe UTF-8 and Rune Manipulation

An efficient Go implementation for reversing strings while ensuring the integrity of multibyte characters (UTF-8).

Although reversing a string seems fundamentally simple, this repository practically demonstrates core Go memory and encoding principles. It showcases a properly structured approach to safely manipulating strings, avoiding common pitfalls with byte slices and demonstrating O(N) algorithmic efficiency.

## 🧠 Architecture & Design Decisions

This codebase implements the following practices:

* **UTF-8 Handling (Rune Conversion):** In Go, strings are read-only byte slices. Reversing bytes directly corrupts accented characters (like `á` or `ç`), which occupy more than one byte. The solution converts the string into a slice of `runes` (int32), ensuring Unicode code points are handled as whole units, preserving character integrity.

* **O(N) Performance (Two Pointers Technique):** Instead of redundantly appending to a new slice (which causes costly memory reallocations and quadratic complexity), the algorithm uses a two-pointer approach (`i` at the start, `j` at the end). The reversal is done in-place through parallel assignment, iterating only up to the middle of the slice.

* **Conscious Limitations (Grapheme Clusters):** By strictly using the standard library, the solution correctly reverses runes but will intentionally flip the internal positions of Regional Indicators (like those making up the 🇧🇷 flag emoji). A fully Grapheme-Cluster-proof implementation would require a complex state machine or external packages.

* **Idiomatic Go:** Correct use of parallel assignment inside the `for` loop to swap values without requiring temporary variables.

# 🚀 Running the Project

The project was built using only the Go standard library, with no external dependencies required.

**Prerequisites:** Go 1.18+ installed.

1. Clone the repository:
```bash
git clone [https://github.com/CasimiroDev/stringinverter-challenge](https://github.com/CasimiroDev/stringinverter-challenge)
```

2. Navigate to the project folder:
```bash
cd go-string-reverser
```

3. Run the application:
```bash
go run main.go
```

# 🔮 Next Steps (To-Do)

- [ ] Unit Tests (TDD): Implement a test suite using the native `testing` package, validating edge cases such as empty strings, single characters, and palindromes.

- [ ] Grapheme Cluster Support: Integrate the `github.com/rivo/uniseg` library to properly handle complex emojis composed of multiple runes joined by Zero-Width Joiners (ZWJ).