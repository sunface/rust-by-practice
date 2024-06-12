# Formatarea ieșirilor

```rust,editable,ignore,mdbook-runnable
fn main() {
    // În general, `{}` va fi înlocuit automat cu orice argumente.
    // Acestea vor fi transformate în șiruri de caractere.
    println!("{} days", 31);

    // Fără un sufix, 31 devine un i32. Puteți schimba tipul lui 31
    // furnizând un sufix. Numărul 31i64, de exemplu, are tipul i64.

    // Există diverse modele opționale cu care acest lucru funcționează.
    // Pot fi utilizate argumente poziționale.
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // La fel de bine pot fi folosite argumente denumite.
    println!("{subject} {verb} {object}",
             object="the lazy dog",
             subject="the quick brown fox",
             verb="jumps over");

    // Se pot specifica formate speciale după un `:`.
    println!("{} of {:b} people know binary, the other half doesn't", 1, 2);

    // Se poate alinia textul la dreapta cu o lățime specificată.
    // Acesta va afișa "     1". 5 spații albe și un "1".
    println!("{number:>width$}", number=1, width=6);

    // Se pot adăuga zerouri suplimentare la numere.
    // Acesta va afișa "000001".
    println!("{number:0>width$}", number=1, width=6);

    // Rust se asigură chiar și că se folosește numărul corect de argumente.
    println!("My name is {0}, {1} {0}", "Bond");
    // FIXME ^ Adăugați argumentul lipsă: "James"

    // Creați o structură numită `Structure` care conține un `i32`.
    #[allow(dead_code)]
    struct Structure(i32);

    // Cu toate acestea, tipurile personalizate, cum ar fi această structură, necesită manipulări mai complicate.
    // Acest lucru nu va funcționa.
    println!("This struct `{}` won't print...", Structure(3));
    // FIXME ^ Comentați această linie.

    // Începând cu Rust 1.58 și mai sus, puteți captura direct argumentul din
    // variabila înconjurătoare. La fel ca în exemplul de mai sus, acesta va afișa
    // "     1". 5 spații albe și un "1".
    let number: f64 = 1.0;
    let width: usize = 6;
    println!("{number:>width$}");
}
```

[std::fmt][fmt] conține multe [traits][traits] care guvernează afișarea
textului. Forma de bază a două dintre cele mai importante este prezentată mai jos:

* fmt::Debug: Utilizează marcajul {:?}. Formatează textul în scopuri de depanare.
* fmt::Display: Utilizează marcajul {}. Formatează textul într-un mod mai elegant și prietenos pentru utilizator.

Aici, am folosit fmt::Display deoarece biblioteca standard oferă implementări pentru aceste tipuri. Pentru a tipări text pentru tipuri personalizate, sunt necesari mai mulți pași.

Implementarea trait-ului fmt::Display implementează automat trait-ul [ToString], care ne permite să [convertim] tipul în [String][string].
