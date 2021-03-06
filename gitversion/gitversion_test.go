// DO NOT EDIT: Generated by https://github.com/spudtrooper/gitversion
package gitversion

import "testing"

func TestCheckVersionFlagTrue(t *testing.T) {
	*version = true
	if !CheckVersionFlag() {
		t.Fatalf("expected true, got false")
	}
}

func TestCheckVersionFlagFalse(t *testing.T) {
	*version = false
	if CheckVersionFlag() {
		t.Fatalf("expected false, got true")
	}
}
