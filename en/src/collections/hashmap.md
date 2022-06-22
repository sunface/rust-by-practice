# HashMap
Where vectors store values by an integer index, HashMaps store values by key. It is a hash map implemented with quadratic probing and SIMD lookup. By default, `HashMap` uses a hashing algorithm selected to provide resistance against HashDoS attacks.

The default hashing algorithm is currently `SipHash 1-3`, though this is subject to change at any point in the future. While its performance is very competitive for medium sized keys, other hashing algorithms will outperform it for small keys such as integers as well as large keys such as long strings, though those algorithms will typically not protect against attacks such as HashDoS.

The hash table implementation is a Rust port of Google’s [SwissTable](https://abseil.io/blog/20180927-swisstables). The original C++ version of SwissTable can be found [here](https://github.com/abseil/abseil-cpp/blob/master/absl/container/internal/raw_hash_set.h), and this [CppCon talk](https://www.youtube.com/watch?v=ncHmEUmJZf4) gives an overview of how the algorithm works.


### Basic Operations
1. 🌟🌟

```rust,editable

// FILL in the blanks and FIX the erros
use std::collections::HashMap;
fn main() {
    let mut scores = HashMap::new();
    scores.insert("Sunface", 98);
    scores.insert("Daniel", 95);
    scores.insert("Ashley", 69.0);
    scores.insert("Katie", "58");

    // get returns a Option<&V>
    let score = scores.get("Sunface");
    assert_eq!(score, Some(98));

    if scores.contains_key("Daniel") {
        // indexing return a value V
        let score = scores["Daniel"];
        assert_eq!(score, __);
        scores.remove("Daniel");
    }

    assert_eq!(scores.len(), __);

    for (name, score) in scores {
        println!("The score of {} is {}", name, score)
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

    // IMPLEMENT team_map2 in two ways
    // tips: one of the approaches is to use `collect` method
    let teams_map2...

    assert_eq!(teams_map1, teams_map2);

    println!("Success!")
}
```

3. 🌟🌟
```rust,editable

// FILL in the blanks
use std::collections::HashMap;
fn main() {
    // type inference lets us omit an explicit type signature (which
    // would be `HashMap<&str, u8>` in this example).
    let mut player_stats = HashMap::new();

    // insert a key only if it doesn't already exist
    player_stats.entry("health").or_insert(100);

    assert_eq!(player_stats["health"], __);

    // insert a key using a function that provides a new value only if it
    // doesn't already exist
    player_stats.entry("health").or_insert_with(random_stat_buff);
    assert_eq!(player_stats["health"], __);

    // Ensures a value is in the entry by inserting the default if empty, and returns
    // a mutable reference to the value in the entry.
    let health = player_stats.entry("health").or_insert(50);
    assert_eq!(health, __);
    *health -= 50;
    assert_eq!(*health, __);

    println!("Success!")
}

fn random_stat_buff() -> u8 {
    // could actually return some random value here - let's just return
    // some fixed value for now
    42
}
```

### Requirements of HashMap key
Any type that implements the `Eq` and `Hash` traits can be a key in `HashMap`. This includes:

- `bool` (though not very useful since there is only two possible keys)
- `int`, `uint`, and all variations thereof
- `String` and `&str` (tips: you can have a `HashMap` keyed by `String` and call `.get()` with an `&str`)
  
Note that `f32` and `f64` do not implement `Hash`, likely because [floating-point precision](https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems) errors would make using them as hashmap keys horribly error-prone.

All collection classes implement `Eq` and `Hash` if their contained type also respectively implements `Eq` and `Hash`. For example, `Vec<T>` will implement `Hash` if `T`implements `Hash`.

4. 🌟🌟
```rust,editable

// FIX the errors
// Tips: `derive` is usually a good way to implement some common used traits
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
    // Use a HashMap to store the vikings' health points.
    let vikings = HashMap::from([
        (Viking::new("Einar", "Norway"), 25),
        (Viking::new("Olaf", "Denmark"), 24),
        (Viking::new("Harald", "Iceland"), 12),
    ]);

    // Use derived implementation to print the status of the vikings.
    for (viking, health) in &vikings {
        println!("{:?} has {} hp", viking, health);
    }
}
```

### Capacity
Like vectors, HashMaps are growable, but HashMaps can also shrink themselves when they have excess space. You can create a `HashMap` with a certain starting capacity using `HashMap::with_capacity(uint)`, or use `HashMap::new()` to get a HashMap with a default initial capacity (recommended).

#### Example
```rust,editable

use std::collections::HashMap;
fn main() {
    let mut map: HashMap<i32, i32> = HashMap::with_capacity(100);
    map.insert(1, 2);
    map.insert(3, 4);
    // indeed ,the capacity of HashMap is not 100, so we can't compare the equality here.
    assert!(map.capacity() >= 100);

    // Shrinks the capacity of the map with a lower limit. It will drop
    // down no lower than the supplied limit while maintaining the internal rules
    // and possibly leaving some space in accordance with the resize policy.

    map.shrink_to(50);
    assert!(map.capacity() >= 50);

    // Shrinks the capacity of the map as much as possible. It will drop
    // down as much as possible while maintaining the internal rules
    // and possibly leaving some space in accordance with the resize policy.
    map.shrink_to_fit();
    assert!(map.capacity() >= 2);
    println!("Success!")
}
```

### Ownership
For types that implement the `Copy` trait, like `i32` , the values are copied into `HashMap`. For owned values like `String`, the values will be moved and `HashMap` will be the owner of those values.

5. 🌟🌟
```rust,editable
// FIX the errors with least changes
// DON'T remove any code line
use std::collections::HashMap;
fn main() {
  let v1 = 10;
  let mut m1 = HashMap::new();
  m1.insert(v1, v1);
  println!("v1 is still usable after inserting to hashmap : {}", v1);

  let v2 = "hello".to_string();
  let mut m2 = HashMap::new();
  // ownership moved here
  m2.insert(v2, v1);
    
  assert_eq!(v2, "hello");

  println!("Success!")
}
```

### Third-party Hash libs
If the performance of `SipHash 1-3` doesn't meet your requirements, you can find replacements in crates.io or github.com.

The usage of third-party hash looks like this:
```rust
use std::hash::BuildHasherDefault;
use std::collections::HashMap;
// introduce a third party hash function
use twox_hash::XxHash64;


let mut hash: HashMap<_, _, BuildHasherDefault<XxHash64>> = Default::default();
hash.insert(42, "the answer");
assert_eq!(hash.get(&42), Some(&"the answer"));
```

