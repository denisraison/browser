package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type MatchMediaSuite struct {
	ScriptHostSuite
}

func NewMatchMediaSuite(h html.ScriptEngine) *MatchMediaSuite {
	return &MatchMediaSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *MatchMediaSuite) TestMatchMediaIsFunction() {
	s.Expect(s.Eval("typeof window.matchMedia")).To(Equal("function"))
}

func (s *MatchMediaSuite) TestReturnsMediaQueryListLike() {
	s.Expect(s.Eval(`window.matchMedia("(prefers-color-scheme: dark)").matches`)).To(BeFalse())
}

func (s *MatchMediaSuite) TestExposesQueryString() {
	s.Expect(s.Eval(`window.matchMedia("(min-width: 600px)").media`)).To(Equal("(min-width: 600px)"))
}

func (s *MatchMediaSuite) TestListenerMethodsAreCallable() {
	s.Expect(s.Eval(`
		const l = window.matchMedia("(min-width: 0px)");
		const cb = () => {};
		l.addListener(cb);
		l.removeListener(cb);
		l.addEventListener("change", cb);
		l.removeEventListener("change", cb);
		true
	`)).To(BeTrue())
}
