# 'println!' și 'format!'
Afișarea este gestionată de o serie de ['macro-uri'][macros] definite în ['std::fmt'][fmt],
Printre acestea se numără:

* format!: scrie text formatat într-un [String][string]
* print!: la fel ca format!, dar textul este tipărit pe consolă (io::stdout).
* println!: la fel ca print!, dar se adaugă o linie nouă.
* eprint!: la fel ca format!, dar textul este tipărit pe eroarea standard (io::stderr).
* eprintln!: la fel ca eprint!, dar se adaugă o linie nouă.

Toate parsează textul în același mod. În plus, Rust verifică corectitudinea formatului la momentul compilării.

## `format!`
1.🌟
```rust,editable

fn main() {
    let s1 = "hello";
    /* Completează spațiul liber */
    let s = format!(__);
    assert_eq!(s, "hello, world!");
}
```

## `print!`, `println!`
2.🌟
```rust,editable

fn main() {
   /* Completați spațiile libere pentru a afișa:
   Hello world, I am 
   Sunface!
   */
   __("hello world, ");
   __("I am");
   __("Sunface!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
