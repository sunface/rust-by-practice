1.

```rust
fn main() {
    let x = 5;
    // fill the blank
    let p = &x;
 
    println!("the memory address of x is {:p}", p); // one possible output: 0x16fa3ac84
}
```

2.

```rust
fn main() {
    let x = 5;
    let y = &x;

    // modify this line only
    assert_eq!(5, *y);
}
```

3.

```rust
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(&s)
}

fn borrow_object(s: &String) {}
```

4.

```rust
fn main() {
    let mut s = String::from("hello, ");

    push_str(&mut s)
}

fn push_str(s: &mut String) {
    s.push_str("world")
}
```

5.

```rust
fn main() {
    let mut s = String::from("hello, ");

    // fill the blank to make it work
    let p = &mut s;
    
    p.push_str("world");
}
```

6.

```rust
fn main() {
    let c = '中';

    let r1 = &c;
    // fill the blank，dont change other code
    let ref r2 = c;

    assert_eq!(*r1, *r2);
    
    // check the equality of the two address strings
    assert_eq!(get_addr(r1),get_addr(r2));
}

// get memory address string
fn get_addr(r: &char) -> String {
    format!("{:p}", r)
}
```

7.

```rust
fn main() {
    let s = String::from("hello");

    let r1 = &s;
    let r2 = &s;

    println!("{}, {}", r1, r2);
}
```

8.

```rust
fn main() {
    //fix error by modifying this line
    let mut s = String::from("hello, ");

    borrow_object(&mut s)
}

fn borrow_object(s: &mut String) {}
```

9.

```rust
fn main() {
    let mut s = String::from("hello, ");

    borrow_object(&s);
    
    s.push_str("world");
}

fn borrow_object(s: &String) {}
```

10.

```rust
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    r1.push_str("world");
    let r2 = &mut s;
    r2.push_str("!");
    
    // println!("{}",r1);
}
```

11.

```rust
fn main() {
    let mut s = String::from("hello, ");

    let r1 = &mut s;
    let r2 = &mut s;

    // add one line below to make a compiler error: cannot borrow `s` as mutable more than once at a time
    // you can't use r1 and r2 at the same time
    r1.push_str("world");
}
```