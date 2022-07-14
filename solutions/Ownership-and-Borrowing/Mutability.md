# Mutability can be changed when ownership is transferred.
# 6. 🌟
```rs
fn main() {
    let s = String::from("hello, ");
    
    // Modify this line only !
    let s1 = s;

    s1.push_str("world");

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let s = String::from("hello, ");
    
    // Modify this line only !
    let mut s1 = s;

    s1.push_str("world");

    println!("Success!");
}
```
# 7. 🌟🌟🌟
```rs
fn main() {
    let x = Box::new(5);
    
    let ...      // Implement this line, dont change other lines!
    
    *y = 4;
    
    assert_eq!(*x, 5);

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let x = Box::new(5);
    
    let mut y = x.clone();      // Implement this line, dont change other lines!
    
    *y = 4;
    
    assert_eq!(*x, 5);

    println!("Success!");
}
```