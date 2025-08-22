# GoF Design Patterns in Go

This repository contains implementations of **Gang of Four (GoF) Design Patterns** in **Go**. 

## Patterns Covered
The patterns are grouped into three categories: **Creational**, **Structural**, **Behavioral**

### Creational
- [X] Singleton
- [X] Factory
- [X] Abstract Factory
- [X] Builder

### Structural
- [X] Adapter
- [X] Bridge
- [X] Decorator
- [X] Facade
- [X] Proxy

### Behavioral
- [X] Strategy
- [X] Observer
- [X] Command
- [X] Mediator
- [X] Memento
- [X] State
- [X] Template
- [X] Visitor
- [X] Chain of Responsibility

## How to Run
Each pattern has its own `main.go` entrypoint located under `cmd/{category}/{pattern}/`.

Run a specific pattern with:

```bash
# Behavioral → Strategy
go run ./cmd/behavioral/strategy/main.go

# Creational → Singleton
go run ./cmd/creational/singleton/main.go

# Structural → Proxy
go run ./cmd/structural/proxy/main.go
```

## Requirements
- Go 1.24.5 or later
- No other external dependencies



## Resources
- **Design Patterns: Elements of Reusable Object-Oriented Software** by Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides  
  [Google Books](https://books.google.com.tr/books/about/Design_Patterns.html?id=6oHuKQe3TjQC&redir_esc=y)

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler, Kent Beck, John Brant, William Opdyke, Don Roberts
  [Google Books](https://books.google.com.tr/books/about/Refactoring.html?id=HmrDHwgkbPsC&redir_esc=y)

- **Refactoring Guru**
  [Site](https://refactoring.guru/)