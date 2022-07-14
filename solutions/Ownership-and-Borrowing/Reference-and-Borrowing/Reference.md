# 1. 🌟
```rs
fn main() {
   let x = 5;
   // Fill the blank
   let p = __;

   println!("the memory address of x is {:p}", p); // One possible output: 0x16fa3ac84
}
```
# Solution:
```rs

fn main() {
   let x = 5;
   // Fill the blank
   let p = &x;

   println!("the memory address of x is {:p}", p); // One possible output: 0x16fa3ac84
}
```
# 2. 🌟
```rs
fn main() {
    let x = 5;
    let y = &x;

    // Modify this line only
    assert_eq!(5, y);

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let x = 5;
    let y = &x;

    // Modify this line only
    assert_eq!(5, *y);

    println!("Success!");
}
```
# 3. 🌟
```rs
// Fix error
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(s);

    println!("Success!");
}

fn borrow_object(s: &String) {}
```
# Solution:
```rs
// Fix error
fn main() {
    let mut s = &String::from("hello, ");

    borrow_object(s);

    println!("Success!");
}

fn borrow_object(s: &String) {}
```
# 4. 🌟
```rs
// Fix error
fn main() {
    let mut s = String::from("hello, ");

    push_str(s);

    println!("Success!");
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```
# Solution:
```rs
// Fix error
fn main() {
    let mut s = String::from("hello, ");

    push_str(&mut s);

    println!("Success!");
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```
# 5. 🌟🌟
```rs
fn main() {
    let mut s = String::from("hello, ");

    // Fill the blank to make it work
    let p = __;
    
    p.push_str("world");

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let mut s = String::from("hello, ");

    // Fill the blank to make it work
    let mut p = s;
    
    p.push_str("world");

    println!("Success!");
}
```