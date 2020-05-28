# utils
Collect some go functions with good performance and convenient.

like Stings and byte convertion. generator.

There is a better suggestion to match the [validator](https://github.com/go-playground/validator) package

## Install

```go
go get github.com/slclub/utils
```

## function list

- strings
  - [strings bytes convenient](https://github.com/slclub/utils/blob/master/bytesconv/bytesconv.go) [writer@gin](https://github.com/gin-gonic/gin/tree/master/internal/bytesconv)
  
```go 
import("github.com/slclub/utils/bytesconv") 
```  

- generator number
  - [grange](https://github.com/slclub/utils/blob/master/grange.go)

- reflect functions
  - [_FUNC_NAME](https://github.com/slclub/utils/blob/master/reflect.go)
  - [_IS_FUNC](https://github.com/slclub/utils/blob/master/reflect.go)
  - [_IS_STRUCT](https://github.com/slclub/utils/blob/master/reflect.go)
