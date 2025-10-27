# Heap

push faster 4%
pop faster 39%

```txt
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCustomHeap_Init-12          918           1214760 ns/op          802854 B/op       2 allocs/op
BenchmarkStdHeap_Init-12             860           1388620 ns/op          802840 B/op       2 allocs/op
BenchmarkCustomHeap_Push-12     49751402                25.86 ns/op           42 B/op       0 allocs/op
BenchmarkStdHeap_Push-12        50084911                26.98 ns/op           50 B/op       0 allocs/op
BenchmarkCustomHeap_Pop-12      10286382               121.1 ns/op             0 B/op       0 allocs/op
BenchmarkStdHeap_Pop-12          5971455               199.6 ns/op            15 B/op       1 allocs/op
```