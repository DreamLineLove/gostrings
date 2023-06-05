# To Add
- Write an initial quiz program that allows you to pick which topic you would like to be quizzed on.
- Maybe use the flags library for it?
- Write a universal package with helper functions that allow you to construct the quiz in every repo. The flag might be a required argument in one of the functions. But it need not necessarily be part of the library.
- Write a production-ready (to be consumed by other devs) version of this quiz package
- Could be called the *CLI Quiz Maker* perhaps ;)

# Feature Ideas to Include
- Command-line flag/ user-input chooses from these game modes:
    - Single player with point-tracking (with end)/ no point-tracking (user manually ends)
    - Multiplayer with point-tracking and winner determination (with end)
        - Everyone takes turn answering
        - Knower answers (the number of questions in the question set should be divisible by the number of players!)
- Multiplayer mode could use a golang map to keep track of the player-score value-pairs and alternate each player each question until the end of the question set is reached.
