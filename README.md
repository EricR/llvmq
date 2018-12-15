# llvmq
A query tool for LLVM IR.

## Usage

```
$ llvmq testdata/coreutils/test/stat.ll "ll.Funcs.forEach(function(func) { console.log(func.GlobalIdent.GlobalName) })"
2018/12/15 14:11:30 Parsing testdata/coreutils/test/stat.ll
2018/12/15 14:11:30 Parsed in 0.231463s
2018/12/15 14:11:30 Executing query
usage
llvm.dbg.declare
llvm.dbg.value
dcgettext
__fprintf_chk
__printf_chk
fputs_unlocked
llvm.lifetime.start.p0i8
llvm.memcpy.p0i8.p0i8.i64
strcmp
setlocale
strncmp
llvm.lifetime.end.p0i8
exit
main
bindtextdomain
textdomain
localeconv
strlen
atexit
[...]
2018/12/15 14:11:30 Query executed in 0.069387s

undefined
```

# Credits

Big thanks to:
- [robertkrimen/otto](https://github.com/robertkrimen/otto) for the embedded JavaScript interpreter.
- [llir/llvm](https://github.com/llir/llvm) for the LLVM IR parser.