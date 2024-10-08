package environment

import (
	"errors"
	"strings"

	"github.com/lbenetti/cisco_exporter/rpc"
	"github.com/lbenetti/cisco_exporter/util"
)

// Parse parses cli output using textfsm and tries to find all temperature, power & fan related data
func (c *environmentCollector) Parse(ostype string, output string) ([]EnvironmentItem, error) {
	if ostype != rpc.IOSXE && ostype != rpc.IOS {
		return nil, errors.New("'show environment' is not implemented for " + ostype)
	}
	items := []EnvironmentItem{}

	results_temp, err := util.ParseTextfsm(templ_temp, output)
	if err != nil {
		return items, errors.New("Error parsing via templ_temp: " + err.Error())
	}
	results_power, err := util.ParseTextfsm(templ_power, output)
	if err != nil {
		return items, errors.New("Error parsing via templ_power: " + err.Error())
	}
	results_fan, err := util.ParseTextfsm(templ_fan, output)
	if err != nil {
		return items, errors.New("Error parsing via templ_fan: " + err.Error())
	}
	for _, result := range results_temp {
		location := result["LOCATION"].(string)
		sensor := result["SENSOR"].(string)
		state := strings.ToLower(strings.TrimSpace(result["STATE"].(string)))
		state_ok := state == "normal" || state == "good" || state == "ok" || state == "green"
		x := EnvironmentItem{
			Name:        strings.TrimSpace(location + " " + sensor),
			IsTemp:      true,
			OK:          state_ok,
			Status:      state,
			Temperature: util.Str2float64(result["VALUE"].(string)),
		}
		items = append(items, x)
	}
	for _, result := range results_power {
		location := result["LOCATION"].(string)
		sensor := result["SENSOR"].(string)
		state := strings.ToLower(strings.TrimSpace(result["STATE"].(string)))
		state_ok := state == "normal" || state == "good" || state == "ok" || state == "green"
		x := EnvironmentItem{
			Name:    strings.TrimSpace(location + " " + sensor),
			IsPower: true,
			OK:      state_ok,
			Status:  state,
			Power:   util.Str2float64(result["VALUE"].(string)),
		}
		items = append(items, x)
	}
	for _, result := range results_fan {
		location := result["LOCATION"].(string)
		sensor := result["SENSOR"].(string)
		state := strings.ToLower(strings.TrimSpace(result["STATE"].(string)))
		state_ok := state == "normal" || state == "good" || state == "ok" || state == "green"
		x := EnvironmentItem{
			Name:   strings.TrimSpace(location + " " + sensor),
			IsFan:  true,
			OK:     state_ok,
			Status: state,
			Fan:    util.Str2float64(result["VALUE"].(string)),
		}
		items = append(items, x)
	}
	return items, nil
}
