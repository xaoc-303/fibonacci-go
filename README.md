# fibonacci-go

[![Build Status](https://travis-ci.org/xaoc-303/fibonacci-go.svg?branch=master)](https://travis-ci.org/xaoc-303/fibonacci-go)

## recursive if-else

| v | # | 30 | 35 | 40 | 45 |
| --- | --- | --- | --- | --- | --- |
| 1.13.1 | [Go](./main.go) (compiled) | 0.00591578 | 0.05948975 | 0.53472273 | 5.80004869 |
| 1.13.1 | [Go](./main.go) | 0.00559842 | 0.05122938 | 0.53362809 | 5.75826140 |
| | [Total](https://github.com/xaoc-303/fibonacci) | | | | |

## optimization

| v | # | 30 | 35 | 40 | 45 |
| --- | --- | --- | --- | --- | --- |
| 1.13.1 | [Go](./main.go) (compiled) | 0.00000024 | 0.00000019 | 0.00000023 | 0.00000021 |
| 1.13.1 | [Go](./main.go) | 0.00000020 | 0.00000022 | 0.00000021 | 0.00000022 |
| | [Total](https://github.com/xaoc-303/fibonacci) | | | | |

#### setting console command 'time'
```
export TIMEFMT=$'\nreal\t%*E\nuser\t%*U\nsys\t%*S'
```

#### run

```
time go run main.go f1 30
time go run main.go f2 30

go build main.go
time ./main f1 30
time ./main f2 30
```

```
go test
```
