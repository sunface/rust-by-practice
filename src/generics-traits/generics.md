# Generics

### Functions
1. 🌟🌟🌟
```rust,editable

// fill in the blanks to make it work
struct A;          // Concrete type `A`.
struct S(A);       // Concrete type `S`.
struct SGen<T>(T); // Generic type `SGen`.

fn reg_fn(_s: S) {}

fn gen_spec_t(_s: SGen<A>) {}

fn gen_spec_i32(_s: SGen<i32>) {}

fn generic<T>(_s: SGen<T>) {}

fn main() {
    // Using the non-generic functions
    reg_fn(__);          // Concrete type.
    gen_spec_t(__);   // Implicitly specified type parameter `A`.
    gen_spec_i32(__); // Implicitly specified type parameter `i32`.

    // Explicitly specified type parameter `char` to `generic()`.
    generic::<char>(__);

    // Implicitly specified type parameter `char` to `generic()`.
    generic(__);

    println!("Success!")
}
```

2. 🌟🌟 A function call with explicitly specified type parameters looks like: `fun::<A, B, ...>()`.
```rust,editable

// implement the generic function below
fn sum

fn main() {
    assert_eq!(5, sum(2i8, 3i8));
    assert_eq!(50, sum(20, 30));
    assert_eq!(2.46, sum(1.23, 1.23));

    println!("Success!")
}
```


### Struct and `impl`

3. 🌟
```rust,editable

// implement struct Point to make it work


fn main() {
    let integer = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };

    println!("Success!")
}
```

4. 🌟🌟
```rust,editable

// modify this struct to make the code work
struct Point<T> {
    x: T,
    y: T,
}

fn main() {
    // DON'T modify here
    let p = Point{x: 5, y : "hello".to_string()};

    println!("Success!")
}
```

5. 🌟🌟
```rust,editable

// add generic for Val to make the code work, DON'T modify the code in `main`
struct Val {
    val: f64,
}

impl Val {
    fn value(&self) -> &f64 {
        &self.val
    }
}


fn main() {
    let x = Val{ val: 3.0 };
    let y = Val{ val: "hello".to_string()};
    println!("{}, {}", x.value(), y.value());
}
```

### Method
6. 🌟🌟🌟 

```rust,editable
struct Point<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point<T, U> {
    // implement mixup to make it work, DON'T modify other code
    fn mixup
}

fn main() {
    let p1 = Point { x: 5, y: 10 };
    let p2 = Point { x: "Hello", y: '中'};

    let p3 = p1.mixup(p2);

    assert_eq!(p3.x, 5);
    assert_eq!(p3.y, '中');

    println!("Success!")
}
```

7. 🌟🌟
```rust,editable

// fix the errors to make the code work
struct Point<T> {
    x: T,
    y: T,
}

impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

fn main() {
    let p = Point{x: 5, y: 10};
    println!("{}",p.distance_from_origin())
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it

