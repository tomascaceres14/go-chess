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
I felt the need to challenge myself, to dive into unknown territory and see what I could come up to. At the beginning, I wasn't very sure how regular chess engines work, so I started with what was the most intuitive for me at the time: OOP and a 2D matrix containing the pieces.
After some time, I decided to do some research, only to realize that this isn't how a regular chess engine is made. The 2D Matrix datastructure, although intuitive, it's not the most efficient nor performant option. But I was already long into development, so I left it as it was.

Anyway, it still works pretty well and response fairly fast so it's not a big issue.

Feel free to clone/fork the repo and give it a try :)
