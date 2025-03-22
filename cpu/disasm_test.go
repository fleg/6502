package cpu

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisasm(t *testing.T) {
	bin, err := os.ReadFile("../test_suites/disasm/code.bin")
	assert.NoError(t, err)

	expected := "" +
		"0000: 48        pha\n" +
		"0001: 68        pla\n" +
		"0002: 6a        ror a\n" +
		"0003: a910      lda #$10\n" +
		"0005: a510      lda $10\n" +
		"0007: 9410      sty $10,x\n" +
		"0009: b610      ldx $10,y\n" +
		"000b: f002      beq *+4\n" +
		"000d: d0fa      bne *-4\n" +
		"000f: 4ccdab    jmp $abcd\n" +
		"0012: 9d0130    sta $3001,x\n" +
		"0015: 390140    and $4001,y\n" +
		"0018: 6ccdab    jmp ($abcd)\n" +
		"001b: a140      lda ($40,x)\n" +
		"001d: b140      lda ($40),y\n"

	actual := Disasm(bin)
	assert.Equal(t, expected, actual)
}
