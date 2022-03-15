1.
```rust
#[derive(Debug)]
struct Structure(i32);

fn main() {
    // Types in std and Rust have implemented the fmt::Debug trait
    println!("{:?} months in a year.", 12);

    println!("Now {:?} will print!", Structure(3));
}
```
2.
```rust
#[derive(Debug)]
struct Person {
    name: String,
    age: u8
}

fn main() {
    let person = Person { name:  "Sunface".to_string(), age: 18 };

    println!("{:#?}", person);
}
```

3.
```rust
use std::fmt;

struct Structure(i32);

struct Deep(Structure);
impl fmt::Debug for Deep {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:?}", self.0.0)
    }
}

fn main() {    
    // The problem with `derive` is there is no control over how
    // the results look. What if I want this to just show a `7`?

    /* Make it output: Now 7 will print! */
    println!("Now {:?} will print!", Deep(Structure(7)));
}
```

4
```rust
use std::fmt;

struct Point2D {
    x: f64,
    y: f64,
}

impl fmt::Display for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Display: {} + {}i", self.x, self.y)
    }
}

impl fmt::Debug for Point2D {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "Debug: Complex {{ real: {:?}, imag: {:?} }}", self.x, self.y)
    }
}

fn main() {

    let point = Point2D { x: 3.3, y: 7.2 };
    println!("{}", point);

    println!("{:?}", point);
}
```