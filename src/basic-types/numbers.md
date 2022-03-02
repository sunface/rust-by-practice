# Numbers

### Integer

1. 🌟 

> Tips: If we don't explicitly give one type to a varible, then the compiler will infer one for us

```rust,editable

// remove something to make it work
fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    y = x;
    
    let z = 10; // type of z ? 
}
```

2. 🌟
```rust,editable

//  fill the blank
fn main() {
    let v: u16 = 38_u8 as __;
}
```

3. 🌟🌟🌟  

> Tips: If we don't explicitly give one type to a varible, then the compiler will infer one for us

```rust,editable

// modify `assert_eq!` to make it work
fn main() {
    let x = 5;
    assert_eq!("u32".to_string(), type_of(&x));
}

// get the type of given variable, return a string representation of the type  , e.g "i8", "u8", "i32", "u32"
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```

4. 🌟🌟 
```rust,editable

// fill the blanks to make it work
fn main() {
    assert_eq!(i8::MAX, __); 
    assert_eq!(u8::MAX, __); 
}
```

5. 🌟🌟 
```rust,editable

// fix errors and panics to make it work
fn main() {
   let v1 = 251_u8 + 8;
   let v2 = i8::checked_add(251, 8).unwrap();
   println!("{},{}",v1,v2);
}
```

6. 🌟🌟
```rust,editable

// modify `assert!` to make it work
fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1579);
}
```


### Floating-Point
7. 🌟

```rust,editable

//  replace ? with your answer
fn main() {
    let x = 1_000.000_1; // ?
    let y: f32 = 0.12; // f32
    let z = 0.01_f64; // f64
}
```

1. 🌟🌟 make it work in two distinct ways

```rust,editable

fn main() {
    assert!(0.1+0.2==0.3);
}
```

### Range
9. 🌟🌟 two goals: 1. modify `assert!` to make it work 2. make `println!` output: 97 - 122

```rust,editable
fn main() {
    let mut sum = 0;
    for i in -3..2 {
        sum += i
    }

    assert!(sum == -3);

    for c in 'a'..='z' {
        println!("{}",c);
    }
}
```

10. 🌟🌟 
```rust,editable

// fill the blanks
use std::ops::{Range, RangeInclusive};
fn main() {
    assert_eq!((1..__), Range{ start: 1, end: 5 });
    assert_eq!((1..__), RangeInclusive::new(1, 5));
}
```

### Computations

11. 🌟 
```rust,editable

// fill the blanks and fix the errors
fn main() {
    // Integer addition
    assert!(1u32 + 2 == __);

    // Integer subtraction
    assert!(1i32 - 2 == __);
    assert!(1u8 - 2 == -1); 
    
    assert!(3 * 50 == __);

    assert!(9.6 / 3.2 == 3.0); // error ! make it work

    assert!(24 % 5 == __);
    // Short-circuiting boolean logic
    assert!(true && false == __);
    assert!(true || false == __);
    assert!(!true == __);

    // Bitwise operations
    println!("0011 AND 0101 is {:04b}", 0b0011u32 & 0b0101);
    println!("0011 OR 0101 is {:04b}", 0b0011u32 | 0b0101);
    println!("0011 XOR 0101 is {:04b}", 0b0011u32 ^ 0b0101);
    println!("1 << 5 is {}", 1u32 << 5);
    println!("0x80 >> 2 is 0x{:x}", 0x80u32 >> 2);
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it