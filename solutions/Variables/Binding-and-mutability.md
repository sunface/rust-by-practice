# Question: 1. 🌟 A variable can be used only if it has been initialized.
```rs
// Fix the error below with least amount of modification to the code
fn main() {
    let x: i32; // Uninitialized but used, ERROR !
    let y: i32; // Uninitialized but also unused, only a Warning !

    assert_eq!(x, 5);
    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let x: i32 = 5; // Initialized and used
    let y: i32; // Uninitialized but also unused, only a Warning !

    assert_eq!(x, 5);
    println!("Success!");
}
```

# Question: 2. 🌟 Use mut to mark a variable as mutable.
```rust

// Fill the blanks in the code to make it compile
fn main() {
    let __ =  1;
    __ += 2; 
    
    assert_eq!(x, 3);
    println!("Success!");
}
```
# Solution:
```rs
// Fill the blanks in the code to make it compile
fn main() {
    let mut x =  1;
    x += 2; 
    
    assert_eq!(x, 3);
    println!("Success!");
}
```