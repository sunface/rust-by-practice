# String
`std::string::String` is a UTF-8 encoded, growable string. It is the most common string type we used in daily dev, it also has ownership over the string contents.

### Basic operations
1. 🌟🌟
```rust,editable

// FILL in the blanks and FIX errors
// 1. Don't use `to_string()`
// 2. Dont't add/remove any code line
fn main() {
    let mut s: String = "hello, ";
    s.push_str("world".to_string());
    s.push(__);

    move_ownership(s);

    assert_eq!(s, "hello, world!");

    println!("Success!")
}

fn move_ownership(s: String) {
    println!("ownership of \"{}\" is moved here!", s)
}
```

### String and &str
A `String` is stored as a vector of bytes (`Vec<u8>`), but guaranteed to always be a valid UTF-8 sequence. `String` is heap allocated, growable and not null terminated.

`&str` is a slice (`&[u8]`) that always points to a valid UTF-8 sequence, and can be used to view into a String, just like `&[T]` is a view into `Vec<T>`.

2. 🌟🌟
```rust,editable
// FILL in the blanks
fn main() {  
   // get a slice of String with reference: String -> &str 
   let mut s = String::from("hello, world");

   let slice1: &str = __; // in two ways
   assert_eq!(slice1, "hello, world");

   let slice2 = __;
   assert_eq!(slice2, "hello");

   let slice3: __ = __; 
   slice3.push('!');
   assert_eq!(slice3, "hello, world!");

   println!("Success!")
}
```

### UTF-8 & Indexing
Strings are always valid UTF-8. This has a few implications:

- the first of which is that if you need a non-UTF-8 string, consider [OsString](https://doc.rust-lang.org/stable/std/ffi/struct.OsString.html). It is similar, but without the UTF-8 constraint. 
- The second implication is that you cannot index into a String

Indexing is intended to be a constant-time operation, but UTF-8 encoding does not allow us to do this. Furthermore, it’s not clear what sort of thing the index should return: a byte, a codepoint, or a grapheme cluster. The bytes and chars methods return iterators over the first two, respectively.

2. 🌟🌟🌟 You can't use index to access a char in a string, but you can use slice `&s1[start..end]`.

```rust,editable

// FIX errors
fn main() {
    let s = String::from("hello, 世界");
    let slice1 = s[0]; //tips: `h` only takes 1 byte in UTF8 format
    assert_eq!(slice1, "h");

    let slice2 = &s[3..5];// tips: `中`  takes 3 bytes in UTF8 format
    assert_eq!(slice2, "世");

    println!("Success!")
}
```


3. 🌟🌟🌟
> Tips: maybe you need `from_utf8` method

```rust,editable

// FILL in the blanks
fn main() {
    let mut s = String::new();
    __;

    // some bytes, in a vector
    let v = vec![104, 101, 108, 108, 111];

    // Turn a bytes vector into a String
    // We know these bytes are valid, so we'll use `unwrap()`.
    let s1 = __;
    
    
    assert_eq!(s, s1);

    println!("Success!")
}
```

### Representation





### Common methods

