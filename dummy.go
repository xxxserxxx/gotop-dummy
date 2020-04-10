package dummy

import (
	"math/rand"
	"time"

	"github.com/xxxserxxx/gotop/v3/devices"
)

func init() {
	devices.RegisterTemp(updateTemp)
	devices.RegisterMem(updateMem)
	devices.RegisterCPU(updateUsage)
}

func updateTemp(temps map[string]int) map[string]error {
	temps["Crazy1"] = rand.Intn(50) + 50
	temps["Crazy2"] = rand.Intn(50) + 50
	return nil
}

func updateMem(mems map[string]devices.MemoryInfo) map[string]error {
	m := uint64(rand.Intn(80e8))
	max := uint64(80e8)
	mems["Dum1"] = devices.MemoryInfo{
		Total:       max,
		Used:        m,
		UsedPercent: (float64(m) / float64(max)) * 100,
	}
	return nil
}

func updateUsage(cpus map[string]int, _ time.Duration, _ bool) map[string]error {
	cpus["Wopper01"] = rand.Intn(100)
	cpus["Wopper02"] = rand.Intn(100)
	return nil
}
