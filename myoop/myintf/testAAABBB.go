package myintf

import (
	"go-study/myoop/myintf/a"
	"go-study/myoop/myintf/b"
)

/*
如果回圈引用就会报以下错误：
import cycle not allowed
package main
	imports go-study/myoop/myintf
	imports go-study/myoop/myintf/a
	imports go-study/myoop/myintf/b
	imports go-study/myoop/myintf/a
*/
func TestAAABBBCCC() {
	obj := a.AAA{"AAA"}
	obj1 := b.BBB{"BBB"}
	obj.DisAAA(&obj1)
	obj1.DisBBB(&obj)
}
