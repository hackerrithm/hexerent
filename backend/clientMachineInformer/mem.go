package clientMachineInformer

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// MemoryData does stuff
func MemoryData() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v,\n Free:%v,\n UsedPercent:%f%%\n\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Printf("\n\n\n")
}
