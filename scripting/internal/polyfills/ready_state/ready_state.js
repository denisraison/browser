// document.readyState
//
// Mirrors the spec lifecycle: "loading" during initial script execution,
// "interactive" after DOMContentLoaded, "complete" after load. Without this,
// the property is undefined and libraries that branch on it (htmx 4, jQuery's
// ready, et al.) take the wrong path during bootstrap.
//
// State is flipped by intercepting EventTarget.prototype.dispatchEvent rather
// than addEventListener, because the document instance is recreated during
// parseReader after polyfills run, so a listener attached at polyfill time
// would target a stale document.

(function () {
  "use strict";
  if (typeof Document === "undefined") return;
  if (Object.getOwnPropertyDescriptor(Document.prototype, "readyState")) return;

  let state = "loading";
  Object.defineProperty(Document.prototype, "readyState", {
    configurable: true,
    get() { return state; },
  });

  if (typeof EventTarget !== "undefined" && EventTarget.prototype.dispatchEvent) {
    const orig = EventTarget.prototype.dispatchEvent;
    EventTarget.prototype.dispatchEvent = function (ev) {
      const t = ev && ev.type;
      if (t === "DOMContentLoaded" && state === "loading") state = "interactive";
      else if (t === "load" && state !== "complete") state = "complete";
      return orig.call(this, ev);
    };
  }
})();
