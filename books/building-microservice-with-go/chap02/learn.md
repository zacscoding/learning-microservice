# Ch02. 좋은 API 디자인 하기

> ## Restful design

**URI 형식**  

```
URI = schema "://" authority "/" path [ "?" query] ["#" fragment]
URI = http://server.com/path?query=1#document
```  

- 슬래시는 리소스 사이의 계층적 관계를 나타내는 데사용
- URI의 마지막에 슬래시가 포함돼서는 안 됨
- 가독성을 높이기 위해 하이픈(-)을 사용
- 밑줄 문자(underscore)는 URI에서 사용 X
- URI의 경로 부분은 대소문자를 구문하므로 소문자를 사용하는 것이 좋음  

**REST 서비스를 위한 URI 경로 설계**  
; 경로는 문서(document), 컬렉션(collection), 저장소(store), 컨트롤러(controller)로 구분  


> 컬렉션(Collection)


```
// 컬렉션 이름은 복수 명사
GET	/cats		-> 모든 고양이가 컬렉션에 들어 있음
GET	/cats/1		-> 1번 고양이를 위한 하나의 문서
```

> 문서(Document)  

; 문서는 DB의 행(row)과 비슷한 하나의 객체를 가리키는 리소스  
하나의 문서는 하위 문서 OR 같은 하위 리소스(child resource)를 가질 수 있음  


```
GET	/cats/1				-> 1번 고양이를 위한 하나의 문서
GET	/cats/1/kittens		-> 1번 고양이의 모든 새끼 고양이들(kittens)
GET	/cats/1/kittens/1	-> 1번 고양이의 1번 새끼 고양이
```  

> 컨트롤러(Controller)  

; 컨트롤러 리소스(Controller resource)는 프로시저와 비슷하지만 리소스를 표준CRUD  
기능에 매핑 할 수 없는 경우에 사용  
=> 컨트롤러에 이름을 정의할 때는 항상 동사를 사용


```
POST	/cats/1/feed			-> 1번 고양이에게 먹이 주기
POST	/cats/1/feed?food=fish	-> 1번 고양이에게 물고기를 먹이로 주기
```

> 저장소(Store)  

저장소는 클라이언트가 관리하는 리소스 저장소이며 클라이언트가 리소스를 추가,  
검색 및 삭제할 수 있게 함  

```
// id가 2인 새 고양이 추가
PUT	/cats/2
```

**HTTP 동사**  

- GET
;GET 메소드는 리소스를 검색하는 데 사용(변경 작업X)  
- POST  
;컬렉션에 새로운 리소스를 만들거나 컨트롤러를 실행하는 데 사용   
=> 일반적으로 비멱등(non-idempotent, 반복 수행 시 매번 변경)  
- PUT  
; 변경 가능한 리소스를 업데이트 하는 데 사용  
=> 항상 리소스 식별정보(resource locator)를 포함해야 함  
- PATCH  
; 부분 업데이트를 수행하는 데 사용  
- DELETE  
; 리소르를 제거하려는 경우에 사용 (리소스의 ID를 전달하여)  
- HEAD  
; 클라이언트는 본문 없이 리소스에 대한 헤더만 검색하려는 경우 HEAD 동사를 사용  
- OPTIONS  
; 클라이언트가 서버의 리소스에 대해 수행 가능한 동작을 알아보려고 할 때 사용  
=> 일반적으로 Allow 헤더를 리턴  

**URI query design**  

- 페이징 처리(Paging)
- 필터링 (Filtering)
- 정렬 (Sorting)  

**응답 코드**  

**2xx Success(성공)**  
; 클라이언트의 요청이 성공적으로 수신되고 이해됐음을 나타냄

- 200 OK  
  - GET : 요청된 리소스에 해당하는 엔티티(entity)  
  - HEAD : 메시지 본문 없이 요청된 리소스에 해당하는 헤더 필드(field)
  - POST : 처리 결과를 설명하거나 포함하고 있는 엔티티  

