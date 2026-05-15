// document.activeElement + element.focus / blur
// Tracks the most recent focus() call and exposes it via activeElement so
// keyboard-trap and autofocus-aware code sees a consistent answer instead
// of undefined. blur() no-ops.

(function () {
  "use strict";
  let active = null;
  try {
    const nativeFocus = HTMLElement.prototype.focus;
    HTMLElement.prototype.focus = function () {
      active = this;
      if (typeof nativeFocus === "function") {
        try {
          nativeFocus.call(this);
        } catch (_) {}
      }
    };
  } catch (_) {}
  if (typeof HTMLElement.prototype.blur !== "function") {
    HTMLElement.prototype.blur = function () {};
  }
  try {
    const desc = Object.getOwnPropertyDescriptor(Document.prototype, "activeElement");
    if (!desc || desc.get === undefined) {
      Object.defineProperty(Document.prototype, "activeElement", {
        configurable: true,
        get() {
          return active || this.body || null;
        },
      });
    }
  } catch (_) {}
})();
