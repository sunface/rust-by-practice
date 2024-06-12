# Struct

### Tipurile de structuri
1. 🌟 Trebuie să precizăm valori concrete pentru fiecare dintre câmpurile din struct.
```rust,editable

// Remediați eroarea
struct Person {
    name: String,
    age: u8,
    hobby: String
}
fn main() {
    let age = 30;
    let p = Person {
        name: String::from("sunface"),
        age,
    };

    println!("Success!");
} 
```

2. 🌟 Structura unității nu are câmpuri. Poate fi util atunci când trebuie să implementați o trăsătură pe un anumit tip, dar nu aveți date pe care doriți să le stocați în tipul în sine.
```rust,editable

struct Unit;
trait SomeTrait {
    // ...Unele comportamente definite aici.
}

// Nu ne pasă de ce câmpuri sunt în Unitate, dar ne pasă de comportamentele acesteia.
// Așadar, folosim o structură fără câmpuri și implementăm unele comportamente pentru aceasta
impl SomeTrait for Unit {  }
fn main() {
    let u = Unit;
    do_something_with_unit(u);

    println!("Success!");
} 

// Completați spațiul liber pentru a face codul să funcționeze
fn do_something_with_unit(u: __) {   }
```

3. 🌟🌟🌟 Structura tuplurilor arată similar cu tuplurile, are un sens adăugat pe care numele structurii o oferă, dar nu are câmpuri denumite. Este util atunci când vrei să dai un nume întregului tuplu, dar nu-ți pasă de numele câmpurilor.

```rust,editable

// Remediați eroarea și completați spațiile libere
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);
fn main() {
    let v = Point(__, __, __);
    check_color(v);

    println!("Success!");
}   

fn check_color(p: Color) {
    let (x, _, _) = p;
    assert_eq!(x, 0);
    assert_eq!(p.1, 127);
    assert_eq!(__, 255);
 }
```

### Operarea pe structuri

4. 🌟 Puteți face un întreg struct mutabil atunci când îl instanțiați, dar Rust nu ne permite să marchem doar anumite câmpuri ca mutabile.

```rust,editable

// Completați spațiul liber și remediați eroarea fără a adăuga/elimina o nouă linie
struct Person {
    name: String,
    age: u8,
}
fn main() {
    let age = 18;
    let p = Person {
        name: String::from("sunface"),
        age,
    };

    // Cum poți să crezi că fața solară are doar 18 ani?
    p.age = 30;

    // Completați spațiul liber
    __ = String::from("sunfei");

    println!("Success!");
}
```

5. 🌟 Folosind *field init sintaxa scurtă* pentru a reduce repetițiile.
```rust,editable

// Completați spațiul liber
struct Person {
    name: String,
    age: u8,
}
fn main() {
    println!("Success!");
} 

fn build_person(name: String, age: u8) -> Person {
    Person {
        age,
        __
    }
}
```

6. 🌟 Puteți crea instanță dintr-o altă instanță cu *sintaxă de actualizare a structurii*
```rust,editable

// Completați spațiul liber pentru ca codul să funcționez
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}
fn main() {
    let u1 = User {
        email: String::from("someone@example.com"),
        username: String::from("sunface"),
        active: true,
        sign_in_count: 1,
    };

    let u2 = set_email(u1);

    println!("Success!");
} 

fn set_email(u: User) -> User {
    User {
        email: String::from("contact@im.dev"),
        __
    }
}
```

### Afișați structurile
7. 🌟🌟 Putem folosi `#[derive(Debug)]` pentru a [face un struct imprimabil](https://doc.rust-lang.org/book/ch05-02-example-structs.html?highlight= %23%5Bderive(Debug)%5D#adding-useful-functionality-with-derived-traits).

```rust,editable

// Completați spațiile libere pentru ca codul să funcționeze
#[__]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let scale = 2;
    let rect1 = Rectangle {
        width: dbg!(30 * scale), // Imprimați informațiile de depanare în stderr și atribuiți valoarea `30 * scale` la `width`
        height: 50,
    };

    dbg!(&rect1); // Imprimă informațiile de depanare în stderr

    println!(__, rect1); // Imprimă informațiile de depanare în stdout
}
```

### Mutare parțială
În cadrul destructurarii unei singure variabile, pot fi utilizate în același timp atât legările de model prin mutare, cât și prin referință. Acest lucru va duce la o mutare parțială a variabilei, ceea ce înseamnă că părți ale variabilei vor fi mutate în timp ce celelalte părți rămân. Într-un astfel de caz, variabila părinte nu poate fi utilizată ulterior ca un întreg, totuși părțile care sunt doar referite (și nu sunt mutate) pot fi încă folosite.

#### Exemplu
```rust,editable

fn main() {
    #[derive(Debug)]
    struct Person {
        name: String,
        age: Box<u8>,
    }

    let person = Person {
        name: String::from("Alice"),
        age: Box::new(20),
    };

    // „name” este mutat din persoană, dar se face referire la „age”.
    let Person { name, ref age } = person;

    println!("The person's age is {}", age);

    println!("The person's name is {}", name);

    // Eroare! împrumut de valoare parțial mutată: se produce mutarea parțială a variabilei „person”.
    //println!("The person struct is {:?}", person);

    // `person` nu poate fi folosit, dar `person.age` poate fi folosit deoarece nu este mutat
    println!("The person's age from person struct is {}", person.age);
}
```


#### Exerciții

8. 🌟🌟
```rust,editable

// Remediați erorile pentru ca acest cod să funcționeze
#[derive(Debug)]
struct File {
    name: String,
    data: String,
}
fn main() {
    let f = File {
        name: String::from("readme.md"),
        data: "Rust By Practice".to_string()
    };

    let _name = f.name;

    // modifică NUMAI această linie
    println!("{}, {}, {:?}",f.name, f.data, f);
} 
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.