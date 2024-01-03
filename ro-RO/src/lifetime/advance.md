# Advance în durata de viață

## Limitele de tipuri
La fel ca tipurile generice pot fi limitate, duratele de viață pot fi, de asemenea, limitate după cum urmează:
- T: `'a` - toate referințele din 'T' trebuie să supraviețuiască duratei de viață `'a`
- T: Trait + `'a`:' 'T' trebuie să implementeze trăsătura Trait, iar toate referințele din 'T' trebuie să supraviețuiască `'a`

**Exemplu**
```rust,editable
use std::fmt::Debug; // Trăsătură pentru a limita.

#[derive(Debug)]
struct Ref<'a, T: 'a>(&'a T);
/// `Ref` conține o referință la un tip generic `T` cu o durată de viață necunoscută `'a`.
// `T` este limitat astfel încât orice *referințe* în `T` trebuie să supraviețuiască `'a`.
// În plus, durata de viață a `Ref` nu poate depăși `'a`.

// Funcție generică care tipărește utilizând trăsătura `Debug`.
fn print<T>(t: T) where
    T: Debug {
    println!("`print`: t is {:?}", t);
}

// Aici se ia o referință la `T` unde `T` implementează `Debug` și toate *referințele*
// din `T` supraviețuiesc lui `'a`. În plus, `'a` trebuie să supraviețuiască funcției.
fn print_ref<'a, T>(t: &'a T) where
    T: Debug + 'a {
    println!("`print_ref`: t is {:?}", t);
}

fn main() {
    let x = 7;
    let ref_x = Ref(&x);

    print_ref(&ref_x);
    print(ref_x);
}
```

1. 🌟
```rust,editable
/* Adăugați durata de viață la structură:
1. `r` și `s` trebuie să aibă durate de viață diferite
2. durata de viață a `s` este mai mare decât cea a 'r'
*/
struct DoubleRef<T> {
    r: &T,
    s: &T
}
fn main() {
    println!("Success!")
}
```


2. 🌟🌟
```rust,editable
/* Adăugați limite de trăsături pentru a face codul să funcționeze */
struct ImportantExcerpt<'a> {
    part: &'a str,
}

impl<'a, 'b> ImportantExcerpt<'a> {
    fn announce_and_return_part(&'a self, announcement: &'b str) -> &'b str {
        println!("Attention please: {}", announcement);
        self.part
    }
}

fn main() {
    println!("Success!")
}
```

3. 🌟🌟
```rust,editable
/* Adăugați limite de trăsături pentru a face codul să funcționeze */
fn f<'a, 'b>(x: &'a i32, mut y: &'b i32) {
    y = x;                      
    let r: &'b &'a i32 = &&0;   
}

fn main() {
    println!("Success!")
}
```

## HRTB (Higher-ranked trait bounds)
Limitele de tip pot fi mai înalte decât duratele de viață. Aceste limite specifică faptul că o limită este valabilă pentru toate duratele de viață. De exemplu, o limită cum ar fi `for<'a> &'a T: PartialEq<i32>` ar necesita o implementare ca:

```rust
impl<'a> PartialEq<i32> for &'a T {
    // ...
}
```

ași apoi ar putea fi folosită pentru a compara un `&'a` `T` cu orice durată de viață față de un `i32`.

Doar o limită mai înaltă poate fi folosită aici, deoarece durata de viață a referinței este mai scurtă decât orice posibil parametru de durată de viață al funcției.

4. 🌟🌟🌟
```rust,editable
/* Adăugați HRTB pentru a face codul să funcționeze! */
fn call_on_ref_zero<'a, F>(f: F) where F: Fn(&'a i32) {
    let zero = 0;
    f(&zero);
}

fn main() {
    println!("Success!");
}
```
## NLL (Non-Lexical Lifetime)
Înainte de a explica NLL, să vedem mai întâi un cod:
```rust
fn main() {
   let mut s = String::from("hello");

    let r1 = &s;
    let r2 = &s;
    println!("{} and {}", r1, r2);

    let r3 = &mut s;
    println!("{}", r3);
}
```

