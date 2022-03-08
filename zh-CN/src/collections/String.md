# String
`std::string::String` 是 UTF-8 编码、可增长的动态字符串. 它也是我们日常开发中最常用的字符串类型，同时对于它所拥有的内容拥有所有权。

### 基本操作
1. 🌟🌟
```rust,editable

// 填空并修复错误
// 1. 不要使用 `to_string()`
// 2. 不要添加/删除任何代码行
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
虽然 `String` 的底层是 `Vec<u8>` 也就是字节数组的形式存储的，但是它是基于 UTF-8 编码的字符序列。`String` 分配在堆上、可增长且不是以 `null` 结尾。

而 `&str` 是[切片引用](https://course.rs/confonding/slice.html)类型( `&[u8]` )，指向一个合法的 UTF-8 字符序列，总之，`&str` 和 `String` 的关系类似于 `&[T]` 和 `Vec<T>` 。

如果大家想了解更多，可以看看[易混淆概念解析 - &str 和 String](https://course.rs/confonding/string.html)。


2. 🌟🌟
```rust,editable
// 填空
fn main() {  
   let mut s = String::from("hello, world");

   let slice1: &str = __; // 使用两种方法
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
use utf_slice;
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
    // We know these bytes are valid, so we'll use `unwrap()`.
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

    // story has nineteen bytes
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