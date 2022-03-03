# Associated function & Method

## Examples
```rust,editable
struct Point {
    x: f64,
    y: f64,
}

// Implementation block, all `Point` associated functions & methods go in here
impl Point {
    // This is an "associated function" because this function is associated with
    // a particular type, that is, Point.
    //
    // Associated functions don't need to be called with an instance.
    // These functions are generally used like constructors.
    fn origin() -> Point {
        Point { x: 0.0, y: 0.0 }
    }

    // Another associated function, taking two arguments:
    fn new(x: f64, y: f64) -> Point {
        Point { x: x, y: y }
    }
}

struct Rectangle {
    p1: Point,
    p2: Point,
}

impl Rectangle {
    // This is a method
    // `&self` is sugar for `self: &Self`, where `Self` is the type of the
    // caller object. In this case `Self` = `Rectangle`
    fn area(&self) -> f64 {
        // `self` gives access to the struct fields via the dot operator
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        // `abs` is a `f64` method that returns the absolute value of the
        // caller
        ((x1 - x2) * (y1 - y2)).abs()
    }

    fn perimeter(&self) -> f64 {
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        2.0 * ((x1 - x2).abs() + (y1 - y2).abs())
    }

    // This method requires the caller object to be mutable
    // `&mut self` desugars to `self: &mut Self`
    fn translate(&mut self, x: f64, y: f64) {
        self.p1.x += x;
        self.p2.x += x;

        self.p1.y += y;
        self.p2.y += y;
    }
}

// `Pair` owns resources: two heap allocated integers
struct Pair(Box<i32>, Box<i32>);

impl Pair {
    // This method "consumes" the resources of the caller object
    // `self` desugars to `self: Self`
    fn destroy(self) {
        // Destructure `self`
        let Pair(first, second) = self;

        println!("Destroying Pair({}, {})", first, second);

        // `first` and `second` go out of scope and get freed
    }
}

fn main() {
    let rectangle = Rectangle {
        // Associated functions are called using double colons
        p1: Point::origin(),
        p2: Point::new(3.0, 4.0),
    };

    // Methods are called using the dot operator
    // Note that the first argument `&self` is implicitly passed, i.e.
    // `rectangle.perimeter()` === `Rectangle::perimeter(&rectangle)`
    println!("Rectangle perimeter: {}", rectangle.perimeter());
    println!("Rectangle area: {}", rectangle.area());

    let mut square = Rectangle {
        p1: Point::origin(),
        p2: Point::new(1.0, 1.0),
    };

    // Error! `rectangle` is immutable, but this method requires a mutable
    // object
    //rectangle.translate(1.0, 0.0);
    // TODO ^ Try uncommenting this line

    // Okay! Mutable objects can call mutable methods
    square.translate(1.0, 1.0);

    let pair = Pair(Box::new(1), Box::new(2));

    pair.destroy();

    // Error! Previous `destroy` call "consumed" `pair`
    //pair.destroy();
    // TODO ^ Try uncommenting this line
}
```

## Exercises

### Method
1. 🌟🌟 Methods are similar to functions: declare with `fn`, have parameters and a return value.  Unlike functions, methods are defined within the context of a struct (or an enum or a trait object), and their first parameter is always `self`, which represents the instance of the struct the method is being called on.
```rust,editable
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // complete the area method which return the area of a Rectangle
    fn area
}

fn main() {
    let rect1 = Rectangle { width: 30, height: 50 };

    assert_eq!(rect1.area(), 1500);
}
```

2. 🌟🌟 `self` will take the ownership of current struct instance, however, `&self` will only borrow a reference from the instance

```rust,editable
// Only fill in the blanks, DON'T remove any line!
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    pub fn show_state(__)  {
        println!("the current state is {}", __.color);
    }
}
fn main() {
    let light = TrafficLight{
        color: "red".to_owned(),
    };
    // Don't take the ownership of `light` here
    light.show_state();
    // ..otherwise, there will be an error below
    println!("{:?}", light);
}
```
3. 🌟🌟 The `&self` is actually short for `self: &Self`. Within an `impl` block, the type `Self` is an alias for the type that the `impl` block is for. Methods must have a parameter named `self` of type `Self` for their first parameter, so Rust lets you abbreviate this with only the name `self` in the first parameter spot. 
```rust,editable
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // using `Self` to fill in the blank
    pub fn show_state(__)  {
        println!("the current state is {}", self.color);
    }

    // fill in the blank, DON'T use any variants of `Self`
    pub fn change_state() {
        self.color = "green".to_string()
    }
}
fn main() {}
```


### Associated function

4. 🌟🌟  All functions defined within an `impl` block are called associated functions because they’re associated with the type named after the `impl`. We can define associated functions that don’t have `self` as their first parameter (and thus are not methods) because they don’t need an instance of the type to work with.

```rust,editable
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // 1. implement a assotiated function `new`,
    // 2. it will return a TrafficLight contains color "red"
    // 3. must use `Self`, DONT use `TrafficLight` in fn signatures or body
    pub fn new() 

    pub fn get_state(&self) -> &str {
        &self.color
    }
}

fn main() {
    let light = TrafficLight::new();
    assert_eq!(light.get_state(), "red");
}
```

### Multiple `impl` blocks
5. 🌟 Each struct is allowed to have multiple impl blocks. 
```rust,editable

struct Rectangle {
    width: u32,
    height: u32,
}

// using multiple `impl` blocks to rewrite the code below
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}


fn main() {}
```

### Enums
6. 🌟🌟🌟 We can also implement methods for enums.

```rust,editable

#[derive(Debug)]
enum TrafficLightColor {
    Red,
    Yellow,
    Green,
}

// implement TrafficLightColor with a method
impl TrafficLightColor {
    
}

fn main() {
    let c = TrafficLightColor::Yellow;

    assert_eq!(c.color(), "yellow");

    println!("{:?}",c);
}
```

## Practice

@todo