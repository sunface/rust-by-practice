# 方法和关联函数

## 示例
```rust,editable
struct Point {
    x: f64,
    y: f64,
}

// `Point` 的关联函数都放在下面的 `impl` 语句块中
impl Point {
    // 关联函数的使用方法跟构造器非常类似
    fn origin() -> Point {
        Point { x: 0.0, y: 0.0 }
    }

    // 另外一个关联函数，有两个参数
    fn new(x: f64, y: f64) -> Point {
        Point { x: x, y: y }
    }
}

struct Rectangle {
    p1: Point,
    p2: Point,
}

impl Rectangle {
    // 这是一个方法
    // `&self` 是 `self: &Self` 的语法糖
    // `Self` 是当前调用对象的类型，对于本例来说 `Self` = `Rectangle`
    fn area(&self) -> f64 {
        // 使用点操作符可以访问 `self` 中的结构体字段
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

  
        // `abs` 是一个 `f64` 类型的方法，会返回调用者的绝对值
        ((x1 - x2) * (y1 - y2)).abs()
    }

    fn perimeter(&self) -> f64 {
        let Point { x: x1, y: y1 } = self.p1;
        let Point { x: x2, y: y2 } = self.p2;

        2.0 * ((x1 - x2).abs() + (y1 - y2).abs())
    }

    // 该方法要求调用者是可变的，`&mut self` 是 `self: &mut Self` 的语法糖
    fn translate(&mut self, x: f64, y: f64) {
        self.p1.x += x;
        self.p2.x += x;

        self.p1.y += y;
        self.p2.y += y;
    }
}

// `Pair` 持有两个分配在堆上的整数
struct Pair(Box<i32>, Box<i32>);

impl Pair {
    // 该方法会拿走调用者的所有权
    // `self` 是 `self: Self` 的语法糖
    fn destroy(self) {
        let Pair(first, second) = self;

        println!("Destroying Pair({}, {})", first, second);

        // `first` 和 `second` 在这里超出作用域并被释放
    }
}

fn main() {
    let rectangle = Rectangle {
        // 关联函数的调用不是通过点操作符，而是使用 `::`
        p1: Point::origin(),
        p2: Point::new(3.0, 4.0),
    };

    // 方法才是通过点操作符调用
    // 注意，这里的方法需要的是 `&self` 但是我们并没有使用 `(&rectangle).perimeter()` 来调用，原因在于：
    // 编译器会帮我们自动取引用
    //  `rectangle.perimeter()` === `Rectangle::perimeter(&rectangle)`
    println!("Rectangle perimeter: {}", rectangle.perimeter());
    println!("Rectangle area: {}", rectangle.area());

    let mut square = Rectangle {
        p1: Point::origin(),
        p2: Point::new(1.0, 1.0),
    };


    // 错误！`rectangle` 是不可变的，但是这个方法要求一个可变的对象
    //rectangle.translate(1.0, 0.0);
    // TODO ^ 试着反注释此行，看看会发生什么

    // 可以！可变对象可以调用可变的方法
    square.translate(1.0, 1.0);

    let pair = Pair(Box::new(1), Box::new(2));

    pair.destroy();

    // Error! 上一个 `destroy` 调用拿走了 `pair` 的所有权
    //pair.destroy();
    // TODO ^ 试着反注释此行
}
```

## Exercises

### Method
1. 🌟🌟 方法跟函数类似：都是使用 `fn` 声明，有参数和返回值。但是与函数不同的是，方法定义在结构体的上下文中(枚举、特征对象也可以定义方法)，而且方法的第一个参数一定是 `self` 或其变体 `&self` 、`&mut self`，`self` 代表了当前调用的结构体实例。

```rust,editable
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    // 完成 area 方法，返回矩形 Rectangle 的面积
    fn area
}

fn main() {
    let rect1 = Rectangle { width: 30, height: 50 };

    assert_eq!(rect1.area(), 1500);
}
```

2. 🌟🌟 `self` 会拿走当前结构体实例(调用对象)的所有权，而 `&self` 却只会借用一个不可变引用，`&mut self` 会借用一个可变引用

```rust,editable
// 只填空，不要删除任何代码行!
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    pub fn show_state(__)  {
        println!("the current state is {}", __.color);
    }
}
fn main() {
    let light = TrafficLight{
        color: "red".to_owned(),
    };
    // 不要拿走 `light` 的所有权
    light.show_state();
    // 否则下面代码会报错
    println!("{:?}", light);
}
```
3. 🌟🌟  `&self` 实际上是 `self: &Self` 的缩写或者说语法糖
```rust,editable
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // 使用 `Self` 填空
    pub fn show_state(__)  {
        println!("the current state is {}", self.color);
    }

    // 填空，不要使用 `Self` 或其变体
    pub fn change_state(__) {
        self.color = "green".to_string()
    }
}
fn main() {}
```


### Associated function

4. 🌟🌟  所有定义在 `impl` 语句块中的函数被称为关联函数，因为它们跟当前类型关联在一起。我们可以定义没有 `self` 作为其第一个参数（因此不是方法）的关联函数，它们不需要该类型的实例来处理。这种关联函数往往可以用于作为构造函数：初始化一个实例对象。

```rust,editable
#[derive(Debug)]
struct TrafficLight {
    color: String,
}

impl TrafficLight {
    // 1. 实现下面的关联函数 `new`,
    // 2. 该函数返回一个 TrafficLight 实例，包含 `color` "red"
    // 3. 该函数必须使用 `Self` 作为类型，不能在签名或者函数体中使用 `TrafficLight`
    pub fn new() 

    pub fn get_state(&self) -> &str {
        &self.color
    }
}

fn main() {
    let light = TrafficLight::new();
    assert_eq!(light.get_state(), "red");
}
```

### 多个 `impl` 语句块
5. 🌟 每一个结构体允许拥有多个 `impl` 语句块
```rust,editable

struct Rectangle {
    width: u32,
    height: u32,
}

// 使用多个 `impl` 语句块重写下面的代码
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}


fn main() {}
```

### Enums
6. 🌟🌟🌟 我们还可以为枚举类型定义方法

```rust,editable

#[derive(Debug)]
enum TrafficLightColor {
    Red,
    Yellow,
    Green,
}

// 为 TrafficLightColor 实现所需的方法
impl TrafficLightColor {
    
}

fn main() {
    let c = TrafficLightColor::Yellow;

    assert_eq!(c.color(), "yellow");

    println!("{:?}",c);
}
```

## Practice

@todo


> 你可以在[这里](https://github.com/sunface/rust-by-practice/blob/master/solutions/method.md)找到答案(在 solutions 路径下) 
