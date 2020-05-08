package requester

import "testing"

func Test_Hatch(t *testing.T) {
	hatch := NewHatch("marcus@hatchify.co", "password")

	if err := hatch.login(); err != nil {
		t.Fatal(err)
	}

	if _, err := hatch.getUser(); err != nil {
		t.Fatal(err)
	}
}
