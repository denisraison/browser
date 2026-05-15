// document.startViewTransition
// Runs the update callback synchronously and resolves immediately. There's
// no transition; this only keeps callers that gate on the API existing from
// taking a degraded "no transition" branch that disables their swap entirely.

(function () {
  "use strict";
  if (typeof Document === "undefined") return;
  if (typeof Document.prototype.startViewTransition === "function") return;
  Document.prototype.startViewTransition = function (cb) {
    const done = Promise.resolve();
    try {
      if (typeof cb === "function") cb();
    } catch (e) {
      return {
        ready: Promise.reject(e),
        finished: Promise.reject(e),
        updateCallbackDone: Promise.reject(e),
        skipTransition() {},
      };
    }
    return {
      ready: done,
      finished: done,
      updateCallbackDone: done,
      skipTransition() {},
    };
  };
})();
