package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ComputedStyleSuite struct {
	ScriptHostSuite
}

func NewComputedStyleSuite(h html.ScriptEngine) *ComputedStyleSuite {
	return &ComputedStyleSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ComputedStyleSuite) TestExists() {
	s.Expect(s.Eval("typeof window.getComputedStyle")).To(Equal("function"))
}

func (s *ComputedStyleSuite) TestReturnsObject() {
	s.Expect(s.Eval(`typeof window.getComputedStyle(document.body)`)).To(Equal("object"))
}

func (s *ComputedStyleSuite) TestPropertyAccessYieldsEmptyString() {
	s.Expect(s.Eval(`window.getComputedStyle(document.body).display`)).To(Equal(""))
}

func (s *ComputedStyleSuite) TestGetPropertyValueYieldsEmptyString() {
	s.Expect(s.Eval(`window.getComputedStyle(document.body).getPropertyValue("color")`)).To(Equal(""))
}
