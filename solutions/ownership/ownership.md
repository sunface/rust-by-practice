1.

```rust
fn main() {
    let x = String::from("hello, world");
    let y = x.clone();
    println!("{},{}",x,y);
}
```

```rust
fn main() {
    let x = "hello, world";
    let y = x;
    println!("{},{}",x,y);
}
```

```rust
fn main() {
    let x = &String::from("hello, world");
    let y = x;
    println!("{},{}",x,y);
}
```
```rust
fn main() {
    let x = String::from("hello, world");
    let y = x.as_str();
    println!("{},{}",x,y);
}
```

2.

```rust
// Don't modify code in main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Only modify the code below!
fn take_ownership(s: String) -> String {
    println!("{}", s);
    s
}
```

3.

```rust
fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // convert String to Vec
    let _s = s.as_bytes();
    s
}
```

```rust
fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    s
}
```

4.

```rust
fn main() {
    let s = String::from("hello, world");

    print_str(s.clone());

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```

```rust
 fn main() {
     let s = String::from("hello, world");
     print_str(&s);
     println!("{}", s);
 }
 fn print_str(s: &String)  {
     println!("{}",s)
 }
```

5.

```rust
fn main() {
    let x = (1, 2, (), "hello");
    let y = x;
    println!("{:?}, {:?}", x, y);
}
```

6.

```rust
fn main() {
    let s = String::from("hello, ");
    
    // modify this line only !
    let mut s1 = s;

    s1.push_str("world")
}
```

7.

```rust
fn main() {
    let x = Box::new(5);
    
    let mut y = Box::new(3);       // implement this line, dont change other lines!
    
    *y = 4;
    
    assert_eq!(*x, 5);
}
```

8.

```rust
fn main() {
    let t = (String::from("hello"), String::from("world"));
 
    let _s = t.0;
 
    // modify this line only, don't use `_s`
    println!("{:?}", t.1);
 }
```

9.

```rust
fn main() {
    let t = (String::from("hello"), String::from("world"));

    // fill the blanks
    let (s1, s2) = t.clone();

    println!("{:?}, {:?}, {:?}", s1, s2, t); // -> "hello", "world", ("hello", "world")
}
```

```rust
fn main() {
    let t = (String::from("hello"), String::from("world"));

    // fill the blanks
    let (ref s1, ref s2) = t;

    println!("{:?}, {:?}, {:?}", s1, s2, t); // -> "hello", "world", ("hello", "world")
}
```
