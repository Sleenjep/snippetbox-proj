# 1 Изначальная структура веб-приложения на Go

```bash
go mod init
```

go.mod

```bash
go env
```

# 2. Основы веб-приложений на Golang

server mux в golang

```go
w http.ResponseWriter, r *http.Request
```

```go
mux := http.NewServeMux()
mux.HandleFunc("/", home)
```

```go
w.Write([]byte("message"))
```

nginx, apache

```bash
go run main.go
```

# 3. Маршрутизация HTTP-запросов используя ServeMux

фиксированные и многоуровневые пути

```go
if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
}
```

```go
http.Handle()
http.HandleFunc()
```

DefaultServeMux

```go
var DefaultServeMux = NewServeMux()
```

проблемы с использованием дефолтного server mux

особенности проверки url-шаблонов в server mux

# 4. Настройка HTTP заголовков веб-приложения

```go
if r.Method != http.MethodPost {
    w.WriteHeader(405)
    w.Write([]byte("GET-метод запрещён!"))
    return
}
```

```go
w.Header().Set("Allow", http.MethodPost)
```

```go
http.Error(w, "Метод запрещён!", http.StatusMethodNotAllowed)
```

```go
w.Header().Set("Cache-Control", "public, max-age=31536000")
 
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")
 
w.Header().Del("Cache-Control")
 
w.Header().Get("Cache-Control")
```

```go
http.DetectContentType()
```

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```

```go
textproto.CanonicalMIMEHeaderKey()
```

```go
w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
```

```go
w.Header()["Date"] = nil
```

```go
id, err := strconv.Atoi(r.URL.Query().Get("id"))
if err != nil || id < 1 {
    http.NotFound(w, r)

    return
}
```
