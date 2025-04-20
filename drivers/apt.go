package drivers

import "unipkg/internal/executil"

type apt struct{}

func (a *apt) Install(pkg string) error {
	return executil.Run("sudo", "apt", "install", "-y", pkg)
}
func (a *apt) Remove(pkg string) error {
	return executil.Run("sudo", "apt", "remove", "-y", pkg)
}
func (a *apt) Update(pkg string) error { return nil }
func (a *apt) Search(term string) error { return nil }