# Iterator
Modelul iterator permite efectuarea unor sarcini pe o secvență de elemente pe rând. Un iterator este responsabil pentru logica iterării prin fiecare element și determinarea momentului în care secvența s-a încheiat.

## 'for' și iterator
```rust
fn main() {
    let v = vec![1, 2, 3];
    for x in v {
        println!("{}",x)
    }
}
```

În codul de mai sus, puteți considera for ca un simplu ciclu, dar de fapt acesta iterează printr-un iterator.

În mod implicit, 'for' va aplica 'into_iter' la colecție, transformând-o într-un iterator. Ca rezultat, codul de mai jos este echivalent cu cel anterior:
```rust
fn main() {
    let v = vec![1, 2, 3];
    for x in v.into_iter() {
        println!("{}",x)
    }
}
```


1. 🌟

```rust,editable
/* Refactorizați următorul cod folosind iteratori */
fn main() {
    let arr = [0; 10];
    for i in 0..arr.len() {
        println!("{}",arr[i]);
    }
}
```

2. 🌟 Una dintre cele mai simple modalități de a crea un iterator este utilizarea notației de interval: 'a..b'.
```rust,editable
/* Completați spațiul liber */
fn main() {
    let mut v = Vec::new();
    for n in __ {
       v.push(n);
    }

    assert_eq!(v.len(), 100);
}
```

## Metoda next
Toate iteratoarele implementează un trăsătura numită 'Iterator' definită în biblioteca standard:
```rust
pub trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self::Item>;

    // Metode cu implementări implicite eliberate
}
```

Și putem apela metoda 'next' direct pe iteratoare.


3. 🌟🌟

```rust,editable
/* Completați spațiile libere și rezolvați erorile.
Folosiți două modalități dacă este posibil */
fn main() {
    let v1 = vec![1, 2];

    assert_eq!(v1.next(), __);
    assert_eq!(v1.next(), __);
    assert_eq!(v1.next(), __);
}
```

## 'into_iter', 'iter' și 'iter_mut'

În secțiunea anterioară, am menționat că 'for' va aplica 'into_iter' la colecție și o va transforma într-un iterator. Cu toate acestea, aceasta nu este singura modalitate de a converti colecțiile în iteratoare.

- 'into_iter', 'iter' și 'iter_mut', toate acestea pot converti o colecție într-un iterator, dar în moduri diferite.

- 'into_iter' consumă colecția; o dată ce colecția a fost consumată, nu mai este disponibilă pentru reutilizare, deoarece stăpânirea sa a fost mutată în cadrul buclei.
- 'iter' împrumută fiecare element al colecției prin fiecare iterație, lăsând astfel colecția neatinsă și disponibilă pentru reutilizare după buclă.
- 'iter_mut' împrumută mutabil fiecare element al colecției, permițând modificarea colecției pe loc.


4. 🌟

```rust,editable
/* Faceți-l să funcționeze */
fn main() {
    let arr = vec![0; 10];
    for i in arr {
        println!("{}", i);
    }

    println!("{:?}",arr);
}
```


5. 🌟

```rust,editable
/* Completați spațiul liber */
fn main() {
    let mut names = vec!["Bob", "Frank", "Ferris"];

    for name in names.__{
        *name = match name {
            &mut "Ferris" => "There is a rustacean among us!",
            _ => "Hello",
        }
    }

    println!("names: {:?}", names);
}
```


6. 🌟🌟

```rust,editable
/* Completați spațiul liber */
fn main() {
    let mut values = vec![1, 2, 3];
    let mut values_iter = values.__;

    if let Some(v) = values_iter.__{
        __
    }

    assert_eq!(values, vec![0, 2, 3]);
}
```


## Crearea propriilor noștri iteratori
Nu putem doar crea iteratori din tipurile de colecție, ci și putem crea iteratori prin implementarea trăsăturii 'Iterator' pe propriile noastre tipuri.

**Exemplu**
```rust
struct Counter {
    count: u32,
}

impl Counter {
    fn new() -> Counter {
        Counter { count: 0 }
    }
}

impl Iterator for Counter {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        if self.count < 5 {
            self.count += 1;
            Some(self.count)
        } else {
            None
        }
    }
}

fn main() {
    let mut counter = Counter::new();

    assert_eq!(counter.next(), Some(1));
    assert_eq!(counter.next(), Some(2));
    assert_eq!(counter.next(), Some(3));
    assert_eq!(counter.next(), Some(4));
    assert_eq!(counter.next(), Some(5));
    assert_eq!(counter.next(), None);
}
```


