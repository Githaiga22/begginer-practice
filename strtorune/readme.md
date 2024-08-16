## Convert a String to Runes

### Question:
 Write a Go program that takes the string "Hello, 世界" as input and converts it into a slice of runes. Print each rune and its corresponding Unicode code point in the format: Rune: H, Code Point: 72.

Step 1: Understand the Problem

- Goal: Convert a string into runes and print each rune with its Unicode code point.
-  Concepts: In Go, a rune is an alias for int32, and it represents a Unicode code point. Strings in Go are a sequence of bytes, but when dealing with Unicode characters (like those in "世界"), you should use runes.

Step 2: Writing Pseudo-Code

Let's outline what the program should do:

- Define the string "Hello, 世界".
-   Convert the string into a slice of runes.
-   Iterate over the slice of runes.
-   For each rune, print the rune itself and its Unicode code point.

output
```bash
Rune: H, Code Point: U+0048
Rune: e, Code Point: U+0065
Rune: l, Code Point: U+006C
Rune: l, Code Point: U+006C
Rune: o, Code Point: U+006F
Rune: ,, Code Point: U+002C
Rune:  , Code Point: U+0020
Rune: 世, Code Point: U+4E16
Rune: 界, Code Point: U+754C
```

### Comparison with fmt.Printf vs. println:

 #### println(r): Prints the integer value of the rune (its Unicode code point as a decimal integer).

 terminals output
 ```bash
 72
101
108
108
111
44
32
19990
30028
```

  What Each Number Represents:
```  
  72: The Unicode code point for 'H'.
  101: The Unicode code point for 'e'.
  108: The Unicode code point for 'l' (appears twice for both 'l' characters).
  111: The Unicode code point for 'o'.
  44: The Unicode code point for ','.
  32: The Unicode code point for the space character ' '.
  19990: The Unicode code point for '世'.
  30028: The Unicode code point for '界'.
  ```
#### fmt.Printf("%c, %U\n", r, r): Prints the character (%c) and the Unicode code point in hexadecimal format (%U).

1. "%c": Printing the Rune as a Character

    - %c is a verb in Go's fmt package that formats the output as a single character.
    - When you use %c with a rune, it prints the actual character that the rune represents.
    - For example, if r is the rune for the letter 'H', %c will print H.

2. "Code Point: %U": Printing the Unicode Code Point

    - %U is another verb that formats the output as a Unicode code point.
    - It prints the Unicode code point in the format U+XXXX, where XXXX is the hexadecimal value of the code point.
    - For example, if r is the rune for the letter 'H', %U will print U+0048, which is the Unicode code point for 'H'.

3. \n: New Line

    - \n is a newline character, which moves the cursor to the next line after printing the current line.
    - This ensures that each rune and its code point are printed on separate lines.

4. "r, r": Passing the Same Rune Twice

    - The format string expects two values to be passed:
     -  The first r is for %c, to print the character.
      - The second r is for %U, to print the Unicode code point.
    - By passing r twice, we print both the character and its corresponding Unicode code point.

### Key Differences:

1.  Format of the Unicode Code Point:
    - fmt.Printf("%U"): Displays the Unicode code point in hexadecimal format (base 16) with a U+ prefix.
      - Example: U+0048 for 'H'.
       - println(r): Displays the Unicode code point as a decimal integer (base 10).
    -   Example: 72 for 'H'.

2.  Representation of Characters:
 -   fmt.Printf("%c"): Explicitly shows the character itself, allowing you to see what the rune represents.
  -    Example: 'H'.
   -   println(r): Does not directly show the character; it only shows the code point value as an integer.


