// Element.prototype.getBoundingClientRect + getClientRects
// Returns an all-zero DOMRect; the in-process browser has no layout. Range
// gets the same treatment so selection-positioning code doesn't throw.

(function () {
  "use strict";
  function zeroRect() {
    return {
      x: 0,
      y: 0,
      width: 0,
      height: 0,
      top: 0,
      right: 0,
      bottom: 0,
      left: 0,
      toJSON() {
        return this;
      },
    };
  }
  if (typeof Element.prototype.getBoundingClientRect !== "function") {
    Element.prototype.getBoundingClientRect = function () {
      return zeroRect();
    };
  }
  if (typeof Element.prototype.getClientRects !== "function") {
    Element.prototype.getClientRects = function () {
      return [];
    };
  }
  if (typeof Range !== "undefined" && Range.prototype && typeof Range.prototype.getBoundingClientRect !== "function") {
    Range.prototype.getBoundingClientRect = function () {
      return zeroRect();
    };
    Range.prototype.getClientRects = function () {
      return [];
    };
  }
})();
