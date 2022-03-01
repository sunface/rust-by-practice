# match, if let

### match
🌟🌟
```rust,editable

// fill the blanks
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
        __  => { // matching South or North here
            println!("South or North");
        },
        _ => println!(__),
    };
}
```

🌟🌟 match is an expression, so we can use it in assignments
```rust,editable

fn main() {
    let boolean = true;

    // fill the blank with an match expression:
    //
    // boolean = true, binary = 1
    // boolean = false, binary = 0
    let binary = __;

    assert_eq!(binary, 1);
}
```

🌟🌟 using match to get the data an enum variant holds
```rust,editable

// fill in the blanks
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
} 

fn show_message(msg: Message) {
    match msg {
        __ => { // matches  Message::Move
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
[`matches!`](https://doc.rust-lang.org/stable/core/macro.matches.html) looks like `match`, but can do something different

🌟🌟
```rust,editable

fn main() {
    let alphabets = ['a', 'E', 'Z', '0', 'x', '9' , 'Y'];

    // fill the blank with `matches!` to make the code work
    for ab in alphabets {
        assert!(__)
    }
} 
```

🌟🌟
```rust,editable

enum MyEnum {
    Foo,
    Bar
}

fn main() {
    let mut count = 0;

    let v = vec![MyEnum::Foo,MyEnum::Bar,MyEnum::Foo];
    for e in v {
        if e == MyEnum::Foo { // fix the error with changing only this line
            count += 1;
        }
    }

    assert_eq!(count, 2);
}
```

### if let