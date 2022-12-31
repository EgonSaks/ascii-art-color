# ASCII art color

This program paints string based on letters and indexes.

+ You can run `go run . --color=red "hello"` or `go run . --color=red "hello" :` for fully coloring hello to red.

+ You can also choose what letter from `h` `e` `l` `l` `o` you want to color just by adding after the input string letter or letters with spaces between each other e.g. `go run . --color=red "hello" e o ` so you will have `e` and `o` painted red. 

+ You can choose part of the word or sentence to be colored by using slicing after the input string e.g.

    + `3:` for coloring all letters from the 3rd letter, including 3rd.
    + `:3` for coloring all letters till 3rd letter, including 3rd.
    + `2:4` for coloring all letters from 2nd till 4th, including 2nd and 4th.
    + `:` coloring all letters, same as writing just string without any arguments

To test it with test cases run `sh test.sh`

![hello](/img/hello.png)

You can choose between colors for letters or also background.

#### Available colors/keywords:

```
black
red
green
yellow
orange
blue
magenta
cyan
white
brightBlack
brightRed
brightGreen
brightYellow
brightBlue
brightMagenta
brightCyan
brightWhite
backgroundBlack
backgroundRed
backgroundGreen
backgroundYellow
backgroundBlue
backgroundMagenta
backgroundCyan
backgroundWhite
backgroundBrightBlack
backgroundBrightRed
backgroundBrightGreen
backgroundBrightYellow
backgroundBrightBlue
backgroundBrightMagenta
backgroundBrightCyan
backgroundBrightWhite
```
Thank you :)
