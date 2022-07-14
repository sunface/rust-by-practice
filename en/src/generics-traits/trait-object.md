# Trait Object
In [traits chapter](https://practice.rs/generics-traits/traits.html#returning-types-that-implement-traits) we have seen that we can't use `impl Trait` when returning multiple types.

Another limitation of arrays is that they can only store elements of one type. Using enums is not a bad solution when we have a fixed set of types at compile time, but trait objects would be more flexible and powerful.

## Returning Traits with dyn
The Rust compiler needs to know how much space a function's return type requires. Because the different implementations of a trait probably uses different amounts of memory, functions need to either return a concrete type or the same type when using  `impl Trait`, or return a trait object with `dyn`.

1. 🌟🌟🌟
```rust,editable

trait Bird {
    fn quack(&self) -> String;
}

struct Duck;
impl Duck {
    fn swim(&self) {
        println!("Look, the duck is swimming")
    }
}
struct Swan;
impl Swan {
    fn fly(&self) {
        println!("Look, the duck.. oh sorry, the swan is flying")
    }
}

impl Bird for Duck {
    fn quack(&self) -> String{
        "duck duck".to_string()
    }
}

impl Bird for Swan {
    fn quack(&self) -> String{
        "swan swan".to_string()
    }
}

fn main() {
    // FILL in the blank.
    let duck = __;
    duck.swim();

    let bird = hatch_a_bird(2);
    // This bird has forgotten how to swim, so below line will cause an error.
    // bird.swim();
    // But it can quak.
    assert_eq!(bird.quack(), "duck duck");

    let bird = hatch_a_bird(1);
    // This bird has forgotten how to fly, so below line will cause an error.
    // bird.fly();
    // But it can quak too.
    assert_eq!(bird.quack(), "swan swan");

    println!("Success!");
}   

// IMPLEMENT this function.
fn hatch_a_bird...

```
## Array with trait objects
2. 🌟🌟
```rust,editable 
trait Bird {
    fn quack(&self);
}

struct Duck;
impl Duck {
    fn fly(&self) {
        println!("Look, the duck is flying")
    }
}
struct Swan;
impl Swan {
    fn fly(&self) {
        println!("Look, the duck.. oh sorry, the swan is flying")
    }
}

impl Bird for Duck {
    fn quack(&self) {
        println!("{}", "duck duck");
    }
}

impl Bird for Swan {
    fn quack(&self) {
        println!("{}", "swan swan");
    }
}

fn main() {
    // FILL in the blank to make the code work.
    let birds __;

    for bird in birds {
        bird.quack();
        // When duck and swan turn into Birds, they all forgot how to fly, only remember how to quack.
        // So, the code below will cause an error.
        // bird.fly();
    }
}
```


## `&dyn` and `Box<dyn>`

3. 🌟🌟
```rust,editable

// FILL in the blanks.
trait Draw {
    fn draw(&self) -> String;
}

impl Draw for u8 {
    fn draw(&self) -> String {
        format!("u8: {}", *self)
    }
}

impl Draw for f64 {
    fn draw(&self) -> String {
        format!("f64: {}", *self)
    }
}

fn main() {
    let x = 1.1f64;
    let y = 8u8;

    // Draw x.
    draw_with_box(__);

    // Draw y.
    draw_with_ref(&y);

    println!("Success!");
}

fn draw_with_box(x: Box<dyn Draw>) {
    x.draw();
}

fn draw_with_ref(x: __) {
    x.draw();
}
```

## Static and Dynamic dispatch
When we use trait bounds on generics, the compiler generates nongeneric implementations of functions and methods for each concrete type that we use in place of a generic type parameter. The code that results from monomorphization is doing static dispatch, which is when the compiler knows what method you’re calling at compile time. 

When we use trait objects, Rust must use dynamic dispatch. The compiler doesn’t know all the types that might be used with the code that is using trait objects, so it doesn’t know which method implemented on which type to call. Instead, at runtime, Rust uses the pointers inside the trait object to know which method to call. There is a runtime cost when this lookup happens that doesn’t occur with static dispatch. Dynamic dispatch also prevents the compiler from choosing to inline a method’s code, which in turn prevents some optimizations. 

However, we do get extra flexibility when using dynamic dispatch.

4. 🌟🌟
```rust,editable

trait Foo {
    fn method(&self) -> String;
}

impl Foo for u8 {
    fn method(&self) -> String { format!("u8: {}", *self) }
}

impl Foo for String {
    fn method(&self) -> String { format!("string: {}", *self) }
}

// IMPLEMENT below with generics.
fn static_dispatch...

// Implement below with trait objects.
fn dynamic_dispatch...

fn main() {
    let x = 5u8;
    let y = "Hello".to_string();

    static_dispatch(x);
    dynamic_dispatch(&y);

    println!("Success!");
}
```

## Object safe
You can only make object-safe traits into trait objects. A trait is object safe if all the methods defined in the trait have the following properties:

- The return type isn’t `Self`.
- There are no generic type parameters.

5. 🌟🌟🌟🌟
```rust,editable

// Use at least two approaches to make it work.
// DON'T add/remove any code line.
trait MyTrait {
    fn f(&self) -> Self;
}

impl MyTrait for u32 {
    fn f(&self) -> Self { 42 }
}

impl MyTrait for String {
    fn f(&self) -> Self { self.clone() }
}

fn my_function(x: Box<dyn MyTrait>)  {
    x.f()
}

fn main() {
    my_function(Box::new(13_u32));
    my_function(Box::new(String::from("abc")));

    println!("Success!");
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it :)
