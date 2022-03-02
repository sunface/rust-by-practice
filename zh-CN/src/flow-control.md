# 流程控制

### if/else
1. 🌟 
```rust,editable

// 填空
fn main() {
    let n = 5;

    if n < 0 {
        println!("{} is negative", n);
    } __ n > 0 {
        println!("{} is positive", n);
    } __ {
        println!("{} is zero", n);
    }
} 
```

2. 🌟🌟 if/else 可以用作表达式来进行赋值
```rust,editable

// 修复错误
fn main() {
    let n = 5;

    let big_n =
        if n < 10 && n > -10 {
            println!(" 数字太小，先增加 10 倍再说");

            10 * n
        } else {
            println!("数字太大，我们得让它减半");

            n / 2.0 ;
        }

    println!("{} -> {}", n, big_n);
} 
```

### for
3. 🌟 `for in` 可以用于迭代一个迭代器，例如序列 `a..b`.

```rust,editable

fn main() {
    for n in 1..=100 { // 修改此行，让代码工作
        if n == 100 {
            panic!("NEVER LET THIS RUN")
        }
    }
} 
```


4. 🌟🌟 
```rust,editable

// 修复错误，不要新增或删除代码行
fn main() {
    let names = [String::from("liming"),String::from("hanmeimei")];
    for name in names {
        // do something with name...
    }

    println!("{:?}", names);

    let numbers = [1, 2, 3];
    // numbers中的元素实现了 Copy，因此无需转移所有权
    for n in numbers {
        // do something with name...
    }
    
    println!("{:?}", numbers);
} 
```

5. 🌟
```rust,editable
fn main() {
    let a = [4,3,2,1];

    // 通过索引和值的方式迭代数组 `a` 
    for (i,v) in a.__ {
        println!("第{}个元素是{}",i+1,v);
    }
}
```

### while
6. 🌟🌟 当条件为 true 时，`while` 将一直循环

```rust,editable

// 填空，让最后一行的  println! 工作 !
fn main() {
    // 一个计数值
    let mut n = 1;

    // 当条件为真时，不停的循环
    while n __ 10 {
        if n % 15 == 0 {
            println!("fizzbuzz");
        } else if n % 3 == 0 {
            println!("fizz");
        } else if n % 5 == 0 {
            println!("buzz");
        } else {
            println!("{}", n);
        }


        __;
    }

    println!("n 的值是 {}, 循环结束",n);
}
```

### continue and break
7. 🌟 使用 `break` 可以跳出循环
```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n == 66 {
           __
       }
       n += 1;
    }

    assert_eq!(n, 66);
}
```

8. 🌟🌟 `continue` 会结束当次循环并立即开始下一次循环
```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut n = 0;
    for i in 0..=100 {
       if n != 66 {
           n+=1;
           __;
       }
       
       __
    }

    assert_eq!(n, 66);
}
```

### loop 

9. 🌟🌟 loop 一般都需要配合 `break` 或 `continue` 一起使用。

```rust,editable

// 填空，不要修改其它代码
fn main() {
    let mut count = 0u32;

    println!("Let's count until infinity!");

    // 无限循环
    loop {
        count += 1;

        if count == 3 {
            println!("three");

            // 跳过当此循环的剩余代码
            __;
        }

        println!("{}", count);

        if count == 5 {
            println!("OK, that's enough");

            __;
        }
    }

    assert_eq!(count, 5);
}
```

10. 🌟🌟 loop 是一个表达式，因此我们可以配合 `break` 来返回一个值
```rust,editable

// 填空
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            __;
        }
    };

    assert_eq!(result, 20);
}
```

11. 🌟🌟🌟 当有多层循环时，你可以使用 `continue` 或 `break` 来控制外层的循环。要实现这一点，外部的循环必须拥有一个标签 `'label`, 然后在 `break` 或 `continue` 时指定该标签

```rust,editable

// 填空
fn main() {
    let mut count = 0;
    'outer: loop {
        'inner1: loop {
            if count >= 20 {
                // 这只会跳出 inner1 循环
                break 'inner1; // 这里使用 `break` 也是一样的
            }
            count += 2;
        }

        count += 5;

        'inner2: loop {
            if count >= 30 {
                break 'outer;
            }

            continue 'outer;
        }
    }

    assert!(count == __)
}
```

> 你可以在[这里](https://github.com/sunface/rust-by-practice)找到答案(在 solutions 路径下) 