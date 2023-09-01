1.

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rect1 = Rectangle { width: 30, height: 50 };

    assert_eq!(rect1.area(), 1500);
}
```

2.

```rust
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    pub fn show_state(&self) {
        println!("the current state is {}", self.color);
    }
}

fn main() {
    let light = TrafficLight {
        color: "red".to_owned(),
    };
    // Don't take the ownership of `light` here
    light.show_state();
    // ..otherwise, there will be an error below
    println!("{:?}", light);
}
```

3.

```rust
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // using `Self` to fill in the blank
    pub fn show_state(self: &Self) {
        println!("the current state is {}", self.color);
    }

    // fill in the blank, DON'T use any variants of `Self`
    pub fn change_state(&mut self) {
        self.color = "green".to_string()
    }
}

fn main() {}
```

4.

```rust
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // 1. implement a associated function `new`,
    // 2. it will return a TrafficLight contains color "red"
    // 3. must use `Self`, DONT use `TrafficLight`
    pub fn new() -> Self {
        Self {
            color: "red".to_string()
        }
    }

    pub fn get_state(&self) -> &str {
        &self.color
    }
}

fn main() {
    let light = TrafficLight::new();
    assert_eq!(light.get_state(), "red");
}
```

5.

```rust
struct Rectangle {
    width: u32,
    height: u32,
}

// rewrite Rectangle to use multiple `impl` blocks
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

fn main() {}
```

6.

```rust
#[derive(Debug)]
enum TrafficLightColor {
    Red,
    Yellow,
    Green,
}

// implement TrafficLightColor with a method
impl TrafficLightColor {
    fn color(&self) -> String {
        match *self {
            Self::Red => "red".to_string(),
            Self::Yellow => "yellow".to_string(),
            Self::Green => "green".to_string(),
        }
    }
}

fn main() {
    let c = TrafficLightColor::Yellow;

    assert_eq!(c.color(), "yellow");

    println!("{:?}", c);
}
```
