package drivers

import "fmt"

type snap struct{}

func (s *snap) Install(pkg string) error { return fmt.Errorf("snap not implemented") }
func (s *snap) Remove(pkg string) error { return fmt.Errorf("snap not implemented") }
func (s *snap) Update(pkg string) error { return nil }
func (s *snap) Search(term string) error { return nil }