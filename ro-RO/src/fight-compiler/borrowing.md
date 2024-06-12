# Împrumut (Borrowing)

1. 🌟🌟
```rust,editable
// REZOLVĂ eroarea fără a elimina nicio linie de cod:
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
> Puteți găsi soluțiile [aici](https://github.com/sunface/rust-by-practice) (în cadrul căii soluțiilor), dar folosiți-le doar atunci când aveți nevoie. :)
