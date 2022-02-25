# 字符、布尔、单元类型

### 字符
🌟
```rust, editable

use std::mem::size_of_val;
fn main() {
    let c1 = 'a';
    assert_eq!(size_of_val(&c1),1); 

    let c2 = '中';
    assert_eq!(size_of_val(&c2),3); 
} 
```

🌟
```rust, editable

fn main() {
    let c1 = "中";
    print_char(c1);
} 

fn print_char(c : char) {
    println!("{}", c);
}
```

### Bool
🌟
```rust, editable

// 让  println! 工作
fn main() {
    let _f: bool = false;

    let t = true;
    if !t {
        println!("hello, world");
    }
} 
```

🌟
```rust, editable

fn main() {
    let f = true;
    let t = true && false;
    assert_eq!(t, f);
}
```


### Unit type
🌟🌟
```rust,editable

// 让代码工作，但不要修改 `implicitly_ret_unit` !
fn main() {
    let _v: () = ();

    let v = (2, 3);
    assert_eq!(v, implicitly_ret_unit())
}

fn implicitly_ret_unit() {
    println!("I will returen a ()")
}

// 不要使用下面的函数，它只用于演示！
fn explicitly_ret_unit() -> () {
    println!("I will returen a ()")
}
```

🌟🌟 单元类型占用的内存大小是多少？
```rust,editable

// 让代码工作：修改 `assert!` 中的 `4` 
use std::mem::size_of_val;
fn main() {
    let unit: () = ();
    assert!(size_of_val(&unit) == 4);
}
```