package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/yuin/gopher-lua/parse"

	"github.com/traefik/yaegi/interp"
	lua "github.com/yuin/gopher-lua"
	"go.starlark.net/resolve"
	"go.starlark.net/starlark"
)


func Benchmark_glua_Add_single_ctx(b *testing.B) {
	src := `
function bench_gauss(n)
	for i=1,n do
		local acc = 0
		for x=0,91999 do
			acc = acc + x
		end
	end
end
	`
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(src); err != nil {
		b.Error(err)
		return
	}
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("bench_gauss"),
		NRet: 0,
		Protect: true,
	}, lua.LNumber(b.N)); err != nil {
		b.Error(err)
	}
}

func Benchmark_starlark_Add_single_ctx(b *testing.B) {
	src := `
def bench_gauss(n):
    for _ in range(n):
        acc = 0
        for x in range(92000):
			acc = acc + x
	`
	thread := new(starlark.Thread)
	globals, err := starlark.ExecFile(thread, "", src, nil)
	if err != nil {
		b.Error(err)
		return
	}
	fn := globals["bench_gauss"]
	_, err = starlark.Call(thread, fn, starlark.Tuple{starlark.MakeInt(b.N)}, nil)
	if err != nil {
		b.Error(err)
	}
}

func Benchmark_yaegi_Add_single_ctx(b *testing.B) {
	src := `
func benchGauss(n int) {
	for i := 0; i < n; i++ {
		acc := 0
		for x := 0; x < 92000; x++ {
			acc = acc + x
		}
	}
}	
`
	i := interp.New(interp.Options{})

	_, err := i.Eval(src)
	if err != nil {
		b.Error(err)
	}

	_, err = i.Eval(fmt.Sprintf("benchGauss(%d)", b.N))
	if err != nil {
		b.Error(err)
	}
}

func Benchmark_go_Add(b *testing.B) {
	fn := func(n int) {
		for i := 0; i < n; i++ {
			acc := 0
			for x := 0; x < 92000; x++ {
				acc = acc + x
			}
		}
	}
	fn(b.N)
}

func Benchmark_glua_Add_multi_ctx(b *testing.B) {
	src := `
function bench_gauss(n)
	for i=1,n do
		local acc = 0
		for x=0,91999 do
			acc = acc + x
		end
	end
end
	`
	name := "bench-gauss"
	chunk, err := parse.Parse(strings.NewReader(src), name)
    if err != nil {
        b.Error(err)
		return
    }
    proto, err := lua.Compile(chunk, name)
	if err != nil {
		b.Error(err)
		return
	}
	for n := 0; n < b.N; n++ {
		L := lua.NewState()
		lfunc := L.NewFunctionFromProto(proto)
		L.Push(lfunc)
		L.PCall(0, lua.MultRet, nil)
		if err := L.CallByParam(lua.P{
			Fn: L.GetGlobal("bench_gauss"),
			NRet: 0,
			Protect: true,
		}, lua.LNumber(100)); err != nil {
			b.Error(err)
		}
		L.Close()
	}
}

