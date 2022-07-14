# 7. 🌟
```rs

//  Replace ? with your answer
fn main() {
    let x = 1_000.000_1; // ?
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64

    println!("Success!");
}
```
# Solution:
```rs
//  Replace ? with your answer
fn main() {
    let x = 1_000.000_1; // f64
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64

    println!("Success!");
}
```
# 8. 🌟🌟 Make it work in two distinct ways
```rs
fn main() {
    assert!(0.1 + 0.2 == 0.3);

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
let x:f32 = 0.1 + 0.2;
    assert!(x == 0.3);

    println!("Success!");
}
```
or
```rs
fn main() {
    assert!((0.1_f32+0.2_f32) == 0.3);

    println!("Success!");
}
```