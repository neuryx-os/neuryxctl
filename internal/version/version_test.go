package version

import "testing"

func TestDefaultVersion(t *testing.T) {
	if got := String(); got != "dev" {
		t.Fatalf("String() = %q, esperado %q", got, "dev")
	}
}

func TestOverride(t *testing.T) {
	old := Version
	t.Cleanup(func() { Version = old })

	Version = "1.2.3"
	if got := String(); got != "1.2.3" {
		t.Fatalf("String() = %q, esperado %q", got, "1.2.3")
	}
}
