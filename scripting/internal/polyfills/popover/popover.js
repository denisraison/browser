// HTML Popover API: showPopover / hidePopover / togglePopover, reflected
// `popover` IDL attribute, and a minimal ToggleEvent constructor. Open
// state is tracked via the `data-popover-open` attribute since the
// in-process browser has no top-layer or style engine. Toggle events are
// dispatched so listener-based libraries (templui popover.js, etc) react.

(function () {
  "use strict";

  if (!("popover" in HTMLElement.prototype)) {
    Object.defineProperty(HTMLElement.prototype, "popover", {
      configurable: true,
      enumerable: true,
      get() {
        const v = this.getAttribute("popover");
        if (v === null) return null;
        const lc = v.toLowerCase();
        if (lc === "" || lc === "auto") return "auto";
        if (lc === "manual") return "manual";
        if (lc === "hint") return "hint";
        return "auto";
      },
      set(v) {
        if (v === null) this.removeAttribute("popover");
        else this.setAttribute("popover", String(v));
      },
    });
  }

  function fireToggle(el, oldState, newState) {
    let evt;
    try {
      evt = new Event("toggle", { bubbles: false, cancelable: false });
    } catch (_) {
      evt = { type: "toggle" };
    }
    evt.oldState = oldState;
    evt.newState = newState;
    if (typeof el.dispatchEvent === "function") {
      try {
        el.dispatchEvent(evt);
      } catch (_) {}
    }
  }

  if (typeof HTMLElement.prototype.showPopover !== "function") {
    HTMLElement.prototype.showPopover = function () {
      if (this.getAttribute("popover") === null) return;
      if (this.hasAttribute("data-popover-open")) return;
      this.setAttribute("data-popover-open", "");
      fireToggle(this, "closed", "open");
    };
    HTMLElement.prototype.hidePopover = function () {
      if (!this.hasAttribute("data-popover-open")) return;
      this.removeAttribute("data-popover-open");
      fireToggle(this, "open", "closed");
    };
    HTMLElement.prototype.togglePopover = function (force) {
      const open = this.hasAttribute("data-popover-open");
      const want = typeof force === "boolean" ? force : !open;
      if (want && !open) this.showPopover();
      else if (!want && open) this.hidePopover();
      return this.hasAttribute("data-popover-open");
    };
  }

  if (typeof window.ToggleEvent === "undefined") {
    window.ToggleEvent = function ToggleEvent(type, init) {
      const e = new Event(type, init || {});
      e.oldState = (init && init.oldState) || "";
      e.newState = (init && init.newState) || "";
      return e;
    };
  }
})();
