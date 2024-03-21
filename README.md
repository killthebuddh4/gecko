## Gecko

Gecko is an experimental functional programming language and interpreter for radically distributed applications. Every local node in a Gecko program can be transparently replaced with a remote node. Remote nodes are accessed via RPC over the [XMTP](https://xmtp.org) protocol.

XMTP is an end-to-end-encrypted messaging protocol which uses Ethereum addresses as identities. By using these decentralized, open protocols for _message passing_ (in the [Smalltalk sense](https://en.wikipedia.org/wiki/Message_passing)) we may be able to bootstrap a computing environment with unheard levels of collaboration and composability.

Gecko is heavily influenced by [Scheme](https://www.scheme.org/), [Smalltalk](https://en.wikipedia.org/wiki/Smalltalk) and [Ruby](https://www.ruby-lang.org/en/).

_This project and documentation are both under heavy development. To see how things are going, please see the [roadmap](#roadmap)._

#### 👋 Say hi!

Thanks for checking out the project! If you think it's interesting I'd love to hear from you (I'd love to hear criticism too, actually). The best ways to reach out are probably [Discord](https://discord.gg/wG9rEmw8), [Twitter](https://twitter.com/killthebuddha_), or [an issue](https://github.com/killthebuddh4/gecko/issues/new).

## Contents

- [Gecko](#gecko)
    - [👋 Say hi!](#-say-hi)
- [Contents](#contents)
- [Why?](#why)
- [The language](#the-language)
- [Syntax and semantics, overview](#syntax-and-semantics-overview)
- [Syntax and semantics, reference](#syntax-and-semantics-reference)
    - [Blocks](#blocks)
    - [Variables](#variables)
    - [Values](#values)
    - [Lambdas, parameters, and arguments](#lambdas-parameters-and-arguments)
    - [Predicates, operators, and literals](#predicates-operators-and-literals)
    - [Branching](#branching)
    - [Arrays](#arrays)
    - [Maps](#maps)
    - [Strings](#strings)
    - [Input and Output](#input-and-output)
    - [Signals and exceptions](#signals-and-exceptions)
- [Run a Gecko script](#run-a-gecko-script)
- [Testing](#testing)
- [Roadmap](#roadmap)
    - [Phase 1, minimal language core](#phase-1-minimal-language-core)
    - [Phase 2, the rest of the core](#phase-2-the-rest-of-the-core)
    - [Phase 3, XMTP RPC](#phase-3-xmtp-rpc)
    - [Phase 4, nice to haves](#phase-4-nice-to-haves)


## Why?

A few reasons (one intuitive, one practical, and one speculative):

1. The intuitive reason is that decentralized, end-to-end-encrypted messaging for [message passing](https://en.wikipedia.org/wiki/Message_passing) simply _feels incredibly powerful_.
2. The practical reason is that service discovery via public keys makes for an extremely simple software distribution mechanism, at least for "toy" software. No account creation, no DNS, no hosting. In certain cases you don't even need a web server because XMTP can run in the browser. This could be useful in certain situations e.g. in the classroom.
3. We're heading towards a world where a large fraction (all, eventually) of software is encoded by, or generated on the fly, by an AI model. In that world the API may overtake the library as the primary method for sharing software.

## The language

Gecko is a dynamically-typed, lexically-scoped, expression-oriented, interpreted, functional programming language with a Ruby-like syntax.

__Some design goals__

- Beginner friendly
- Uncluttered syntax (inspired by Ruby)
- Semantic simplicity (inspired by Go)
- Application-oriented (less "general purpose" than, say, Python)

## Syntax and semantics, overview

__Expressions__

In Gecko everything is an expression that returns a value. The kinds of expressions are:

- variable definition
- function definition
- function call
- block
- predicate
- literal

__Values and variables__

Values can be bound to variables. Variables are referenced by name. All values are immutable but variables can be re-bound to new values. Gecko includes the following primitive types:

- string
- number
- boolean
- nil
- array
- dict
- function

Some examples:

```

# string

"I'm a string"

# number

10.0

# boolean

true

# nil

nil

# array

array
  1
  2
  "test"
end

# dict

dict
  "a" => 10
  "b" => array 1 2 3 end
  "c" => "why?"
end

# function

(n: Number d: Number -> Number)
  return n / d 
end
```

__Functions__

Every function has a signature. A signature is a list of named parameters, their schemas (optional), and the function's return schema (optional). All language built-ins are functions, a new function is defined by writing its signature and body. For example:

```
(car: Car to_speed: Mph -> Mph)
  if
    test
      to_speed > 50
    end

    # crash
    else
      # set the car's speed to the new speed
    end
  end

  return to_speed
end
```

Functions are anonymous and must be bound to a variable if we want to call it. For example:

```
def
  name
    "accelerate"
  end

  value
    (car: Car to_speed: Mph -> Mph)
      if to_speed > 50
        # crash
      else
        # set the car's speed to the new speed
      end

      return to_speed
    end
  end
end
```

After a function is defined it can be called by naming it and assigning values to its parameters. For example:

```
accelerate
  car
    my_moms_car
  end

  to_speed
    100
  end
end
```

If we want to reference a function without calling it we can prefix its name with an `&`. For example:

```
def
  name
    "double"
  end

  value
    (n: Number -> Number)
      return 2 * n
    end
  end
end

for
  array
    my_values
  end

  do
    &double
  end
end
```

__A note on schemas__

A schema in Gecko is a function used to validate values, especially values which are the arguments to or return from a fucntion. If you're a TypeScript programmer and have used [zod](https://github.com/colinhacks/zod) they should look familiar.

When a function is called, each named argument is passed to its corresponding type. Each type is a parser function that validates the argument or throws. If a parameter's type is not specified, then `Identity` is used, which always succeeds. When a function returns a value to the caller, the function's return type is used to first parse the value.

__Comments__

Comments begin with a `#` and continue until the end of the line. Whitespace is ignored (except as token separators).

__Evaluation__

All expressions, with a few exceptions, in a Gecko program tree are evaluated according to a basic depth-first tree walk. The children of the conditional expressions `and`, `or`, `if`, `when`, and `switch` may be (in certain obvious cases) skipped. Additionally, Gecko provides a single mechanism for parallel execution of children expressions via the `parallel` keyword.

## Syntax and semantics, reference

__Please note that this section has gone stale. For current examples of syntax and usage please see the [examples](./examples/) or [tests](./tests/).__

#### Blocks

A __block__ is a sequence of expressions delimited by a _keyword_ and `end`. A
__keyword__ determines its block's behavior or semantics. Most of the language's
keywords will be described throughout the rest of this section but you can also
find a comprehensive, runnable example in
[examples.core.fly](examples.core.fly).

The simplest block is the `do` block:

_`do expression* end`_

The expressions are evaluated in order and the value of the last expression is
returned.

```text
do
  puts "hey" end

  2

  do
    3 + 4
  end
end
```

#### Variables

A Gecko __variable__ is an expression that resolves to a _value_ by referencing
it. A variable is defined using a `def` block and re-defined using a `let`
block. After a variable is defined it can be referenced in any expression.

_`def identifier expression end`_

Defines a variable with the given identifier. The variable resolves to the value
of the expression. Variables are _lexically scoped_. If the variable is already
defined in the local scope, it is an error. If the variable is defined in an
outer scope, it will be _shadowed_ in the local scope.

```text
def surname "smith" end
```

_`let identifier expression end`_

Re-defines an existing variable with the given identifier. The variable resolves
to the value of the expression. If the variable does not already exist, it is
an error.

```
def val "hi" end
let val "goodbye" end
```

- [ ] Namespace declaration and resolution.

#### Values

Every value is a _string_, _number_, _array_, _map_, _lambda_, or _nil_.

A __string__ is created by enclosing characters in quotes.

```text
"I am string"
```

A __number__ is created by writing it out in decimal notation. All numbers are
represented as floats internally.

```text
1
0.1
10.0
```

There is no _boolean_ type in Gecko. All "boolean" operators take _number_
operands and treat `0` as false-y and any other number as truth-y. All other
values cause errors when used as a boolean.

An __array__ is created using the `array` block and is a number-indexed list of
values. See the [arrays](#arrays) section for more details on arrays.

A __map__ is created using the `map` block and is a string-keyed dictionary of
values. See the [maps](#maps) section for more details on maps.

A __lambda__ is created using the `fn` block and can be thought of as a
parameterized _do_ block or "anonymous function". See the [lambdas](#lambdas)
section for more details on lambdas.

#### Lambdas, parameters, and arguments

A __lambda__ is a "parameterized block" that is not evaluated until each time it
is called. A lambda can have zero or more _parameters_. A __parameter__ is a
name that is defined each time the lambda is called. Parameters are declared
between `|` characters. If the lambda takes zero parameters, the `|` characters
must be omitted. The  __arguments__ to the lambda are the values of the
expressions in the calling block (using the `.` keyword) bound to the lambda's
parameters.

_`fn (|identifier+|)? expression end`_

When the lambda expression is evaluated, it creates a lambda. The key difference
between a lambda expression and other expressions is that its subexpressions are
evaluated only when the lambda is called. The lambda can take zero or more
parameters. If the lambda takes zero parameters, the `|` characters must be
omitted.

_`. expression* end`_

Calls the lambda expression. Each subexpression is evaluated and bound to the
lambda's parameters. The lambda is then evaluated, returning the value of its
last subexpression.

```text
def add
  # parameters are a and b
  fn |a b|
    a + b
  end
end

.add
  # arguments are 8 and 3, bound to a and b
  2 * 4
  3
end

map
  array 1 2 3 end

  fn |n i|
    n + i
  end
end
```

#### Predicates, operators, and literals

A __predicate__ is an expression involving an _operator_ and _operands_. See the
[operators](#operators) section for more details on each operator. An
__operand__ is either a _predicate_ or a _literal_. A __literal__ is an
expression without subexpressions (string, number, boolean, variable). A
predicate evaluates to a _number_ (because an operator evaluates to a number).

_Because predicates cannot include blocks they cannot include function calls.
This is somewhat cumbersome to us human programmers, forcing us to write many
instances of trivial indirection, but I think we'll see strong benefits for code
generation and program synthesis because it will make parse trees simpler. Maybe
not, we'll see._

```text
# Not predicates.

fn
  std.write "hi" end
end

def val "hi" end

# Predicates.

val

val == "goodbye"

10 > 0 # => 1

100 / 20 # => 5

!val
```
#### Branching

The key difference between branching expressions and other expressions is that
their subexpression are evaluated conditionally. The specific behavior of which
subexpressions are evaluated depends on the keyword.

_Note that branching expressions are not predicates, they may return any value._

_`if number expression expression end`_

If the number is truth-y, the first expression is evaluated. Otherwise, the
second expression is evaluated. The value of the last evaluated expression is
returned.


_`and (number expression)+ end`_

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated. If any of the subexpressions evaluate to a false-y value,
`nil` is returned. Otherwise, the value of the last subexpression is returned.

_`or (number expression)+ end`_

For each pair of subexpressions, if the first evaluates to a truth-y value, the
second is evaluated and returned. If none of the subexpressions evaluate to a
truth-y value, `nil` is returned.

_`while number expression+ end`_

While the first expression evaluates to a truth-y value, the rest of the expressions
are evaluated. The value of the last subexpression is returned.

#### Arrays

_`array expression* end`_

Creates an array whose values are the values of the subexpressions. The array is 
returned.

_`array.read array number end`_

The value of the array at the index of the number is returned.

_`array.write array number expression end`_

Clones the array and sets the value at the index of the number to the value of
the expression. The cloned array is returned.

_`array.for array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
The value of the last evaluated lambda is returned.

_`array.map array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the result of each lambda call is returned.

_`array.filter array lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's first parameter and the index bound to the lambda's second parameter.
An array whose values are the values for which the lambda call returned a
truth-y value is returned.

_`array.reduce array expression lambda end`_

For each value in the array, the lambda is called with the value bound to the
lambda's second parameter and the index bound to the lambda's third parameter.
When the lambda is called for the first value in the array, the first parameter
is bound to the value of expression. For each subsequent value in the array, the
first parameter is bound to the value returned by the previous lambda call. The
value of the last evaluated lambda is returned.

_`array.push array expression end`_

Clones the array and appends the value of the expression to the cloned array. 
The cloned array is returned.

_`array.pop array end`_

Clones the array and removes the last value from the cloned array. The cloned
array is returned.

_`array.unshift array expression end`_

Clones the array and prepends the value of the expression to the cloned array.
The cloned array is returned.

_`array.shift array end`_

Clones the array and removes the first value from the cloned array. The cloned
array is returned.

_`array.reverse array end`_

Clones the array and reverses the order of the values in the cloned array. The
cloned array is returned.

_`array.sort array lambda end`_

Clones the array and sorts the values in the cloned array according to the value
returned by the lambda. The lambda takes two parameters, the values of which are
the values in the array. The lambda returns a negative number if the first value
should be sorted before the second, a positive number if the first value should
be sorted after the second, and `0` if the values are equal. The cloned (sorted)
array is returned.

_`array.segment array number number end`_

Clones the array and returns a new array whose values are the values of the
cloned array between the first index and the second index (exclusive). The
cloned array is returned.

_`array.splice array number array end`_

Clones the first array and divides it in half at the index of the number. It
appends the values of the second array to the first half, and then appends the
second half to the result. The result is returned.

#### Maps

_`map (string expression)* end`_

Creates a map whose keys are the strings and whose values are the values of
the expressions. The map is returned.

_`map.read map string end`_

The value of the map at the key of the string is returned.

_`map.write map string expression end`_

Clones the map and sets the value at the key of the string to the value of
the expression. The cloned map is returned.

_`map.delete map array end`_

The array is an array of strings. Clones the map and deletes the keys of the
strings from the cloned map. The cloned map is returned.

_`map.extract map array end`_

The array is an array of strings. Returns a map whose keys are the keys of
the strings and whose values are the values of the keys of the strings in the
map. The new map is returned.

_`map.merge map map end`_

Clones the first map and then for each kv pair in the second map, sets the
value of the cloned map at the key of the kv pair to the value of the kv
pair. Returns the cloned map.

_`map.keys map end`_

An array whose values are the keys of the map is returned.

_`map.values map end`_

An array whose values are the values of the map is returned.

#### Strings

_`split string end`_

Returns an array whose values are the characters in the string.

_`concat string+ end`_

Returns a string whose value is the concatenation of the values of the strings.

_`substring string number number end`_

Returns a string whose value is the substring of the string between the first
index and the second index (exclusive).

#### Input and Output

#### Signals and exceptions

__TODO__

## Run a Gecko script

_Requires `go` 1.21 or higher. Learn how to install `go` [here](https://go.dev/doc/install)._

```bash
go run . <path to Gecko source>
```

Try running the examples:

```bash
for file in examples/*.fly; do
  go run . $file
done
```

## Testing

 You can run the tests with:

 ```bash
 ./test.sh
 ```

 The goal is to have tests commensurate with the maturity of the project and its
 components. The near term goal is to have something like 100% coverage for the
 core language keywords. Basically, this means "all of the keywords and
 operators". We'll do this incrementally, in phases.

## Roadmap

#### Phase 1, minimal language core

- [x] design and implement the architecture
  - [x] lex
  - [x] parse
  - [x] eval
- [ ] keywords
  - [x] for arrays
  - [x] for strings
  - [x] for dicts
  - [x] for control flow
  - [x] for functions
  - [x] for variables
  - [x] for predicates
  - [x] for io
    - [x] fs
    - [x] env
- [x] errors

#### Phase 2, the rest of the core

- [ ] namespaces
  - [ ] export
  - [ ] import
- [x] syntax highlighting

#### Phase 3, XMTP RPC

- [ ] golang XMTP client port
- [ ] `brpc` implementation using XMTP client

#### Phase 4, nice to haves

- [ ] language server protocol implementation
- [ ] repl
- [ ] tail call optimization
- [ ] basic static analysis
  - [ ] auto format
  - [ ] linter
  - [ ] schema -> typechecking?
