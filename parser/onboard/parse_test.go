package onboard_test

import (
	"dmidecode/parser/onboard"
	"dmidecode/smbios"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   41,
			Length: 11,
			Handle: 10498,
		},
		Formatted: []byte{0x1, 0x85, 0x3, 0x0, 0x0, 0x2, 0x0},
		Strings:   []string{"Integrated NIC 3"},
	}
)

func TestParse(t *testing.T) {
	bios, err := onboard.Parse(s)
	if assert.NoError(t, err) {
		t.Log(bios)
	}
}
