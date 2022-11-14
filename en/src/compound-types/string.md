# String
The type of string literal `"hello, world"` is `&str`, e.g `let s: &str = "hello, world"`.


### Str and &str
1. 🌟 We can't use `str` type in normal ways, but we can use `&str`.

```rust,editable

// Fix error without adding new line
fn main() {
    let s: str = "hello, world";

    println!("Success!");
}
```


2. 🌟🌟 We can only use `str` by boxed it, `&` can be used to convert `Box<str>` to `&str` 

```rust,editable

// Fix the error with at least two solutions
fn main() {
    let s: Box<str> = "hello, world".into();
    greetings(s)
}

fn greetings(s: &str) {
    println!("{}",s)
}
```

### String
`String` type is defined in std and stored as a vector of bytes (Vec<u8>), but guaranteed to always be a valid UTF-8 sequence. String is heap allocated, growable and not null terminated.

3. 🌟
```rust,editable

// Fill the blank
fn main() {
    let mut s = __;
    s.push_str("hello, world");
    s.push('!');

    assert_eq!(s, "hello, world!");

    println!("Success!");
}
```

4. 🌟🌟🌟
```rust,editable

// Fix all errors without adding newline
fn main() {
    let s = String::from("hello");
    s.push(',');
    s.push(" world");
    s += "!".to_string();

    println!("{}", s);
}
```

5. 🌟🌟 `replace` can be used to replace substring
```rust,editable

// Fill the blank
fn main() {
    let s = String::from("I like dogs");
    // Allocate new memory and store the modified string there
    let s1 = s.__("dogs", "cats");

    assert_eq!(s1, "I like cats");

    println!("Success!");
}
```

More `String` methods can be found under [String](https://doc.rust-lang.org/std/string/struct.String.html) module.

6. 🌟🌟 You can only concat a `String` with `&str`, and `String`'s ownership can be moved to another variable.

```rust,editable

// Fix errors without removing any line
fn main() {
    let s1 = String::from("hello,");
    let s2 = String::from("world!");
    let s3 = s1 + s2; 
    assert_eq!(s3, "hello,world!");
    println!("{}", s1);
}
```

### &str and String
Opposite to the seldom using of `str`, `&str` and `String` are used everywhere!

7. 🌟🌟 `&str` can be converted to `String` in two ways
```rust,editable

// Fix error with at least two solutions
fn main() {
    let s = "hello, world";
    greetings(s)
}

fn greetings(s: String) {
    println!("{}", s)
}
```

8. 🌟🌟 We can use `String::from` or `to_string` to convert a `&str` to `String`

```rust,editable

// Use two approaches to fix the error and without adding a new line
fn main() {
    let s = "hello, world".to_string();
    let s1: &str = s;

    println!("Success!");
}
```

### String escapes
9. 🌟 
```rust,editable
fn main() {
    // You can use escapes to write bytes by their hexadecimal values
    // Fill the blank below to show "I'm writing Rust"
    let byte_escape = "I'm writing Ru\x73__!";
    println!("What are you doing\x3F (\\x3F means ?) {}", byte_escape);

    // ...Or Unicode code points.
    let unicode_codepoint = "\u{211D}";
    let character_name = "\"DOUBLE-STRUCK CAPITAL R\"";

    println!("Unicode character {} (U+211D) is called {}",
                unicode_codepoint, character_name );

    let long_string = "String literals
                        can span multiple lines.
                        The linebreak and indentation here \
                         can be escaped too!";
    println!("{}", long_string);
}
```

10. 🌟🌟🌟 Sometimes there are just too many characters that need to be escaped or it's just much more convenient to write a string out as-is. This is where raw string literals come into play.

```rust,editable

/* Fill in the blank and fix the errors */
fn main() {
    let raw_str = r"Escapes don't work here: \x3F \u{211D}";
    assert_eq!(raw_str, "Escapes don't work here: ? ℝ");

    // If you need quotes in a raw string, add a pair of #s
    let quotes = r#"And then I said: "There is no escape!""#;
    println!("{}", quotes);

    // If you need "# in your string, just use more #s in the delimiter.
    // You can use up to 65535 #s.
    let delimiter = r###"A string with "# in it. And even "##!"###;
    println!("{}", delimiter);

    let long_delimiter = __;
    assert_eq!(long_delimiter, "Hello, \"##\"");

    println!("Success!");
}
```

### Byte string
Want a string that's not UTF-8? (Remember, str and String must be valid UTF-8). Or maybe you want an array of bytes that's mostly text? Byte strings to the rescue!

**Example**:
```rust,editable
use std::str;

fn main() {
    // Note that this is not actually a `&str`
    let bytestring: &[u8; 21] = b"this is a byte string";

    // Byte arrays don't have the `Display` trait, so printing them is a bit limited
    println!("A byte string: {:?}", bytestring);

    // Byte strings can have byte escapes...
    let escaped = b"\x52\x75\x73\x74 as bytes";
    // ...But no unicode escapes
    // let escaped = b"\u{211D} Is not allowed";
    println!("Some escaped bytes: {:?}", escaped);


    // Raw byte strings work just like raw strings
    let raw_bytestring = br"\u{211D} is not escaped here";
    println!("{:?}", raw_bytestring);

    // Converting a byte array to `str` can fail
    if let Ok(my_str) = str::from_utf8(raw_bytestring) {
        println!("And the same as text: '{}'", my_str);
    }

    let _quotes = br#"You can also use "fancier" formatting, \
                    like with normal raw strings"#;

    // Byte strings don't have to be UTF-8
    let shift_jis = b"\x82\xe6\x82\xa8\x82\xb1\x82\xbb"; // "ようこそ" In SHIFT-JIS

    // But then they can't always be converted to `str`
    match str::from_utf8(shift_jis) {
        Ok(my_str) => println!("Conversion successful: '{}'", my_str),
        Err(e) => println!("Conversion failed: {:?}", e),
    };
}
```

A more detailed listing of the ways to write string literals and escape characters is given in the ['Tokens' chapter](https://doc.rust-lang.org/reference/tokens.html) of the Rust Reference.

### String index
11. 🌟🌟🌟 You can't use index to access a char in a string, but you can use slice `&s1[start..end]`.

```rust,editable

fn main() {
    let s1 = String::from("hi,中国");
    let h = s1[0]; // Modify this line to fix the error, tips: `h` only takes 1 byte in UTF8 format
    assert_eq!(h, "h");

    let h1 = &s1[3..5]; // Modify this line to fix the error, tips: `中`  takes 3 bytes in UTF8 format
    assert_eq!(h1, "中");

    println!("Success!");
}
```

### Operate on UTF8 string
12. 🌟
```rust,editable

fn main() {
    // Fill the blank to print each char in "你好，世界"
    for c in "你好，世界".__ {
        println!("{}", c)
    }
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

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it
