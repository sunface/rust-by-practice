# 切片
Slices are similar to arrays, but their length is not known at compile time, so you can't use slice directly.

🌟🌟 Here, both `[i32]` and `str` are slice types, but directly using it will cause errors. You have to use the reference of the slice instead: `&[i32]`, `&str`.
```rust,editable

// fix the errors, DON'T add new lines!
fn main() {
    let arr = [1, 2, 3];
    let s1: [i32] = arr[0..2];

    let s2: str = "hello, world" as str;
}
```

A slice reference is a two-word object, for simplicity reasons, from now on we will use slice instead of `slice reference`.  The first word is a pointer to the data, and the second word is the length of the slice. The word size is the same as usize, determined by the processor architecture eg 64 bits on an x86-64. Slices can be used to borrow a section of an array, and have the type signature `&[T]`.

🌟🌟🌟
```rust,editable

fn main() {
    let arr: [char; 3] = ['中', '国', '人'];

    let slice = &arr[..2];
    
    // modify '6' to make it work
    // TIPS: slice( reference ) IS NOT an array, if it is an array, then `assert!` will passed: each of the two UTF-8 chars '中' and '国'  occupies 3 bytes, 2 * 3 = 6
    assert!(std::mem::size_of_val(&slice) == 6);
}
```

🌟🌟
```rust,editable

fn main() {
  let arr: [i32; 5] = [1, 2, 3, 4, 5];
  // fill the blanks to make the code work
  let slice: __ = __;
  assert_eq!(slice, &[2, 3, 4]);
}
```

### string slices
🌟 
```rust,editable

fn main() {
    let s = String::from("hello");

    let slice1 = &s[0..2];
    // fill the blank to make the code work
    let slice2 = &s[__];

    assert_eq!(slice1, slice2);
}
```

🌟
```rust,editable

fn main() {
    let s = "你好，世界";
    // modify this line to make the code work
    let slice = &s[0..2];

    assert!(slice == "你");
}
```

🌟🌟 `&String` can be implicitly converted into `&str`.
```rust,editable

// fix errors
fn main() {
    let mut s = String::from("hello world");

    // here, &s is `&String` type, but `first_word` need a `&str` type.
    // it works because `&String` implicitly be converted to `&str, If you want know more ,this is called `Deref` 
    let word = first_word(&s);

    s.clear(); // error!

    println!("the first word is: {}", word);
}
fn first_word(s: &str) -> &str {
    &s[..1]
}
```
