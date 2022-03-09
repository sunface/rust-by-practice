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

3. 🌟🌟
```rust,editable

// Question: how many heap allocations are happend here ?
// Your answer: 
fn main() {  
    // Create a String type based on `&str`
    // the type of string literals is `&str`
   let s: String = String::from("hello, world!");

   // create a slice point to String `s`
   let slice: &str = &s;

   // create a String type based on the recently created slice
   let s: String = slice.to_string();

   assert_eq!(s, "hello, world!");

   println!("Success!")
}
```

### UTF-8 & Indexing
Strings are always valid UTF-8. This has a few implications:

- the first of which is that if you need a non-UTF-8 string, consider [OsString](https://doc.rust-lang.org/stable/std/ffi/struct.OsString.html). It is similar, but without the UTF-8 constraint. 
- The second implication is that you cannot index into a String

Indexing is intended to be a constant-time operation, but UTF-8 encoding does not allow us to do this. Furthermore, it’s not clear what sort of thing the index should return: a byte, a codepoint, or a grapheme cluster. The bytes and chars methods return iterators over the first two, respectively.

4. 🌟🌟🌟 You can't use index to access a char in a string, but you can use slice `&s1[start..end]`.

```rust,editable

// FILL in the blank and FIX errors
fn main() {
    let s = String::from("hello, 世界");
    let slice1 = s[0]; //tips: `h` only takes 1 byte in UTF8 format
    assert_eq!(slice1, "h");

    let slice2 = &s[3..5];// tips: `中`  takes 3 bytes in UTF8 format
    assert_eq!(slice2, "世");
    
    // iterate all chars in s
    for (i, c) in s.__ {
        if i == 7 {
            assert_eq!(c, '世')
        }
    }

    println!("Success!")
}
```


#### utf8_slice
You can use [utf8_slice](https://docs.rs/utf8_slice/1.0.0/utf8_slice/fn.slice.html) to slice UTF8 string, it can index chars instead of bytes.

**Example**
```rust
use utf8_slice;
fn main() {
   let s = "The 🚀 goes to the 🌑!";

   let rocket = utf8_slice::slice(s, 4, 5);
   // Will equal "🚀"
}
```


5. 🌟🌟🌟
> Tips: maybe you need `from_utf8` method

```rust,editable

// FILL in the blanks
fn main() {
    let mut s = String::new();
    __;

    // some bytes, in a vector
    let v = vec![104, 101, 108, 108, 111];

    // Turn a bytes vector into a String
    let s1 = __;
    
    
    assert_eq!(s, s1);

    println!("Success!")
}
```

### Representation
A String is made up of three components: a pointer to some bytes, a length, and a capacity. 

The pointer points to an internal buffer String uses to store its data. The length is the number of bytes currently stored in the buffer( always stored on the heap ), and the capacity is the size of the buffer in bytes. As such, the length will always be less than or equal to the capacity.

6. 🌟🌟 If a String has enough capacity, adding elements to it will not re-allocate
```rust,editable

// modify the code below to print out: 
// 25
// 25
// 25
// Here, there’s no need to allocate more memory inside the loop.
fn main() {
    let mut s = String::new();

    println!("{}", s.capacity());

    for _ in 0..2 {
        s.push_str("hello");
        println!("{}", s.capacity());
    }

    println!("Success!")
}
```

7. 🌟🌟🌟
```rust,editable

// FILL in the blanks
use std::mem;

fn main() {
    let story = String::from("Rust By Practice");

    // Prevent automatically dropping the String's data
    let mut story = mem::ManuallyDrop::new(story);

    let ptr = story.__();
    let len = story.__();
    let capacity = story.__();

    assert_eq!(16, len);

    // We can re-build a String out of ptr, len, and capacity. This is all
    // unsafe because we are responsible for making sure the components are
    // valid:
    let s = unsafe { String::from_raw_parts(ptr, len, capacity) };

    assert_eq!(*story, s);

    println!("Success!")
}
```


### Common methods
More exercises of String methods can be found [here](../std/String.md).

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it