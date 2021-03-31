package smartrg

import (
	"regexp"
	"strings"
)

func MactoUpper(mac string) string {
	reg, _ := regexp.Compile("[^a-fA-f0-9]+")
	mac = reg.ReplaceAllString(mac, "")
	mac = strings.ToUpper(mac)
	return mac
}
