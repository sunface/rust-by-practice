# Trăsături Avansate

## Tipuri asociate
Utilizarea "Tipurilor asociate" îmbunătățește în mod general citibilitatea codului prin mutarea tipurilor interne local într-o trăsătură ca tipuri de ieșire. De exemplu:
```rust
pub trait CacheableItem: Clone + Default + fmt::Debug + Decodable + Encodable {
  type Address: AsRef<[u8]> + Clone + fmt::Debug + Eq + Hash;
  fn is_null(&self) -> bool;
}
```

Utilizarea tipului Adresa este mult mai clară și mai convenabilă decât 'AsRef<[u8]> + Clone + fmt::Debug + Eq + Hash'.

1. 🌟🌟🌟
```rust,editable

struct Container(i32, i32);

// UTILIZAREA tipurilor asociate pentru a reimplementa trăsătura Contine.
// trăsătură Contine {
//    tip A;
//    tip B;

trait Contains<A, B> {
    fn contains(&self, _: &A, _: &B) -> bool;
    fn first(&self) -> i32;
    fn last(&self) -> i32;
}

impl Contains<i32, i32> for Container {
    fn contains(&self, number_1: &i32, number_2: &i32) -> bool {
        (&self.0 == number_1) && (&self.1 == number_2)
    }
    // Preia primul număr.
    fn first(&self) -> i32 { self.0 }

    // Preia ultimul număr.
    fn last(&self) -> i32 { self.1 }
}

fn difference<A, B, C: Contains<A, B>>(container: &C) -> i32 {
    container.last() - container.first()
}

fn main() {
    let number_1 = 3;
    let number_2 = 10;

    let container = Container(number_1, number_2);

    println!("Does container contain {} and {}: {}",
        &number_1, &number_2,
        container.contains(&number_1, &number_2));
    println!("First number: {}", container.first());
    println!("Last number: {}", container.last());
    
    println!("The difference is: {}", difference(&container));
}
```

## Parametri de Tip Generic Implicit
Atunci când utilizăm parametri de tip generic, putem specifica un tip concret implicit pentru tipul generic. Acest lucru elimină necesitatea ca implementatorii trăsăturii să specifice un tip concret dacă tipul implicit funcționează.

2. 🌟🌟
```rust,editable

use std::ops::Sub;

#[derive(Debug, PartialEq)]
struct Point<T> {
    x: T,
    y: T,
}

// COMPLETAȚI spațiile libere în trei moduri: două dintre ele folosesc parametrii de tip generic implicit, cealaltă nu.
// Observați că implementarea folosește tipul asociat `Output`.
impl __ {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        Point {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

fn main() {
    assert_eq!(Point { x: 2, y: 3 } - Point { x: 1, y: 0 },
        Point { x: 1, y: 3 });

    println!("Success!");
}
```

## Sintaxă Complet Calificată
Nimic în Rust nu împiedică o trăsătură să aibă o metodă cu același nume ca o altă metodă a unei alte trăsături, iar Rust nu vă împiedică să implementați ambele trăsături pe un singur tip. De asemenea, este posibil să implementați o metodă direct pe tip cu același nume ca metodele din trăsături.

Atunci când apelăm metode cu același nume, trebuie să folosim Sintaxa Complet Calificată.

#### Exemplu
```rust,editable
trait UsernameWidget {
    // Preia numele de utilizator selectat din acest widget
    fn get(&self) -> String;
}

trait AgeWidget {
    // Preia vârsta selectată din acest widget
    fn get(&self) -> u8;
}

// O formă cu un WidgetNumeUtilizator(UsernameWidget), cât și un WidgetVârstă(AgeWidget).
struct Form {
    username: String,
    age: u8,
}

impl UsernameWidget for Form {
    fn get(&self) -> String {
        self.username.clone()
    }
}

impl AgeWidget for Form {
    fn get(&self) -> u8 {
        self.age
    }
}

fn main() {
    let form = Form{
        username: "rustacean".to_owned(),
        age: 28,
    };

    // Dacă decomentați această linie, veți primi o eroare care spune
    // "s-au găsit multiple `get`". Pentru că, în cele din urmă, există mai multe metode
    // denumite `get`.
    // println!("{}", form.get());
    
    let username = UsernameWidget::get(&form);
    assert_eq!("rustacean".to_owned(), username);
    let age = AgeWidget::get(&form); // Puteți folosi și `<Form as AgeWidget>::get`
    assert_eq!(28, age);

    println!("Success!");
}
```

