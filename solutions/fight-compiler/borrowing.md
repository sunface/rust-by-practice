1.

```rust
struct Test {
    list: Vec<i32>,
    a: i32
}

impl Test {
    pub fn new() -> Self {
        Test { list:vec![1,2,3,4,5,6,7], a:0 }
    }

    pub fn run(&mut self) {
        for i in self.list.clone().iter() {
            self.do_something(*i)
        }

    }

    pub fn do_something(&mut self, n: i32) {
        self.a = n;
    }
}

fn main() {}
```
