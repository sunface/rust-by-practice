# Match, if let

### Match
1. 🌟🌟
```rust,editable

// Fill the blanks
enum Direction {
    East,
    West,
    North,
    South,
}

fn main() {
    let dire = Direction::South;
    match dire {
        Direction::East => println!("East"),
        __  => { // Matching South or North here
            println!("South or North");
        },
        _ => println!(__),
    };
}
```

2. 🌟🌟 Match is an expression, so we can use it in assignments.
```rust,editable

fn main() {
    let boolean = true;

    // Fill the blank with a match expression:
    //
    // boolean = true => binary = 1
    // boolean = false =>  binary = 0
    let binary = __;

    assert_eq!(binary, 1);

    println!("Success!");
}
```

3. 🌟🌟 Using match to get the data an enum variant holds.
```rust,editable

// Fill in the blanks
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn main() {
    let msgs = [
        Message::Quit,
        Message::Move{x:1, y:3},
        Message::ChangeColor(255,255,0)
    ];

    for msg in msgs {
        show_message(msg)
    }

    println!("Success!");
} 

fn show_message(msg: Message) {
    match msg {
        __ => { // match  Message::Move
            assert_eq!(a, 1);
            assert_eq!(b, 3);
        },
        Message::ChangeColor(_, g, b) => {
            assert_eq!(g, __);
            assert_eq!(b, __);
        }
        __ => println!("no data in these variants")
    }
}
```

### matches!
[`matches!`](https://doc.rust-lang.org/stable/core/macro.matches.html) looks like `match`, but can do something different.

4. 🌟🌟
```rust,editable

fn main() {
    let alphabets = ['a', 'E', 'Z', '0', 'x', '9' , 'Y'];

    // Fill the blank with `matches!` to make the code work
    for ab in alphabets {
        assert!(__)
    }

    println!("Success!");
} 
```

5. 🌟🌟
```rust,editable

enum MyEnum {
    Foo,
    Bar
}

fn main() {
    let mut count = 0;

    let v = vec![MyEnum::Foo,MyEnum::Bar,MyEnum::Foo];
    for e in v {
        if e == MyEnum::Foo { // Fix the error by changing only this line
            count += 1;
        }
    }

    assert_eq!(count, 2);

    println!("Success!");
}
```

### If let
For some cases, when matching enums, `match` is too heavy. We can use `if let` instead.

6. 🌟 
```rust,editable

fn main() {
    let o = Some(7);

    // Remove the whole `match` block, using `if let` instead 
    match o {
        Some(i) => {
            println!("This is a really long string and `{:?}`", i);

            println!("Success!");
        }
        _ => {}
    };
}
```

7. 🌟🌟
```rust,editable

// Fill in the blank
enum Foo {
    Bar(u8)
}

fn main() {
    let a = Foo::Bar(1);

    __ {
        println!("foobar holds the value: {}", i);

        println!("Success!");
    }
}
```

8. 🌟🌟
```rust,editable

enum Foo {
    Bar,
    Baz,
    Qux(u32)
}

fn main() {
    let a = Foo::Qux(10);

    // Remove the codes below, using `match` instead 
    if let Foo::Bar = a {
        println!("match foo::bar")
    } else if let Foo::Baz = a {
        println!("match foo::baz")
    } else {
        println!("match others")
    }
}
```

### Shadowing
9. 🌟🌟
```rust,editable

// Fix the errors in-place
fn main() {
    let age = Some(30);
    if let Some(age) = age { // Create a new variable with the same name as previous `age`
       assert_eq!(age, Some(30));
    } // The new variable `age` goes out of scope here
    
    match age {
        // Match can also introduce a new shadowed variable
        Some(age) =>  println!("age is a new variable, it's value is {}",age),
        _ => ()
    }
 }
```


> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it