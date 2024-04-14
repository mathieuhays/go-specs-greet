package interactions

import (
	"github.com/mathieuhays/go-specs-greet/specifications"
	"testing"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.GreetAdapter(Curse),
	)
}
