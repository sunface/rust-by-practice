# 7. 🌟
```rs
// Remove something to make it work
// Don't remove a whole line !
fn main() {
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {}", r1, r2);

    println!("Success!");
}
```
# Solution:
```rs
// Remove something to make it work
// Don't remove a whole line !
fn main() {
    let mut s = &String::from("hello");

    let r1 = s;
    let r2 = s;

    println!("{}, {}", r1, r2);

    println!("Success!");
}
```
