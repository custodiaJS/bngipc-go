package bngipcgo

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func getUptime() (int64, error) {
	switch runtime.GOOS {
	case "linux":
		// Uptime aus /proc/uptime lesen
		out, err := exec.Command("cat", "/proc/uptime").Output()
		if err != nil {
			return 0, err
		}
		uptimeStr := strings.Split(string(out), " ")[0]
		uptimeSeconds, err := strconv.ParseFloat(uptimeStr, 64)
		if err != nil {
			return 0, err
		}
		return int64(uptimeSeconds), nil

	case "darwin", "freebsd":
		// Uptime auf macOS oder BSD lesen
		out, err := exec.Command("sysctl", "-n", "kern.boottime").Output()
		if err != nil {
			return 0, err
		}
		bootTimeStr := strings.Split(strings.TrimSpace(string(out)), " ")[3]
		bootTimeStr = strings.Trim(bootTimeStr, ",")
		bootTimeUnix, err := strconv.ParseInt(bootTimeStr, 10, 64)
		if err != nil {
			return 0, err
		}
		return bootTimeUnix, nil

	default:
		return 0, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}
