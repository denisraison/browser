// window.matchMedia
// Always reports no match. Listeners are accepted but never fired since the
// in-process browser has no viewport, color scheme, or reduced-motion signal.

(function () {
  "use strict";
  if (typeof window.matchMedia === "function") return;
  window.matchMedia = function (query) {
    return {
      media: String(query),
      matches: false,
      onchange: null,
      addListener() {},
      removeListener() {},
      addEventListener() {},
      removeEventListener() {},
      dispatchEvent() {
        return false;
      },
    };
  };
})();
