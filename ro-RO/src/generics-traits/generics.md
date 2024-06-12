# Generice

### Funcții
1. 🌟🌟🌟
```rust,editable

// Completați spațiile libere pentru a face codul să funcționeze
struct A;          // Tip concret `A`.
struct S(A);       // Tip concret `S`.
struct SGen<T>(T); // Tip generic `SGen`.

fn reg_fn(_s: S) {}

fn gen_spec_t(_s: SGen<A>) {}

fn gen_spec_i32(_s: SGen<i32>) {}

fn generic<T>(_s: SGen<T>) {}

fn main() {
    // Folosind funcțiile non-generice
    reg_fn(__);          // Tip concret.
    gen_spec_t(__);   // Parametru de tip `A` specificat implicit.
    gen_spec_i32(__); // Parametru de tip `i32` specificat implicit.

    // Parametru de tip `char` specificat explicit pentru `generic()`.
    generic::<char>(__);

    // Parametru de tip `char` specificat implicit pentru `generic()`.
    generic(__);

    println!("Success!");
}
```

2. 🌟🌟 Un apel de funcție cu parametri de tip specificați explicit arată astfel: 'fun::<A, B, ...>()'.
```rust,editable

// Implementați funcția generică mai jos.
fn sum

fn main() {
    assert_eq!(5, sum(2i8, 3i8));
    assert_eq!(50, sum(20, 30));
    assert_eq!(2.46, sum(1.23, 1.23));

    println!("Success!");
}
```


### Structură și impl

3. 🌟
```rust,editable

// Implementați structura Punct pentru a face codul să funcționeze.


fn main() {
    let integer = Point { x: 5, y: 10 };
    let float = Point { x: 1.0, y: 4.0 };

    println!("Success!");
}
```

4. 🌟🌟
```rust,editable

// Modificați această structură pentru a face codul să funcționeze
struct Point<T> {
    x: T,
    y: T,
}

fn main() {
    // NU modificați acest cod.
    let p = Point{x: 5, y : "hello".to_string()};

    println!("Success!");
}
```

5. 🌟🌟
```rust,editable

// Adăugați generic pentru Val pentru a face codul să funcționeze, NU modificați codul din `main`.
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

### Metodă
6. 🌟🌟🌟 

```rust,editable
struct Point<T, U> {
    x: T,
    y: U,
}

impl<T, U> Point<T, U> {
    // Implementați mixup pentru a face codul să funcționeze, NU modificați alte coduri.
    fn mixup
}

fn main() {
    let p1 = Point { x: 5, y: 10 };
    let p2 = Point { x: "Hello", y: '中'};

    let p3 = p1.mixup(p2);

    assert_eq!(p3.x, 5);
    assert_eq!(p3.y, '中');

    println!("Success!");
}
```

7. 🌟🌟
```rust,editable

// Corectați erorile pentru a face codul să funcționeze.
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
    println!("{}",p.distance_from_origin());
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
