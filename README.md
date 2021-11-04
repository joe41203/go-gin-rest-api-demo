# go-gin-rest-api-demo

## 初期設定
```
$ git clone https://github.com/joe41203/go-gin-rest-api-demo.git
$ cd go-gin-rest-api-demo
$ docker build . -t go-gin-rest-api-demo:latest
```

## ginの起動
```
$ docker run -p 8080:8080 go-gin-rest-api-demo:latest
```

## 動作確認

Home
```
$ curl http://localhost:8080
```

Post
```
$ curl -d "HelloGin" -X POST http://localhost:8080/post
```

QueryParameters
```
$ curl http://localhost:8080/query\?name=golang\&age=12
```

PathParameters
```
$ curl http://localhost:8080/path/golang/12
```
