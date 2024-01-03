# Proprietate (Ownership)

1. 🌟🌟 
```rust,editable

fn main() {
    // Folosiți cât mai multe abordări posibile pentru a face ca acest cod să funcționeze
    let x = String::from("hello, world");
    let y = x;
    println!("{},{}",x,y);
}
```

2. 🌟🌟
```rust,editable
// Nu modificați codul în funcția main!
fn main() {
    let s1 = String::from("hello, world");
    let s2 = take_ownership(s1);

    println!("{}", s2);
}

// Modificați doar codul de mai jos!
fn take_ownership(s: String) {
    println!("{}", s);
}
```


3. 🌟🌟
```rust,editable

fn main() {
    let s = give_ownership();
    println!("{}", s);
}

// Modificați doar codul de mai jos!
fn give_ownership() -> String {
    let s = String::from("hello, world");
    // Convertiți String în Vec
    let _s = s.into_bytes();
    s
}
```

4. 🌟🌟
```rust,editable
// Corectați eroarea fără a elimina linia de cod
fn main() {
    let s = String::from("hello, world");

    print_str(s);

    println!("{}", s);
}

fn print_str(s: String)  {
    println!("{}",s)
}
```

5. 🌟🌟 
```rust,editable
// Nu folosiți clone, folosiți copy în schimb
fn main() {
    let x = (1, 2, (), "hello".to_string());
    let y = x.clone();
    println!("{:?}, {:?}", x, y);
}
```

#### Mutabilitate
Mutabilitatea poate fi schimbată atunci când proprietatea este transferată.

6. 🌟
```rust,editable

fn main() {
    let s = String::from("hello, ");
    
    // Modificați doar această linie!
    let s1 = s;

    s1.push_str("world");

    println!("Success!");
}
```

7. 🌟🌟🌟
```rust,editable

fn main() {
    let x = Box::new(5);
    
    let ...      // Implementați această linie. Nu schimbați alte linii!
    
    *y = 4;
    
    assert_eq!(*x, 5);

    println!("Success!");
}
```

### Mutare Parțială
În cadrul destrămării unei variabile unice, atât legăturile de tip mutabil, cât și cele de tip referință pot fi folosite în același timp. Acest lucru va rezulta într-o mutare parțială a variabilei, ceea ce înseamnă că părți ale variabilei vor fi mutate în timp ce alte părți rămân. Într-un astfel de caz, variabila părinte nu poate fi folosită ulterior ca întreg, însă părțile care sunt doar referențiate (și nu mutate) pot fi în continuare utilizate.

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

    // `name` este mutat din person, dar `age` este referențiată
    let Person { name, ref age } = person;

    println!("The person's age is {}", age);

    println!("The person's name is {}", name);

    // Eroare! împrumut de valoare parțial mutată: `person` mutare parțială apare
    //println!("The person struct is {:?}", person);

    // `person`nu poate fi folosită, dar `person.age` poate fi folosită deoarece nu este mutată
    println!("The person's age from person struct is {}", person.age);
}
```

#### Exerciții

8. 🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

   let _s = t.0;

   // Modificați doar această linie, nu folosiți `_s`
   println!("{:?}", t);
}
```

9. 🌟🌟
```rust,editable

fn main() {
   let t = (String::from("hello"), String::from("world"));

    // Completați spațiile libere
    let (__, __) = __;

    println!("{:?}, {:?}, {:?}", s1, s2, t); // -> "hello", "world", ("hello", "world")
}
```


> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)