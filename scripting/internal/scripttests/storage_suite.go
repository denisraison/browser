package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type StorageSuite struct {
	ScriptHostSuite
}

func NewStorageSuite(h html.ScriptEngine) *StorageSuite {
	return &StorageSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *StorageSuite) TestLocalStorageExists() {
	s.Expect(s.Eval("typeof localStorage")).To(Equal("object"))
}

func (s *StorageSuite) TestSessionStorageExists() {
	s.Expect(s.Eval("typeof sessionStorage")).To(Equal("object"))
}

func (s *StorageSuite) TestSetAndGetItem() {
	s.Expect(s.Eval(`
		localStorage.setItem("k", "v");
		localStorage.getItem("k")
	`)).To(Equal("v"))
}

func (s *StorageSuite) TestGetItemMissingReturnsNull() {
	s.Expect(s.Eval(`localStorage.getItem("no-such-key")`)).To(BeNil())
}

func (s *StorageSuite) TestRemoveItem() {
	s.Expect(s.Eval(`
		localStorage.setItem("x", "1");
		localStorage.removeItem("x");
		localStorage.getItem("x")
	`)).To(BeNil())
}

func (s *StorageSuite) TestLength() {
	s.Expect(s.Eval(`
		localStorage.clear();
		localStorage.setItem("a", "1");
		localStorage.setItem("b", "2");
		localStorage.length
	`)).To(BeEquivalentTo(2))
}

func (s *StorageSuite) TestClear() {
	s.Expect(s.Eval(`
		localStorage.setItem("a", "1");
		localStorage.clear();
		localStorage.length
	`)).To(BeEquivalentTo(0))
}

func (s *StorageSuite) TestKey() {
	s.Expect(s.Eval(`
		localStorage.clear();
		localStorage.setItem("only", "val");
		localStorage.key(0)
	`)).To(Equal("only"))
}

func (s *StorageSuite) TestKeyOutOfRange() {
	s.Expect(s.Eval(`
		localStorage.clear();
		localStorage.key(0)
	`)).To(BeNil())
}

func (s *StorageSuite) TestPropertyAccessSyntax() {
	s.Expect(s.Eval(`
		localStorage["myKey"] = "hello";
		localStorage["myKey"]
	`)).To(Equal("hello"))
}

func (s *StorageSuite) TestCoercesToString() {
	s.Expect(s.Eval(`
		localStorage.setItem("n", 42);
		typeof localStorage.getItem("n")
	`)).To(Equal("string"))
}

func (s *StorageSuite) TestLocalAndSessionAreIndependent() {
	s.Expect(s.Eval(`
		localStorage.clear();
		sessionStorage.clear();
		localStorage.setItem("x", "local");
		sessionStorage.getItem("x")
	`)).To(BeNil())
}
