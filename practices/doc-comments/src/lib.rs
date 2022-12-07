//! # Doc comments
//! 
//! A library for showing how to use doc comments

pub mod compute;

/// Add one to the given value and return a new value
///
/// # Examples
///
/// ```
/// let arg = 5;
/// let answer = doc_comments::add_one(arg);
///
/// assert_eq!(6, answer);
/// ```
pub fn add_one(x: i32) -> i32 {
    x + 1
}



/** Add two to the given value and return a new value

# Examples

```
let arg = 5;
let answer = doc_comments::add_two(arg);

assert_eq!(7, answer);
```
*/
pub fn add_two(x: i32) -> i32 {
    x + 2
}


/// Add three to the given value and return a [`Option`] type
pub fn add_three(x: i32) -> Option<i32> {
    Some(x + 3)
}

mod a {
    /// Add four to the given value and return a [`Option`] type
    /// [`crate::MySpecialFormatter`]
    pub fn add_four(x: i32) -> Option<i32> {
        Some(x + 4)
    }
}

struct MySpecialFormatter;


