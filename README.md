# Gecko
Gecko is a RegEx engine implemented in Go. The grammar implemented is currently a subset of a full RegEx grammar as specified below.

## Grammar Specification
```
RegExpr         ::= Union | SimpleExpr
Union           ::= SimpleExpr "|" RegExpr
SimpleExpr      ::= Concatenation | BasicExpr
Concatenation   ::= BasicExpr SimpleExpr
BasicExpr       ::= Star | Plus | Question | Element
Star            ::= Element "*"
Plus            ::= Element "+"
Question        ::= Element "?"
Element         ::= Character | Group | Set
Group           ::= "(" RegExpr ")"
Escape          ::= "\" <literal character>
Set             ::= PositiveSet | NegativeSet
PositiveSet     ::= "[" SetItems "]"
NegativeSet     ::= "[^" SetItems "]"
SetItems        ::= SetItem SetItems
SetItem         ::= Range | Character
Range           ::= Character "-" Character
Character       ::= Escape | <literal character>
```

## Project Structure

### Core
The core package contains the code for the engine itself in the form of a lexer, parser and compiler. It also provides a struct `Compiler` which acts as a simple API for the engine.

### Front End
The frontend package contains the code for a web front-end, written using React, which aims to act as something of a teaching aid. 

Currently this presents a visualisation for the underlying finite-state machine implementation for a particular regular expression.

This will be expanded at some point in the near future to to include the ability to "step-through" a particular pattern with a particular input and to also display the underlying abstract syntax tree of a given regular expression.

### Server
Contains all code for serving content over HTTP. This includes a simple server to run the front-end and a back-end server which exposes the underlying core functionality as a REST API.

### Cmd
Collects all the main files which form various executables in a single package. This currently consists of:
 1. A simple command-line runner for the main engine, allowing matching of inputs against a given pattern
 2. A runner for the back-end server, providing a REST API for the core functionality
 3. A runner for the front-end server
 4. An "app" runner which packages the back-end server and the front-end server into a single executable allowing them to be run together as a pseudo-desktop app.
