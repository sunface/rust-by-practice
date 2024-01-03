# Enum
1. 🌟🌟 Enum-urile pot fi create cu un discriminator explicit.

```rust,editable

// Rezolvați erorile
enum Number {
    Zero,
    One,
    Two,
}

enum Number1 {
    Zero = 0,
    One,
    Two,
}

// Enum de tip C
enum Number2 {
    Zero = 0.0,
    One = 1.0,
    Two = 2.0,
}


fn main() {
    // O variantă enum poate fi convertită la un întreg folosind as
    assert_eq!(Number::One, Number1::One);
    assert_eq!(Number1::One, Number2::One);

    println!("Success!");
} 
```

2. 🌟 Fiecare variantă de enum poate deține propriile sale date.
```rust,editable

// Completați spațiul liber
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn main() {
    let msg1 = Message::Move{__}; // Instanțiere cu x = 1, y = 2
    let msg2 = Message::Write(__); // Instanțiere cu "hello, world!"

    println!("Success!");
} 
```

3. 🌟🌟 Putem obține datele pe care o variantă enum le deține prin potrivirea de modele.
```rust,editable

// Completați spațiul liber și remediați eroarea
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn main() {
    let msg = Message::Move{x: 1, y: 2};

    if let Message::Move{__} = msg {
        assert_eq!(a, b);
    } else {
        panic!("NEVER LET THIS RUN！");
    }

    println!("Success!");
} 
```

4. 🌟🌟 

```rust,editable

// Completați spațiul liber și remediați erorile
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn main() {
    let msgs: __ = [
        Message::Quit,
        Message::Move{x:1, y:3},
        Message::ChangeColor(255,255,0)
    ];

    for msg in msgs {
        show_message(msg)
    }
} 

fn show_message(msg: Message) {
    println!("{}", msg);
}
```

5. 🌟🌟 Deoarece nu există null în Rust, trebuie să folosim enum-ul Option<T> pentru a gestiona cazurile în care valoarea lipsește.
```rust,editable

// Completați spațiul liber pentru a face funcționarea println.
// Adăugați, de asemenea, un cod pentru a preveni rularea panic.
fn main() {
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);

    if let __ = six {
        println!("{}", n);

        println!("Success!");
    } 
        
    panic!("NEVER LET THIS RUN！");
} 

fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        __ => None,
        __ => Some(i + 1),
    }
}
```


6. 🌟🌟🌟🌟 Implementați o `linked list`` prin intermediul enum-urilor.

```rust,editable

use crate::List::*;

enum List {
    // Cons: Structură tuplu care înconjoară un element și un pointer către următorul nod
    Cons(u32, Box<List>),
    // Nil: Un nod care semnifică sfârșitul listei legate
    Nil,
}

// Metodele pot fi atașate unui enum
impl List {
    // Creați o listă goală
    fn new() -> List {
        // `Nil` are tipul `List`
        Nil
    }

    // Consumați o listă și returnați aceeași listă cu un element nou la începutul ei
    fn prepend(self, elem: u32) -> __ {
        // Cons are, de asemenea, tipul List
        Cons(elem, Box::new(self))
    }

    // Returnați lungimea listei
    fn len(&self) -> u32 {
        // self trebuie potrivit, deoarece comportamentul acestei metode depinde de variantă self
        // self are tipul &List, și *self are tipul List, potrivirea pe un tip concret T este preferată față de o potrivire pe o referință &T
        // După Rust 2018, puteți utiliza self aici și tail (fără ref) mai jos, de asemenea,
        // Rust va deduce &-urile și ref-ul pentru voi.
        // Vezi https://doc.rust-lang.org/edition-guide/rust-2018/ownership-and-lifetimes/default-match-bindings.html
        match *self {
            // Nu putem prelua proprietatea cozii, deoarece self este împrumutat;
            // În schimb, preluați o referință la coadă
            Cons(_, ref tail) => 1 + tail.len(),
            // Caz de bază: O listă goală are lungimea zero
            Nil => 0
        }
    }

    // Returnați reprezentarea listei ca un șir (alocat pe heap)
    fn stringify(&self) -> String {
        match *self {
            Cons(head, __ tail) => {
                // format! este similar cu print!, dar
                // returnează un șir alocat pe heap în loc să imprime la consolă
                format!("{}, {}", head, tail.__())
            },
            Nil => {
                format!("Nil")
            },
        }
    }
}

fn main() {
    // Creează o listă simplu înlănțuită goală
    let mut list = List::new();

    // Adaugă elemente înainte (la începutul listei)
    list = list.prepend(1);
    list = list.prepend(2);
    list = list.prepend(3);

    // Afișează starea finală a listei
    println!("linked list has length: {}", list.len());
    println!("{}", list.stringify());
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.