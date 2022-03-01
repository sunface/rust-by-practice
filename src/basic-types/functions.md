# Functions
🌟🌟🌟
```rust,editable

fn main() {
    // don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(1, 2);

    assert_eq!(s, 3);
}

fn sum(x, y: i32) {
    x + y;
}
```


🌟🌟
```rust,editable
fn main() {
   print();
}

// replace i32 with another type
fn print() -> i32 {
   println!("hello,world");
}
```


🌟🌟

```rust,editable
fn main() {
    never_return();
}

fn never_return() -> ! {
    // implement this function, don't modify fn signatures
    
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it