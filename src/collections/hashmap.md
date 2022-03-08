# HashMap
Where vectors store values by an integer index, HashMaps store values by key. It is a hash map implemented with quadratic probing and SIMD lookup. By default, `HashMap` uses a hashing algorithm selected to provide resistance against HashDoS attacks.

The default hashing algorithm is currently `SipHash 1-3`, though this is subject to change at any point in the future. While its performance is very competitive for medium sized keys, other hashing algorithms will outperform it for small keys such as integers as well as large keys such as long strings, though those algorithms will typically not protect against attacks such as HashDoS.

The hash table implementation is a Rust port of Google’s [SwissTable](https://abseil.io/blog/20180927-swisstables). The original C++ version of SwissTable can be found [here](https://github.com/abseil/abseil-cpp/blob/master/absl/container/internal/raw_hash_set.h), and this [CppCon talk](https://www.youtube.com/watch?v=ncHmEUmJZf4) gives an overview of how the algorithm works.


### Basic Operations
1. 🌟🌟

```rust,editbale
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
 `HashMap` keys can be booleans, integers, strings, or any other type that implements the Eq and Hash traits. 


### Capacity
Like vectors, HashMaps are growable, but HashMaps can also shrink themselves when they have excess space. You can create a HashMap with a certain starting capacity using HashMap::with_capacity(uint), or use HashMap::new() to get a HashMap with a default initial capacity (recommended).