# Go Unit Test

- [https://www.joinc.co.kr/w/man/12/golang/TDD](https://www.joinc.co.kr/w/man/12/golang/TDD) 을 읽고 정리한 내용입니다.

```
|--- handler
|   \--- handler.go
|   \--- handler_test.go
|--- math
|   \--- math.go
|   \--- math_test.go
|--- main.go
```

## 1. testing 패키지

- testing 패키지는 Go에 내장되어 있는 테스트 프레임워크다.

- 함수가 실행 된 결과가 예측한 결과와 맞아 떨어지는 지를 검사하는 방식으로 테스트를 진행한다. t.Fatal(), t.Fail() 등을 이용해서 테스트를 제어할 수 있다.
- `FailNow()` : 테스트 함수를 즉시 종료하고 다음 테스트 함수를 실행한다.
- `Fatal()` : 로그를 출력하는 것을 제외하고 FailNow()와 동일하다.
- `Fail()` : 테스트가 실패하더라도 함수를 종료하지 않고 다음 코드를 계속 실행한다.
- `Error()` : 로그를 출력하는 걸 제외하고 Fail 메서드와 같은 일을 한다.
- `Errorf()` : 형식화된 로그를 출력한다. Fail 메서드와 같은 일을 한다.
- `Log()` : 테스트 로그를 출력한다.
- `Logf()` : 형식화된 테스트 로그를 출력한다.
- `Failed()` : 실패하더라도 레포트하지 않는다.

## 2. HTTP 핸들러 테스트

- HTTP 핸들러의 경우 웹 서버를 구현 해야 하기 때문에, 메서드보다 테스트가 까다롭다.
  - net/http/httptest 패키지를 이용해서 테스트를 진행한다.
  - httptest를 이용하면 루프백(127.0.0.1)에 바인드 되는 서버를 구현 할 수 있다.
  - 이 후, net/http에서 제공하는 클라이언트 메서드들을 이용하여 서버/클라이언트 모드에서 테스트를 진행 할 수 있다.

- httptest 패키지는 테스트를 위해서 내장된 웹 서버를 실행한다.
  - 핸들러 등록, 데이터베이스 연결과 같이 서비스를 위해서 필요한 자원들을 초기화해야 한다.
  - 로그를 표준출력 할 수 없기 때문에 파일(test.log)에 access log를 남기도록 한다.

## 3. 유닛테스트 실행

- go test 명령으로 테스트를 수행하면 된다.
  - 명령을 실행한 패키지에서 *_test.go 파일을 찾아 유닛테스트를 수행한다.

> go test -v

![1](/img/gotest1.png)

- 테스트 커버리지 확인
  - go test 명령에 옵션을 추가하여 커버리지를 확인할 수 있다.

> go test -cover  
> go test -coverprofile=coverage.out

![2](/img/gotest2.png)
![3](/img/gotest3.png)

- go tool 명령으로 커버리지 파일을 HTML로 상세하게 볼 수 있다.
  - 테스트가 적용되지 않은 부분을 표시해준다.

> go tool cover -html=coverage.out

![4](/img/gotest4.png)

- 외부 패키지(gocov)를 이용하여 퀄리티 높은 커버리지 레포트를 생성해 보자.
  - 패키지를 다운로드 받고, gocov 명령으로 커버리지 결과를 json 파일로 저장한다.
  - gocov-html은 json 파일의 내용으로 보기 편한 HTML 레포트를 작성해 준다.

> go get github.com/axw/gocov/gocov  
> go get github.com/matm/gocov-html  
> gocov test > handler.json  
> gocov-html handler.json > handler.html

![5](/img/gotest5.png)

![6](/img/gotest6.png)

![7](/img/gotest7.png)
