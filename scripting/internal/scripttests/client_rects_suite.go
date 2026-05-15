package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ClientRectsSuite struct {
	ScriptHostSuite
}

func NewClientRectsSuite(h html.ScriptEngine) *ClientRectsSuite {
	return &ClientRectsSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ClientRectsSuite) TestGetBoundingClientRectExists() {
	s.Expect(s.Eval("typeof document.body.getBoundingClientRect")).To(Equal("function"))
}

func (s *ClientRectsSuite) TestReturnsZeroRect() {
	s.Expect(s.Eval(`
		const r = document.body.getBoundingClientRect();
		r.x === 0 && r.y === 0 && r.width === 0 && r.height === 0 &&
		r.top === 0 && r.right === 0 && r.bottom === 0 && r.left === 0
	`)).To(BeTrue())
}

func (s *ClientRectsSuite) TestRectIsJSONSerializable() {
	s.Expect(s.Eval(`JSON.stringify(document.body.getBoundingClientRect()).includes('"width":0')`)).To(BeTrue())
}

func (s *ClientRectsSuite) TestGetClientRectsReturnsEmpty() {
	s.Expect(s.Eval(`document.body.getClientRects().length`)).To(BeEquivalentTo(0))
}
