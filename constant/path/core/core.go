// 내부 API 경로 정의 패키지
package core

type InternalApi struct {
	Location LocationPath
}

type LocationPath struct {
	FindEvent string
}
