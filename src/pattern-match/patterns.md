# Patterns

🌟🌟 Using `|` to match several values.
```rust,editable

fn main() {}
fn match_number(n: i32) {
    match n {
        // match a single value
        1 => println!("One!"),
        // fill in the blank with `|`, DON'T use `..` ofr `..=`
        __ => println!("match 2 -> 5"),
        // match an inclusive range
        6..=10 => {
            println!("match 6 -> 10")
        },
        _ => {
            println!("match 11 -> +infinite")
        }
    }
}
```

