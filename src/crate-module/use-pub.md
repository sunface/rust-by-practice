# use and pub
1. 🌟 We can bring two types of the same name into the same scope with use, but you need `as` keyword.

```rust,editable
use std::fmt::Result;
use std::io::Result;

fn main() {}
```

2. 🌟🌟 If we are using multiple items defined in the same crate or module, then listing each item on its own line will take up too much verticall space.

```rust,editable

// FILL in the blank in two ways
// DON'T add new code line
use std::collections::__;

fn main() {
    let _c1:HashMap<&str, i32> = HashMap::new();
    let mut c2 = BTreeMap::new();
    c2.insert(1, "a");
    let _c3: HashSet<i32> = HashSet::new();
}
```

### Re-exporting names with `pub use`
3. 🌟🌟🌟 In our recently created package `hello-package`, add something to make the below code work
```rust,editable
fn main() {
    assert_eq!(hello_package::hosting::seat_at_table(), "sit down please");
     assert_eq!(hello_package::eat_at_restaurant(),"yummy yummy!");
}
```


### pub(in Crate) 
Sometimes we want an item only be public to a certain crate, then we can use the `pub(in Crate)` syntax.

#### Example
```rust,editable
pub mod a {
    pub const I: i32 = 3;

    fn semisecret(x: i32) -> i32 {
        use self::b::c::J;
        x + J
    }

    pub fn bar(z: i32) -> i32 {
        semisecret(I) * z
    }
    pub fn foo(y: i32) -> i32 {
        semisecret(I) + y
    }

    mod b {
        pub(in crate::a) mod c {
            pub(in crate::a) const J: i32 = 4;
        }
    }
}
```

### Full Code
The full code of `hello-package` is [here](https://github.com/sunface/rust-by-practice/tree/master/practices/hello-package).