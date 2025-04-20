package distro

import (
	"bufio"
	"os"
	"strings"
)

// Detect reads /etc/os-release to get distro ID
func Detect() string {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return ""
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		if strings.HasPrefix(l, "ID=") {
			return strings.Trim(strings.SplitN(l, "=", 2)[1], "\"")
		}
	}
	return ""
}