# Statements and Expressions

### Examples
```rust,editable
fn main() {
    let x = 5u32;

    let y = {
        let x_squared = x * x;
        let x_cube = x_squared * x;

        // This expression will be assigned to `y`
        x_cube + x_squared + x
    };

    let z = {
        // The semicolon suppresses this expression and `()` is assigned to `z`
        2 * x;
    };

    println!("x is {:?}", x);
    println!("y is {:?}", y);
    println!("z is {:?}", z);
}
```

### Exercises
🌟🌟
```rust,editable
// make it work with two ways: both modify the inner {}
fn main() {
   let v = {
       let mut x = 1;
       x += 2
   };

   assert_eq!(v, 3);
}
```

🌟
```rust,editable

fn main() {
   let v = (let x = 3);

   assert!(v == 3);
}
```

🌟
```rust,editable

fn main() {}

fn sum(x: i32, y: i32) -> i32 {
    x + y;
}
```