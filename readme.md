# go-design-pattern

go 设计模式实现，包含 23 种常见的设计模式实现

同时这也是 [极客时间-设计模式之美](https://time.geekbang.org/column/intro/250) 的笔记

## Roadmap

持续更新中，预计一周更新 2 ~ 3 种设计模式，预计到 202010 月底前更新完成

## 创建型模式

### 常用

- [单例模式(Singleton Design Pattern)](./01_singleton) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/singleton.html)
  - [饿汉式](./01_singleton/singleton.go)
  - [懒汉式(延迟加载)](./01_singleton/singleton_lazy.go)
- [工厂模式(Factory Design Pattern)](./02_factory) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/factory.html)
  - [简单工厂](./02_factory/021_simple_factory)
  - [工厂方法](./02_factory/022_factory_method)
  - [抽象工厂(不常用)](./02_factory/023_abstract_factory)
  - [DI 容器](./02_factory/024_di)
- [建造者模式(Builder Design Pattern)](./03_builder) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/builder.html)

### 不常用

- [原型模式(Prototype Design Pattern)](./04_prototype) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/prototype.html)
  
  不常用只是相对而言，在 js 当中十分普遍，在业务代码中其实也时常会用到这种思想

## 结构型模式

### 常用

- [代理模式(Proxy Design Pattern)](./05_proxy) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/proxy.html)
  - 静态代理
  - 动态代理(go generate 实现)
- [桥接模式(Bridge Design Pattern)](./06_bridge) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/bridge.html)
- [装饰器模式(Decorator Design Pattern)](./07_decorator) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/decorator.html)
- [适配器模式(Adapter Design Pattern)](./08_adapter) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/adapter.html)

### 不常用

- [门面模式(Facade Design Pattern)](./09_facade) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/facade.html)
- [组合模式(Composite Design Pattern)](./10_composite) [![](https://img.shields.io/badge/BLOG-%E7%82%B9%E5%87%BB%E6%9F%A5%E7%9C%8B-success?style=flat&cacheSeconds=360000)](https://lailin.xyz/post/composite.html)
- 享元模式

## 行为型模式

### 常用

- 观察者模式
- 模板模式
- 简略模式
- 责任链模式
- 迭代器模式
- 状态模式

### 不常用

- 访问者模式
- 备忘录模式
- 命令模式
- 解释器模式
- 中介模式
