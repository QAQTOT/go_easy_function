package GoEasyFunction

import (
	"github.com/QAQTOT/go_easy_function/easy_request"
	"testing"
)

func TestCarbon(t *testing.T) {

	type Person struct {
		Name    string
		Age     int
		Address string
	}
	//form, err := easy_request.Get("https://www.baidu.com", "/", Person{
	//	Name:    "John Doe",
	//	Age:     30,
	//	Address: "123 Main St",
	//})

	res, err := easy_request.Get("https://www.baidu.com", "", map[string]string{
		"Name":    "John，， Doe哈哈哈",
		"Age":     "30！！",
		"Address": "123 Main St",
	})

	t.Log(res, err)
	//t.Log(quick_func.SubString("111212ads", 0, -1))
	//
	//t.Log(carbon.Now().GetDateTimeString())
	//t.Log(carbon.Now().GetDateString())
	//t.Log(carbon.Now().GetTimeString())
	//
	//t.Log(carbon.Now().GetUnixTimeStamp())
	//t.Log(carbon.Now().GetUnixMicroTimeStamp())
	//t.Log(carbon.Now().GetUnixNanoTimeStamp())

}
