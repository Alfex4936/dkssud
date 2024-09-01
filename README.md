# dkssud: QWERTY 한영 매핑 프로그램

<p align="center">
  <img width="300" src="https://github.com/user-attachments/assets/c09a64c8-277d-4a56-ab25-f93f2d226fb4" alt="GDG"/></br>
  <span style="font-size:12px; color:grey;">logo from Google (GDG)</span>
</p>

---

[![GoDoc](https://pkg.go.dev/badge/github.com/Alfex4936/dkssud.svg)](https://pkg.go.dev/github.com//Alfex4936/dkssud)
[![codecov](https://codecov.io/gh/Alfex4936/dkssud/graph/badge.svg?token=PYJT7QQ4BW)](https://codecov.io/gh/Alfex4936/dkssud)

QWERTY 키보드용 한국어/영어 간 매핑 프로그램입니다.

QWERTY 키보드에서 한글을 입력하거나, 반대로 영어로 변환할 수 있습니다.

이 라이브러리는 [gksdudaovld 한영매핑](https://github.com/ForestHouse2316/gksdudaovld) Python 라이브러리에서 영감을 받아 제작되었습니다.


> [!NOTE]
> dkssud 은 '안녕'을 영어로 치면 나옵니다.

## 소개

`dkssud` 패키지를 사용하면 다음과 같이 QWERTY 키보드 입력을 한글로 변환할 수 있습니다:

```go
import "github.com/Alfex4936/dkssud"

func main() {
    result := dkssud.QwertyToHangul("dkssud")
    fmt.Println(result) // 출력: "안녕"
}
```

위와 같이 한국어를 영어로, 또는 영어를 한국어로 바꿀 수 있습니다.

## 설치

```bash
go get github.com/Alfex4936/dkssud
```

## 사용 예시

### QWERTY -> 한글 변환

QWERTY 키보드 입력을 한글로 변환하는 간단한 예제입니다:

```go
import "github.com/Alfex4936/dkssud"

func main() {
    hangul := dkssud.QwertyToHangul("rkskekfk")
    fmt.Println(hangul) // 출력: "가나다라"

    hangul := dkssud.QwertyToHangul("rjRlRkwldii")
    fmt.Println(hangul) // 출력: "거끼까지야ㅑ"
}
```

### 한글 -> QWERTY 변환

한글을 QWERTY 키보드 입력으로 변환하는 예제입니다:

```go
import "github.com/Alfex4936/dkssud"

func main() {
    qwerty := dkssud.HangulToQwerty("안녕하세요")
    fmt.Println(qwerty) // 출력: "dkssudgktpdy"

    qwerty = dkssud.HangulToQwerty("뮻ㅇ")
    fmt.Println(qwerty) // 출력: "abcd"
}
```

