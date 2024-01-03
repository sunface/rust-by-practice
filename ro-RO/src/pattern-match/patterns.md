# Pattern-uri (Patterns)

1. 🌟🌟 Folosiți `|` pentru a potrivi mai multe valori, utilizați `..=` pentru a potrivi o gamă inclusivă.
```rust,editable

fn main() {}
fn match_number(n: i32) {
    match n {
        // Potriviți o singură valoare
        1 => println!("One!"),
        // Completați spațiile libere cu `|`, NU folosiți `..` sau `..=`
        __ => println!("match 2 -> 5"),
        // Potriviți o gamă inclusivă
        6..=10 => {
            println!("match 6 -> 10")
        },
        _ => {
            println!("match -infinite -> 0 or 11 -> +infinite")
        }
    }
}
```

2. 🌟🌟🌟 Operatorul `@` ne permite să creăm o variabilă care reține o valoare, în același timp când testăm dacă acea valoare se potrivește cu un model.
```rust,editable

struct Point {
    x: i32,
    y: i32,
}

fn main() {
    // Completați spațiile libere pentru a face ca p să se potrivească cu al doilea braț
    let p = Point { x: __, y: __ };

    match p {
        Point { x, y: 0 } => println!("On the x axis at {}", x),
        // Al doilea braț
        Point { x: 0..=5, y: y@ (10 | 20 | 30) } => println!("On the y axis at {}", y),
        Point { x, y } => println!("On neither axis: ({}, {})", x, y),
    }
}
```

3. 🌟🌟🌟

```rust,editable

// Remediați erorile
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

4. 🌟🌟 Un gard de potrivire este o condiție suplimentară specificată după model într-un braț de potrivire, care trebuie, de asemenea, să se potrivească, împreună cu potrivirea modelului, pentru ca acel braț să fie ales.
```rust,editable

// Completați spațiile libere pentru a face codul să funcționeze, `split` TREBUIE să fie utilizat
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

5. 🌟🌟 Ignorarea părților rămase ale valorii cu `..`
```rust,editable

// Completați spațiile libere pentru a face codul să funcționeze
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

6. 🌟🌟 Utilizarea modelului `&mut V` pentru a potrivi o referință mutabilă necesită foarte multă atenție, datorită faptului că `V` este o valoare după potrivire.

```rust,editable

// REMEDIAȚI eroarea cu cel mai mic număr de modificări
// NU eliminați nici o linie de cod
fn main() {
    let mut v = String::from("hello,");
    let r = &mut v;

    match r {
       &mut value => value.push_str(" world!") 
    }
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
