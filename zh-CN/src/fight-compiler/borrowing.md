# 借用

1. 🌟🌟
```rust,editable
// 不删除任何代码，修复错误
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
> 参考答案：<https://github.com/sunface/rust-by-practice/blob/master/solutions/fight-compiler/borrowing.md>（solutions 路径），仅在需要时查看。

