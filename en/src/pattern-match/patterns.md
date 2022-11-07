# Patterns

1. 🌟🌟 Use `|` to match several values, use `..=` to match an inclusive range.
```rust,editable

fn main() {}
fn match_number(n: i32) {
    match n {
        // Match a single value
        1 => println!("One!"),
        // Fill in the blank with `|`, DON'T use `..` or `..=`
        __ => println!("match 2 -> 5"),
        // Match an inclusive range
        6..=10 => {
            println!("match 6 -> 10")
        },
        _ => {
            println!("match -infinite -> 0 or 11 -> +infinite")
        }
    }
}
```

2. 🌟🌟🌟 The `@` operator lets us create a variable that holds a value, at the same time we are testing that value to see whether it matches a pattern.
```rust,editable

struct Point {
    x: i32,
    y: i32,
}

fn main() {
    // Fill in the blank to let p match the second arm
    let p = Point { x: __, y: __ };

    match p {
        Point { x, y: 0 } => println!("On the x axis at {}", x),
        // Second arm
        Point { x: 0..=5, y: y@ (10 | 20 | 30) } => println!("On the y axis at {}", y),
        Point { x, y } => println!("On neither axis: ({}, {})", x, y),
    }
}
```

3. 🌟🌟🌟

```rust,editable

// Fix the errors
enum Message {
    Hello { id: i32 },
}

fn main() {
    let msg = Message::Hello { id: 5 };

    match msg {
        Message::Hello {
            id:  3..=7,
        } => println!("Found an id in range [3, 7]: {}", id),
        Message::Hello { id: newid@10 | 11 | 12 } => {
            println!("Found an id in another range [10, 12]: {}", newid)
        }
        Message::Hello { id } => println!("Found some other id: {}", id),
    }
}
```

4. 🌟🌟 A match guard is an additional if condition specified after the pattern in a match arm that must also match, along with the pattern matching, for that arm to be chosen.
```rust,editable

// Fill in the blank to make the code work, `split` MUST be used
fn main() {
    let num = Some(4);
    let split = 5;
    match num {
        Some(x) __ => assert!(x < split),
        Some(x) => assert!(x >= split),
        None => (),
    }

    println!("Success!");
}
```

5. 🌟🌟 Ignoring remaining parts of the value with `..`
```rust,editable

// Fill the blank to make the code work
fn main() {
    let numbers = (2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048);

    match numbers {
        __ => {
           assert_eq!(first, 2);
           assert_eq!(last, 2048);
        }
    }

    println!("Success!");
}
```

6. 🌟🌟 Using pattern `&mut V` to match a mutable reference needs you to be very careful, due to `V` being a value  after matching.
```rust,editable

// FIX the error with least changing
// DON'T remove any code line
fn main() {
    let mut v = String::from("hello,");
    let r = &mut v;

    match r {
       &mut value => value.push_str(" world!") 
    }
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it
