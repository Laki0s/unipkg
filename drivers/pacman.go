package drivers

import "fmt"

type pacman struct{}

func (p *pacman) Install(pkg string) error { return fmt.Errorf("pacman not implemented") }
func (p *pacman) Remove(pkg string) error { return fmt.Errorf("pacman not implemented") }
func (p *pacman) Update(pkg string) error { return nil }
func (p *pacman) Search(term string) error { return nil }