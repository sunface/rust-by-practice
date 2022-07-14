## A scope is the range within the program for which the item is valid.

# 3. 🌟
```rs

// Fix the error below with least amount of modification
fn main() {
    let x: i32 = 10;
    {
        let y: i32 = 5;
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y); 
}
```
# Solution:
```rs
// Fix the error below with least amount of modification
fn main() {
    let x: i32 = 10;
    let y: i32 = 10; 
    {
        let y: i32 = 5;
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y); 
}
```
# output:
```
The value of x is 10 and value of y is 5
The value of x is 10 and value of y is 10
```

# 4. 🌟🌟
```rs

// Fix the error with the use of define_x
fn main() {
    println!("{}, world", x); 
}

fn define_x() {
    let x = "hello";
}
```
# Solution:
```rs
fn main() {
let x = define_x(); 
    println!("{}, world", x); 
}

fn define_x() -> &'static str{
    let x = "hello";
    x
}
```