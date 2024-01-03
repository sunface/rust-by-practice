# Match, if let

### Match
1. 🌟🌟
```rust,editable

// Completați spațiile libere
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
        __  => { // Se potrivește cu Sud sau Nord aici
            println!("South or North");
        },
        _ => println!(__),
    };
}
```

2. 🌟🌟 `Match` este o expresie, așadar o putem folosi în asignări.
```rust,editable

fn main() {
    let boolean = true;

    // Completați spațiile libere cu o expresie match:
    //
    // boolean = true => binary = 1
    // boolean = false =>  binary = 0
    let binary = __;

    assert_eq!(binary, 1);

    println!("Success!");
}
```

3. 🌟🌟 Utilizarea match pentru a obține datele pe care o variantă enum o deține.
```rust,editable

// Completați spațiile libere
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
        __ => { // se potrivește cu Message::Move
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
[`matches!`](https://doc.rust-lang.org/stable/core/macro.matches.html) pare a fi match, dar poate face ceva diferit.

4. 🌟🌟
```rust,editable

fn main() {
    let alphabets = ['a', 'E', 'Z', '0', 'x', '9' , 'Y'];

    // Completați spațiile libere cu `matches!` pentru a face codul să funcționeze
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
        if e == MyEnum::Foo { // Remediați eroarea schimbând doar această linie
            count += 1;
        }
    }

    assert_eq!(count, 2);

    println!("Success!");
}
```

### If let
Pentru unele cazuri, atunci când potrivirea enumurilor este prea grea, putem folosi `if let`.

6. 🌟 
```rust,editable

fn main() {
    let o = Some(7);

    // Eliminați blocul `match` întreg, folosind `if let` în schimb
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

// Completați spațiile libere
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

    // Eliminați codurile de mai jos, folosind `match` în schimb 
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

// Remediați erorile în locul lor
fn main() {
    let age = Some(30);
    if let Some(age) = age { // Creați o nouă variabilă cu același nume ca și `age` anterior
       assert_eq!(age, Some(30));
    } // Noua variabilă `age` iese din scop aici
    
    match age {
        // Match poate introduce, de asemenea, o nouă variabilă umbrită
        Some(age) =>  println!("age is a new variable, it's value is {}",age),
        _ => ()
    }
 }
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)