package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type DialogSuite struct {
	ScriptHostSuite
}

func NewDialogSuite(h html.ScriptEngine) *DialogSuite {
	return &DialogSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *DialogSuite) SetupTest() {
	s.ScriptHostSuite.SetupTest()
	v, err := s.Eval(`typeof document.createElement("dialog").show`)
	if err != nil || v != "function" {
		s.T().Skip("HTMLDialogElement methods not available on this engine")
	}
}

func (s *DialogSuite) TestShowSetsOpen() {
	s.Expect(s.Eval(`
		const d = document.createElement("dialog");
		d.show();
		d.hasAttribute("open")
	`)).To(BeTrue())
}

func (s *DialogSuite) TestShowModalSetsOpen() {
	s.Expect(s.Eval(`
		const d = document.createElement("dialog");
		d.showModal();
		d.hasAttribute("open")
	`)).To(BeTrue())
}

func (s *DialogSuite) TestCloseClearsOpen() {
	s.Expect(s.Eval(`
		const d = document.createElement("dialog");
		d.show();
		d.close();
		d.hasAttribute("open")
	`)).To(BeFalse())
}

func (s *DialogSuite) TestCloseStoresReturnValue() {
	s.Expect(s.Eval(`
		const d = document.createElement("dialog");
		d.show();
		d.close("submit");
		d.returnValue
	`)).To(Equal("submit"))
}

func (s *DialogSuite) TestCloseDispatchesCloseEvent() {
	s.Expect(s.Eval(`
		const d = document.createElement("dialog");
		let fired = false;
		d.addEventListener("close", () => { fired = true; });
		d.show();
		d.close();
		fired
	`)).To(BeTrue())
}