- 201 Created (생성)
  - 요청이 성공하고 새 엔티티가 생성된 경우 응답을 보냄  
  - 일반적으로 API는 응답과 함께 생성된 엔티티의 위치가 있는 Location 헤더를 리턴    
  ```
  201 Created
  Location : https://api.kittens.com/v1/kittens/123dfdf1111
  ```  
- 204 No Content(내용 없음)  
  - 클라이언트의 요청이 성공적으로 처리됐음을 나타냄(본문은 없음)
  - e.g) DELETE 요청에 대한 응답

**3xx Redirection(리다이렉션, 경로 재지정)**  
; 클라이언트가 요청을 완료하기 위해 추가 조치를 취해야 함을 나타내는 상태 코드  
클래스를 나타냄  
대부분 CDN 및 기타 콘텐츠 리디렉션 기법에서 사용되지만, 304 코드는 클라이언트에게  
의미론적 피드백을 제공하기 위해 API를 설계 할 때 유용이 사용  

- 301 Moved Permanently (영구적 이동)  
  - 요청한 리소스가 영구적으로 다른 위치로 이동됐음을 클라이언트에게 알려줌
  - 일반적인 경우가 아닌 예외적인 상황에서만 사용해야 함  
  (클라이언트는 묵시적으로implicitly 301 리다이렉션을 따르지 않으며 기능을 구현하는 것이  
    고객의 복잡성을 늘릴 수 있음)
- 304 Not Modified (변경 없음)
  - 일반적으로 CDN 또는 캐싱(caching) 서버에서 사용되며 API에 대한 마지막 호출 이후  
  응답이 변경되지 않았음을 나타냄  
  - 대역폭(bandwidth)을 절약하기 위해 설계됐으며 요청이 본문을 반환하지 않지만  
  Content-Location 및 Expires 헤더를 리턴함  

**4xx Client Error(클라이언트 에러)**  
; 서버가 아닌 클라이언트로 인해 발생한 에러의 경우 서버는 4xx 응답을 리턴하고  
항상 에러에 대한 자세한 내용을 제공하는 엔티티를 리턴  

- 400 Bad Request(잘못 된 요청)
  - 잘못된 형식의 요청 또는 도메인 유효성 검사 실패로 인해 클라이언트가  
  요청을 이해할 수 없음을 나타냄
- 401 Unahuthorized(권한 없음)
  - 요청이 사용자 인증을 요구하고 리소스를 요청하는데 사용할 수 있는 챌린지  
  (challenge, 인증을 요청하는 데 필요한 값)를 포함하는 WWW-Authenticate 헤더를  
  포함하고 있음을 나타낸다.
  - 사용자가 필수 자격 증명(required credential)을 WWW-Authenticate 헤더에 포함해  
  요청한 경우, 응답에 관련 진단 정보(diagnostic information)를 가지고 있는 에러 객체가  
  포함돼야 함
- 403 Forbidden(접근 금지)
  - 서버가 요청을 이해했지만 요청에 대한 실행을 거부하는 것을 의미  
  - 인증되지 않은 사용자가 리소스에 대해 잘못된 수준으로 접근을 요청했기 때문 일 수 있음
  - 접근할 수 없다는 사실을 공개하고 싶지 않으면 404 Not found 상태를 리턴할 수 도 있음
- 404 Not Found(찾을 수 없음)
  - 서버가 요청된 URI와 일치하는 것을 찾지 못했음을 나타냄  
- 405 Method Not Allowed(허용되지 않은 메소드)  
  - 요청에 지정된 메소드가 URI로 표시된 리소스에 허용되지 않음을 의미    
- 408 Request Timeout(요청 시간 초과)
  - 서버의 대기 시간 내에ㅐ 클라이언트가 요청을 보내지 않은 상태  

**5xx Server Error(서버 오류)**  
; 무언가 문제가 생긴 상태  
=> 영구적이거나 일시적인 에러에 대한 설명이 포함된 에러 엔티티를 응답과 함께 리턴  
=> 보안상 스택 트레이스나 에러에 대한 내부 정보는 실제로 시스템을 손상시키는 데 활용 될 수 있으므로
매우 일반적인 내용만 리턴

