package drivers

import (
	"fmt"
)

// Driver interface
type Driver interface {
	Install(pkg string) error
	Remove(pkg string) error
	Update(pkg string) error
	Search(term string) error
}

// Select picks native driver by distro ID
func Select(dist string) (Driver, error) {
	switch dist {
	case "ubuntu", "debian":
		return &apt{}, nil
	case "fedora", "centos", "rhel":
		return &dnf{}, nil
	case "arch", "manjaro":
		return &pacman{}, nil
	}
	return nil, fmt.Errorf("no native driver for %s", dist)
}

// Fallback returns apt as a default stub
func Fallback() (Driver, error) {
	return &apt{}, nil
}