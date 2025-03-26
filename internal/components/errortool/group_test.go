package errortool

import (
	"github.com/pkg/errors"
	"testing"
)

func TestDefine(t *testing.T) {
	d := DefaultGroup()
	d.Error("aaa")
	//d.Error("bbb")
	g1 := d.Group()
	g1.Error("ccc")
	g1_1 := g1.Group()
	g1_1.Error("ccc_1")

	g2 := d.CustomGroup("g2")
	g2.Error("ddd")
	g2_1 := g2.Group()
	g2_1.CustomError("ddd_1", "ddd_1")

	t.Logf("%+v", d.Codes())
}

func TestWrap(t *testing.T) {
	d := DefaultGroup()
	group_a := d.CustomGroup("a")
	aaa := group_a.Error("masg aaa")

	group_b := d.CustomGroup("b")
	bbb := group_b.CustomError("12345", "masg bbb")

	a_func := func() Error {
		return aaa.Trace("i'm aaa")
	}

	ccc := bbb.TraceWrap(a_func(), "i'm bbb")

	t.Logf("%+v", ccc)
}

func TestIs(t *testing.T) {
	d := DefaultGroup()
	group_a := d.CustomGroup("a")
	aaa := group_a.Error("masg aaa")

	group_b := d.CustomGroup("b")
	bbb := group_b.CustomError("12345", "masg bbb")

	a_func := func() Error {
		return aaa.Trace("i'm aaa")
	}

	ccc := bbb.TraceWrap(a_func(), "i'm bbb")

	t.Logf("%+v", ccc)
	t.Logf("Is: %v", ccc.Is(aaa))
	t.Logf("Is: %v", ccc.Is(bbb))
}

func TestErrorIs(t *testing.T) {
	errorGroup := DefaultGroup()
	errorGroupA := errorGroup.CustomGroup("a")
	errAAA := errorGroupA.Error("masg aaa")

	errorGroupB := errorGroup.CustomGroup("b")
	errBBB := errorGroupB.CustomError("12345", "masg bbb")

	errCCC := func() error {
		return errAAA
	}()

	errDDD := errors.New("masg ddd")

	t.Logf("Is: %v", errors.Is(errAAA, errAAA))
	t.Logf("Is: %v", errors.Is(errAAA, errBBB))
	t.Logf("Is: %v", errors.Is(errAAA, errCCC))
	t.Logf("Is: %v", errors.Is(errAAA, errDDD))
}
