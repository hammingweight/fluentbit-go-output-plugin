package main

import (
	"C"
	"fmt"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)
import "time"

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	output.FLBPluginRegister(ctx, "noopplugin", "this plugin does nothing")
	return output.FLB_OK
}

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	dec := output.NewDecoder(data, int(length))
	for {
		ret, ts, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		t := ts.(output.FLBTime).Time

		r := make(map[string]string)
		for k, v := range record {
			key := k.(string)

			var val string
			switch v := v.(type) {
			case []uint8:
				val = string(v)
			case string:
				val = v
			case nil:
				val = ""
			default:
				val = fmt.Sprintf("%v", v)
			}
			r[key] = val
		}
		fmt.Printf("%s: %s %s\n", C.GoString(tag), t.Format(time.RFC3339), r)
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
