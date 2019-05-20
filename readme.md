# Olivier's Reference Sheet
If I google something more than once, it means that I did not take the proper time to learn everything about this
thing. I am creating this sheet so that I don't fall in the bad habbit of googling everything. I will format this
by categories/Technologies

## Content
1. [R-language](## R-language)
  1. cut {base}
2. [GoLang](## GoLang)
  1. Import
3. [Python](##Python)
  1. How to Slice Lists/Arrays and Tuples in Python

## R-language
### cut {base}
cut divides the range of x into intervals and codes the values in x according to which interval they fall. The leftmost interval corresponds to level one, the next leftmost to level two and so on. logical, indicating if an ‘x[i]’ equal to
    the lowest (or highest, for <code>right = FALSE</code>) ‘breaks’
    value should be included.
    If a <code>labels</code> parameter is specified, its values are used to name
      the factor levels.  If none is specified, the factor level labels are
      constructed as <code>"(b1, b2]"</code>, <code>"(b2, b3]"</code> etc. for
      <code>right = TRUE</code> and as <code>"[b1, b2)"</code>, … if <code>right =
        FALSE</code>.
      In this case, <code>dig.lab</code> indicates the minimum number of digits
      should be used in formatting the numbers <code>b1</code>, <code>b2</code>, ….
      A larger value (up to 12) will be used if needed to distinguish
      between any pair of endpoints: if this fails labels such as
      <code>"Range3"</code> will be used.  Formatting is done by
      <code><a rd-options="" href="/link/formatC?package=base&amp;version=3.6.0" data-mini-rdoc="base::formatC">formatC<a><code>
```
cut(x, …)
# S3 method for default
cut(x, breaks, labels = NULL,
    include.lowest = FALSE, right = TRUE, dig.lab = 3,
    ordered_result = FALSE, …)
```
Example how to cut date into multiple quarter
```
OrderCons$Period<-cut(OrderCons$ORDER_CREATION_DATE ,
                    breaks = ymd(c( "2018-04-28","2018-07-29","2018-10-28",
                                  "2019-01-27","2019-04-28")),
                    labels =c( "Q4FY18","Q1FY19","Q2FY19","Q3FY19"))
```
#### Reference
1. [rdocumentation](https://www.rdocumentation.org/packages/base/versions/3.6.0/topics/cut)
2. [stackoverflow](https://stackoverflow.com/questions/45201474/customize-quarterly-dates-on-r)

## GoLang

### Import

Anatomy of and Import declarations
```
ImportDeclaration = "import" ImportSpec
ImportSpec        = [ "." | "_" | Identifier ] ImportPath
```
- Identifier is any valid identifier which will be used in qualified identifiers
- ImportPath is string literal (raw or interpreted)
- String literal used in import specification (each import declaration contains one or more import specification) tells what package to import. This string is called import path. According to language spec it depends on the implementation how import path (string) is interpreted but in real life it’s path relative package’s vendor directory or `go env GOPATH`/src (more about [GOPATH](https://golang.org/doc/code.html#GOPATH)).

```Go
import (
    "math"
    m "math"
    . "math"
    _ "math"
)
```

- `import  m "math" `: It’s possible to pass custom package name for import
- `import . "math` another option which allows to access exported identifier without qualified identifier. Useful in testing
- Golang’s compiler will yell if package is imported and not used
-  `import _ "math"`: doesn’t require to use package math in importing file but init function(s) from imported package will be executed anyway (package and it dependencies will be initialized). It’s useful if we’re interested only in bootstrapping work done by imported package but we don’t reference any exported identifiers from it.

#### Circular Import
Go specification explicitly forbids circular imports — when package imports itself indirectly. The most obvious case is when package a imports package b and package b in turn imports package a
```
> go build
can't load package: import cycle not allowed
package github.com/mlowicki/a
    imports github.com/mlowicki/b
    imports github.com/mlowicki/a
```
#### Reference
1. [Medium Article](https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff)
2. [Init Function in GO](https://medium.com/golangspec/init-functions-in-go-eac191b3860a)

## Python
### How to Slice Lists/Arrays and Tuples in Python

So you've got an list, tuple or array and you want to get specific sets of sub-elements from it, without any long, drawn out for loops? **Slicing** can not only be used for lists, tuples or arrays, but custom data structures as well, with the slice object

Example

```python
>>> a = [1, 2, 3, 4, 5, 6, 7, 8]
```

```python
>>> a[1:4]
[2, 3, 4]
```
There is also an optional second clause that we can add that allows us to set how the list's index will increment between the indexes that we've set.
```python
>>> a[1:4:2]
[2, 4]
```

 That last colon tells Python that we'd like to choose our slicing increment. By default, Python sets this increment to **1**, but that extra colon at the end of the numbers allows us to specify what we want it to be.

 ```python
>>> a[::-1]
[8, 7, 6, 5, 4, 3, 2, 1]
 ```
 Lists have a default bit of functionality when slicing. If there is no value before the first colon, it means to start at the beginning index of the list. If there isn't a value after the first colon, it means to go all the way to the end of the list. This saves us time so that we don't have to manually specify len(a) as the ending index.

 Okay. And that -1 I snuck in at the end? It means to increment the index every time by -1, meaning it will traverse the list by going backwards. If you wanted the even indexes going backwards, you could skip every second element and set the iteration to -2. Simple.

There is another form of syntax you can use that may be easier to understand. There are objects in Python called slice objects and the can be used in place of the colon syntax above.

```python
>>> a = [1, 2, 3, 4, 5]
>>> sliceObj = slice(1, 3)
>>> a[sliceObj]
[2, 3]
```

### Reference

1. [Python Central](https://www.pythoncentral.io/how-to-slice-listsarrays-and-tuples-in-python/)
