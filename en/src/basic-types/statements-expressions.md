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
        2 * x
    };

    println!("x is {:?}", x); // 5 
    println!("y is {:?}", y); // 25
    println!("z is {:?}", z); // 155
}
```

### Exercises
1. 🌟🌟
```rust,editable
// Make it work with two ways
fn main() {
   let v = {
       let mut x = 1;
       x += 2;
       x
   };

   assert_eq!(v, 3);

   println!("Success!");
}

// Make it work with two ways
fn main() {
   let v = {
       let mut x = 3;
       x
   };

   assert_eq!(v, 3);

   println!("Success!");
}
```

2. 🌟
```rust,editable

fn main() {
   let v = {let x = 3; x};

   assert!(v == 3);

   println!("Success!");
}
```

3. 🌟
```rust,editable

fn main() {
    let s = sum(1 , 2);
    assert_eq!(s, 3);

    println!("Success!");
}

fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it