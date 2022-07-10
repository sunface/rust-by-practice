# Others

### Convert any type to String
To convert any type to `String`, you can simply the `ToString` trait for that type. Rather than doing that directly, you should implement the `fmt::Display` trait which will automatically provides `ToString` and also allows you to print the type with `println!`.

1. 🌟🌟
```rust,editable
use std::fmt;

struct Point {
    x: i32,
    y: i32,
}

impl fmt::Display for Point {
    // IMPLEMENT fmt method
}

fn main() {
    let origin = Point { x: 0, y: 0 };
    // FILL in the blanks
    assert_eq!(origin.__, "The point is (0, 0)");
    assert_eq!(format!(__), "The point is (0, 0)");

    println!("Success!")
}
```

### Parse a String
2. 🌟🌟🌟 We can use `parse` method to convert a `String` into a `i32` number, this is becuase `FromStr` is implemented for `i32` type in standard library: `impl FromStr for i32`
```rust,editable
// To use `from_str` method, you needs to introduce this trait into the current scope.
use std::str::FromStr;
fn main() {
    let parsed: i32 = "5".__.unwrap();
    let turbo_parsed = "10".__.unwrap();
    let from_str = __.unwrap();
    let sum = parsed + turbo_parsed + from_str;
    assert_eq!(sum, 35);

    println!("Success!")
}
``` 


3. 🌟🌟 We can also implement the `FromStr` trait for our custom types
```rust,editable
use std::str::FromStr;
use std::num::ParseIntError;

#[derive(Debug, PartialEq)]
struct Point {
    x: i32,
    y: i32
}

impl FromStr for Point {
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let coords: Vec<&str> = s.trim_matches(|p| p == '(' || p == ')' )
                                 .split(',')
                                 .collect();

        let x_fromstr = coords[0].parse::<i32>()?;
        let y_fromstr = coords[1].parse::<i32>()?;

        Ok(Point { x: x_fromstr, y: y_fromstr })
    }
}
fn main() {
    // FILL in the blanks in two ways
    // DON'T change code anywhere else 
    let p = __;
    assert_eq!(p.unwrap(), Point{ x: 3, y: 4} );

    println!("Success!")
}
```

### Deref
You can find all the examples and exercises of the `Deref` trait [here](https://practice.rs/smart-pointers/deref.html).

### transmute
`std::mem::transmute` is a **unsafe function** can be used to reinterprets the bits of a value of one type as another type. Both of the original and the result types must have the same size and neither of them can be invalid.

`transmute` is semantically equivalent to a bitwise move of one type into another. It copies the bits from the source value into the destination value, then forgets the original, seems equivalent to C's `memcpy` under the hood.

So, **`transmute` is incredibly unsafe !** The caller has to ensure all the safes himself!

#### Examples
1. `transmute` can be used to turn a  pointer into a function pointer, this is not portable on machines where function pointer and data pointer have different sizes.

```rust,editable
fn foo() -> i32 {
    0
}

fn main() {
    let pointer = foo as *const ();
    let function = unsafe {
        std::mem::transmute::<*const (), fn() -> i32>(pointer)
    };
    assert_eq!(function(), 0);
}
```

2. Extending a lifetime or shortening the lifetime of an invariant is an advanced usage of `transmute`, yeah, **very unsafe Rust!**.
```rust,editable
struct R<'a>(&'a i32);
unsafe fn extend_lifetime<'b>(r: R<'b>) -> R<'static> {
    std::mem::transmute::<R<'b>, R<'static>>(r)
}

unsafe fn shorten_invariant_lifetime<'b, 'c>(r: &'b mut R<'static>)
                                             -> &'b mut R<'c> {
    std::mem::transmute::<&'b mut R<'static>, &'b mut R<'c>>(r)
}
```

3. Rather than using `transmute`, you can use some alternatives instead.
```rust,editable
fn main() {
    /*Turning raw bytes(&[u8]) to u32, f64, etc.: */
    let raw_bytes = [0x78, 0x56, 0x34, 0x12];

    let num = unsafe { std::mem::transmute::<[u8; 4], u32>(raw_bytes) };

    // use `u32::from_ne_bytes` instead
    let num = u32::from_ne_bytes(raw_bytes);
    // or use `u32::from_le_bytes` or `u32::from_be_bytes` to specify the endianness
    let num = u32::from_le_bytes(raw_bytes);
    assert_eq!(num, 0x12345678);
    let num = u32::from_be_bytes(raw_bytes);
    assert_eq!(num, 0x78563412);

    /*Turning a pointer into a usize: */
    let ptr = &0;
    let ptr_num_transmute = unsafe { std::mem::transmute::<&i32, usize>(ptr) };

    // Use an `as` cast instead
    let ptr_num_cast = ptr as *const i32 as usize;

    /*Turning an &mut T into an &mut U: */
    let ptr = &mut 0;
    let val_transmuted = unsafe { std::mem::transmute::<&mut i32, &mut u32>(ptr) };

    // Now, put together `as` and reborrowing - note the chaining of `as`
    // `as` is not transitive
    let val_casts = unsafe { &mut *(ptr as *mut i32 as *mut u32) };

    /*Turning an &str into a &[u8]: */
    // this is not a good way to do this.
    let slice = unsafe { std::mem::transmute::<&str, &[u8]>("Rust") };
    assert_eq!(slice, &[82, 117, 115, 116]);

    // You could use `str::as_bytes`
    let slice = "Rust".as_bytes();
    assert_eq!(slice, &[82, 117, 115, 116]);

    // Or, just use a byte string, if you have control over the string
    // literal
    assert_eq!(b"Rust", &[82, 117, 115, 116]);
}
```