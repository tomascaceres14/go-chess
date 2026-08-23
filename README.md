# Go Chess

**A custom from-scratch implementation of a chess engine written in Go.**

This project was born from a mix of curiosity and passion for coding and chess (and the Go programming language, obviously).

The project consists of two main parts:

## 1. The Engine
A fully-functional, object-oriented chess engine that can be used as a standalone Go module into any project.
The engine implements all of the major chess rules, including:

- Legal move generation for every piece.
- Captures.
- Check / checkmate detection.
- Pins and move filtering.
- Special moves such as en passant and castling.
- There's still missing all of the tie possibilities but I'll get there.
 
## 2. The Game (in progress)
A graphical chess game built using Ebitengine to showcase the chess working engine with a not-so-attractive visual interface.

---

### The learning process
I wanted to challenge myself by exploring unfamiliar territory and building a chess engine from scratch. I started with an object-oriented design using a 2D matrix to represent the board and pieces, as it was the most intuitive approach for modeling the game logic.
Anyway, it still works pretty well and response fairly fast so it's not a big issue.

Feel free to clone/fork the repo and give it a try :)
