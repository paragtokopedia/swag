package api

import (
	"github.com/gin-gonic/gin"
)

type Foo struct {
	Field1 string
}
type Bar struct {
	Field2 string
}

type FooBar struct {
	Foo
	Bar
}

type FooBarPointer struct {
	*Foo
	*Bar
}

// @Description get Foo
// @ID get-foo
// @Accept json
// @Produce json
// @Success 200 {object} api.Foo
// @Router /testapi/get-foo [get]
func GetFoo(c *gin.Context) {
	//write your code
	var _ = Foo{}
}

// @Description get Bar
// @ID get-bar
// @Accept json
// @Produce json
// @Success 200 {object} api.Bar
// @Router /testapi/get-bar [get]
func GetBar(c *gin.Context) {
	//write your code
	var _ = Bar{}
}

// @Description get FooBar
// @ID get-foobar
// @Accept json
// @Produce json
// @Success 200 {object} api.FooBar
// @Router /testapi/get-foobar [get]
func GetFooBar(c *gin.Context) {
	//write your code
	var _ = FooBar{}
}

// @Description get FooBarPointer
// @ID get-foobar-pointer
// @Accept json
// @Produce json
// @Success 200 {object} api.FooBarPointer
// @Router /testapi/get-foobar-pointer [get]
func GetFooBarPointer(c *gin.Context) {
	//write your code
	var _ = FooBarPointer{}
}
