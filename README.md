# Go Script Lang Benchmark

```
git clone github.com/akkuman/go_script_lang_benchmark
cd go_script_lang_benchmark
go test -benchmem -benchtime 20s -bench .
```

## Benchmark

### Windows

- Windows10 21H1
- go version go1.17.5 windows/amd64

```
goos: windows
goarch: amd64
pkg: github.com/akkuman/go_script_lang_benchmark
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
Benchmark_glua_Add_single_ctx-12                    6070           3987641 ns/op         2205896 B/op       8616 allocs/op
Benchmark_starlark_Add_single_ctx-12                1616          14836361 ns/op         6331483 B/op     316324 allocs/op
Benchmark_yaegi_Add_single_ctx-12                   7512           3127616 ns/op              24 B/op          2 allocs/op
Benchmark_go_Add-12                              1000000             21776 ns/op               0 B/op          0 allocs/op
Benchmark_glua_Add_multi_ctx-12                       60         400304143 ns/op       220766320 B/op    862445 allocs/op
Benchmark_starlark_Add_multi_ctx-12                   15        1462908447 ns/op       633148154 B/op  31632452 allocs/op
Benchmark_yaegi_Add_multi_ctx-12                      75         308568796 ns/op           26924 B/op        455 allocs/op
Benchmark_starlark_fib_single_ctx-12                2980           8041286 ns/op         2276678 B/op      54728 allocs/op
Benchmark_yaegi_fib_single_ctx-12                   2059          12162389 ns/op         9229806 B/op     208115 allocs/op
Benchmark_go_fib-12                               829292             29116 ns/op               0 B/op          0 allocs/op
Benchmark_starlark_fib_multi_ctx-12                 2989           8049483 ns/op         2278558 B/op      54755 allocs/op
Benchmark_glua_fib_multi_ctx-12                     6823           3107600 ns/op          181784 B/op        781 allocs/op
Benchmark_yaegi_fib_multi_ctx-12                    2089          12202257 ns/op         9247725 B/op     208214 allocs/op
PASS
ok      github.com/akkuman/go_script_lang_benchmark     316.026s
```

### Linux

- Ubuntu 18.04.4 LTS
-

```
goos: linux
goarch: amd64
pkg: github.com/akkuman/go_script_lang_benchmark
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
Benchmark_glua_Add_single_ctx-6             5420           4713832 ns/op         2205903 B/op       8616 allocs/op
Benchmark_starlark_Add_single_ctx-6         2320          10166468 ns/op         3387441 B/op     132323 allocs/op
Benchmark_yaegi_Add_single_ctx-6            7185           3293777 ns/op              24 B/op          2 allocs/op
Benchmark_go_Add-6                        512067             46732 ns/op               0 B/op          0 allocs/op
Benchmark_glua_Add_multi_ctx-6                48         466571296 ns/op       220765344 B/op    862434 allocs/op
Benchmark_starlark_Add_multi_ctx-6            24        1000723251 ns/op       338744804 B/op  13232330 allocs/op
Benchmark_yaegi_Add_multi_ctx-6               70         329575685 ns/op           25535 B/op        444 allocs/op
Benchmark_starlark_fib_single_ctx-6         3375           7049034 ns/op         1751301 B/op      21892 allocs/op
Benchmark_yaegi_fib_single_ctx-6            1388          18160846 ns/op         9219985 B/op     208115 allocs/op
Benchmark_go_fib-6                        744727             30648 ns/op               0 B/op          0 allocs/op
Benchmark_starlark_fib_multi_ctx-6          3370           7073438 ns/op         1753182 B/op      21919 allocs/op
Benchmark_glua_fib_multi_ctx-6              7570           3180979 ns/op          180410 B/op        771 allocs/op
Benchmark_yaegi_fib_multi_ctx-6             1221          17804050 ns/op         9233767 B/op     208202 allocs/op
PASS
ok      github.com/akkuman/go_script_lang_benchmark        317.636s
```
