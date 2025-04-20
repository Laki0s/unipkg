package drivers

import "fmt"

type dnf struct{}

func (d *dnf) Install(pkg string) error { return fmt.Errorf("dnf not implemented") }
func (d *dnf) Remove(pkg string) error { return fmt.Errorf("dnf not implemented") }
func (d *dnf) Update(pkg string) error { return nil }
func (d *dnf) Search(term string) error { return nil }