# Numbers

### Integer

🌟 remove something to make it work

> Tips: If we don't explicitly give one type to a varible, then the compiler will infer one for us
```rust,editable

fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    y = x;
    
    let z = 10; // type of z ? 
}
```


🌟🌟🌟  modify `assert_eq!` to make it work

> Tips: If we don't explicitly give one type to a varible, then the compiler will infer one for us

```rust,editable

fn main() {
    let x = 5;
    assert_eq!("u32".to_string(), type_of(&x));
}

// get the type of given variable, return a string representation of the type  , e.g "i8", "u8", "i32", "u32"
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```

🌟🌟 fill the blanks to make it work
```rust,editable

fn main() {
    assert_eq!(i8::MAX, __); 
    assert_eq!(u8::MAX, __); 
}
```

🌟🌟 fix errors and panics to make it work
```rust,editable

fn main() {
   let v1 = 251_u8 + 8;
   let v2 = i8::checked_add(251, 8).unwrap();
   println!("{},{}",v1,v2);
}
```

🌟🌟🌟 modify `assert!` to make it work
```rust,editable

fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1579);
}
```

### Floating-Point
🌟 replace ? with your answer

```rust,editable

fn main() {
    let x = 1_000.000_1; // ?
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64
}
```
🌟🌟🌟 use two ways to make it work

> Tips: 1. abs 2. f32

```rust,editable

fn main() {
    assert!(0.1+0.2==0.3);
}
```