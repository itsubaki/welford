# welford

 * Welford algorithm for Go


```go
w := welford.New()
w.Add(1)
w.Add(2, 3, 4, 5)
w.Add([]float64{6, 7, 8, 9, 10}...)

fmt.Println(w.N())
fmt.Println(w.Mean())
fmt.Println(w.Var())
fmt.Println(w.Stddev())

// Output:
// 10
// 5.5
// 9.166666666666666
// 3.0276503540974917
```