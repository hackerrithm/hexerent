package clientMachineInformer

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// MemoryData analyses system data such as memory usage
func MemoryData() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("\n\n\t\tTotal: %v,\n\t\tFree:%v,\n\t\tUsedPercent:%f%%\n\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

	fmt.Printf("\n\n\n")
}
