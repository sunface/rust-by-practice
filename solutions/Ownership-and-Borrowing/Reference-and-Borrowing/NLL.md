# 10. 🌟🌟
```rs

// Comment one line to make it work
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    r1.push_str("world");
    let r2 = &mut s;
    r2.push_str("!");
    
    println!("{}",r1);
}
```
# Solution:
```rs
// Comment one line to make it work
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    r1.push_str("world");
    // let r2 = &mut s;
    r1.push_str("!");
    
    println!("{}",r1);
}
```
# 11. 🌟🌟
```rs
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    let r2 = &mut s;

    // Add one line below to make a compiler error: cannot borrow `s` as mutable more than once at a time
    // You can't use r1 and r2 at the same time
}
```
# Solution:
```rs
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    let r2 = &mut s;

    // Add one line below to make a compiler error: cannot borrow `s` as mutable more than once at a time
    print!("{}",r2)
    // You can't use r1 and r2 at the same time
}
```