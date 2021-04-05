package bdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Behavior Driven Development

BDD in GO

	https://github.com/smartystreets/goconvey

安装

go get -u github.com/smartystreets/goconvey/convey

启动 WEB UI

$GOPATH/bin/goconvey
*/
func TestSpec(t *testing.T) {

	Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("When add the tow numbers", func() {
			c := a + b

			Convey("Then the result is still even", func() {

				So(c%2, ShouldEqual, 0)

			})

		})
	})
}
