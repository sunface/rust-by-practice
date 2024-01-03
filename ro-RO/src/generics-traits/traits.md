# Trăsături
O trăsătură spune compilatorului Rust despre funcționalitatea pe care un anumit tip o are și pe care o poate împărtăși cu alte tipuri. Putem folosi trăsături pentru a defini comportament comun într-un mod abstract. Putem utiliza constrângeri de trăsături pentru a specifica că un tip generic poate fi orice tip care are anumit comportament.

> Notă: Trăsăturile sunt similare cu interfețele în alte limbaje, deși există unele diferențe.

## Exemple
```rust,editable

struct Sheep { naked: bool, name: String }

trait Animal {
    // Semnătura funcției asociate; `Self` se referă la tipul implementatorului.
    fn new(name: String) -> Self;

    // Semnăturile metodei; acestea vor returna un șir de caractere.
    fn name(&self) -> String;
    
    fn noise(&self) -> String;

    // Trăsăturile pot oferi definiții implicite pentru metode.
    fn talk(&self) {
        println!("{} says {}", self.name(), self.noise());
    }
}

impl Sheep {
    fn is_naked(&self) -> bool {
        self.naked
    }

    fn shear(&mut self) {
        if self.is_naked() {
            // Metodele implementatorului pot folosi metodele trăsăturii ale implementatorului.
            println!("{} is already naked...", self.name());
        } else {
            println!("{} gets a haircut!", self.name);

            self.naked = true;
        }
    }
}

// Implementați trăsătura `Animal` pentru `Sheep`.
impl Animal for Sheep {
    // `Self` este tipul implementator: `Sheep`.
    fn new(name: String) -> Sheep {
        Sheep { name: name, naked: false }
    }

    fn name(&self) -> String {
        self.name.clone()
    }

    fn noise(&self) -> String {
        if self.is_naked() {
            "baaaaah?".to_string()
        } else {
            "baaaaah!".to_string()
        }
    }
    
    // Metodele implicite ale trăsăturii pot fi suprascrise.
    fn talk(&self) {
        // De exemplu, putem adăuga o contemplare scurtă.
        println!("{} pauses briefly... {}", self.name, self.noise());
    }
}

fn main() {
    // Este necesară anotarea tipului în acest caz.
    let mut dolly: Sheep = Animal::new("Dolly".to_string());
    // TODO ^ Încercați să eliminați anotările de tip.

    dolly.talk();
    dolly.shear();
    dolly.talk();
}
```

## Exercises
1. 🌟🌟
```rust,editable

// Completați cele două blocuri `impl` pentru a face codul să funcționeze.
// NU modificați codul din `main`.
trait Hello {
    fn say_hi(&self) -> String {
        String::from("hi")
    }

    fn say_something(&self) -> String;
}

struct Student {}
impl Hello for Student {
}
struct Teacher {}
impl Hello for Teacher {
}

fn main() {
    let s = Student {};
    assert_eq!(s.say_hi(), "hi");
    assert_eq!(s.say_something(), "I'm a good student");

    let t = Teacher {};
    assert_eq!(t.say_hi(), "Hi, I'm your new teacher");
    assert_eq!(t.say_something(), "I'm not a bad teacher");

    println!("Success!");
}
```