func Benchmark_starlark_Add_multi_ctx(b *testing.B) {
	src := `
def bench_gauss(n):
    for _ in range(n):
        acc = 0
        for x in range(92000):
            acc = acc + x
	`
	thread := new(starlark.Thread)
	globals, err := starlark.ExecFile(thread, "", src, nil)
	if err != nil {
		b.Error(err)
		return
	}
	fn := globals["bench_gauss"]
	for n := 0; n < b.N; n++ {
		_, err = starlark.Call(new(starlark.Thread), fn, starlark.Tuple{starlark.MakeInt(100)}, nil)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_yaegi_Add_multi_ctx(b *testing.B) {
	src := `
func benchGauss(n int) {
	for i := 0; i < n; i++ {
		acc := 0
		for x := 0; x < 92000; x++ {
			acc = acc + x
		}
	}
}
`
	i := interp.New(interp.Options{})
	prog, err := i.Compile(src)
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		itp := interp.New(interp.Options{})
		itp.Execute(prog)
		_, err = i.Eval("benchGauss(100)")
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_starlark_fib_single_ctx(b *testing.B) {
	src := `
def fib(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    return fib(n - 2) + fib(n - 1)
`
	resolve.AllowRecursion = true
	thread := new(starlark.Thread)
	globals, err := starlark.ExecFile(thread, "", src, nil)
	if err != nil {
		b.Error(err)
		return
	}
	fn := globals["fib"]
	for n := 0; n < b.N; n++ {
		_, err = starlark.Call(thread, fn, starlark.Tuple{starlark.MakeInt(20)}, nil)
		if err != nil {
			b.Error(err)
		}
	}
	resolve.AllowRecursion = false
}

// func Benchmark_glua_fib_single_ctx(b *testing.B) {
// 	src := `
// function fib(n)
//     if n == 0 then
//         return 0
//     elseif n == 1 then
//         return 1
//     end
//     return fib(n-1) + fib(n-2)
// end
// 	`
// 	L := lua.NewState()
// 	defer L.Close()
// 	if err := L.DoString(src); err != nil {
// 		b.Error(err)
// 		return
// 	}
// 	for n := 0; n < b.N; n++ {
// 		if err := L.CallByParam(lua.P{
// 			Fn: L.GetGlobal("fib"),
// 			NRet: 0,
// 			Protect: true,
// 		}, lua.LNumber(20)); err != nil {
// 			b.Error(err)
// 		}
// 		L.Close()
// 	}
// }


func Benchmark_yaegi_fib_single_ctx(b *testing.B) {
	src := `
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
`
	i := interp.New(interp.Options{})
	_, err := i.Eval(src)
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		_, err = i.Eval("fib(20)")
		if err != nil {
			b.Error(err)
		}
	}
}

func _fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return _fib(n-1) + _fib(n-2)
	}
}

func Benchmark_go_fib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_fib(20)
	}
}

func Benchmark_starlark_fib_multi_ctx(b *testing.B) {
	src := `
def fib(n):
    if n == 0:
        return 0
    elif n == 1:
        return 1
    return fib(n - 2) + fib(n - 1)
	`
	resolve.AllowRecursion = true
	thread := new(starlark.Thread)
	globals, err := starlark.ExecFile(thread, "", src, nil)
	if err != nil {
		b.Error(err)
		return
	}
	fn := globals["fib"]
	for n := 0; n < b.N; n++ {
		_, err = starlark.Call(new(starlark.Thread), fn, starlark.Tuple{starlark.MakeInt(20)}, nil)
		if err != nil {
			b.Error(err)
		}
	}
	resolve.AllowRecursion = false
}

func Benchmark_glua_fib_multi_ctx(b *testing.B) {
	src := `
function fib(n)
    if n == 0 then
        return 0
    elseif n == 1 then
        return 1
    end
    return fib(n-1) + fib(n-2)
end
	`
	name := "fin-test"
	chunk, err := parse.Parse(strings.NewReader(src), name)
    if err != nil {
        b.Error(err)
		return
    }
    proto, err := lua.Compile(chunk, name)
	if err != nil {
		b.Error(err)
		return
	}
	for n := 0; n < b.N; n++ {
		L := lua.NewState()
		lfunc := L.NewFunctionFromProto(proto)
		L.Push(lfunc)
		L.PCall(0, lua.MultRet, nil)
		if err := L.CallByParam(lua.P{
			Fn: L.GetGlobal("fib"),
			NRet: 0,
			Protect: true,
		}, lua.LNumber(20)); err != nil {
			b.Error(err)
		}
		L.Close()
	}
}


func Benchmark_yaegi_fib_multi_ctx(b *testing.B) {
	src := `
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
`
	i := interp.New(interp.Options{})
	prog, err := i.Compile(src)
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		itp := interp.New(interp.Options{})
		itp.Execute(prog)
		_, err = i.Eval("fib(20)")
		if err != nil {
			b.Error(err)
		}
	}
}