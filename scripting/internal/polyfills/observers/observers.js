// IntersectionObserver / ResizeObserver / PerformanceObserver
// Constructible no-ops. observe/unobserve/disconnect are inert; callbacks
// never fire (the in-process browser has no viewport, layout, or perf
// timeline). Library code that feature-tests these constructors will take
// the modern path instead of a broken fallback.

(function () {
  "use strict";
  function makeObserverStub(name) {
    function Stub(callback, options) {
      this._cb = callback;
      this._opts = options;
    }
    Stub.prototype.observe = function () {};
    Stub.prototype.unobserve = function () {};
    Stub.prototype.disconnect = function () {};
    Stub.prototype.takeRecords = function () {
      return [];
    };
    Object.defineProperty(Stub, "name", { value: name });
    return Stub;
  }
  if (typeof window.IntersectionObserver === "undefined") {
    window.IntersectionObserver = makeObserverStub("IntersectionObserver");
  }
  if (typeof window.ResizeObserver === "undefined") {
    window.ResizeObserver = makeObserverStub("ResizeObserver");
  }
  if (typeof window.PerformanceObserver === "undefined") {
    window.PerformanceObserver = makeObserverStub("PerformanceObserver");
  }
})();
