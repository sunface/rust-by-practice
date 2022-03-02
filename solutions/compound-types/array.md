1.
```rust
fn main() {
    let arr: [i32; 5] = [1, 2, 3, 4, 5];

    assert!(arr.len() == 5);
}
```

2.
```rust
fn main() {
    // we can ignore parts of the array type or even the whole type, let the compiler infer it for us
    let arr0 = [1, 2, 3];
    let arr: [_; 3] = ['a', 'b', 'c'];
    
    // Arrays are stack allocated, `std::mem::size_of_val` return the bytes which array occupies
    // A char takes 4 byte in Rust: Unicode char
    assert!(std::mem::size_of_val(&arr) == 12);
}
```

3.
```rust
fn main() {
    let list: [i32; 100] = [1; 100];

    assert!(list[0] == 1);
    assert!(list.len() == 100);
}
```

4.
```rust
fn main() {
    // fix the error
    let _arr = [1, 2, 3];
}
```

5.
```rust
fn main() {
    let arr = ['a', 'b', 'c'];
    
    let ele = arr[0];

    assert!(ele == 'a');
}
```

6.
```rust
fn main() {
    let names = [String::from("Sunfei"), "Sunface".to_string()];
    
    // `get` returns an Option<T>, it's safe to use
    let name0 = names.get(0).unwrap();

    // but indexing is not safe
    let _name1 = &names[1];
}
```