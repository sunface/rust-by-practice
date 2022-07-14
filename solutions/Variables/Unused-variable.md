# Unused variables
## 7. Fix the warning below with :
- 🌟 Only one solution
- 🌟🌟 Two distinct solutions
```
Note: none of the solutions is to remove the line let x = 1
```
```rs
fn main() {
    let x = 1; 
}

// Warning: unused variable: `x`
```
# Solution:
```rs
fn main() {
    let x = 1; 
    print!("{}",x);
}
```