7. 🌟🌟🌟

```rust,editable
struct Fibonacci {
    curr: u32,
    next: u32,
}

// Implementați `Iterator` pentru `Fibonacci`.
// Trăsătura `Iterator` necesită doar definirea unei metode pentru elementul `next`.
impl Iterator for Fibonacci {
    // Putem să ne referim la acest tip folosind Self::Item
    type Item = u32;
    
    /* Implementați metoda next */
    fn next(&mut self)
}

// Returnează un generator de secvență Fibonacci
fn fibonacci() -> Fibonacci {
    Fibonacci { curr: 0, next: 1 }
}

fn main() {
    let mut fib = fibonacci();
    assert_eq!(fib.next(), Some(1));
    assert_eq!(fib.next(), Some(1));
    assert_eq!(fib.next(), Some(2));
    assert_eq!(fib.next(), Some(3));
    assert_eq!(fib.next(), Some(5));
}
```

## Metode care consumă iteratorul
Trăsătura Iterator are un număr de metode cu implementări implicite furnizate de biblioteca standard.

### Adaptoare de consum
Unele dintre aceste metode apelează metoda next pentru a utiliza iteratorul, astfel sunt numite adaptoare de consum.


8. 🌟🌟

```rust,edtiable

/* Completați spațiul liber și rezolvați erorile */
fn main() {
    let v1 = vec![1, 2, 3];

    let v1_iter = v1.iter();

    // Metoda sum va lua stăpânirea iteratorului și va trece prin elemente apelând repetat metoda next
    let total = v1_iter.sum();

    assert_eq!(total, __);

    println!("{:?}, {:?}",v1, v1_iter);
}
```


#### Colectare
În afară de conversia unei colecții într-un iterator, putem de asemenea să collectăm valorile rezultate într-o colecție, collect va consuma iteratorul.


9. 🌟🌟

```rust,editable
/* Faceți-l să funcționeze */
use std::collections::HashMap;
fn main() {
    let names = [("sunface",18), ("sunfei",18)];
    let folks: HashMap<_, _> = names.into_iter().collect();

    println!("{:?}",folks);

    let v1: Vec<i32> = vec![1, 2, 3];

    let v2 = v1.iter().collect();

    assert_eq!(v2, vec![1, 2, 3]);
}
```


###  Adaptoare de iterator
Metodele care vă permit să schimbați un iterator în alt iterator sunt cunoscute sub numele de *adaptoare de iterator*. Puteți concatena mai multe adaptoare de iterator pentru a efectua acțiuni complexe într-un mod lizibil.

Dar deoarece **toate iteratoarele sunt leneșe**, trebuie să apelați unul dintre adaptoarele de consum pentru a obține rezultate din apelurile la adaptoarele de iterator.


10. 🌟🌟

```rust,editable
/* Completați spațiile libere */
fn main() {
    let v1: Vec<i32> = vec![1, 2, 3];

    let v2: Vec<_> = v1.iter().__.__;

    assert_eq!(v2, vec![2, 3, 4]);
}
```


11. 🌟🌟

```rust
/* Completați spațiile libere */
use std::collections::HashMap;
fn main() {
    let names = ["sunface", "sunfei"];
    let ages = [18, 18];
    let folks: HashMap<_, _> = names.into_iter().__.collect();

    println!("{:?}",folks);
}
```


#### Utilizarea închiderilor în adaptoarele de iterator


12. 🌟🌟 

```rust
/* Completați spațiile libere */
#[derive(PartialEq, Debug)]
struct Shoe {
    size: u32,
    style: String,
}

fn shoes_in_size(shoes: Vec<Shoe>, shoe_size: u32) -> Vec<Shoe> {
    shoes.into_iter().__.collect()
}

fn main() {
    let shoes = vec![
        Shoe {
            size: 10,
            style: String::from("sneaker"),
        },
        Shoe {
            size: 13,
            style: String::from("sandal"),
        },
        Shoe {
            size: 10,
            style: String::from("boot"),
        },
    ];

    let in_my_size = shoes_in_size(shoes, 10);

    assert_eq!(
        in_my_size,
        vec![
            Shoe {
                size: 10,
                style: String::from("sneaker")
            },
            Shoe {
                size: 10,
                style: String::from("boot")
            },
        ]
    );
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