Bazat pe cunoștințele noastre actuale, acest cod va cauza o eroare datorită încălcării regulilor de împrumut în Rust.

Dar dacă îl executați cu `cargo run`, totul va fi în regulă, așa că ce se întâmplă aici?

Capacitatea compilatorului de a determina că o referință nu este folosită într-un punct înainte de sfârșitul domeniului de valabilitate se numește **Non-Lexical Lifetimes** (**NLL** în scurt).

Cu această capacitate, compilatorul știe când este ultima dată când o referință este utilizată și optimizează regulile de împrumut pe baza acestei cunoștințe.

```rust
let mut u = 0i32;
let mut v = 1i32;
let mut w = 2i32;

// durata de viață a `a` = α ∪ β ∪ γ
let mut a = &mut u;     // --+ α. durata de viață a `&mut u`  --+ durata lexicală "a" a `&mut u`,`&mut u`, `&mut w` și `a`
use(a);                 //   |                            |
*a = 3; // <-----------------+                            |
...                     //                                |
a = &mut v;             // --+ β. durata de viață a `&mut v`
use(a);                 //   |                            |
*a = 4; // <-----------------+                            |
...                     //                                |
a = &mut w;             // --+ γ. durata de viață a `&mut w`
use(a);                 //   |                            |
*a = 5; // <-----------------+ <--------------------------+
```

## Reborrow
După învățarea NLL, putem înțelege ușor reborrow acum.

**Exemplu**
```rust
#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn move_to(&mut self, x: i32, y: i32) {
        self.x = x;
        self.y = y;
    }
}

fn main() {
    let mut p = Point { x: 0, y: 0 };
    let r = &mut p;
    // Aici vine reborrow
    let rr: &Point = &*r;

    println!("{:?}", rr); //Reborrow se termină aici, NLL introdus

    // Reborrow s-a terminat, acum putem continua să folosim `r`
    r.move_to(10, 10);
    println!("{:?}", r);
}
```


5. 🌟🌟
```rust,editable
/* Faceți-l să funcționeze prin reordonarea unor coduri */
fn main() {
    let mut data = 10;
    let ref1 = &mut data;
    let ref2 = &mut *ref1;

    *ref1 += 1;
    *ref2 += 2;

    println!("{}", data);
}
```


## Durată de viață nelimitată
Vezi mai multe informații în [Nomicon - Unbounded Lifetimes](https://doc.rust-lang.org/nomicon/unbounded-lifetimes.html).


## More elision rules

```rust
impl<'a> Reader for BufReader<'a> {
    // 'a nu este folosit în metodele următoare
}

// poate fi scris ca:
impl Reader for BufReader<'_> {
    
}
```

```rust
// Rust 2015
struct Ref<'a, T: 'a> {
    field: &'a T
}

// Rust 2018
struct Ref<'a, T> {
    field: &'a T
}
```


## Un exercițiu dificil

6. 🌟🌟🌟🌟
```rust,editable
/* Faceți-l să funcționeze */
struct Interface<'a> {
    manager: &'a mut Manager<'a>
}

impl<'a> Interface<'a> {
    pub fn noop(self) {
        println!("interface consumed");
    }
}

struct Manager<'a> {
    text: &'a str
}

struct List<'a> {
    manager: Manager<'a>,
}

impl<'a> List<'a> {
    pub fn get_interface(&'a mut self) -> Interface {
        Interface {
            manager: &mut self.manager
        }
    }
}

fn main() {
    let mut list = List {
        manager: Manager {
            text: "hello"
        }
    };

    list.get_interface().noop();

    println!("Interface should be dropped here and the borrow released");

    use_list(&list);
}

fn use_list(list: &List) {
    println!("{}", list.manager.text);
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
