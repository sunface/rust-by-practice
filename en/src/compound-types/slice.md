# slice todo
Slices are similar to arrays, but their length is not known at compile time. Instead, a slice is a two-word object, the first word is a pointer to the data, and the second word is the length of the slice. The word size is the same as usize, determined by the processor architecture eg 64 bits on an x86-64. Slices can be used to borrow a section of an array, and have the type signature &[T].


```rust,editable
fn main() {
    // we can ignore the array type, let the compiler infer it for us
    let arr: [_; 3] = ['a', 'b', 'c'];

    let arr1 = &arr[..2];
    
    // Arrays are stack allocated 
    // A char takes 4 byte in Rust: Unicode char
    println!("array occupies {} bytes", std::mem::size_of_val(&arr1));

    
}
```

```rust,editable
 // The trimmed string is a slice to the original string, hence no new
// allocation is performed
let chars_to_trim: &[char] = &[' ', ','];
let trimmed_str: &str = string.trim_matches(chars_to_trim);
println!("Used characters: {}", trimmed_str);
```

