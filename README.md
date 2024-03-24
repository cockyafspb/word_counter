# Приложение для подсчета вхождений слова по URL
Флаги:
* -f filename
* -w word_to_find
* -p workers_count\
Пример запуска:
```
Count for https://golang.org/doc/install: 43
Invalid url: https://
Invalid url: abcdefg
Count for https://golang.org/doc/cmd: 32
Count for https://golang.org/: 52       
Count for https://golang.org/doc/: 104  
Count for https://golang.org/project/: 42
Count for https://golang.org/pkg/bufio/: 28
Count for https://golang.org/pkg/compress/gzip/: 28
Invalid url: 1234565
Count for https://golang.org/ref/spec: 95
Count for https://golang.org/pkg/encoding/: 29
Count for https://golang.org/pkg/builtin/: 35
Count for https://golang.org/pkg/: 88
Count for https://golang.org/pkg/errors/: 28
Count for https://golang.org/pkg/bytes/: 40
Count for https://gobyexample.com/: 5
Invalid url: golang.org
Count for https://gobyexample.com/select: 3
Count for https://gobyexample.com/channel-buffering: 1
Count for https://blog.golang.org/protobuf-apiv2: 42
Count for https://blog.golang.org/go1.14: 31
Count for https://golang.org/pkg/io/ioutil/: 39
Count for https://blog.golang.org/pipelines: 37
Count for https://gobyexample.com/channels: 1
Count for https://gobyexample.com/structs: 2
Count for https://gobyexample.com/interfaces: 3
Count for https://golang.org/pkg/fmt/: 36
Count for https://gobyexample.com/goroutines: 3
Count for https://golang.org/pkg/net/http/: 41
Count for https://golang.org/pkg/flag/: 29
Total: 917
```
