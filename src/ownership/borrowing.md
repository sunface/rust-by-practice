# Reference and Borrowing

### Reference
1. 🌟
```rust,editable

fn main() {
   let x = 5;
   // fill the blank
   let p = __;

   println!("the memory address of x is {:p}", p); // one possible output: 0x16fa3ac84
}
```

2. 🌟
```rust,editable

fn main() {
    let x = 5;
    let y = &x;

    // modify this line only
    assert_eq!(5, y);

    println!("Success!")
}
```

3. 🌟
```rust,editable

// fix error
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(s);

    println!("Success!")
}

fn borrow_object(s: &String) {}
```

4. 🌟
```rust,editable

// fix error
fn main() {
    let mut s = String::from("hello, ");

    push_str(s);

    println!("Success!")
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```

5. 🌟🌟
```rust,editable

fn main() {
    let mut s = String::from("hello, ");

    // fill the blank to make it work
    let p = __;
    
    p.push_str("world");

    println!("Success!")
}
```

#### ref
`ref` can be used to take references to a value, similar to `&`.

6. 🌟🌟🌟
```rust,editable

fn main() {
    let c = '中';

    let r1 = &c;
    // fill the blank，dont change other code
    let __ r2 = c;

    assert_eq!(*r1, *r2);
    
    // check the equality of the two address strings
    assert_eq!(get_addr(r1),get_addr(r2));

    println!("Success!")
}

// get memory address string
fn get_addr(r: &char) -> String {
    format!("{:p}", r)
}
```

### Borrowing rules
7. 🌟
```rust,editable

// remove something to make it work
// don't remove a whole line !
fn main() {
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {}", r1, r2);

    println!("Success!")
}
```

#### Mutablity
8. 🌟 Error: Borrow a immutable object as mutable
```rust,editable

fn main() {
    //fix error by modifying this line
    let  s = String::from("hello, ");

    borrow_object(&mut s)

    println!("Success!")
}

fn borrow_object(s: &mut String) {}
```

9. 🌟🌟 Ok: Borrow a mutable object as immutable
```rust,editable

// this code has no errors!
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(&s);
    
    s.push_str("world");

    println!("Success!")
}

fn borrow_object(s: &String) {}
```

### NLL
10. 🌟🌟
```rust,editable

// comment one line to make it work
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    r1.push_str("world");
    let r2 = &mut s;
    r2.push_str("!");
    
    println!("{}",r1);
}
```

11. 🌟🌟
```rust,editable

fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    let r2 = &mut s;

    // add one line below to make a compiler error: cannot borrow `s` as mutable more than once at a time
    // you can't use r1 and r2 at the same time
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it