// crypto.randomUUID + crypto.getRandomValues
// Math.random()-backed: not cryptographically secure, but enough for client
// IDs / idempotency keys that UI libraries generate.

(function () {
  "use strict";
  if (typeof globalThis.crypto === "undefined") {
    globalThis.crypto = {};
  }
  if (typeof globalThis.crypto.getRandomValues !== "function") {
    globalThis.crypto.getRandomValues = function (buf) {
      for (let i = 0; i < buf.length; i++) {
        buf[i] = Math.floor(Math.random() * 256);
      }
      return buf;
    };
  }
  if (typeof globalThis.crypto.randomUUID !== "function") {
    globalThis.crypto.randomUUID = function () {
      const b = new Uint8Array(16);
      globalThis.crypto.getRandomValues(b);
      b[6] = (b[6] & 0x0f) | 0x40;
      b[8] = (b[8] & 0x3f) | 0x80;
      const h = Array.from(b, (x) => x.toString(16).padStart(2, "0"));
      return (
        h.slice(0, 4).join("") +
        "-" +
        h.slice(4, 6).join("") +
        "-" +
        h.slice(6, 8).join("") +
        "-" +
        h.slice(8, 10).join("") +
        "-" +
        h.slice(10, 16).join("")
      );
    };
  }
})();
