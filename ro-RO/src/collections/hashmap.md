# HashMap
În timp ce vectorii stochează valori printr-un index întreg, HashMaps stochează valori prin intermediul unei chei. Este o hartă de tip hash implementată cu sondaj pătratic și căutare SIMD. În mod implicit, HashMap utilizează un algoritm de hash selectat pentru a oferi rezistență împotriva atacurilor HashDoS.

Algoritmul implicit de hash este în prezent SipHash 1-3, cu toate că acest lucru poate suferi modificări în orice moment în viitor. În timp ce performanța sa este foarte competitivă pentru chei de dimensiuni medii, alte algoritme de hash îl vor depăși pentru chei mici, cum ar fi cele întregi, precum și pentru chei mari, cum ar fi șiruri lungi, deși aceste algoritme nu vor oferi, de obicei, protecție împotriva atacurilor precum HashDoS.

Implementarea tabelului de hash este o adaptare Rust a SwissTable de la Google. Versiunea originală în C++ a SwissTable poate fi găsită aici, iar acest discurs CppCon oferă o prezentare a modului în care funcționează algoritmul.


### Operații de bază
1. 🌟🌟

```rust,editable

// UMPLEȚI spațiile goale si REZOLVAȚI erorile
use std::collections::HashMap;
fn main() {
    let mut scores = HashMap::new();
    scores.insert("Sunface", 98);
    scores.insert("Daniel", 95);
    scores.insert("Ashley", 69.0);
    scores.insert("Katie", "58");

    // `Get` întoarce Option<&V>
    let score = scores.get("Sunface");
    assert_eq!(score, Some(98));

    if scores.contains_key("Daniel") {
        // Indexarea returnează o valoare V
        let score = scores["Daniel"];
        assert_eq!(score, __);
        scores.remove("Daniel");
    }

    assert_eq!(scores.len(), __);

    for (name, score) in scores {
        println!("The score of {} is {}", name, score);
    }
}
```

2. 🌟🌟
```rust,editable

use std::collections::HashMap;
fn main() {
    let teams = [
        ("Chinese Team", 100),
        ("American Team", 10),
        ("France Team", 50),
    ];

    let mut teams_map1 = HashMap::new();
    for team in &teams {
        teams_map1.insert(team.0, team.1);
    }

    // IMPLEMENTAȚI team_map2 in două moduri
    // Sfat: o posibilă abordare ar fi folosirea metodei `collect`
    let teams_map2...

    assert_eq!(teams_map1, teams_map2);

    println!("Success!");
}
```

3. 🌟🌟
```rust,editable

// UMPLEȚI spațiile goale
use std::collections::HashMap;
fn main() {
    // Inferența de tip ne permite să omitem o semnătură explicită de tip 
    // (care ar fi HashMap<&str, u8> în acest exemplu).
    let mut player_stats = HashMap::new();

    // Introduce o cheie doar dacă aceasta nu există deja.
    player_stats.entry("health").or_insert(100);

    assert_eq!(player_stats["health"], __);

    // Introduce o cheie folosind o funcție care furnizează o valoare nouă 
    // doar dacă aceasta nu există deja.
    player_stats.entry("health").or_insert_with(random_stat_buff);
    assert_eq!(player_stats["health"], __);

    // Se asigură că o valoare este în intrare prin introducerea valorii implicite dacă este goală 
    // și returnează o referință mutabilă la valoarea din intrare.
    let health = player_stats.entry("health").or_insert(50);
    assert_eq!(health, __);
    *health -= 50;
    assert_eq!(*health, __);

    println!("Success!");
}

fn random_stat_buff() -> u8 {
    // Ar putea returna efectiv o valoare aleatoare aici
    // să returnăm pur și simplu o valoare fixă pentru moment.
    42
}
```

### Cerințe pentru cheile HashMap
Orice tip care implementează trăsăturile Eq și Hash poate fi o cheie în HashMap. Aceasta include:

- bool (deși nu foarte util deoarece există doar două chei posibile)
- int, uint și toate variațiile acestora
- String și &str (sugestie: poți avea o HashMap indexată cu String și apela .get() cu un &str)

