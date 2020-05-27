package main
 
//필요 패키지 임포트
import (
    "fmt"
    "log"
    "net/http"
    "strings"
)
 
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    //Get 파라미터 및 정보 출력
    fmt.Println("default : ", r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("param : ", r.Form["test_param"])
    //Parameter 전체 출력
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    //기본 출력
    fmt.Fprintf(w, "Golang WebServer Working!")
}
 
func main() {
    //기본 Url 핸들러 메소드 지정
    http.HandleFunc("/", defaultHandler)
    //서버 시작
    err := http.ListenAndServe(":9090", nil)
    //예외 처리
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    } else {
        fmt.Println("ListenAndServe Started! -> Port(9090)")
    }
}