#### Exerciții
3. 🌟🌟
```rust,editable
trait Pilot {
    fn fly(&self) -> String;
}

trait Wizard {
    fn fly(&self) -> String;
}

struct Human;

impl Pilot for Human {
    fn fly(&self) -> String {
        String::from("This is your captain speaking.")
    }
}

impl Wizard for Human {
    fn fly(&self) -> String {
        String::from("Up!")
    }
}

impl Human {
    fn fly(&self) -> String {
        String::from("*waving arms furiously*")
    }
}

fn main() {
    let person = Human;

    assert_eq!(__, "This is your captain speaking.");
    assert_eq!(__, "Up!");

    assert_eq!(__, "*waving arms furiously*");

    println!("Success!");
}
```

## Supertrăsături (Supertraits)
Uneori, s-ar putea să aveți nevoie ca o trăsătură să utilizeze funcționalitatea altei trăsături (ca "moștenirea" în alte limbaje). În acest caz, trebuie să vă bazați pe trăsătura dependentă pentru a fi implementată, fiind o 'supertrăsătură' a trăsăturii pe care o implementați.

4. 🌟🌟🌟
```rust,editable

trait Person {
    fn name(&self) -> String;
}

// Persoană este o supertrăsătură a Studentului.
// Implementarea Studentului vă cere să implementați și Persoana.
trait Student: Person {
    fn university(&self) -> String;
}

trait Programmer {
    fn fav_language(&self) -> String;
}

// CompSciStudent (student de informatică) este o subtrăsătură a ambelor Programator 
// și Student. Implementarea CompSciStudent vă cere să implementați ambele supertrăsături.
trait CompSciStudent: Programmer + Student {
    fn git_username(&self) -> String;
}

fn comp_sci_student_greeting(student: &dyn CompSciStudent) -> String {
    format!(
        "My name is {} and I attend {}. My favorite language is {}. My Git username is {}",
        student.name(),
        student.university(),
        student.fav_language(),
        student.git_username()
    )
}

struct CSStudent {
    name: String,
    university: String,
    fav_language: String,
    git_username: String
}

// IMPLEMENTAȚI trăsăturile necesare pentru StudCS pentru a face codul să funcționeze
impl ...

fn main() {
    let student = CSStudent {
        name: "Sunfei".to_string(),
        university: "XXX".to_string(),
        fav_language: "Rust".to_string(),
        git_username: "sunface".to_string()
    };

    // COMPLETAȚI spațiul liber
    println!("{}", comp_sci_student_greeting(__));
}
```

## Reguli Orfane
Nu putem implementa trăsături externe pe tipuri externe. De exemplu, nu putem implementa trăsătura 'Display' pentru 'Vec<T>' în propria noastră create, deoarece 'Display' și 'Vec<T>' sunt definite în biblioteca standard și nu sunt locale în cadrul createi noastre.

Această restricție este adesea numită regula orfan, numită astfel deoarece tipul părinte nu este prezent. Această regulă asigură că codul altor persoane nu poate afecta codul vostru și invers.

Este posibil să se ocolească această restricție folosind șablonul newtype, care implică crearea unui nou tip într-o structură de tupluri.

5. 🌟🌟
```rust,editable
use std::fmt;

// DEFINIȚI un șablon newtype `Frumos` aici


impl fmt::Display for Pretty {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "\"{}\"", self.0.clone() + ", world")
    }
}

fn main() {
    let w = Pretty("hello".to_string());
    println!("w = {}", w);
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice)(under the solutions path), but only use it when you need it :)