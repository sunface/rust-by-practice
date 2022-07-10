# Functions
1. 🌟🌟🌟
```rust,editable

fn main() {
    // Don't modify the following two lines!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);

    println!("Success!");
}

fn sum(x: i32, y: i32) -> i32 {
    x + y
}
```


2. 🌟
```rust,editable
fn main() {
   print();
}

// Replace i32 with another type
fn print() -> () {
   println!("Success!");
}
```


3. 🌟🌟🌟

```rust,editable
// Solve it in two ways
// DON'T let `println!` works
fn main() {
    never_return();
}

fn never_return() -> ! {
    // Implement this function, don't modify the fn signatures
    panic!("hehe")
}


// Solve it in two ways
// DON'T let `println!` works
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
        thread::sleep(time::Duration::from_secs(5))
    }
}
```

### Diverging functions 
Diverging functions never return to the caller, so they may be used in places where a value of any type is expected.

4. 🌟🌟
```rust,editable
There are three 3 types of panic statements:-
1. unimplemented!()
2. panic!()
3. 
loop {
    std::thread::sleep(std::time::Duration::from_secs(1))
}

fn main() {
    println!("Success!");
}

fn get_option(tp: u8) -> Option<i32> {
    match tp {
        1 => {
            // TODO
        }
        _ => {
            // TODO
        }
    };
    
    // Rather than returning a None, we use a diverging function instead
    never_return_fn()
}

// IMPLEMENT this function in THREE ways
fn never_return_fn() -> ! {
    unimplemented!()
}


2nd solution:-

fn main() {
    println!("Success!");
}

fn get_option(tp: u8) -> Option<i32> {
    match tp {
        1 => {
            // TODO
        }
        _ => {
            // TODO
        }
    };
    
    // Rather than returning a None, we use a diverging function instead
    never_return_fn()
}

// IMPLEMENT this function in THREE ways
fn never_return_fn() -> ! {
    panic!()
}

3rd solution:-

fn main() {
    println!("Success!");
}

fn get_option(tp: u8) -> Option<i32> {
    match tp {
        1 => {
            // TODO
        }
        _ => {
            // TODO
        }
    };
    
    // Rather than returning a None, we use a diverging function instead
    never_return_fn()
}

// IMPLEMENT this function in THREE ways
fn never_return_fn() -> ! {
    loop {
        std::thread::sleep(std::time::Duration::from_secs(1))
    }
}


```

5. 🌟🌟
```rust,editable
fn main() {
    // FILL in the blank
    // let b = false;

    let v = match false {
        true => 1,
        // Diverging functions can also be used in match expression
        false => {
            println!("Success!");
            panic!("we have no value for `false`, but we can panic")
        }
    };

    println!("{}", v);
    println!("Excercise Failed if printing out this line!");
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it