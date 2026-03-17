# 格式化输出

```rust,editable,ignore,mdbook-runnable
fn main() {
    // `{}` 会被参数自动替换并转成字符串
    println!("{} days", 31);

    // 默认 31 是 i32，可用后缀改类型，比如 31i64

    // 可以使用位置参数
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // 也可以使用具名参数
    println!("{subject} {verb} {object}",
             object="the lazy dog",
             subject="the quick brown fox",
             verb="jumps over");

    // `:` 后可指定特殊格式
    println!("{} of {:b} people know binary, the other half doesn't", 1, 2);

    // 右对齐并指定宽度，输出 "     1"（5 个空格 + "1"）
    println!("{number:>width$}", number=1, width=6);

    // 用 0 填充，输出 "000001"
    println!("{number:0>width$}", number=1, width=6);

    // Rust 会检查参数个数是否匹配
    println!("My name is {0}, {1} {0}", "Bond");
    // FIXME ^ 补上缺失的参数 "James"

    // 定义一个包含 i32 的结构体
    #[allow(dead_code)]
    struct Structure(i32);

    // 自定义类型需要实现格式化，否则不会工作
    println!("This struct `{}` won't print...", Structure(3));
    // FIXME ^ 注释掉这一行

    // 1.58+ 支持直接捕获外部变量
    // 同样输出 "     1"（5 个空格 + "1"）
    let number: f64 = 1.0;
    let width: usize = 6;
    println!("{number:>width$}");
}
```

[`std::fmt`][fmt] 提供了一组控制文本展示的 [`traits`][traits]，两类常用基石：

* `fmt::Debug`：使用 `{:?}`，偏调试输出
* `fmt::Display`：使用 `{}`，偏用户友好输出

这里用 `fmt::Display` 是因为标准库已为这些类型实现它；自定义类型要打印则需要额外步骤。

为类型实现 `fmt::Display` 也会自动实现 [`ToString`]，可用于把类型 [转换][convert] 成 [`String`][string]。

