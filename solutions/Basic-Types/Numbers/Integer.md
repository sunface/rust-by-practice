# 1. 🌟

```
Tips: If we don't explicitly assign a type to a variable, then the compiler will infer one for us.
```
```rs

// Remove something to make it work
fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    y = x;
    
    let z = 10; // Type of z ? 

    println!("Success!");
}
```
# Solution
```rs
fn main() {
    let x: i32 = 5;
    let mut y: u32 = 5;

    // y = x;
    
    let z = 10; // Type of z ? 

    println!("Success!");
}
```

# 2. 🌟
```rs

//  Fill the blank
fn main() {
    let v: u16 = 38_u8 as __;

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let v: u16 = 38_u8 as u16;

    println!("Success!");
}
```
# 3. 🌟🌟🌟
```
Tips: If we don't explicitly assign a type to a variable, then the compiler will infer one for us.
```
```rs
// Modify `assert_eq!` to make it work
fn main() {
    let x = 5;
    assert_eq!("u32".to_string(), type_of(&x));

    println!("Success!");
}

// Get the type of given variable, return a string representation of the type  , e.g "i8", "u8", "i32", "u32"
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```
# Solution:
```rs
fn main() {
    let x = 5;
    assert_eq!("i32".to_string(), type_of(&x));

    println!("Success!");
}

// Get the type of given variable, return a string representation of the type  , e.g "i8", "u8", "i32", "u32"
fn type_of<T>(_: &T) -> String {
    format!("{}", std::any::type_name::<T>())
}
```
# 4. 🌟🌟
```rs
// Fill the blanks to make it work
fn main() {
    assert_eq!(i8::MAX, __); 
    assert_eq!(u8::MAX, __); 

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    assert_eq!(i8::MAX, 127); 
    assert_eq!(u8::MAX, 255); 

    println!("Success!");
}
```
# 5. 🌟🌟
```rs
// Fix errors and panics to make it work
fn main() {
   let v1 = 251_u8 + 8;
   let v2 = i8::checked_add(251, 8).unwrap();
   println!("{},{}",v1,v2);
}
```
# Solution:
```rs
fn main() {
   let v1 = 247_u8 + 8;
   let v2 = i8::checked_add(119, 8).unwrap();
   println!("{},{}",v1,v2);
}
```
# 6. 🌟🌟
```rs
// Modify `assert!` to make it work
fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1579);

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let v = 1_024 + 0xff + 0o77 + 0b1111_1111;
    assert!(v == 1597);

    println!("Success!");
}
```
