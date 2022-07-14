# Ownership
# 1. 🌟🌟
```rs
fn main() {
    // Use as many approaches as you can to make it work
    let x = String::from("hello, world").split(' ');
    let y = &x[1];
    println!("{},{}",x,y);
}
```
# Solution:
```rs
fn main() {
    // Use as many approaches as you can to make it work
    let x = String::from("hello, world");
    let y = x.clone();
    println!("{},{}",x,y);
}
```
or
```rs
fn main() {
    // Use as many approaches as you can to make it work
    let x = String::from("hello, world");
    let y = &x;
    println!("{},{}",x,y);
}
```
or
```rs
fn main() {
    // Use as many approaches as you can to make it work
    let x = String::from("hello");
    let y = String::from(" world");
    println!("{},{}",x,y);
}
```
# 2. 🌟🌟
```rs
// Don't modify code in main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Only modify the code below!
fn take_ownership(s: String)-> String {
    // println!("{}", s);
    s
}
```
# Solution:
```rs
// Don't modify code in main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Only modify the code below!
fn take_ownership(s: String)-> String {
    // println!("{}", s);
    s
}
```
# 3. 🌟🌟
```rs
fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // Convert String to Vec
    // let _s = s.into_bytes();
    s
}
```
# Solution:
```rs

fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // Convert String to Vec
    // let _s = s.into_bytes();
    s
}
```
# 4. 🌟🌟
```rs
// Fix the error without removing code line
fn main() {
    let s = String::from("hello, world");

    print_str(s);

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```
# Solution:
```rs
// Fix the error without removing code line
fn main() {
    let mut s = String::from("hello, world");

    print_str(s.clone());

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```
# 5. 🌟🌟
```rs
// Don't use clone ,use copy instead
fn main() {
    let x = (1, 2, (), "hello".to_string());
    let y = x.clone();
    println!("{:?}, {:?}", x, y);
}
```
# Solution:
```rs
// Don't use clone ,use copy instead
fn main() {
    let x = (1, 2, (), "hello".to_string());
    let y = &x;
    println!("{:?}, {:?}", x, y);
}
```
