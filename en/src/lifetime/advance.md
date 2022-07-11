# advanced lifetime

## Trait Bounds
Just like generic types can be bounded, lifetimes can also be bounded as below:
- `T: 'a`，all references in `T` must outlive the lifetime `'a`
- `T: Trait + 'a`: `T` must implement trait `Trait` and all references in `T` must outlive `'a`

**Example**
```rust,editable
use std::fmt::Debug; // Trait to bound with.

#[derive(Debug)]
struct Ref<'a, T: 'a>(&'a T);
// `Ref` contains a reference to a generic type `T` that has
// an unknown lifetime `'a`. `T` is bounded such that any
// *references* in `T` must outlive `'a`. Additionally, the lifetime
// of `Ref` may not exceed `'a`.

// A generic function which prints using the `Debug` trait.
fn print<T>(t: T) where
    T: Debug {
    println!("`print`: t is {:?}", t);
}

// Here a reference to `T` is taken where `T` implements
// `Debug` and all *references* in `T` outlive `'a`. In
// addition, `'a` must outlive the function.
fn print_ref<'a, T>(t: &'a T) where
    T: Debug + 'a {
    println!("`print_ref`: t is {:?}", t);
}

fn main() {
    let x = 7;
    let ref_x = Ref(&x);

    print_ref(&ref_x);
    print(ref_x);
}
```

1、🌟
```rust,editable
/* Annotate struct with lifetime:
1. `r` and `s` must has different lifetimes
2. lifetime of `s` is bigger than that of 'r'
*/
struct DoubleRef<T> {
    r: &T,
    s: &T
}
fn main() {
    println!("Success!")
}
```


2、🌟🌟
```rust,editable
/* Adding trait bounds to make it work */
struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a, 'b> ImportantExcerpt<'a> {
    fn announce_and_return_part(&'a self, announcement: &'b str) -> &'b str {
        println!("Attention please: {}", announcement);
        self.part
    }
}

fn main() {
    println!("Success!")
}
```

3、🌟🌟
```rust,editable
/* Adding trait bounds to make it work */
fn f<'a, 'b>(x: &'a i32, mut y: &'b i32) {
    y = x;                      
    let r: &'b &'a i32 = &&0;   
}

fn main() {
    println!("Success!")
}
```

## HRTB(Higher-ranked trait bounds)
Type bounds may be higher ranked over lifetimes. These bounds specify a bound is true for all lifetimes. For example, a bound such as `for<'a> &'a T: PartialEq<i32>` would require an implementation like: 

```rust
impl<'a> PartialEq<i32> for &'a T {
    // ...
}
```

and could then be used to compare a `&'a T` with any lifetime to an `i32`.

Only a higher-ranked bound can be used here, because the lifetime of the reference is shorter than any possible lifetime parameter on the function。

4、🌟🌟🌟
```rust
/* Adding HRTB to make it work!*/
fn call_on_ref_zero<'a, F>(f: F) where F: Fn(&'a i32) {
    let zero = 0;
    f(&zero);
}

fn main() {
    println!("Success!")
}
```
## NLL (Non-Lexical Lifetime)
Before explaining NLL, let's see some code first:
```rust
fn main() {
   let mut s = String::from("hello");

    let r1 = &s;
    let r2 = &s;
    println!("{} and {}", r1, r2);

    let r3 = &mut s;
    println!("{}", r3);
}
```

Based on our current knowledge, this code will cause en error due to violating the borrowing rules in Rust.

But if you `cargo run` it, then everything will be ok, so what's going on here?

The ability of the compiler to tell that a reference is no longer used at a point before the end of the scope, is called **Non-Lexical Lifetimes** (**NLL** for short).

With this ability the compiler knows when is the last time that a reference is used and optimizing the borrowing rules based on this knowledge.

```rust
let mut u = 0i32;
let mut v = 1i32;
let mut w = 2i32;

// lifetime of `a` = α ∪ β ∪ γ
let mut a = &mut u;     // --+ α. lifetime of `&mut u`  --+ lexical "lifetime" of `&mut u`,`&mut u`, `&mut w` and `a`
use(a);                 //   |                            |
*a = 3; // <-----------------+                            |
...                     //                                |
a = &mut v;             // --+ β. lifetime of `&mut v`    |
use(a);                 //   |                            |
*a = 4; // <-----------------+                            |
...                     //                                |
a = &mut w;             // --+ γ. lifetime of `&mut w`    |
use(a);                 //   |                            |
*a = 5; // <-----------------+ <--------------------------+
```

## Reborrow
After learning NLL, we can easily understand reborrow now.

**Example**
```rust
#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn move_to(&mut self, x: i32, y: i32) {
        self.x = x;
        self.y = y;
    }
}

fn main() {
    let mut p = Point { x: 0, y: 0 };
    let r = &mut p;
    // Here comes the reborrow
    let rr: &Point = &*r;

    println!("{:?}", rr); // Reborrow ends here, NLL introduced

    // Reborrow is over, we can continue using `r` now
    r.move_to(10, 10);
    println!("{:?}", r);
}
```


5、🌟🌟
```rust,editable
/* Make it work by reordering some code */
fn main() {
    let mut data = 10;
    let ref1 = &mut data;
    let ref2 = &mut *ref1;

    *ref1 += 1;
    *ref2 += 2;

    println!("{}", data);
}
```


## Unbound lifetime
See more info in [Nomicon - Unbounded Lifetimes](https://doc.rust-lang.org/nomicon/unbounded-lifetimes.html).


## More elision rules

```rust
impl<'a> Reader for BufReader<'a> {
    // 'a is not used in the following methods
}

// can be writing as :
impl Reader for BufReader<'_> {
    
}
```

```rust
// Rust 2015
struct Ref<'a, T: 'a> {
    field: &'a T
}

// Rust 2018
struct Ref<'a, T> {
    field: &'a T
}
```


## A difficult exercise

6、🌟🌟🌟🌟
```rust
/* Make it work */
struct Interface<'a> {
    manager: &'a mut Manager<'a>
}

impl<'a> Interface<'a> {
    pub fn noop(self) {
        println!("interface consumed");
    }
}

struct Manager<'a> {
    text: &'a str
}

struct List<'a> {
    manager: Manager<'a>,
}

impl<'a> List<'a> {
    pub fn get_interface(&'a mut self) -> Interface {
        Interface {
            manager: &mut self.manager
        }
    }
}

fn main() {
    let mut list = List {
        manager: Manager {
            text: "hello"
        }
    };

    list.get_interface().noop();

    println!("Interface should be dropped here and the borrow released");

    use_list(&list);
}

fn use_list(list: &List) {
    println!("{}", list.manager.text);
}
```