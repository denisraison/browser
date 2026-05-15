package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type WebAnimationsSuite struct {
	ScriptHostSuite
}

func NewWebAnimationsSuite(h html.ScriptEngine) *WebAnimationsSuite {
	return &WebAnimationsSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *WebAnimationsSuite) TestAnimateExists() {
	s.Expect(s.Eval("typeof document.body.animate")).To(Equal("function"))
}

func (s *WebAnimationsSuite) TestAnimateReturnsAnimationLike() {
	s.Expect(s.Eval(`
		const a = document.body.animate([], 0);
		typeof a.play === "function" && typeof a.cancel === "function" && a.playState === "finished"
	`)).To(BeTrue())
}

func (s *WebAnimationsSuite) TestFinishedAndReadyResolveImmediately() {
	s.Assert().NoError(s.RunScript(`
		const a = document.body.animate([], 0);
		window.__waResult = "";
		Promise.all([a.finished, a.ready]).then(() => { window.__waResult = "done"; });
	`))
	s.Expect(s.Eval(`window.__waResult`)).To(Equal("done"))
}
