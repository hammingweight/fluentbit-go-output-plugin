package main

import (
	"C"
	"encoding/json"
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
		jsonString, _ := json.Marshal(r)
		fmt.Printf("%s: %s %s\n", C.GoString(tag), ts.(output.FLBTime).Time.Format(time.RFC3339), jsonString)
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
