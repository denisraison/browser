// document.readyState
//
// Mirrors the spec lifecycle: "loading" during initial script execution,
// "interactive" after DOMContentLoaded, "complete" after load. Without this,
// the property is undefined and libraries that branch on it (htmx 4, jQuery's
// ready, et al.) take the wrong path during bootstrap.

(function () {
  "use strict";
  if (typeof Document === "undefined") return;
  if (Object.getOwnPropertyDescriptor(Document.prototype, "readyState")) return;

  let state = "loading";
  Object.defineProperty(Document.prototype, "readyState", {
    configurable: true,
    get() { return state; },
  });

  if (typeof document !== "undefined" && typeof document.addEventListener === "function") {
    document.addEventListener("DOMContentLoaded", () => { state = "interactive"; });
  }
  if (typeof window !== "undefined" && typeof window.addEventListener === "function") {
    window.addEventListener("load", () => { state = "complete"; });
  }
})();
