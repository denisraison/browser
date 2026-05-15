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

func (s *ReadyStateSuite) TestReadyStateCompleteAfterParse() {
	// After parseReader returns, document.body exists, so readyState reads
	// as "complete".
	s.Expect(s.Eval(`document.readyState`)).To(Equal("complete"))
}

func (s *ReadyStateSuite) TestReadyStateTypeofString() {
	s.Expect(s.Eval(`typeof document.readyState`)).To(Equal("string"))
}
