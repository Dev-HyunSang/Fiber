# Fiber

## Fiber란 무엇인가요?
> Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.  

[Fiber](https://gofiber.io/)는 [Node.js Express](https://expressjs.com/ko/)을 보고 영감을 받아서 개발된 FastHTTP 기반의 웹 프레임워크 입니다.  
간단한 사용법에 대해서는 [Fiber Doce](https://docs.gofiber.io/)를 살펴 보세요.
### 왜 사용하나요?
Go 모듈 중에서 가장 빠른 HTTP 클라이언트인 FastHTTP 패키지를 사용합니다. FastHTTP와 `net/http`에 대한 비교 벤치 마킹 결과 FastHTTP가 `net/http`보다 10배 정도 빠릅니다.  

## 사용법
```shell
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/template/html
# and
go mod download
```