# 12. 🌟
```rs
fn main() {
    // Fill the blank to print each char in "你好，世界"
    for c in "你好，世界".__ {
        println!("{}", c)
    }
}
```
# Solution:
```rs
fn main() {
    // Fill the blank to print each char in "你好，世界"
    for c in "你好，世界".chars() {
        println!("{}", c)
    }
}
```
# utf8_slice
## You can use utf8_slice to slice UTF8 string, it can index chars instead of bytes.

# Example
## use utf8_slice;
```rs
fn main() {
    let s = "The 🚀 goes to the 🌑!";

    let rocket = utf8_slice::slice(s, 4, 5);
    // Will equal "🚀"
}
```
