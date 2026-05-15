// window.getComputedStyle
// Returns a CSSStyleDeclaration-like proxy that yields "" for every property
// (the in-process browser has no layout/style engine). Enough surface for
// libraries like floating-ui to read properties without throwing.

(function () {
  "use strict";
  if (typeof window.getComputedStyle === "function") return;
  const blank = new Proxy(
    {
      getPropertyValue() {
        return "";
      },
      getPropertyPriority() {
        return "";
      },
      setProperty() {},
      removeProperty() {
        return "";
      },
      item() {
        return "";
      },
      length: 0,
    },
    {
      get(target, prop) {
        if (prop in target) return target[prop];
        if (typeof prop === "symbol") return undefined;
        return "";
      },
    },
  );
  window.getComputedStyle = function () {
    return blank;
  };
})();
