# Array
The type of array is `[T; Lengh]`, as you can see, array's lengh is part of their type signature. So their length must be known at compile time.

For example, you cant initialized an array as below:
```rust
fn init_arr(n: i32) {
    let arr = [1; n];
}
```

This will cause an error, because the compile have no idea of the exact size of the array in compile time.

🌟 
```rust,editable

fn main() {
    // fill the blank with proper array type
    let arr: __ = [1, 2, 3, 4, 5];

    // modify below to make it work
    assert!(arr.len() == 4);
}
```

🌟🌟
```rust,editable

fn main() {
    // we can ignore parts of the array type or even the whole type, let the compiler infer it for us
    let arr0 = [1, 2, 3];
    let arr: [_; 3] = ['a', 'b', 'c'];
    
    // fill the blank
    // Arrays are stack allocated, `std::mem::size_of_val` return the bytes which array occupies
    // A char takes 4 byte in Rust: Unicode char
    assert!(std::mem::size_of_val(&arr) == __);
}
```

🌟 All elements in an array can be initialized to the same value at once.

```rust,editable

fn main() {
    // fill the blank
    let list: [i32; 100] = __ ;

    assert!(list[0] == 1);
    assert!(list.len() == 100);
}
```

🌟 All elements in an array must be of the same type
```rust,editable

fn main() {
    // fix the error
    let _arr = [1, 2, '3'];
}
```

🌟 Indexing starts at 0.
```rust,editable

fn main() {
    let arr = ['a', 'b', 'c'];
    
    let ele = arr[1]; // only modify this line to make the code work!

    assert!(ele == 'a');
}
```

🌟 Out of bounds indexing causes `panic`.
```rust,editable

// fix the error
fn main() {
    let names = [String::from("Sunfei"), "Sunface".to_string()];
    
    // `get` returns an Option<T>, it's safe to use
    let name0 = names.get(0).unwrap();

    // but indexing is not safe
    let _name1 = &names[2];
}

```

