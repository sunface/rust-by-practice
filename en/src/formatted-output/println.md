# println! and format!
Printing is handled by a series of [`macros`][macros] defined in [`std::fmt`][fmt]
Some of which include:

* `format!`: write formatted text to [`String`][string]
* `print!`: same as `format!` but the text is printed to the console (io::stdout).
* `println!`: same as `print!` but a newline is appended.
* `eprint!`: same as `format!` but the text is printed to the standard error (io::stderr).
* `eprintln!`: same as `eprint!`but a newline is appended.

All parse text in the same fashion. As a plus, Rust checks format correctness at compile time.

## `format!`
1.🌟
```rust,editable

fn main() {
    let s1 = "hello";
    /* Fill in the blank */
    let s = format!(__);
    assert_eq!(s, "hello, world!");
}
```

## `print!`, `println!`
2.🌟
```rust,editable

fn main() {
   /* Fill in the blanks to make it print:
   Hello world, I am 
   Sunface!
   */
   __("hello world, ");
   __("I am");
   __("Sunface!");
}
```