- 500 Internal Server Error (내부 서버 에러)
  - 계획대로 진행되지 않았음을 나타내는 일반적인 에러 메시지
- 503 Service Unavailable(서비스 이용 불가)
  - 일시적인 과부하 또는 유지 관리로 인해 현재 서버를 사용할 수 없는 상태  

**HTTP 헤더**  
; 요청 헤더는 HTTP 요청 및 응답 프로세스에서 매우 중요한 요소로, 표준적인 접근 방식을  
구현하면 사용자가 한 API에서 다른 API로 전환하는 데 도움이 된다  
=> RFC 7231 https://tools.ietf.org/html/rfc7231  

- **표준 요청 헤더**  
  - 요청 헤더는 요청 및 API 응답에 대한 추가 정보를 제공  
  -
- **Authorization - 문자열**  

TODO 정리하기  

---  

## RPC(Remote Procedure Call)  
; 원격 프로시저 호출의 약자로 원격 장비에 있는 함수나 메소드를 실행하는 방법  

### Gob  
https://golang.org/pkg/encoding/gob/  
; Go 프로세스 사이의 통신을 용이하게 하기 위해 특별히 고안되었으며  
Protocol Buffer 같은 것보다 사용하기 쉬우면서 좀 더 효율적일 수 있는 것을 만들려는  
아이디어를 바탕으로 설계돼었음(다른 언어 사이의 통신에서는 추가적인 비용 발생)  

> gob 객체 정의  

```
type HelloWorldRequest struct {
  Name string
}
```  

### Thrift  
https://thrift.apache.org
; 페이스북에서 만들었으며 2007년에 공개  
=> 현재는 아파치 재단에서 관리하고 있음

Thrift 의 주요 목표  

- 단순함 : 직관적이고 친숙하며 불필요한 의존성이 없도록 작성
- 투명성 : 다른 언어에서 일반적인 관용구(idiom)는 그대로 준수
- 일관성 : 간결함, 개별 언어를 위한 기능은 핵심 라이브러리가 아닌 확장 기능에 추가  
- 성능 : 성능을 우선으로 추구하며, 우아하믕ㄴ 나중 문제

> thrift 서비스 정의  

```
struct User {
  1: string name,
  2: i32 id,
  3: string email
}

struct Error {
  1: i32 code,
  2: string detail
}

service Users {
  Error createUser(1: User user)
}
```  

### Protocol Buffer  
https://developers.google.com/protocol-buffers/
; 구글에서 만들었으며 최근 세 번째 버전이 나옴  
=> 생성자(C로 작성)가 10개 이상의 언어에 대해 클라이언트 및 서버 스텁을 읽고  
생성할 수 있는 DSL을 제공하는 방식을 취함  
=> 기본적인 10개의 언어는 구글에서 관리하며 Go, Java, C, NodeJS 용 JS가 여기에 포함  
=> 플러그형(pluggable) 아키텍처이므로 RPC를 포함한 모든 종류의 엔드 포인트를 생성하기  
위한 플러그인을 작성할 수 있음  

> Protocol BUffer 서비스 정의  

```
service Users {
  rpc CreateUser (User) return (Error) {}
}

message User {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;
}

message Error {
  optional code int32 = 1
  optional detail string = 2
}
```  

### JSON-RPC  
http://www.jsonrpc.org/specification
; RPC용 객체를 표현하는 표준 방식으로 JSON을 사용하려는 시도  

> JSON-RPC로 직렬화된 요청  

```
{
  "jsonrpc": "2.0",
  "method" : "Users.v1.CreateUser",
  "params" : {
    "name" : "Nic Jackson",
    "id" : 1234
  },
  "id" : 1
}
```  

> JSON-RPC로 직렬화된 응답  

```
{
  "jsonrpc" : "2.0",
  "result": {...}
  "id": 1
}
```
























<br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br /><br />
