// HTMLDialogElement.show / showModal / close
// Reflect open state via the `open` attribute; close dispatches a `close`
// event. No focus management or top-layer behavior.

(function () {
  "use strict";
  if (typeof HTMLDialogElement === "undefined") return;
  const dp = HTMLDialogElement.prototype;
  if (typeof dp.show !== "function") {
    dp.show = function () {
      this.setAttribute("open", "");
    };
  }
  if (typeof dp.showModal !== "function") {
    dp.showModal = function () {
      this.setAttribute("open", "");
    };
  }
  if (typeof dp.close !== "function") {
    dp.close = function (returnValue) {
      this.removeAttribute("open");
      if (typeof returnValue === "string") this.returnValue = returnValue;
      try {
        this.dispatchEvent(new Event("close"));
      } catch (_) {}
    };
  }
})();
