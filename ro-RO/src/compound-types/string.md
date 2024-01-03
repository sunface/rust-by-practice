# Șir de caractere (String)
Tipul literalului de șir "hello, world" este &str, de exemplu let s: &str = "hello, world".


### Str și &str
1. 🌟 Nu putem utiliza tipul str în moduri normale, dar putem utiliza &str.

```rust,editable

// Remedierează eroarea fără a adăuga noi linii
fn main() {
    let s: str = "hello, world";

    println!("Success!");
}
```


2. 🌟🌟 Putem utiliza doar str prin împachetarea acestuia, & poate fi folosit pentru a converti Box<str> în &str.

```rust,editable

// Remediază eroarea cu cel puțin două soluții
fn main() {
    let s: Box<str> = "hello, world".into();
    greetings(s)
}

fn greetings(s: &str) {
    println!("{}",s)
}
```

### String
Tipul String este definit în std și stocat ca un vector de octeți (Vec<u8>), dar este întotdeauna garantat să fie o secvență UTF-8 validă. String este alocat pe stivă, poate crește și nu este terminat cu null.

3. 🌟
```rust,editable

// Completează spațiul gol
fn main() {
    let mut s = __;
    s.push_str("hello, world");
    s.push('!');

    assert_eq!(s, "hello, world!");

    println!("Success!");
}
```

4. 🌟🌟🌟
```rust,editable

// Remediază toate erorile fără a adăuga linii noi
fn main() {
    let s = String::from("hello");
    s.push(',');
    s.push(" world");
    s += "!".to_string();

    println!("{}", s);
}
```

5. 🌟🌟 Funcția replace poate fi folosită pentru a înlocui o subșir.
```rust,editable

// Completează spațiul gol
fn main() {
    let s = String::from("I like dogs");
    // Alocă memorie nouă și stochează șirul modificat acolo
    let s1 = s.__("dogs", "cats");

    assert_eq!(s1, "I like cats");

    println!("Success!");
}
```

