package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type ObserversSuite struct {
	ScriptHostSuite
}

func NewObserversSuite(h html.ScriptEngine) *ObserversSuite {
	return &ObserversSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *ObserversSuite) TestIntersectionObserverConstructs() {
	s.Expect(s.Eval(`new IntersectionObserver(() => {}) instanceof IntersectionObserver`)).To(BeTrue())
}

func (s *ObserversSuite) TestResizeObserverConstructs() {
	s.Expect(s.Eval(`new ResizeObserver(() => {}) instanceof ResizeObserver`)).To(BeTrue())
}

func (s *ObserversSuite) TestPerformanceObserverConstructs() {
	s.Expect(s.Eval(`new PerformanceObserver(() => {}) instanceof PerformanceObserver`)).To(BeTrue())
}

func (s *ObserversSuite) TestObserveDisconnectAreNoOps() {
	s.Expect(s.Eval(`
		const io = new IntersectionObserver(() => {});
		io.observe(document.body);
		io.unobserve(document.body);
		io.disconnect();
		io.takeRecords().length
	`)).To(BeEquivalentTo(0))
}

func (s *ObserversSuite) TestCallbacksDoNotFire() {
	s.Expect(s.Eval(`
		let fired = false;
		const io = new IntersectionObserver(() => { fired = true; });
		io.observe(document.body);
		fired
	`)).To(BeFalse())
}
