# From/Into
The `From` trait allows for a type to define how to create itself from another type, hence providing a very simple mechanism for converting between several types.

The `From` and `Into` traits are inherently linked, and this is actually part of its implementation. It means if we write something like this: `impl From<T> for U`, then we can use 
`let u: U = U::from(T)` or `let u:U = T.into()`.

The `Into` trait is simply the reciprocal of the `From` trait. That is, if you have implemented the `From` trait for your type, then the `Into` trait will be automatically implemented for the same type.

Using the `Into` trait will typically require the type annotations as the compiler is unable to determine this most of the time.

For example, we can easily convert `&str` into `String` :
```rust
fn main() {
    let my_str = "hello";

    // three conversions below all depends on the fact: String implements From<&str>:
    let string1 = String::from(my_str);
    let string2 = my_str.to_string();
    // Explicit type annotation is required here
    let string3: String = my_str.into();
}
```

Because the standard library has already implemented this for us : `impl From<&'_ str> for String` .

Some implementations of `From` trait can be found [here](https://doc.rust-lang.org/stable/std/convert/trait.From.html#implementors).

1. 🌟🌟🌟
```rust,editable

fn main() {
     // impl From<bool> for i32
    let i1:i32 = false.into();
    let i2:i32 = i32::from(false);  
    assert_eq!(i1, i2);
    assert_eq!(i1, 0);

    // FIX the error in two ways
    // 1. impl From<char> for ? , maybe you should check the docs mentiond above to find the answer
    // 2. a keyword from the last chapter
    let i3: i32 = 'a'.into();

    // FIX the error in two ways
    let s: String = 'a' as String;

    println!("Success!");
}
```

### Implement `From` for custom types
2. 🌟🌟
```rust,editable

// From is now included in `std::prelude`, so there is no need to introduce it into the current scope
// use std::convert::From;

#[derive(Debug)]
struct Number {
    value: i32,
}

impl From<i32> for Number {
    // IMPLEMENT `from` method
}

// FILL in the blanks
fn main() {
    let num = __(30);
    assert_eq!(num.value, 30);

    let num: Number = __;
    assert_eq!(num.value, 30);

    println!("Success!");
}
```

3. 🌟🌟🌟 When performing error handling it is often useful to implement `From` trait for our own error type. Then we can use `?` to automatically convert the underlying error type to our own error type.
```rust,editable

use std::fs;
use std::io;
use std::num;

enum CliError {
    IoError(io::Error),
    ParseError(num::ParseIntError),
}

impl From<io::Error> for CliError {
    // IMPLEMENT from method
}

impl From<num::ParseIntError> for CliError {
    // IMPLEMENT from method
}

fn open_and_parse_file(file_name: &str) -> Result<i32, CliError> {
    // ? automatically converts io::Error to CliError
    let contents = fs::read_to_string(&file_name)?;
    // num::ParseIntError -> CliError
    let num: i32 = contents.trim().parse()?;
    Ok(num)
}

fn main() {
    println!("Success!");
}
```


### TryFrom/TryInto
Similar to `From` and `Into`, `TryFrom` and `TryInto` are generic traits for converting between types.

Unlike `From/Into`, `TryFrom` and `TryInto` are used for fallible conversions and return a `Result` instead of a plain value. 

4. 🌟🌟
```rust,editable
// TryFrom and TryInto are included in `std::prelude`, so there is no need to introduce it into the current scope
// use std::convert::TryInto;

fn main() {
    let n: i16 = 256;

    // Into trait has a method `into`,
    // hence TryInto has a method ?
    let n: u8 = match n.__() {
        Ok(n) => n,
        Err(e) => {
            println!("there is an error when converting: {:?}, but we catch it", e.to_string());
            0
        }
    };

    assert_eq!(n, __);

    println!("Success!");
}
```

5. 🌟🌟🌟
```rust,editable
#[derive(Debug, PartialEq)]
struct EvenNum(i32);

impl TryFrom<i32> for EvenNum {
    type Error = ();

    // IMPLEMENT `try_from`
    fn try_from(value: i32) -> Result<Self, Self::Error> {
        if value % 2 == 0 {
            Ok(EvenNum(value))
        } else {
            Err(())
        }
    }
}

fn main() {
    assert_eq!(EvenNum::try_from(8), Ok(EvenNum(8)));
    assert_eq!(EvenNum::try_from(5), Err(()));

    // FILL in the blanks
    let result: Result<EvenNum, ()> = 8i32.try_into();
    assert_eq!(result, __);
    let result: Result<EvenNum, ()> = 5i32.try_into();
    assert_eq!(result, __);

    println!("Success!");
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/type-conversions/from-into.md)(under the solutions path), but only use it when you need it