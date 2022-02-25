# 函数
🌟🌟🌟
```rust,editable

fn main() {
    // 不要修改下面两行代码!
    let (x, y) = (1, 2);
    let s = sum(1, 2);

    assert_eq!(s, 3);
}

fn sum(x, y: i32) {
    x + y;
}
```


🌟🌟
```rust,editable
fn main() {
   print();
}

// 使用另一个类型来替代 i32
fn print() -> i32 {
   println!("hello,world");
}
```


🌟🌟

```rust,editable
fn main() {
    never_return();
}

fn never_return() -> ! {
    // 实现这个函数，不要修改函数签名!
    
}
```