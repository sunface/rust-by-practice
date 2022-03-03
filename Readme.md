# Rust By Practice

This book was designed for easily diving into Rust，and it's very easy to use: All you need to do is to make each exercise comipile without ERRORS and Panics !

  
## Features

- There are three parts in each chapter: examples, exercises and practices
- Covering nearly all aspects of Rust, such as **async/await, threads, sync primitives, optimizing and stand libraries** etc
- Solution for each exercise
- Difficulty from easy to super hard: easy 🌟  medium 🌟🌟 hard 🌟🌟🌟  super hard 🌟🌟🌟🌟
- Both [English](https://practice.rs) and [Chinsese](https://zh.practice.rs) are supported

## Some of our exercises

🌟🌟🌟 Tuple struct looks similar to tuples, it has added meaning the struct name provides but has no named fields. It's useful when you want give the whole tuple a name, but don't care the fields's names.

```rust

// fix the error and fill the blanks
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);
fn main() {
    let v = Point(___, ___, ___);
    check_color(v);
}   

fn check_color(p: Color) {
    let (x, _, _) = p;
    assert_eq!(x, 0);
    assert_eq!(p.1, 127);
    assert_eq!(___, 255);
 }
```

🌟🌟 Within the destructuring of a single variable, both by-move and by-reference pattern bindings can be used at the same time. Doing this will result in a partial move of the variable, which means that parts of the variable will be moved while other parts stay. In such a case, the parent variable cannot be used afterwards as a whole, however the parts that are only referenced (and not moved) can still be used.
```rust

// fix errors to make it work
#[derive(Debug)]
struct File {
    name: String,
    data: String,
}
fn main() {
    let f = File {
        name: String::from("readme.md"),
        data: "Rust By Practice".to_string()
    };

    let _name = f.name;

    // ONLY modify this line
    println!("{}, {}, {:?}",f.name, f.data, f);
} 
```

🌟🌟 A match guard is an additional if condition specified after the pattern in a match arm that must also match, along with the pattern matching, for that arm to be chosen.
```rust,editable

// fill in the blank to make the code work, `split` MUST be used
fn main() {
    let num = Some(4);
    let split = 5;
    match num {
        Some(x) __ => assert!(x < split),
        Some(x) => assert!(x >= split),
        None => (),
    }
}
```
