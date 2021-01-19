package output

import "testing"

// Pecho使用案例
func TestPecho(t *testing.T) {
	Pecho(1, "232", "cdc")
	Pecho("heellp")
}

// Pvar_dump()使用案例
func TestPvar_dump(t *testing.T) {
	Pvar_dump(1, "h")
}

// Pprint使用案例
func TestPprint(t *testing.T) {
	Pprint("hello world")
}
