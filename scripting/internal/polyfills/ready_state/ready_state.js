// document.readyState
//
// Without this, the property is undefined and libraries that branch on it
// (htmx 4, jQuery's ready, et al.) take the wrong path during bootstrap.
//
// Implementation derives state from whether <body> has been parsed. During
// initial script execution in <head>, document.body is null so we report
// "loading"; once parsing reaches <body>, scripts past that point and code
// running after DOMContentLoaded see "complete". This sidesteps the fact
// that the document instance is replaced inside parseReader after polyfills
// have already installed, which makes a listener-driven state machine
// unreliable.

(function () {
  "use strict";
  if (typeof Document === "undefined") return;
  if (Object.getOwnPropertyDescriptor(Document.prototype, "readyState")) return;
  Object.defineProperty(Document.prototype, "readyState", {
    configurable: true,
    get() { return this.body ? "complete" : "loading"; },
  });
})();
