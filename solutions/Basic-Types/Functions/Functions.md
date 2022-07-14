# 1. 🌟🌟🌟
```rs
fn main() {
    // Don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);

    println!("Success!");
}

fn sum(x, y: i32) {
    x + y;
}
```
# Solution:
```rs
fn main() {
    // Don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);

    println!("Success!");
}
// add input types and return type
fn sum(x: i32, y: i32)-> i32 {
    x + y
}
```
# 2. 🌟
```rs
fn main() {
   print();
}

// Replace i32 with another type
fn print() -> i32 {
   println!("Success!");
}
```
# Solution:
```rs
fn main() {
   print();
}

// Replace i32 with another type
fn print() {
   println!("Success!");
}
```
# 3. 🌟🌟🌟
```rs
// Solve it in two ways
// DON'T let `println!` works
fn main() {
    never_return();

    println!("Failed!");
}

fn never_return() -> ! {
    // Implement this function, don't modify the fn signatures
    
}
```
# Solution:
```rs
fn main() {
    never_return();

    // println!("Failed!");
}

// use std::thread;
// use std::time;

fn never_return() -> ! {
    // implement this function, don't modify fn signatures
    panic!("I never return")
}
```
or 
```rs
fn main() {
    never_return();
}

use std::thread;
use std::time;

fn never_return() -> ! {
    // implement this function, don't modify fn signatures
    loop {
        println!("I return nothing");
        // sleeping for 1 second to avoid exhausting the cpu resoucre
        thread::sleep(time::Duration::from_secs(1))
    }
}
```