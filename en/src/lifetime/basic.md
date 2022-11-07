## Lifetime
The compiler uses lifetime to ensure all borrows are valid. Typically, a variable's lifetime begins when it is created and ends when it is destroyed.

## The scope of lifetime
1. 🌟
```rust,editable
/* Annotate the lifetime of `i` and `borrow2` */

// Lifetimes are annotated below with lines denoting the creation
// and destruction of each variable.
// `i` has the longest lifetime because its scope entirely encloses 
// both `borrow1` and `borrow2`. The duration of `borrow1` compared 
// to `borrow2` is irrelevant since they are disjoint.
fn main() {
    let i = 3;                                             
    {                                                    
        let borrow1 = &i; // `borrow1` lifetime starts. ──┐
        //                                                │
        println!("borrow1: {}", borrow1); //              │
    } // `borrow1 ends. ──────────────────────────────────┘
    {                                                    
        let borrow2 = &i; 
                                                        
        println!("borrow2: {}", borrow2);               
    }                                                   
}   
```

2. 🌟🌟

**Example**
```rust
{
    let x = 5;            // ----------+-- 'b
                          //           |
    let r = &x;           // --+-- 'a  |
                          //   |       |
    println!("r: {}", r); //   |       |
                          // --+       |
}                         // ----------+
```


```rust,editable
/* Annotate `r` and `x` as above, and explain why this code fails to compile, in the lifetime aspect. */

fn main() {  
    {
        let r;                // ---------+-- 'a
                              //          |
        {                     //          |
            let x = 5;        // -+-- 'b  |
            r = &x;           //  |       |
        }                     // -+       |
                              //          |
        println!("r: {}", r); //          |
    }                         // ---------+
}
```

## Lifetime annotating
The **borrow checker uses explicit lifetime annotations** to determine how long a reference should be valid. 

But for us users, in most cases, there is no need to annotate the lifetime, because there are several elision rules, before learning these rules, we need to know how to annotate lifetime manually.

#### Function
Ignoring elision rules, lifetimes in function signatures have a few constraints:

- Any reference must have an annotated lifetime
- Any reference being returned must have the same lifetime as one of the inputs or be static

**Example**
```rust,editable
// One input reference with lifetime `'a` which must live
// at least as long as the function.
fn print_one<'a>(x: &'a i32) {
    println!("`print_one`: x is {}", x);
}

// Mutable references are possible with lifetimes as well.
fn add_one<'a>(x: &'a mut i32) {
    *x += 1;
}

// Multiple elements with different lifetimes. In this case, it
// would be fine for both to have the same lifetime `'a`, but
// in more complex cases, different lifetimes may be required.
fn print_multi<'a, 'b>(x: &'a i32, y: &'b i32) {
    println!("`print_multi`: x is {}, y is {}", x, y);
}

// Returning references that have been passed in is acceptable.
// However, the correct lifetime must be returned.
fn pass_x<'a, 'b>(x: &'a i32, _: &'b i32) -> &'a i32 { x }

fn main() {
    let x = 7;
    let y = 9;
    
    print_one(&x);
    print_multi(&x, &y);
    
    let z = pass_x(&x, &y);
    print_one(z);

    let mut t = 3;
    add_one(&mut t);
    print_one(&t);
}
```

3. 🌟
```rust,editable
/* Make it work by adding proper lifetime annotation */
fn longest(x: &str, y: &str) -> &str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {}
```
4. 🌟🌟🌟
```rust,editable
// `'a` must live longer than the function.
// Here, `&String::from("foo")` would create a `String`, followed by a
// reference. Then the data is dropped upon exiting the scope, leaving
// a reference to invalid data to be returned.

/* Fix the error in three ways  */
fn invalid_output<'a>() -> &'a String { 
    &String::from("foo") 
}

fn main() {
}
```

5. 🌟🌟
```rust,editable
// `print_refs` takes two references to `i32` which have different
// lifetimes `'a` and `'b`. These two lifetimes must both be at
// least as long as the function `print_refs`.
fn print_refs<'a, 'b>(x: &'a i32, y: &'b i32) {
    println!("x is {} and y is {}", x, y);
}

/* Make it work */
// A function which takes no arguments, but has a lifetime parameter `'a`.
fn failed_borrow<'a>() {
    let _x = 12;

    // ERROR: `_x` does not live long enough
    let y: &'a i32 = &_x;
    // Attempting to use the lifetime `'a` as an explicit type annotation 
    // inside the function will fail because the lifetime of `&_x` is shorter
    // than `'a` . A short lifetime cannot be coerced into a longer one.
}

