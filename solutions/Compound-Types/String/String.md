# String type is defined in std and stored as a vector of bytes (Vec), but guaranteed to always be a valid UTF-8 sequence. String is heap allocated, growable and not null terminated.
# 3.🌟
```rs
// Fill the blank
fn main() {
    let mut s = __;
    s.push_str("hello, world");
    s.push('!');

    assert_eq!(s, "hello, world!");

    println!("Success!");
}
```
 # Solution:
 ```rs
// Fill the blank
fn main() {
    let mut s = String::from("");
    s.push_str("hello, world");
    s.push('!');

    assert_eq!(s, "hello, world!");

    println!("Success!");
}
```
# 4. 🌟🌟🌟
```rs
// Fix all errors without adding newline
fn main() {
    let  s =  String::from("hello");
    s.push(',');
    s.push(" world");
    s += "!".to_string();

    println!("{}", s);
}
```
# Solution:
```rs
// Fix all errors without adding newline
fn main() {
    let mut s =  String::from("hello");
    s.push(',');
    s.push_str(" world");
    s += "!";

    println!("{}", s);
}
```
# 5. 🌟🌟 replace can be used to replace substring
```rs
// Fill the blank
fn main() {
    let s = String::from("I like dogs");
    // Allocate new memory and store the modified string there
    let s1 = s.__("dogs", "cats");

    assert_eq!(s1, "I like cats");

    println!("Success!");
}
```
# Solution:
```rs
// Fill the blank
fn main() {
    let s = String::from("I like dogs");
    // Allocate new memory and store the modified string there
    let s1 = s.replace("dogs", "cats");

    assert_eq!(s1, "I like cats");

    println!("Success!");
}
```
# 6. 🌟🌟 You can only concat a String with &str, and String's ownership can be moved to another variable.
```rs
// Fix errors without removing any line
fn main() {
    let s1 = String::from("hello,");
    let s2 = String::from("world!");
    let s3 = s1 + s2; 
    assert_eq!(s3,"hello,world!");
    println!("{}",s1);
}
```
# Solution:
```rs
// Fix errors without removing any line
fn main() {
    let s1 = String::from("hello,");
    let s2 = String::from("world!");
    let s3 = s1.clone() + &s2; 
    assert_eq!(s3,"hello,world!");
    println!("{}",s1);
}
```
