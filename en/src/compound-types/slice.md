# slice todo


```rust,editable
 // The trimmed string is a slice to the original string, hence no new
// allocation is performed
let chars_to_trim: &[char] = &[' ', ','];
let trimmed_str: &str = string.trim_matches(chars_to_trim);
println!("Used characters: {}", trimmed_str);
```