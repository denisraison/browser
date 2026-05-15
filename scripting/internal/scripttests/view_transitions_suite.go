package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ViewTransitionsSuite struct {
	ScriptHostSuite
}

func NewViewTransitionsSuite(h html.ScriptEngine) *ViewTransitionsSuite {
	return &ViewTransitionsSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ViewTransitionsSuite) TestExists() {
	s.Expect(s.Eval("typeof document.startViewTransition")).To(Equal("function"))
}

func (s *ViewTransitionsSuite) TestCallbackRunsSynchronously() {
	s.Expect(s.Eval(`
		let ran = false;
		document.startViewTransition(() => { ran = true; });
		ran
	`)).To(BeTrue())
}

func (s *ViewTransitionsSuite) TestReturnsViewTransitionLike() {
	s.Expect(s.Eval(`
		const t = document.startViewTransition(() => {});
		typeof t.finished === "object" && typeof t.ready === "object" && typeof t.skipTransition === "function"
	`)).To(BeTrue())
}

func (s *ViewTransitionsSuite) TestFinishedResolves() {
	s.Assert().NoError(s.RunScript(`
		window.__vtResult = "";
		document.startViewTransition(() => {}).finished.then(() => { window.__vtResult = "done"; });
	`))
	s.Expect(s.Eval(`window.__vtResult`)).To(Equal("done"))
}
