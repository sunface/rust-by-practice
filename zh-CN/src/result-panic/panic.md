# panic!
Rust 中最简单的错误处理方式就是使用 `panic`。它会打印出一条错误信息并打印出栈调用情况，最终结束当前线程:

- 若 panic 发生在 `main` 线程，那程序会随之退出
- 如果是在生成的( spawn )子线程中发生 panic, 那么当前的线程会结束，但是程序依然会继续运行


1. 🌟🌟
```rust,editable

// 填空
fn drink(beverage: &str) {
    if beverage == "lemonade" {
        println!("Success!");
        // 实现下面的代码
        __
     }

    println!("Exercise Failed if printing out this line!");
}

fn main() {
    drink(__);

    println!("Exercise Failed if printing out this line!");
}
```

## 常见的 panic
2. 🌟🌟
```rust,editable
// 修复所有的 panic，让代码工作
fn main() {
    assert_eq!("abc".as_bytes(), [96, 97, 98]);

    let v = vec![1, 2, 3];
    let ele = v[3];
    let ele = v.get(3).unwrap();

    // 大部分时候编译器是可以帮我们提前发现溢出错误，并阻止编译通过。但是也有一些时候，这种溢出问题直到运行期才会出现
    let v = production_rate_per_hour(2);

    divide(15, 0);

    println!("Success!")
}

fn divide(x:u8, y:u8) {
    println!("{}", x / y)
}

fn production_rate_per_hour(speed: u8) -> f64 {
    let cph: u8 = 221;
    match speed {
        1..=4 => (speed * cph) as f64,
        5..=8 => (speed * cph) as f64 * 0.9,
        9..=10 => (speed * cph) as f64 * 0.77,
        _ => 0 as f64,
    }
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed) / 60 as f64) as u32
}
```

### 详细的栈调用信息
默认情况下，栈调用只会展示最基本的信息: 
```shell
thread 'main' panicked at 'index out of bounds: the len is 3 but the index is 99', src/main.rs:4:5
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```

但是有时候，我们还希望获取更详细的信息:

3. 🌟
```shell
## 填空以打印全部的调用栈
## 提示: 你可以在之前的默认 panic 信息中找到相关线索
$ __ cargo run
thread 'main' panicked at 'assertion failed: `(left == right)`
  left: `[97, 98, 99]`,
 right: `[96, 97, 98]`', src/main.rs:3:5
stack backtrace:
   0: rust_begin_unwind
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/std/src/panicking.rs:498:5
   1: core::panicking::panic_fmt
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/panicking.rs:116:14
   2: core::panicking::assert_failed_inner
   3: core::panicking::assert_failed
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/panicking.rs:154:5
   4: study_cargo::main
             at ./src/main.rs:3:5
   5: core::ops::function::FnOnce::call_once
             at /rustc/9d1b2106e23b1abd32fce1f17267604a5102f57a/library/core/src/ops/function.rs:227:5
note: Some details are omitted, run with `RUST_BACKTRACE=full` for a verbose backtrace.
```

### `unwinding` 和 `abort`

当出现 `panic!` 时，程序提供了两种方式来处理终止流程：**栈展开**和**直接终止**。

其中，默认的方式就是 `栈展开`，这意味着 Rust 会回溯栈上数据和函数调用，因此也意味着更多的善后工作，好处是可以给出充分的报错信息和栈调用信息，便于事后的问题复盘。`直接终止`，顾名思义，不清理数据就直接退出程序，善后工作交与操作系统来负责。

对于绝大多数用户，使用默认选择是最好的，但是当你关心最终编译出的二进制可执行文件大小时，那么可以尝试去使用直接终止的方式，例如下面的配置修改 `Cargo.toml` 文件，实现在 `release`模式下遇到 `panic` 直接终止：

```rust
[profile.release]
panic = 'abort'
```


> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/result-panic/panic.md)找到答案(在 solutions 路径下)
