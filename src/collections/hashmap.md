# HashMap
Where vectors store values by an integer index, HashMaps store values by key. `HashMap` keys can be booleans, integers, strings, or any other type that implements the Eq and Hash traits. 

It is also a hash map implemented with quadratic probing and SIMD lookup. By default, `HashMap` uses a hashing algorithm selected to provide resistance against HashDoS attacks.

The default hashing algorithm is currently `SipHash 1-3`, though this is subject to change at any point in the future. While its performance is very competitive for medium sized keys, other hashing algorithms will outperform it for small keys such as integers as well as large keys such as long strings, though those algorithms will typically not protect against attacks such as HashDoS.

The hash table implementation is a Rust port of Google’s [SwissTable](https://abseil.io/blog/20180927-swisstables). The original C++ version of SwissTable can be found [here](https://github.com/abseil/abseil-cpp/blob/master/absl/container/internal/raw_hash_set.h), and this [CppCon talk](https://www.youtube.com/watch?v=ncHmEUmJZf4) gives an overview of how the algorithm works.


### Basic Operations

2. 🌟🌟
```rust,editable
use std::collections::HashMap;
fn main() {
    let teams = vec![
        ("Chinese Team", 100),
        ("American Team", 10),
        ("France Team", 50),
    ];

    let mut teams_map1 = HashMap::new();
    for team in &teams {
        teams_map1.insert(team.0, team.1);
    }

    // IMPLEMENT team_map2 with method collect
    let teams_map2...

    assert_eq!(teams_map1, teams_map2);
}
```

### Capacity
Like vectors, HashMaps are growable, but HashMaps can also shrink themselves when they have excess space. You can create a HashMap with a certain starting capacity using HashMap::with_capacity(uint), or use HashMap::new() to get a HashMap with a default initial capacity (recommended).