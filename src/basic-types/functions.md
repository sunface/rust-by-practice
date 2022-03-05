# Functions
1. 🌟🌟🌟
```rust,editable

fn main() {
    // don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);

    println!("Success!")
}

fn sum(x, y: i32) {
    x + y;
}
```


2. 🌟
```rust,editable
fn main() {
   print();
}

// replace i32 with another type
fn print() -> i32 {
   println!("Success!")
}
```


3. 🌟🌟🌟

```rust,editable
// solve it in two ways
// DON'T let `println!` works
fn main() {
    never_return();

    println!("Failed!")
}

fn never_return() -> ! {
    // implement this function, don't modify the fn signatures
    
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it