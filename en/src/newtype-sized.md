# newtype and Sized

## Newtype
The orphan rule tells us that we are allowed to implement a trait on a type as long as either the trait or the type are local to our crate.

The **newtype pattern** can help us get around this restriction, which involves creating a new type in a **tuple struct**.

1、🌟
```rust,editable
use std::fmt;

/* Define the Wrapper type */
__;

// Display is an external trait
impl fmt::Display for Wrapper {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "[{}]", self.0.join(", "))
    }
}

fn main() {
    // Vec is an external type, so you cannot implement Display trait on Vec type
    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w = {}", w);
}
```

2、🌟 Hide the methods of the orginal type
```rust,editable
/* Make it workd */
struct Meters(u32);

fn main() {
    let i: u32 = 2;
    assert_eq!(i.pow(2), 4);

    let n = Meters(i);
    // The `pow` method is defined on `u32` type, we can't directly call it 
    assert_eq!(n.pow(2), 4);
}
```

3、🌟🌟 The `newtype` idiom gives compile time guarantees that the right type of value is suplied to a program.
```rust,editable
/* Make it work */
struct Years(i64);

struct Days(i64);

impl Years {
    pub fn to_days(&self) -> Days {
        Days(self.0 * 365)
    }
}


impl Days {
    pub fn to_years(&self) -> Years {
        Years(self.0 / 365)
    }
}

// an age verification function that checks age in years, must be given a value of type Years.
fn old_enough(age: &Years) -> bool {
    age.0 >= 18
}

fn main() {
    let age = Years(5);
    let age_days = age.to_days();
    println!("Old enough {}", old_enough(&age));
    println!("Old enough {}", old_enough(&age_days));
}
```

4、🌟🌟
```rust,editable
use std::ops::Add;
use std::fmt::{self, format};

struct Meters(u32);
impl fmt::Display for Meters {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "There are still {} meters left", self.0)
    }
}

impl Add for Meters {
    type Output = Self;

    fn add(self, other: Meters) -> Self {
        Self(self.0 + other.0)
    }
}
fn main() {
    let d = calculate_distance(Meters(10), Meters(20));
    assert_eq!(format!("{}",d), "There are still 30 meters left");
}

/* implement calculate_distance  */
fn calculate_distance
```

## Type alias
The most importance of type alias is to improve the readability of our codes.

```rust
type Thunk = Box<dyn Fn() + Send + 'static>;

let f: Thunk = Box::new(|| println!("hi"));

fn takes_long_type(f: Thunk) {
    // --snip--
}

fn returns_long_type() -> Thunk {
    // --snip--
}
```

```rust
type Result<T> = std::result::Result<T, std::io::Error>;
```

And Unlike newtype, type alias don't create new types, so the following code is valid:
```rust
type Meters = u32;

let x: u32 = 5;
let y: Meters = 5;

println!("x + y = {}", x + y);
```

5、🌟
```rust,editable
enum VeryVerboseEnumOfThingsToDoWithNumbers {
    Add,
    Subtract,
}

/* Fill in the blank */
__

fn main() {
    // We can refer to each variant via its alias, not its long and inconvenient
    // name.
    let x = Operations::Add;
}
```

6、🌟🌟 There are a few preserved alias in Rust, one of which can be used in `impl` blocks.
```rust,editable
enum VeryVerboseEnumOfThingsToDoWithNumbers {
    Add,
    Subtract,
}

impl VeryVerboseEnumOfThingsToDoWithNumbers {
    fn run(&self, x: i32, y: i32) -> i32 {
        match self {
            __::Add => x + y,
            __::Subtract => x - y,
        }
    }
}
```

## DST and unsized type
These concepts are complicated, so we are not going to explain here, but you can find them in [The Book](https://doc.rust-lang.org/book/ch19-04-advanced-types.html?highlight=DST#dynamically-sized-types-and-the-sized-trait).

7、🌟🌟🌟 Array with dynamic length is a Dynamic Sized Type ( DST ), we can't directly use it
```rust,editable
/* Make it work with const generics */
fn my_function(n: usize) -> [u32; usize] {
    [123; n]
}

fn main() {
    let arr = my_function();
    println!("{:?}",arr);
}
```

8、🌟🌟 Slice is unsized type, but the reference of slice is not.
```rust,editable
/* Make it work with slice references */
fn main() {
    let s: str = "Hello there!";

    let arr: [u8] = [1, 2, 3];
}
```

9、🌟🌟 Trait is also a unsized type
```rust,editable
/* Make it work in two ways */
use std::fmt::Display;
fn foobar(thing: Display) {}    

fn main() {
}
```