Observați că f32 și f64 nu implementează Hash, probabil din cauza problemelor de precizie cu numere în virgulă mobilă care ar face folosirea lor ca chei într-o hartă hash extrem de predispusă la erori.

Toate clasele de colecții implementează Eq și Hash dacă tipul lor conținut implementează, de asemenea, Eq și Hash. De exemplu, Vec<T> va implementa Hash dacă T implementează Hash.

4. 🌟🌟
```rust,editable

// CORECTEAZĂ erorile
// Sfaturi: derive este de obicei o modalitate bună de a implementa unele trăsături folosite frecvent
use std::collections::HashMap;

struct Viking {
    name: String,
    country: String,
}

impl Viking {
    /// Creates a new Viking.
    fn new(name: &str, country: &str) -> Viking {
        Viking {
            name: name.to_string(),
            country: country.to_string(),
        }
    }
}

fn main() {
    // Folosește o HashMap pentru a stoca punctele de viață ale vikingilor.
    let vikings = HashMap::from([
        (Viking::new("Einar", "Norway"), 25),
        (Viking::new("Olaf", "Denmark"), 24),
        (Viking::new("Harald", "Iceland"), 12),
    ]);

    // Folosește implementarea derivată pentru a afișa starea vikingilor.
    for (viking, health) in &vikings {
        println!("{:?} has {} hp", viking, health);
    }
}
```

### Capacitate
La fel ca vectorii, HashMaps sunt extensibili, dar acestea se pot și micșora atunci când au spațiu în exces. Puteți crea un HashMap cu o anumită capacitate inițială folosind HashMap::with_capacity(uint), sau puteți utiliza HashMap::new() pentru a obține un HashMap cu o capacitate inițială implicită (recomandată).

#### Exemplu
```rust,editable

use std::collections::HashMap;
fn main() {
    let mut map: HashMap<i32, i32> = HashMap::with_capacity(100);
    map.insert(1, 2);
    map.insert(3, 4);
    // Într-adevăr, capacitatea HashMap nu este 100, așa că nu putem compara egalitatea aici.
    assert!(map.capacity() >= 100);

    // Micșorează capacitatea hărții cu o limită inferioară. 
    // Nu va scădea mai jos decât limita furnizată, păstrând în același timp regulile interne 
    // și, posibil, lăsând spațiu conform politicii de redimensionare.

    map.shrink_to(50);
    assert!(map.capacity() >= 50);
    
    // Micșorează capacitatea hărții cât mai mult posibil.
    // Va scădea cât mai mult posibil, păstrând în același timp regulile interne 
    // și, posibil, lăsând spațiu conform politicii de redimensionare.
    map.shrink_to_fit();
    assert!(map.capacity() >= 2);
    println!("Success!");
}
```

### Proprietate (Ownership)
Pentru tipuri care implementează trăsătura Copy, cum ar fi i32, valorile sunt copiate în HashMap. Pentru valorile deținute precum String, valorile vor fi mutate, iar HashMap va deveni proprietarul acestora.

5. 🌟🌟
```rust,editable

// CORECTEAZĂ erorile cu cel mai mic număr de modificări
// NU șterge nicio linie de cod
use std::collections::HashMap;
fn main() {
  let v1 = 10;
  let mut m1 = HashMap::new();
  m1.insert(v1, v1);
  println!("v1 is still usable after inserting to hashmap : {}", v1);

  let v2 = "hello".to_string();
  let mut m2 = HashMap::new();
  // Proprietate (Ownership) mutată aici
  m2.insert(v2, v1);
    
  assert_eq!(v2, "hello");

  println!("Success!");
}
```

### Biblioteci Hash terțe
Dacă performanța lui SipHash 1-3 nu îndeplinește cerințele dvs., puteți găsi înlocuitori pe crates.io sau github.com.

Utilizarea unei funcționalități de hash terțe arată în felul următor:
```rust
use std::hash::BuildHasherDefault;
use std::collections::HashMap;

// Introdu o funcție de hash terță parte
use twox_hash::XxHash64;


let mut hash: HashMap<_, _, BuildHasherDefault<XxHash64>> = Default::default();
hash.insert(42, "the answer");
assert_eq!(hash.get(&42), Some(&"the answer"));
```

