package main

import (
	"C"
	"fmt"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)

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
	var records []map[string]string
	for {
		ret, ts, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		fmt.Println(ts)
		t := ts.(output.FLBTime).Time
		fmt.Println(t)

		r := make(map[string]string)
		for k, v := range record {
			key := k.(string)

			var val string
			switch v.(type) {
			case []uint8:
				val = string(v.([]uint8))
			case string:
				val = v.(string)
			case nil:
				val = ""
			default:
				val = fmt.Sprintf("%v", v)
			}
			r[key] = val
		}
		r["tag"] = C.GoString(tag)
		records = append(records, r)
		fmt.Println(records)
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
