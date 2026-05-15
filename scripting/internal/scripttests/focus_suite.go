package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type FocusSuite struct {
	ScriptHostSuite
}

func NewFocusSuite(h html.ScriptEngine) *FocusSuite {
	return &FocusSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *FocusSuite) TestActiveElementDefaultsToBody() {
	s.Expect(s.Eval(`document.activeElement === document.body`)).To(BeTrue())
}

func (s *FocusSuite) TestFocusUpdatesActiveElement() {
	s.Expect(s.Eval(`
		const i = document.createElement("input");
		document.body.appendChild(i);
		i.focus();
		document.activeElement === i
	`)).To(BeTrue())
}

func (s *FocusSuite) TestBlurIsCallable() {
	s.Expect(s.Eval(`
		const i = document.createElement("input");
		i.blur();
		true
	`)).To(BeTrue())
}