Mai multe metode pentru tipul String pot fi găsite în cadrul modulului [String](https://doc.rust-lang.org/std/string/struct.String.html).

6. 🌟🌟 Poți concatena doar un String cu un &str, iar proprietatea String-ului poate fi transferată către o altă variabilă.

```rust,editable

// Remediază erorile fără a elimina nicio linie
fn main() {
    let s1 = String::from("hello,");
    let s2 = String::from("world!");
    let s3 = s1 + s2; 
    assert_eq!(s3, "hello,world!");
    println!("{}", s1);
}
```

### &str și String
Contrar utilizării rare a tipului str, &str și String sunt folosite pretutindeni!

7. 🌟🌟 &str poate fi convertit la String în două moduri:
```rust,editable

// Remedierează eroarea cu cel puțin două soluții
fn main() {
    let s = "hello, world";
    greetings(s)
}

fn greetings(s: String) {
    println!("{}", s)
}
```

8. 🌟🌟 Putem folosi String::from sau to_string pentru a converti un &str la String.

```rust,editable

// Folosește două abordări pentru a remedia eroarea și fără a adăuga o linie nouă
fn main() {
    let s = "hello, world".to_string();
    let s1: &str = s;

    println!("Success!");
}
```

### Escapări pentru șiruri (String escapes)
9. 🌟 
```rust,editable
fn main() {
    // Poți utiliza caractere de escape pentru a scrie octeți folosind valorile lor hexadecimale
    // Completează spațiul gol de mai jos pentru a afișa "I'm writing Rust"
    let byte_escape = "I'm writing Ru\x73__!";
    println!("What are you doing\x3F (\\x3F means ?) {}", byte_escape);

    // ...sau puncte de cod Unicode.
    let unicode_codepoint = "\u{211D}";
    let character_name = "\"DOUBLE-STRUCK CAPITAL R\"";

    println!("Unicode character {} (U+211D) is called {}",
                unicode_codepoint, character_name );

    let long_string = "String literals
                        can span multiple lines.
                        The linebreak and indentation here \
                         can be escaped too!";
    println!("{}", long_string);
}
```

10. 🌟🌟🌟 Uneori există prea multe caractere care trebuie evitate sau este mult mai convenabil să scrii un șir așa cum este. Aici intervin literalurile brute de șiruri (raw string literals).

```rust,editable

/* Completează spațiul gol și remediază erorile */
fn main() {
    let raw_str = r"Escapes don't work here: \x3F \u{211D}";
    // Modifică linia de mai sus pentru a face codul să funcționeze
    assert_eq!(raw_str, "Escapes don't work here: ? ℝ");

    // Dacă ai nevoie de ghilimele într-un șir brut, adaugă o pereche de #
    let quotes = r#"And then I said: "There is no escape!""#;
    println!("{}", quotes);

    // Dacă ai nevoie de "# în șirul tău, folosește mai multe # în delimitator.
    // Poți folosi până la 65535 de #.
    let delimiter = r###"A string with "# in it. And even "##!"###;
    println!("{}", delimiter);

    // Completează spațiul gol
    let long_delimiter = __;
    assert_eq!(long_delimiter, "Hello, \"##\"");

    println!("Success!");
}
```

### Șir de octeți (Byte string)
Vrei un șir care nu este UTF-8? (Ține minte, str și String trebuie să fie UTF-8 valide). Sau poate vrei un tablou de octeți care este în mare parte text? Șirurile de octeți te ajută!

**Exemplu**:
```rust,editable
use std::str;

fn main() {
    // Reține că aceasta nu este de fapt un &str.
    let bytestring: &[u8; 21] = b"this is a byte string";

    // Tablourile de octeți nu au trăsătura Display, așa că afișarea lor este puțin limitată
    println!("A byte string: {:?}", bytestring);

    // Șirurile de octeți pot avea caractere de escape pentru octeți...
    let escaped = b"\x52\x75\x73\x74 as bytes";
    // ...Dar fără caractere de escape Unicode
    // let escaped = b"\u{211D} Nu este permis";
    println!("Some escaped bytes: {:?}", escaped);


    // Șirurile brute de octeți funcționează la fel ca și șirurile brute
    let raw_bytestring = br"\u{211D} is not escaped here";
    println!("{:?}", raw_bytestring);

    // Conversia unui tablou de octeți la str poate eșua
    if let Ok(my_str) = str::from_utf8(raw_bytestring) {
        println!("And the same as text: '{}'", my_str);
    }

    let _quotes = br#"You can also use "fancier" formatting, \
                    like with normal raw strings"#;

    // Șirurile de octeți nu trebuie să fie UTF-8
    let shift_jis = b"\x82\xe6\x82\xa8\x82\xb1\x82\xbb"; // "ようこそ" In SHIFT-JIS

    // Dar nu pot fi mereu convertite in `str`
    match str::from_utf8(shift_jis) {
        Ok(my_str) => println!("Conversion successful: '{}'", my_str),
        Err(e) => println!("Conversion failed: {:?}", e),
    };
}
```

O detaliere mai amplă a modalităților de a scrie literale de șir și caractere de escape este dată în [capitolul 'Jetoane'](https://doc.rust-lang.org/reference/tokens.html) 
a Referinței Rust.

### Index de șiruri
11. 🌟🌟🌟 Nu puteți folosi index pentru a accesa un caracter dintr-un șir, dar puteți folosi slice `&s1[start..end]`.

```rust,editable

fn main() {
    let s1 = String::from("hi,中国");
    let h = s1[0]; // Modificați această linie pentru a remedia eroarea, sfaturi: `h` ia doar 1 octet în format UTF8
    assert_eq!(h, "h");

    let h1 = &s1[3..5]; // Modificați această linie pentru a remedia eroarea, sfaturi: `中` are 3 octeți în format UTF8
    assert_eq!(h1, "中");

    println!("Success!");
}
```

### Operați pe șir UTF8
12. 🌟
```rust,editable

fn main() {
    // Completați spațiul liber pentru a imprima fiecare caracter din „你好，世界”
    for c in "你好，世界".__ {
        println!("{}", c)
    }
}
```

#### utf8_slice
Puteți folosi [utf8_slice](https://docs.rs/utf8_slice/1.0.0/utf8_slice/fn.slice.html) pentru a tăia șirul UTF8. Poate indexa caractere în loc de octeți.

**Exemplu**
```rust
use utf8_slice;
fn main() {
    let s = "The 🚀 goes to the 🌑!";

    let rocket = utf8_slice::slice(s, 4, 5);
    // Va fi egal cu „🚀”
}
```

> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie.
