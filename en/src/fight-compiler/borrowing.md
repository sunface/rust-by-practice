# Borrowing

1. 🌟🌟
```rust,editable
// FIX the error without removing any code line
struct test {
    list: Vec<i32>,
    a: i32
}

impl test {
    pub fn new() -> Self {
        test { list:vec![1,2,3,4,5,6,7], a:0 }
    }

    pub fn run(&mut self) {
        for i in self.list.iter() {
            self.do_something(*i)
        }

    }

    pub fn do_something(&mut self, n: i32) {
        self.a = n;
    }
}

fn main() {}
```
> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/fight-compiler/borrowing.md)(under the solutions path), but only use it when you need it :)
