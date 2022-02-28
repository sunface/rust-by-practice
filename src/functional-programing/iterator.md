# Iterator

```rust,editable
// (all the type annotations are superfluous)
// A reference to a string allocated in read only memory
let pangram: &'static str = "the quick brown fox jumps over the lazy dog";
println!("Pangram: {}", pangram);

// Iterate over words in reverse, no new string is allocated
println!("Words in reverse");
for word in pangram.split_whitespace().rev() {
    println!("> {}", word);
}

// Copy chars into a vector, sort and remove duplicates
let mut chars: Vec<char> = pangram.chars().collect();
chars.sort();
chars.dedup();
```