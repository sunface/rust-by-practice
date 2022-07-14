# Slice
Slices are similar to arrays, but their length is not known at compile time, so you can't use slice directly.

1. 🌟🌟 Here, both `[i32]` and `str` are slice types, but directly using it will cause errors. You have to use the reference of the slice instead: `&[i32]`, `&str`.
```rust,editable

// Fix the errors, DON'T add new lines!
fn main() {
    let arr = [1, 2, 3];
    let s1: &[i32] = &arr[0..2];

    let s2: &str = "hello, world";

    println!("Success!");
}
```

A slice reference is a two-word object, for simplicity reasons, from now on we will use slice instead of `slice reference`.  The first word is a pointer to the data, and the second word is the length of the slice. The word size is the same as usize, determined by the processor architecture, eg 64 bits on an x86-64. Slices can be used to borrow a section of an array, and have the type signature `&[T]`.

2. 🌟🌟🌟
```rust,editable

fn main() {
    let arr: [char; 3] = ['中', '国', '人'];

    let slice = &arr[..2];
    
    // Modify '6' to make it work
    // TIPS: slice( reference ) IS NOT an array, if it is an array, then `assert!` will passed: Each of the two UTF-8 chars '中' and '国'  occupies 3 bytes, 2 * 3 = 6
    assert!(std::mem::size_of_val(&slice) == 16);

    println!("Success!");
}
```

3. 🌟🌟
```rust,editable

fn main() {
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    // Fill the blanks to make the code work
    let slice: &[i32] = &arr[1..4];
    assert_eq!(slice, &[2, 3, 4]);

    println!("Success!");
}
```

### String slices
4. 🌟 
```rust,editable

fn main() {
    let s = String::from("hello");

    let slice1 = &s[0..2];
    // Fill the blank to make the code work, DON'T USE 0..2 again
    let slice2 = &s[..2];

    assert_eq!(slice1, slice2);

    println!("Success!");
}
```

5. 🌟
```rust,editable

fn main() {
    let s = "你好，世界";
    // Modify this line to make the code work
    let slice = &s[0..3];

    assert!(slice == "你");

    println!("Success!");
}
```

6. 🌟🌟 `&String` can be implicitly converted into `&str`.
```rust,editable

// Fix errors
fn main() {
    let mut s = String::from("hello world");

    // Here, &s is `&String` type, but `first_word` need a `&str` type.
    // It works because `&String` can be implicitly converted to `&str, If you want know more ,this is called `Deref` 
    let word = first_word(&s);

    println!("the first word is: {}", word);
}
fn first_word(s: &str) -> &str {
    &s[..1]
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it