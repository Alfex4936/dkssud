// Package dkssud는 QWERTY 키보드 입력과 한글 문자 간의 변환을 제공하는 함수들을 포함하고 있습니다.
//
// 이 패키지는 QWERTY 키보드 레이아웃을 기반으로 한국어 텍스트를 해석하거나 생성해야 하는
// 애플리케이션에 유용합니다.
//
// 이 패키지에는 두 가지 주요 함수가 포함되어 있습니다:
//
// - QwertyToHangul: QWERTY 키보드 입력 문자열을 해당하는 한글 문자로 변환합니다.
//
// - HangulToQwerty: 한글 문자를 해당하는 QWERTY 키보드 입력으로 변환합니다.
//
// 사용 예:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/Alfex4936/dkssud"
//	)
//
//	func main() {
//		// QWERTY를 한글로 변환
//		hangul := dkssud.QwertyToHangul("rkskekfk")
//		fmt.Println(hangul) // 출력: "가나다라"
//
//		// 한글을 QWERTY로 변환
//		qwerty := dkssud.HangulToQwerty("가나다라")
//		fmt.Println(qwerty) // 출력: "rkskekfk"
//	}
//
// 이 패키지는 한국어 텍스트 처리에서 일반적으로 발생하는 상황을 처리하도록 설계되었으며,
// 간단하고 사용하기 쉽게 만들어졌습니다.
package dkssud // import "github.com/Alfex4936/dkssud"
