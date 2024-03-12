# Box

1. 🌟
```rust,editable
// Make it work
fn main() {
    // Create a new box `b` that contains the integer 5
    assert_eq!(*b, 5);

    println!("Success!");
}
```

2. 🌟
```rust,editable

// Make it work
fn main() {
    let b = Box::new("Hello");
    print_boxed_string(b);
}

fn print_boxed_string(b : _) {
    println!("{}", b);
}
```

3. 🌟
```rust,editable

// Make it work
fn main() {
    let b1 = Box::new(5);
    let b2 = b1;
    assert_eq!(_, 5);

    println!("Success!");
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it
