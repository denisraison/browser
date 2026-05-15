package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type CustomElementsSuite struct {
	ScriptHostSuite
}

func NewCustomElementsSuite(h html.ScriptEngine) *CustomElementsSuite {
	return &CustomElementsSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *CustomElementsSuite) TestRegistryAccessible() {
	s.Expect(s.Eval("typeof window.customElements")).To(Equal("object"))
}

func (s *CustomElementsSuite) TestDefineRecordsName() {
	s.Expect(s.Eval(`
		class C {}
		customElements.define("x-test", C);
		customElements.get("x-test") === C
	`)).To(BeTrue())
}

func (s *CustomElementsSuite) TestWhenDefinedResolvesForKnown() {
	s.Assert().NoError(s.RunScript(`
		class C {}
		customElements.define("x-known", C);
		window.__ceResult = false;
		customElements.whenDefined("x-known").then(c => { window.__ceResult = c === C; });
	`))
	s.Expect(s.Eval(`window.__ceResult`)).To(BeTrue())
}

func (s *CustomElementsSuite) TestUpgradeIsCallable() {
	s.Expect(s.Eval(`
		customElements.upgrade(document.body);
		true
	`)).To(BeTrue())
}
