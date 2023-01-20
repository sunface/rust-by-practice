# Reference and Borrowing

### Reference
1. 🌟
```rust,editable

fn main() {
   let x = 5;
   // Fill the blank
   let p = __;

   println!("the memory address of x is {:p}", p); // One possible output: 0x16fa3ac84
}
```

2. 🌟
```rust,editable

fn main() {
    let x = 5;
    let y = &x;

    // Modify this line only
    assert_eq!(5, y);

    println!("Success!");
}
```

3. 🌟
```rust,editable

// Fix error
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(s);

    println!("Success!");
}

fn borrow_object(s: &String) {}
```

4. 🌟
```rust,editable

// Fix error
fn main() {
    let mut s = String::from("hello, ");

    push_str(s);

    println!("Success!");
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```

5. 🌟🌟
```rust,editable

fn main() {
    let mut s = String::from("hello, ");

    // Fill the blank to make it work
    let p = __;
    
    p.push_str("world");

    println!("Success!");
}
```

#### Ref
`ref` can be used to take references to a value, similar to `&`.

6. 🌟🌟🌟
```rust,editable

fn main() {
    let c = '中';

    let r1 = &c;
    // Fill the blank，dont change other code
    let __ r2 = c;

    assert_eq!(*r1, *r2);
    
    // Check the equality of the two address strings
    assert_eq!(get_addr(r1),get_addr(r2));

    println!("Success!");
}

// Get memory address string
fn get_addr(r: &char) -> String {
    format!("{:p}", r)
}
```

### Borrowing rules
7. 🌟
```rust,editable

// Remove something to make it work
// Don't remove a whole line !
fn main() {
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {}", r1, r2);

    println!("Success!");
}
```

#### Mutability
8. 🌟 Error: Borrow an immutable object as mutable
```rust,editable

fn main() {
    // Fix error by modifying this line
    let  s = String::from("hello, ");

    borrow_object(&mut s);

    println!("Success!");
}

fn borrow_object(s: &mut String) {}
```

9. 🌟🌟 Ok: Borrow a mutable object as immutable
```rust,editable

// This code has no errors!
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(&s);
    
    s.push_str("world");

    println!("Success!");
}

fn borrow_object(s: &String) {}
```

### NLL
10. 🌟🌟
```rust,editable

// Comment one line to make it work
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

    // Add one line below to make a compiler error: cannot borrow `s` as mutable more than once at a time
    // You can't use r1 and r2 at the same time
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/ownership/borrowing.md)(under the solutions path), but only use it when you need it