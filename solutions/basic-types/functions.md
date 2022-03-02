1.
```rust
fn main() {
    // don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);
}

fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```

2.
```rust
fn main() {
    print();
}

// replace i32 with another type
fn print() -> () {
    println!("hello,world");
}
```

3.
```rust
fn main() {
    never_return();
}

fn never_return() -> ! {
    // implement this function, don't modify fn signatures
    panic!("I return nothing!")
}
```

```rust
fn main() {
    never_return();
}

use std::thread;
use std::time;
fn never_return() -> ! {
    // implement this function, don't modify fn signatures
    loop {
        println!("I return nothing");
        // sleeping for 1 second to avoid exhausting the cpu resoucre
        thread::sleep(time::Duration::from_secs(1))
    }
}
```