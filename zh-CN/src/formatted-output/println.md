# println! 与 format!
[`std::fmt`][fmt] 中有一组用于打印的宏，例如：

* `format!`：写入到 [`String`][string]
* `print!`：同 `format!`，但输出到 stdout
* `println!`：`print!` 并追加换行
* `eprint!`：输出到 stderr
* `eprintln!`：`eprint!` 并追加换行

所有这些宏的格式解析方式一致，且编译期检查格式正确性。

## `format!`
1.🌟
```rust,editable

fn main() {
    let s1 = "hello";
    /* 填空 */
    let s = format!(__);
    assert_eq!(s, "hello, world!");
}
```

## `print!`、`println!`
2.🌟
```rust,editable

fn main() {
   /* 填空，让输出变为:
   Hello world, I am 
   Sunface!
   */
   __("hello world, ");
   __("I am");
   __("Sunface!");
}
```

> 参考答案在这里（solutions 路径）：<https://github.com/sunface/rust-by-practice/blob/master/solutions/formatted-output/println.md>，仅在需要时再看哦 :)

