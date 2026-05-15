package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type PopoverSuite struct {
	ScriptHostSuite
}

func NewPopoverSuite(h html.ScriptEngine) *PopoverSuite {
	return &PopoverSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *PopoverSuite) TestShowPopoverExists() {
	s.Expect(s.Eval("typeof HTMLElement.prototype.showPopover")).To(Equal("function"))
}

func (s *PopoverSuite) TestPopoverAttrReflects() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "auto");
		d.popover
	`)).To(Equal("auto"))
}

func (s *PopoverSuite) TestPopoverAttrCoercesEmptyToAuto() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "");
		d.popover
	`)).To(Equal("auto"))
}

func (s *PopoverSuite) TestPopoverAttrCoercesManual() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "manual");
		d.popover
	`)).To(Equal("manual"))
}

func (s *PopoverSuite) TestShowAndHidePopoverTracksState() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "manual");
		document.body.appendChild(d);
		d.showPopover();
		const open = d.hasAttribute("data-popover-open");
		d.hidePopover();
		const closed = !d.hasAttribute("data-popover-open");
		open && closed
	`)).To(BeTrue())
}

func (s *PopoverSuite) TestShowPopoverWithoutAttrDoesNothing() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.showPopover();
		d.hasAttribute("data-popover-open")
	`)).To(BeFalse())
}

func (s *PopoverSuite) TestTogglePopover() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "manual");
		document.body.appendChild(d);
		d.togglePopover();
		const a = d.hasAttribute("data-popover-open");
		d.togglePopover();
		const b = !d.hasAttribute("data-popover-open");
		a && b
	`)).To(BeTrue())
}

func (s *PopoverSuite) TestToggleEventDispatched() {
	s.Expect(s.Eval(`
		const d = document.createElement("div");
		d.setAttribute("popover", "manual");
		document.body.appendChild(d);
		let states = [];
		d.addEventListener("toggle", (e) => states.push(e.newState));
		d.showPopover();
		d.hidePopover();
		states.join(",")
	`)).To(Equal("open,closed"))
}
