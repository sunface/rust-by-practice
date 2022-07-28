# 函数
1. 🌟🌟🌟
```rust,editable

fn main() {
    // 不要修改下面两行代码!
    let (x, y) = (1, 2);
    let s = sum(x, y);

    assert_eq!(s, 3);
}

fn sum(x, y: i32) {
    x + y;
}
```


2. 🌟🌟
```rust,editable
fn main() {
   print();
}

// 使用另一个类型来替代 i32
fn print() -> i32 {
   println!("hello,world");
}
```


3. 🌟🌟🌟

```rust,editable
// 用两种方法求解
fn main() {
    never_return();
}

fn never_return() -> ! {
    // 实现这个函数，不要修改函数签名!
    
}
```

4. 🌟🌟 发散函数( Diverging function )不会返回任何值，因此它们可以用于替代需要返回任何值的地方
```rust,editable

fn main() {
    println!("Success!");
}

fn get_option(tp: u8) -> Option<i32> {
    match tp {
        1 => {
            // TODO
        }
        _ => {
            // TODO
        }
    };
    
    // 这里与其返回一个 None，不如使用发散函数替代
    never_return_fn()
}

// 使用三种方法实现以下发散函数
fn never_return_fn() -> ! {
    
}
```

5. 🌟🌟
```rust,editable

fn main() {
    // 填空
    let b = __;

    let _v = match b {
        true => 1,
        // 发散函数也可以用于 `match` 表达式，用于替代任何类型的值
        false => {
            println!("Success!");
            panic!("we have no value for `false`, but we can panic")
        }
    };

    println!("Exercise Failed if printing out this line!");
}
```

> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/basic-types/functions.md)找到答案(在 solutions 路径下) 
