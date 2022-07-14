# 1. 🌟 We can't use str type in normal ways, but we can use &str.
```rs
// Fix error without adding new line
fn main() {
    let s: str = "hello, world";

    println!("Success!");
}
```
 # Solution:
 ```rs
// Fix error without adding new line
fn main() {
    let s: &str = "hello, world";

    println!("Success!");
}
```
# 2. 🌟🌟 We can only use str by boxed it, & can be used to convert Box<str> to &str
```rs
// Fix the error with at least two solutions
fn main() {
    let s: Box<str> =  "hello, world".into();
    greetings(s)
}

fn greetings(s: &str) {
    println!("{}",s)
}
```
# Solution:
```rs
// Fix the error with at least two solutions
fn main() {
    let s: &str =  "hello, world".into();
    greetings(s)
}

fn greetings(s: &str) {
    println!("{}",s)
}
```
or
```rs
// Fix the error with at least two solutions
fn main() {
    let s: Box<str> =  "hello, world".into();
    greetings(&s)
}

fn greetings(s: &str) {
    println!("{}",s)
}
```