### Derivare
Compilatorul este capabil să furnizeze implementări de bază pentru unele trăsături prin intermediul atributului `#[derive]`. Pentru mai multe informații, vă rugăm să vizitați [aici](https://doc.rust-lang.org/book/appendix-03-derivable-traits.html).

2. 🌟🌟
```rust,editable

// `Centimeters`, o structură tuplu care poate fi comparată
#[derive(PartialEq, PartialOrd)]
struct Centimeters(f64);

// `Inches`, o structură tuplu care poate fi afișată
#[derive(Debug)]
struct Inches(i32);

impl Inches {
    fn to_centimeters(&self) -> Centimeters {
        let &Inches(inches) = self;

        Centimeters(inches as f64 * 2.54)
    }
}

// ADAUGAȚI unele atribute pentru a face codul să funcționeze!
// NU modificați alt cod!
struct Seconds(i32);

fn main() {
    let _one_second = Seconds(1);

    println!("One second looks like: {:?}", _one_second);
    let _this_is_true = (_one_second == _one_second);
    let _this_is_false = (_one_second > _one_second);

    let foot = Inches(12);

    println!("One foot equals {:?}", foot);

    let meter = Centimeters(100.0);

    let cmp =
        if foot.to_centimeters() < meter {
            "smaller"
        } else {
            "bigger"
        };

    println!("One foot is {} than one meter.", cmp);
}
```


### Operator
În Rust, mulți dintre operatori pot fi suprascriși prin intermediul trăsăturilor. Adică, unele operații pot fi utilizate pentru a realiza diferite sarcini în funcție de argumentele lor de intrare. Acest lucru este posibil pentru că operatorii sunt zahăr sintactic pentru apeluri de metode. De exemplu, operatorul '+' în a + b apelează metoda 'add' (ca în a.add(b)). Această metodă 'add' face parte din trăsătura 'Add'. Prin urmare, operatorul '+' poate fi folosit de orice implementator al trăsăturii 'Add'.

3. 🌟🌟
```rust,editable

use std::ops;

// Implementați funcția `înmulți` pentru a face codul să funcționeze.
// Așa cum s-a menționat mai sus, `+` are nevoie de `T` să implementeze trăsătura `std::ops::Add`.
// Deci, ce se întâmplă cu `*`? Puteți găsi răspunsul aici: https://doc.rust-lang.org/core/ops/
fn multipl

fn main() {
    assert_eq!(6, multiply(2u8, 3u8));
    assert_eq!(5.0, multiply(1.0, 5.0));

    println!("Success!");
}
```

4. 🌟🌟🌟
```rust,editable

// Remediați erorile, NU modificați codul din `main`.
use std::ops;

struct Foo;
struct Bar;

struct FooBar;

struct BarFoo;

// Trăsătura `std::ops::Add` este utilizată pentru a specifica funcționalitatea `+`.
// Aici, facem `Add<Bar>` - trăsătura pentru adunare cu un RHS de tip `Bar`.
// Blocul următor implementează operația: Foo + Bar = FooBar
impl ops::Add<Bar> for Foo {
    type Output = FooBar;

    fn add(self, _rhs: Bar) -> FooBar {
        FooBar
    }
}

impl ops::Sub<Foo> for Bar {
    type Output = BarFoo;

    fn sub(self, _rhs: Foo) -> BarFoo {
        BarFoo
    }
}

fn main() {
    // NU modificați codul de mai jos.
    // Trebuie să derivați unele trăsături pentru FooBar pentru a le face comparabile.
    assert_eq!(Foo + Bar, FooBar);
    assert_eq!(Foo - Bar, BarFoo);

    println!("Success!");
}
```

### Folosirea trăsăturii ca parametru pentru funcții
În loc de un tip concret pentru parametrul elementului, specificăm cuvântul cheie impl și numele trăsăturii. Acest parametru acceptă orice tip care implementează trăsătura specificată. 

5. 🌟🌟🌟
```rust,editable

// Implementați `fn rezumat` pentru a face codul să funcționeze.
// Remediați erorile fără a elimina nicio linie de cod
trait Summary {
    fn summarize(&self) -> String;
}

#[derive(Debug)]
struct Post {
    title: String,
    author: String,
    content: String,
}

impl Summary for Post {
    fn summarize(&self) -> String {
        format!("The author of post {} is {}", self.title, self.author)
    }
}

#[derive(Debug)]
struct Weibo {
    username: String,
    content: String,
}

impl Summary for Weibo {
    fn summarize(&self) -> String {
        format!("{} published a weibo {}", self.username, self.content)
    }
}

fn main() {
    let post = Post {
        title: "Popular Rust".to_string(),
        author: "Sunface".to_string(),
        content: "Rust is awesome!".to_string(),
    };
    let weibo = Weibo {
        username: "sunface".to_string(),
        content: "Weibo seems to be worse than Tweet".to_string(),
    };

    summary(post);
    summary(weibo);

    println!("{:?}", post);
    println!("{:?}", weibo);
}

// Implementați `fn summary` mai jos.

```

### Returnarea tipurilor care implementează trăsături
Putem folosi, de asemenea, sintaxa 'impl Trait' în poziția de returnare pentru a returna o valoare de un anumit tip care implementează o trăsătură.

Cu toate acestea, puteți utiliza 'impl Trait' doar atunci când returnați un singur tip, folosiți Obiectele Trăsăturii atunci când aveți nevoie cu adevărat să returnați mai multe tipuri.

6. 🌟🌟
```rust,editable

struct Sheep {}
struct Cow {}

trait Animal {
    fn noise(&self) -> String;
}

impl Animal for Sheep {
    fn noise(&self) -> String {
        "baaaaah!".to_string()
    }
}

impl Animal for Cow {
    fn noise(&self) -> String {
        "moooooo!".to_string()
    }
}

/ Returnează unele structuri care implementează Animal, dar nu știm care la momentul compilării.
// REMEDIAȚI erorile aici, puteți face o glumă aleatoare, sau puteți utiliza un obiect de trăsătură.
fn random_animal(random_number: f64) -> impl Animal {
    if random_number < 0.5 {
        Sheep {}
    } else {
        Cow {}
    }
}

fn main() {
    let random_number = 0.234;
    let animal = random_animal(random_number);
    println!("You've randomly chosen an animal, and it says {}", animal.noise());
}
```

### Condiție de trăsătură
Sintaxa 'impl Trait' funcționează pentru cazurile simple, dar este de fapt zahăr sintactic pentru o formă mai lungă, numită condiție de trăsătură.

Atunci când lucrați cu generice, parametrii de tip trebuie să utilizeze adesea trăsături ca limite pentru a specifica funcționalitatea pe care un tip o implementează.

7. 🌟🌟
```rust,editable
fn main() {
    assert_eq!(sum(1, 2), 3);
}

// Implementați `fn sum` cu condiție de trăsătură în două moduri.
fn sum<T>(x: T, y: T) -> T {
    x + y
}
```
8. 🌟🌟
```rust,editable

// REMEDIAȚI erorile.
struct Pair<T> {
    x: T,
    y: T,
}

impl<T> Pair<T> {
    fn new(x: T, y: T) -> Self {
        Self {
            x,
            y,
        }
    }
}

impl<T: std::fmt::Debug + PartialOrd> Pair<T> {
    fn cmp_display(&self) {
        if self.x >= self.y {
            println!("The largest member is x = {:?}", self.x);
        } else {
            println!("The largest member is y = {:?}", self.y);
        }
    }
}

struct Unit(i32);

fn main() {
    let pair = Pair{
        x: Unit(1),
        y: Unit(3)
    };

    pair.cmp_display();
}
```

9. 🌟🌟🌟
```rust,editable

// Completează spațiile libere pentru a face codul să funcționeze
fn example1() {
    // T: Trait este modalitatea obișnuită de utilizare.
    // T: Fn(u32) -> u32 specifică că putem trece doar o închidere la T.
    struct Cacher<T: Fn(u32) -> u32> {
        calculation: T,
        value: Option<u32>,
    }

    impl<T: Fn(u32) -> u32> Cacher<T> {
        fn new(calculation: T) -> Cacher<T> {
            Cacher {
                calculation,
                value: None,
            }
        }

        fn value(&mut self, arg: u32) -> u32 {
            match self.value {
                Some(v) => v,
                None => {
                    let v = (self.calculation)(arg);
                    self.value = Some(v);
                    v
                },
            }
        }
    }

    let mut cacher = Cacher::new(|x| x+1);
    assert_eq!(cacher.value(10), __);
    assert_eq!(cacher.value(15), __);
}


fn example2() {
    // Putem folosi și where pentru a construi T.
    struct Cacher<T>
        where T: Fn(u32) -> u32,
    {
        calculation: T,
        value: Option<u32>,
    }

    impl<T> Cacher<T>
        where T: Fn(u32) -> u32,
    {
        fn new(calculation: T) -> Cacher<T> {
            Cacher {
                calculation,
                value: None,
            }
        }

        fn value(&mut self, arg: u32) -> u32 {
            match self.value {
                Some(v) => v,
                None => {
                    let v = (self.calculation)(arg);
                    self.value = Some(v);
                    v
                },
            }
        }
    }

    let mut cacher = Cacher::new(|x| x+1);
    assert_eq!(cacher.value(20), __);
    assert_eq!(cacher.value(25), __);
}



fn main() {
    example1();
    example2();

    println!("Success!");
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)