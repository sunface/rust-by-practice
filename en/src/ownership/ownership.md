# Ownership

🌟🌟
```rust,editable

fn main() {
    // modify this line only! use as many approaches as you can
    let x = String::from("hello, world");
    let y = x;
    println!("{},{}",x,y);
}
```

🌟🌟
```rust,editable
// Don't modify code in main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Only modify the code below!
fn take_ownership(s: String) {
    println!("{}", s);
}
```


🌟🌟
```rust,editable

fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Only modify the code below!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // convert String to Vec
    let _s = s.into_bytes();
    s
}
```

🌟🌟
```rust,editable
// use clone to fix it
fn main() {
    let s = String::from("hello, world");

    print_str(s);

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```

🌟🌟 
```rust, editable
// don't use clone ,use copy instead
fn main() {
    let x = (1, 2, (), "hello");
    let y = x.clone();
    println!("{:?}, {:?}", x, y);
}
```

