package gosi

import (
	"fmt"
	"strconv"

	"github.com/shirou/gopsutil/v3/host"
)

func getTemperatures() (string, error) {
	t, err := host.SensorsTemperatures()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	for _, v := range t {
		if v.SensorKey == "ACPI\\ThermalZone\\THM0_0" {
			return strconv.FormatFloat(v.Temperature, 'f', -1, 64) + "℃", nil
		}
	}
	return "", nil
}
