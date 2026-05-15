package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type CryptoSuite struct {
	ScriptHostSuite
}

func NewCryptoSuite(h html.ScriptEngine) *CryptoSuite {
	return &CryptoSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *CryptoSuite) TestCryptoExists() {
	s.Expect(s.Eval("typeof crypto")).To(Equal("object"))
}

func (s *CryptoSuite) TestRandomUUIDFormat() {
	s.Expect(s.Eval(`/^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/.test(crypto.randomUUID())`)).To(BeTrue())
}

func (s *CryptoSuite) TestRandomUUIDIsUnique() {
	s.Expect(s.Eval(`crypto.randomUUID() !== crypto.randomUUID()`)).To(BeTrue())
}

func (s *CryptoSuite) TestGetRandomValuesFillsBuffer() {
	s.Expect(s.Eval(`
		const b = new Uint8Array(8);
		crypto.getRandomValues(b);
		b.length
	`)).To(BeEquivalentTo(8))
}

func (s *CryptoSuite) TestGetRandomValuesReturnsSameBuffer() {
	s.Expect(s.Eval(`
		const b = new Uint8Array(4);
		crypto.getRandomValues(b) === b
	`)).To(BeTrue())
}
