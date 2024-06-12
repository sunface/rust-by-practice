# Obiecte de Trăsături (Trait Object)
În [capitolul despre trăsături](https://practice.rs/generics-traits/traits.html#returning-types-that-implement-traits) am văzut că nu putem folosi impl Trait atunci când returnăm mai multe tipuri.

O altă limitare a array-urilor este că pot stoca doar elemente de un singur tip. Utilizarea enum-urilor nu este o soluție rea atunci când avem un set fix de tipuri la timpul de compilare, dar obiectele de trăsături ar fi mai flexibile și mai puternice.

## Returnarea Trăsăturilor cu 'dyn'
Compilatorul Rust trebuie să știe cât spațiu necesită tipul de returnare al unei funcții. Deoarece implementările diferite ale unei trăsături folosesc probabil cantități diferite de memorie, funcțiile trebuie să returneze fie un tip concret sau același tip atunci când se folosește 'impl Trait', fie să returneze un obiect de trăsătură cu 'dyn'.

1. 🌟🌟🌟
```rust,editable

trait Bird {
    fn quack(&self) -> String;
}

struct Duck;
impl Duck {
    fn swim(&self) {
        println!("Look, the duck is swimming")
    }
}
struct Swan;
impl Swan {
    fn fly(&self) {
        println!("Look, the duck.. oh sorry, the swan is flying")
    }
}

impl Bird for Duck {
    fn quack(&self) -> String{
        "duck duck".to_string()
    }
}

impl Bird for Swan {
    fn quack(&self) -> String{
        "swan swan".to_string()
    }
}

fn main() {
    // Completați spațiul liber.
    let duck = __;
    duck.swim();

    let bird = hatch_a_bird(2);
    // Această pasăre a uitat să înoate, așa că linia de mai jos va provoca o eroare.
    // pasare.inoata();
    // Dar poate cotcodă.
    assert_eq!(bird.quack(), "duck duck");

    let bird = hatch_a_bird(1);
    // Această pasăre a uitat să zboare, așa că linia de mai jos va provoca o eroare.
    // pasare.zboară();
    // Dar poate cotcodă și ea.
    assert_eq!(bird.quack(), "swan swan");

    println!("Success!");
}   

// IMPLEMENTAȚI această funcție.
fn hatch_a_bird...

```
## Array cu obiecte de trăsături
2. 🌟🌟
```rust,editable 
trait Bird {
    fn quack(&self);
}

struct Duck;
impl Duck {
    fn fly(&self) {
        println!("Look, the duck is flying")
    }
}
struct Swan;
impl Swan {
    fn fly(&self) {
        println!("Look, the duck.. oh sorry, the swan is flying")
    }
}

impl Bird for Duck {
    fn quack(&self) {
        println!("{}", "duck duck");
    }
}

impl Bird for Swan {
    fn quack(&self) {
        println!("{}", "swan swan");
    }
}

fn main() {
    // Completați spațiul liber pentru a face codul să funcționeze.
    let birds __;

    for bird in birds {
        bird.quack();
        // Când rata și lebăda se transformă în Păsări, ele au uitat cum să zboare, au reținut doar cum să cotcodească.
        // Așadar, codul de mai jos va provoca o eroare.
        // bird.fly();
    }
}
```


## `&dyn' și 'Box<dyn>'

3. 🌟🌟
```rust,editable

// Completați spațiile libere.
trait Draw {
    fn draw(&self) -> String;
}

impl Draw for u8 {
    fn draw(&self) -> String {
        format!("u8: {}", *self)
    }
}

impl Draw for f64 {
    fn draw(&self) -> String {
        format!("f64: {}", *self)
    }
}

fn main() {
    let x = 1.1f64;
    let y = 8u8;

    // Desenează x.
    draw_with_box(__);

    // Desenează y.
    draw_with_ref(&y);

    println!("Success!");
}

fn draw_with_box(x: Box<dyn Draw>) {
    x.draw();
}

fn draw_with_ref(x: __) {
    x.draw();
}
```

## Dispatch Static și Dinamic
Când folosim limite de trăsături pe generice, compilatorul generează implementări nongenerice ale funcțiilor și metodelor pentru fiecare tip concret pe care îl folosim în locul unui parametru generic de tip. Codul rezultat din monomorfizare realizează un dispatch static, adică când compilatorul știe ce metodă apelezi în timpul compilării.

Când folosim obiecte de trăsături, Rust trebuie să folosească un dispatch dinamic. Compilatorul nu cunoaște toate tipurile care ar putea fi folosite cu codul care utilizează obiecte de trăsături, așa că nu știe ce metodă implementată pe ce tip să apeleze. În schimb, la runtime, Rust folosește pointerii din obiectul de trăsătură pentru a ști ce metodă să apeleze. Există un cost la runtime atunci când se face această căutare, care nu apare în cazul dispatch-ului static. Dispatch-ul dinamic împiedică, de asemenea, compilatorul să aleagă să încorporeze codul unei metode, ceea ce, la rândul său, împiedică unele optimizări.

Cu toate acestea, obținem flexibilitate suplimentară atunci când folosim dispatch-ul dinamic.

4. 🌟🌟
```rust,editable

trait Foo {
    fn method(&self) -> String;
}

impl Foo for u8 {
    fn method(&self) -> String { format!("u8: {}", *self) }
}

impl Foo for String {
    fn method(&self) -> String { format!("string: {}", *self) }
}

// IMPLEMENTAȚI mai jos cu generice.
fn static_dispatch...

// Implementați mai jos cu obiecte de trăsături.
fn dynamic_dispatch...

fn main() {
    let x = 5u8;
    let y = "Hello".to_string();

    static_dispatch(x);
    dynamic_dispatch(&y);

    println!("Success!");
}
```

## Obiect sigur
Puteți face doar din trăsături care sunt sigure pentru obiecte trăsătură. O trăsătură este sigură pentru obiect dacă toate metodele definite în trăsătură au următoarele proprietăți:

- Tipul de returnare nu este Self.
- Nu există parametri de tip generic.

5. 🌟🌟🌟🌟
```rust,editable

// Folosiți cel puțin două abordări pentru a face codul să funcționeze.
// NU adăugați/eliminați nici o linie de cod.
trait MyTrait {
    fn f(&self) -> Self;
}

impl MyTrait for u32 {
    fn f(&self) -> Self { 42 }
}

impl MyTrait for String {
    fn f(&self) -> Self { self.clone() }
}

fn my_function(x: Box<dyn MyTrait>)  {
    x.f()
}

fn main() {
    my_function(Box::new(13_u32));
    my_function(Box::new(String::from("abc")));

    println!("Success!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
