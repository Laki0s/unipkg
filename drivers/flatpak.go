package drivers

import "fmt"

type flatpak struct{}

func (f *flatpak) Install(pkg string) error { return fmt.Errorf("flatpak not implemented") }
func (f *flatpak) Remove(pkg string) error { return fmt.Errorf("flatpak not implemented") }
func (f *flatpak) Update(pkg string) error { return nil }
func (f *flatpak) Search(term string) error { return nil }