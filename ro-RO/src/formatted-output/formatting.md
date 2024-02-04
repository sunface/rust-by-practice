# Formatare

## Argumente poziționale

1.🌟🌟
```rust,editable
/* Completați spațiile libere */
fn main() {
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob"); // => Alice, this is Bob. Bob, this is Alice
    assert_eq!(format!("{1}{0}", 1, 2), __);
    assert_eq!(format!(__, 1, 2), "2112");
    println!("Success!");
}
```

## Argumente numite

2.🌟🌟
```rust,editable
fn main() {
    println!("{argument}", argument = "test"); // => "test"

    /* Completați spațiile libere */
    assert_eq!(format!("{name}{}", 1, __), "21");
    assert_eq!(format!(__,a = "a", b = 'b', c = 3 ), "a 3 b");
    
    /* Remediați eroarea */
    // Argumentul numit trebuie plasat după celelalte argumente
    println!("{abc} {1}", abc = "def", 2);

    println!("Success!");
}
```

## Umplerea cu spații a șirului (Padding cu string-uri)

3.🌟🌟 În mod implicit, puteți umple șirul cu spații
```rust,editable
fn main() {
    // Următoarele două umplu cu 5 spații
    println!("Hello {:5}!", "x"); // =>  "Hello x    !"  
    println!("Hello {:1$}!", "x", 5); // =>  "Hello x    !"

    /* Completați spațiile libere */
    assert_eq!(format!("Hello __!", 5, "x"), "Hello x    !");
    assert_eq!(format!("Hello __!", "x", width = 5), "Hello x    !");

    println!("Success!");
}
```

4.🌟🌟🌟  Aliniere la stânga, dreapta, umplere cu caractere specificate.
```rust,editable
fn main() {
    // Aliniere la stânga
    println!("Hello {:<5}!", "x"); // => Hello x    !
    // Aliniere la dreapta
    assert_eq!(format!("Hello __!", "x"), "Hello     x!");
    // Aliniere la centru
    assert_eq!(format!("Hello __!", "x"), "Hello   x  !");

    // Aliniere la stânga, umplere cu '&'
    assert_eq!(format!("Hello {:&<5}!", "x"), __);

    println!("Success!");
}
```

5.🌟🌟 Puteți umple numerele cu zerouri suplimentare.
```rust,editable
fn main() {
    println!("Hello {:5}!", 5); // => Hello     5!
    println!("Hello {:+}!", 5); // =>  Hello +5!
    println!("Hello {:05}!", 5); // => Hello 00005!
    println!("Hello {:05}!", -5); // => Hello -0005!

    /* Completați spațiile libere */
    assert!(format!("{number:0>width$}", number=1, width=6) == __);
    
    println!("Success!")
;}
```

## Precizie
6.🌟🌟 Precizie pentru virgulă mobilă
```rust,editable

/* Completați spațiile libere */
fn main() {
    let v = 3.1415926;

    println!("{:.1$}", v, 4); // same as {:.4} => 3.1416 

    assert_eq!(format!("__", v), "3.14");
    assert_eq!(format!("__", v), "+3.14");
    assert_eq!(format!("__", v), "3");

    println!("Success!");
}
```

7.🌟🌟🌟 Lungimea șirului
```rust,editable
fn main() {
    let s = "Hello, world!";

    println!("{0:.5}", s); // => Hello

    assert_eq!(format!("Hello __!", 3, "abcdefg"), "Hello abc!");

    println!("Success!");
}
```   

## Binare, octale, hexazecimale

- format!("{}", foo) -> "3735928559"
- format!("0x{:X}", foo) -> "0xDEADBEEF"
- format!("0o{:o}", foo) -> "0o33653337357"
  
8.🌟🌟
```rust,editable
fn main() {
    assert_eq!(format!("__", 27), "0b11011");
    assert_eq!(format!("__", 27), "0o33");
    assert_eq!(format!("__", 27), "0x1b");
    assert_eq!(format!("__", 27), "0x1B");

    println!("{:x}!", 27); // Hex fără prefix => 1b

    println!("{:#010b}", 27); // Padding binar cu 0, lățime = 10,  => 0b00011011

    println!("Success!");
}
```

## Capturarea mediului
9.🌟🌟🌟
```rust,editable
fn get_person() -> String {
    String::from("sunface")
}

fn get_format() -> (usize, usize) {
    (4, 1)
}


fn main() {
    let person = get_person();
    println!("Hello, {person}!");

    let (width, precision) = get_format();
    let scores = [("sunface", 99.12), ("jack", 60.34)];
    /* Faceți-l să afișeze:
    sunface: 99.1
    jack: 60.3
    */
    for (name, score) in scores {
        println!("{name}: __");
    }
}
```


## Altele

**Exemplu**
```rust,editable
fn main() {
    // Exponent
    println!("{:2e}", 1000000000); // => 1e9
    println!("{:2E}", 1000000000); // => 1E9

    // Adresa pointerului
    let v= vec![1, 2, 3];
    println!("{:p}", v.as_ptr()); // => 0x600002324050

    // Escape
    println!("Hello {{}}"); // => Hello {}
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
