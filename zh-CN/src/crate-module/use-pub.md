# use and pub
1. 🌟 使用 `use` 可以将两个同名类型引入到当前作用域中，但是别忘了 `as` 关键字.

```rust,editable
use std::fmt::Result;
use std::io::Result;

fn main() {}
```

2. 🌟🌟 如果我们在使用来自同一个包或模块中的多个不同项，那么可以通过简单的方式将它们一次性引入进来

```rust,editable

// 使用两种方式填空
// 不要添加新的代码行
use std::collections::__;

fn main() {
    let _c1:HashMap<&str, i32> = HashMap::new();
    let mut c2 = BTreeMap::new();
    c2.insert(1, "a");
    let _c3: HashSet<i32> = HashSet::new();
}
```

### 使用 `pub use` 进行再导出

3. 🌟🌟🌟 在之前创建的`hello-package` 的库包中, 添加一些代码让下面的代码能够正常工作
```rust,editable
fn main() {
    assert_eq!(hello_package::hosting::seat_at_table(), "sit down please");
     assert_eq!(hello_package::eat_at_restaurant(),"yummy yummy!");
}
```


### pub(in Crate) 
有时我们希望某一个项只对特定的包可见，那么就可以使用 `pub(in Crate)` 语法.

#### 示例
```rust,editable
pub mod a {
    pub const I: i32 = 3;

    fn semisecret(x: i32) -> i32 {
        use self::b::c::J;
        x + J
    }

    pub fn bar(z: i32) -> i32 {
        semisecret(I) * z
    }
    pub fn foo(y: i32) -> i32 {
        semisecret(I) + y
    }

    mod b {
        pub(in crate::a) mod c {
            pub(in crate::a) const J: i32 = 4;
        }
    }
}
```

### 完整代码
至此，包与模块章节已经结束，关于 `hello-package` 的完整代码可以在[这里](https://github.com/sunface/rust-by-practice/tree/master/practices/hello-package) 找到.


> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/crate-module/use-pub.md)找到答案(在 solutions 路径下) 