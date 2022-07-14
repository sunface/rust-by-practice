<h1 align="center">Rust语言实战</h1>

<div align="center">
    <img height="150" src="https://github.com/sunface/rust-by-practice/blob/master/en/assets/logo.png?raw=true">
</div>
    
<p align="center">通过有挑战性的示例、练习题、实践项目来提升 Rust 水平，建立从入门学习到上手实战的直通桥梁</p>
    
<div align="center">

[![Stars Count](https://img.shields.io/github/stars/sunface/rust-by-practice?style=flat)](https://github.com/sunface/rust-by-practice/stargazers) [![Forks Count](https://img.shields.io/github/forks/sunface/rust-by-practice.svg?style=flat)](https://github.com/naaive/orange/network/members)
[![LICENSE](https://img.shields.io/badge/license-mit-green?style=flat)](https://github.com/sunface/rust-by-practice/blob/master/LICENSE)
</div>

*Rust语言实战* 的目标是通过大量的实战练习帮助大家更好的学习和上手使用 Rust 语言。书中的练习题非常易于使用：你所需的就是在线完成练习，并让它通过编译。


## 在线阅读

- [https://zh.practice.rs](https://zh.practice.rs)

## 本地运行

我们使用 [mdbook](https://rust-lang.github.io/mdBook/) 构建在线练习题，你也可以下载到本地运行：
```shell
$ cargo install mdbook
$ cd rust-by-practice && mdbook serve 
```
在本地win 10或者linux服务器上运行时，应当使用 -n 参数指定mdbook服务所监听的IP地址（-p 参数指定服务监听的端口，不指定则为默认的3000），以win 10本地运行为例：
```shell
$ mdbook serve -p 8888 -n 127.0.0.1
```
## 特性

部分示例和习题借鉴了 [Rust By Example](https://github.com/rust-lang/rust-by-example), 书中的示例真的非常棒！

尽管它们非常优秀，我们这本书也有自己的秘密武器 :)

- 每个章节分为三个可选部分：示例、练习和实践项目

- 除了示例外，我们还有大量的高质量练习题，你可以在线阅读、修改和编译它们

- 覆盖了 Rust 语言的几乎所有方面：基础语言特性、高级语言特性、async/await 异步编程、多线程、并发原语、性能优化、工具链使用、标准库、数据结构和算法等

- 每一道练习题都提供了解答

- 整体难度相对更高，更加贴近于实战难度: 简单 🌟 , 中等 🌟🌟 , 困难 🌟🌟🌟  , 地狱 🌟🌟🌟🌟

**总之，我们想做的就是解决入门学习后，不知道该如何运用的问题，毕竟对于 Rust 来说，从学习到实战，中间还隔着数个 Go语言 的难度**

## 关于我们

*Rust语言实战* 由 Rust 编程学院倾情打造。

同时我们还提供了一本目前最好也是最用心的开源 Rust 书籍 - [Rust语言圣经](https://github.com/sunface/rust-course)， 适合从入门到精通所有阶段的学习，欢迎大家阅读使用。

对我们来说，来自读者大大的肯定比什么都重要，因此一个 [Github star](https://github.com/sunface/rust-by-practice) 要比一杯咖啡更让我们开心，而且现在它在跳楼打折，无需 998 ， 仅需 0 元钱 :)







