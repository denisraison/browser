package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ReadyStateSuite struct {
	ScriptHostSuite
}

func NewReadyStateSuite(h html.ScriptEngine) *ReadyStateSuite {
	return &ReadyStateSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ReadyStateSuite) TestReadyStateDefinedAfterParse() {
	// After parseReader returns, DOMContentLoaded + load have both been
	// dispatched, so readyState should be "complete".
	s.Expect(s.Eval(`document.readyState`)).To(Equal("complete"))
}

func (s *ReadyStateSuite) TestReadyStateTypeofString() {
	s.Expect(s.Eval(`typeof document.readyState`)).To(Equal("string"))
}
