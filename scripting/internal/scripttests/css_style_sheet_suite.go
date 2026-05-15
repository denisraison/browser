package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type CSSStyleSheetSuite struct {
	ScriptHostSuite
}

func NewCSSStyleSheetSuite(h html.ScriptEngine) *CSSStyleSheetSuite {
	return &CSSStyleSheetSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *CSSStyleSheetSuite) TestGlobalExists() {
	s.Expect(s.Eval("typeof CSSStyleSheet")).To(Equal("function"))
}

func (s *CSSStyleSheetSuite) TestConstruct() {
	s.Expect(s.Eval("new CSSStyleSheet() instanceof CSSStyleSheet")).To(BeTrue())
}

func (s *CSSStyleSheetSuite) TestReplaceSync() {
	s.MustRunScript(`
		const sheet = new CSSStyleSheet();
		sheet.replaceSync("body { color: red; }");
	`)
}

func (s *CSSStyleSheetSuite) TestReplace() {
	s.Expect(s.Eval(`
		const sheet = new CSSStyleSheet();
		sheet.replace("body { color: red; }") instanceof Promise
	`)).To(BeTrue())
}

func (s *CSSStyleSheetSuite) TestInsertRule() {
	s.Expect(s.Eval(`
		const sheet = new CSSStyleSheet();
		sheet.insertRule("body { color: red; }", 0);
		sheet.cssRules.length
	`)).To(BeEquivalentTo(1))
}

func (s *CSSStyleSheetSuite) TestDeleteRule() {
	s.Expect(s.Eval(`
		const sheet = new CSSStyleSheet();
		sheet.insertRule("body { color: red; }", 0);
		sheet.deleteRule(0);
		sheet.cssRules.length
	`)).To(BeEquivalentTo(0))
}

func (s *CSSStyleSheetSuite) TestAdoptedStyleSheets() {
	s.Expect(s.Eval("Array.isArray(document.adoptedStyleSheets)")).To(BeTrue())
}

func (s *CSSStyleSheetSuite) TestAdoptedStyleSheetsAssignable() {
	s.Expect(s.Eval(`
		const sheet = new CSSStyleSheet();
		document.adoptedStyleSheets = [sheet];
		document.adoptedStyleSheets.length
	`)).To(BeEquivalentTo(1))
}
