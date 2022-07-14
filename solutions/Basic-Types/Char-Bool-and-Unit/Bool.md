# 3. 🌟
```rs
// Make println! work
fn main() {
    let _f: bool = false;

    let t = true;
    if !t {
        println!("Success!");
    }
} 
```
# Solution:
```rs
fn main() {
    let _f: bool = false;

    let t = _f;
    if !t {
        println!("Success!");
    }
} 
```
or
```rs
fn main() {
    let _f: bool = false;

    let t = false;
    if !t {
        println!("Success!");
    }
} 
```
# 4. 🌟
```rs

// Make it work
fn main() {
    let f = true;
    let t = true && false;
    assert_eq!(t, f);

    println!("Success!");
}
```
# Solution:
```rs
fn main() {
    let f = true;
    let t = true || false;
    assert_eq!(t, f);

    println!("Success!");
}
```
or
```rs
fn main() {
    let f = true;
    let t = true && true;
    assert_eq!(t, f);

    println!("Success!");
}
```