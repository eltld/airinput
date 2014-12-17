// Package main provides ...
package main

import (
	"time"

	"github.com/netease/airinput/go-airinput"
	"github.com/robertkrimen/otto"
)

var vm = otto.New()

func init() {
	totime := func(msec int64) time.Duration {
		return time.Millisecond * time.Duration(msec)
	}
	vm.Set("tap", func(call otto.FunctionCall) otto.Value {
		x, _ := call.Argument(0).ToInteger()
		y, _ := call.Argument(1).ToInteger()
		msec, _ := call.Argument(2).ToInteger()
		airinput.Tap(int(x), int(y), totime(msec))
		return otto.UndefinedValue()
	})
	vm.Set("drag", func(call otto.FunctionCall) otto.Value {
		x0, _ := call.Argument(0).ToInteger()
		y0, _ := call.Argument(1).ToInteger()
		x1, _ := call.Argument(2).ToInteger()
		y1, _ := call.Argument(3).ToInteger()
		steps, _ := call.Argument(4).ToInteger()
		msec, _ := call.Argument(5).ToInteger()
		airinput.Drag(int(x0), int(y0), int(x1), int(y1), int(steps), totime(msec))
		return otto.UndefinedValue()
	})
	vm.Set("pinch", func(call otto.FunctionCall) otto.Value {
		ax0, _ := call.Argument(0).ToInteger()
		ay0, _ := call.Argument(1).ToInteger()
		ax1, _ := call.Argument(2).ToInteger()
		ay1, _ := call.Argument(3).ToInteger()
		bx0, _ := call.Argument(4).ToInteger()
		by0, _ := call.Argument(5).ToInteger()
		bx1, _ := call.Argument(6).ToInteger()
		by1, _ := call.Argument(7).ToInteger()
		steps, _ := call.Argument(8).ToInteger()
		msec, _ := call.Argument(9).ToInteger()
		airinput.Pinch(
			int(ax0), int(ay0), int(ax1), int(ay1),
			int(bx0), int(by0), int(bx1), int(by1),
			int(steps), totime(msec))
		return otto.UndefinedValue()
	})
}

// //abc = 1 + 2
// console.log("The value of abc is " + abc); // 4
func RunJS(code string) (otto.Value, error) {
	return vm.Run(code)
}