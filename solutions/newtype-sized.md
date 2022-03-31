1、
```rust
use std::fmt;

struct Wrapper(Vec<String>);

impl fmt::Display for Wrapper {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "[{}]", self.0.join(", "))
    }
}

fn main() {
    let w = Wrapper(vec![String::from("hello"), String::from("world")]);
    println!("w = {}", w);
}
```

2、
```rust
struct Meters(u32);

fn main() {
    let i: u32 = 2;
    assert_eq!(i.pow(2), 4);

    let n = Meters(i);
    assert_eq!(n.0.pow(2), 4);
}
```

3、
```rust
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
    println!("Old enough {}", old_enough(&age_days.to_years()));
}
```

4、Sometimes `newtype` pattern can provide extra readability.
```rust
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
fn calculate_distance(d1: Meters, d2: Meters) -> Meters {
    d1 + d2
}
```

5、
```rust
enum VeryVerboseEnumOfThingsToDoWithNumbers {
    Add,
    Subtract,
}

/* Fill in the blank */
type Operations = VeryVerboseEnumOfThingsToDoWithNumbers;

fn main() {
    // We can refer to each variant via its alias, not its long and inconvenient
    // name.
    let x = Operations::Add;
}
```

6、
```rust
enum VeryVerboseEnumOfThingsToDoWithNumbers {
    Add,
    Subtract,
}

impl VeryVerboseEnumOfThingsToDoWithNumbers {
    fn run(&self, x: i32, y: i32) -> i32 {
        match self {
            Self::Add => x + y,
            Self::Subtract => x - y,
        }
    }
}

fn main() {
    
}
```

7、
```rust
fn my_function<const N: usize>() -> [u32; N] {
    [123; N]
}

fn main() {
    let arr = my_function::<5>();
    println!("{:?}",arr);
}
```

8、
```rust
fn main() {
    let s: &str = "Hello there!";

    let arr: &[u8] = &[1, 2, 3];
}
```


9、
```rust
use std::fmt::Display;
fn foobar_1(thing: &dyn Display) {}    
fn foobar_2(thing: Box<dyn Display>) {} 

fn main() {
}
```