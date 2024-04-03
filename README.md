# Match

[![GoDoc](https://godoc.org/github.com/tidwall/match?status.svg)](https://godoc.org/github.com/tidwall/match)

Match is a very simple pattern matcher where '*' matches on any 
number characters and '?' matches on any one character.

## Installing

```
go get -u github.com/tidwall/match
```

## Example

```go
match.Match("hello", "*llo") // true
match.Match("jello", "?ello") // true
match.Match("hello", "h*o") // true

match.IsPattern("name") // false
match.IsPattern("H*o") // true
match.IsPattern("??") // true
```


## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

Match source code is available under the MIT [License](/LICENSE).
