1.

```rust
fn main() {
    let x: i32 = 5; // uninitialized but using, ERROR !
    let y: i32; // uninitialized but also unusing, only warning
    println!("{} is equal to 5", x);
}
```

2.

```rust
fn main() {
    let mut x = 1;
    x += 2;

    println!("{} is equal to 3", x);
}
```

3.

```rust
fn main() {
    let x: i32 = 10;
    let y: i32 = 20;
    {
        let y: i32 = 5;
        println!("The value of x is {} and value of y is {}", x, y);
    }
    println!("The value of x is {} and value of y is {}", x, y); 
}
```

4.

```rust
fn main() {
    let x = define_x();
    println!("{}, world", x);
}

fn define_x() -> String {
    let x = "hello".to_string();
    x
}
```

```rust
fn main() {
    let x = define_x();
    println!("{:?}, world", x);
}

fn define_x() -> &'static str {
    let x = "hello";
    x
}
```

5.

```rust
fn main() {
    let x: i32 = 5;
    {
        let x = 12;
        assert_eq!(x, 12);
    }

    assert_eq!(x, 5);

    let x = 42;
    println!("{}", x); // Prints "42".
}
```

6.

```rust
fn main() {
    let mut x: i32 = 1;
    x = 7;
    // Shadowing and re-binding
    let mut x = x; 
    x += 3;


    let y = 4;
    // Shadowing
    let y = "I can also be bound to text!"; 

    println!("Success!");
}
```
```rust
fn main() {
    let mut x: i32 = 1;
    x = 7;
    // Remove "let x = x;"
    x += 3;


    let y = 4;
    // Shadowing
    let y = "I can also be bound to text!"; 

    println!("Success!");
}
```
7.

```rust
fn main() {
    let _x = 1;
}
```

```rust
#[allow(unused_variables)]
fn main() {
    let x = 1;
}
```

8.

```rust
fn main() {
    let (mut x, y) = (1, 2);
    x += 2;

    assert_eq!(x, 3);
    assert_eq!(y, 2);
}
```

```rust
fn main() {
    let (x, y) = (1, 2);
    let x = 3;

    assert_eq!(x, 3);
    assert_eq!(y, 2);
}
```

9.

```rust
fn main() {
    let (x, y);
    (x, ..) = (3, 4);
    [.., y] = [1, 2];
    // fill the blank to make the code work
    assert_eq!([x, y], [3, 2]);
}
```
