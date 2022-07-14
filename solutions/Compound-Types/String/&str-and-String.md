# Oppsite to the seldom using of str, &str and String are used everywhere!
# 7. 🌟🌟 &str can be converted to String in two ways
```rs
// Fix error with at least two solutions
fn main() {
    let s =  "hello, world";
    greetings(s)
}

fn greetings(s: String) {
    println!("{}",s)
}
```
# Solution:
```rs
// Fix error with at least two solutions
fn main() {
    let s =  "hello, world";
    greetings(s.to_string())
}

fn greetings(s: String) {
    println!("{}",s)
}
```
or
```rs
// Fix error with at least two solutions
fn main() {
    let s = String::from("hello, world");
    greetings(s)
}

fn greetings(s: String) {
    println!("{}",s)
}
```
# 8. 🌟🌟 We can use String::from or to_string to convert a &str to String
```rs 
// Use two approaches to fix the error and without adding a new line
fn main() {
    let s =  "hello, world".to_string();
    let s1: &str = s;

    println!("Success!");
}
```
# Solution:
```rs
// Use two approaches to fix the error and without adding a new line
fn main() {
    let s =  "hello, world".to_string();
    let s1 = s;

    println!("Success!");
}
```
or 
```rs
// Use two approaches to fix the error and without adding a new line
fn main() {
    let s =  String::from("hello, world");
    let s1 = s;

    println!("Success!");
}
```