# ROADMAP

| Status | Feature                              | Description                                                                 |
|--------|--------------------------------------|-----------------------------------------------------------------------------|
| [x]     | Basic board setup                    | Board representation, coordinates, and utility functions                   |
| [x]     | Piece abstraction (`Movable`)        | Interface-based polymorphism for all pieces                                |
| [x]     | `BasePiece` composition              | Common logic shared across pieces (movement, threat generation, etc.)      |
| [x]     | King and Queen implementations       | Custom movement rules and proper string representation                     |
| [x]     | Movement logic (`PossibleMoves`)     | Valid destination squares for each piece                                   |
| [x]     | Threat logic (`PossibleThreats`)     | Squares each piece can attack or control                                   |
| [x]     | Piece tracking per player            | Player owns a list of `Movable` pieces                                     |
| [x]     | Piece position management            | `GetPosition` / `SetPosition`, stored in each piece                        |
| [x]     | Movement validation (basic)          | Prevents invalid movement directions or occupied-friendly squares          |
| [x]     | Piece capture                        | Can remove opposing pieces via `DeletePiece` and board interaction         |
| [x]     | Board cloning / move simulation      | Needed for legal move filtering and check evaluation                       |
| [x]     | Check detection (`IsChecked`)        | Determine if a king is under threat                                        |
| [x]     | Legal move filtering                 | Only allow moves that do not leave the king in check                       |
| [x]     | Checkmate detection                  | Detect if player is in check and has no legal moves                        |
| [x]     | Stalemate detection                  | Detect when there are no legal moves but no check                          |
| [ ]     | Turn management                      | Enforce turn order based on color                                          |
| [ ]     | Full game state (`Game`, `Match`)    | Store global game state, turn tracking, move history, etc.                 |
| [ ]     | En passant                           | Pawn special capture after a double-step move by opponent                  |
| [ ]     | Castling                             | King-side and queen-side, with rule validation (path, check, etc.)         |
| [ ]     | Pawn promotion                       | Allow promotion when reaching final rank                                   |
| [ ]     | Move history                         | Store previous moves for undo, en passant, repetition detection            |
| [ ]     | Repetition / 50-move rule            | Optional draw conditions                                                   |
| [ ]     | Simple AI (random or heuristic)      | Select legal moves for basic bot opponent                                  |