fn main() {
    let (four, nine) = (4, 9);
    
    // Borrows (`&`) of both variables are passed into the function.
    print_refs(&four, &nine);
    // Any input which is borrowed must outlive the borrower. 
    // In other words, the lifetime of `four` and `nine` must 
    // be longer than that of `print_refs`.
    
    failed_borrow();
    // `failed_borrow` contains no references to force `'a` to be 
    // longer than the lifetime of the function, but `'a` is longer.
    // Because the lifetime is never constrained, it defaults to `'static`.
}
```

#### Structs
6. 🌟
```rust,editable
/* Make it work by adding proper lifetime annotation */

// A type `Borrowed` which houses a reference to an
// `i32`. The reference to `i32` must outlive `Borrowed`.
#[derive(Debug)]
struct Borrowed(&i32);

// Similarly, both references here must outlive this structure.
#[derive(Debug)]
struct NamedBorrowed {
    x: &i32,
    y: &i32,
}

// An enum which is either an `i32` or a reference to one.
#[derive(Debug)]
enum Either {
    Num(i32),
    Ref(&i32),
}

fn main() {
    let x = 18;
    let y = 15;

    let single = Borrowed(&x);
    let double = NamedBorrowed { x: &x, y: &y };
    let reference = Either::Ref(&x);
    let number    = Either::Num(y);

    println!("x is borrowed in {:?}", single);
    println!("x and y are borrowed in {:?}", double);
    println!("x is borrowed in {:?}", reference);
    println!("y is *not* borrowed in {:?}", number);
}
```


7. 🌟🌟
```rust,editable
/* Make it work */

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

fn main()
{ 
  /* 'a tied to fn-main stackframe */
  let var_a = 35;
  let example: Example;
  
  {
    /* Lifetime 'b tied to new stackframe/scope */ 
    let var_b = NoCopyType {};
    
    /* fixme */
    example = Example { a: &var_a, b: &var_b };
  }
  
  println!("(Success!) {:?}", example);
}
```


8. 🌟🌟
```rust,editable

#[derive(Debug)]
struct NoCopyType {}

#[derive(Debug)]
#[allow(dead_code)]
struct Example<'a, 'b> {
    a: &'a u32,
    b: &'b NoCopyType
}

/* Fix function signature */
fn fix_me(foo: &Example) -> &NoCopyType
{ foo.b }

fn main()
{
    let no_copy = NoCopyType {};
    let example = Example { a: &1, b: &no_copy };
    fix_me(&example);
    println!("Success!")
}
```

## Method
Methods are annotated similarly to functions.

**Example**
```rust,editable
struct Owner(i32);

impl Owner {
    // Annotate lifetimes as in a standalone function.
    fn add_one<'a>(&'a mut self) { self.0 += 1; }
    fn print<'a>(&'a self) {
        println!("`print`: {}", self.0);
    }
}

fn main() {
    let mut owner = Owner(18);

    owner.add_one();
    owner.print();
}
```

9. 🌟🌟
```rust,editable
/* Make it work by adding proper lifetime annotations */
struct ImportantExcerpt {
    part: &str,
}

impl ImportantExcerpt {
    fn level(&'a self) -> i32 {
        3
    }
}

fn main() {}
```

## Elision
Some lifetime patterns are so comman that the borrow checker will allow you to omit them to save typing and to improve readability.

This is known as **Elision**. Elision exist in Rust only because these patterns are common.

For a more comprehensive understanding of elision, please see [lifetime elision](https://doc.rust-lang.org/book/ch10-03-lifetime-syntax.html#lifetime-elision) in the official book.

10. 🌟🌟
```rust,editable
/* Remove all the lifetimes that can be elided */

fn nput<'a>(x: &'a i32) {
    println!("`annotated_input`: {}", x);
}

fn pass<'a>(x: &'a i32) -> &'a i32 { x }

fn longest<'a, 'b>(x: &'a str, y: &'b str) -> &'a str {
    x
}

struct Owner(i32);

impl Owner {
    // Annotate lifetimes as in a standalone function.
    fn add_one<'a>(&'a mut self) { self.0 += 1; }
    fn print<'a>(&'a self) {
        println!("`print`: {}", self.0);
    }
}

struct Person<'a> {
    age: u8,
    name: &'a str,
}

enum Either<'a> {
    Num(i32),
    Ref(&'a i32),
}

fn main() {}
```