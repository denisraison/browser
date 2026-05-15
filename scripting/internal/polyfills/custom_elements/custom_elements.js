// window.customElements
// The native getter throws "Not implemented". Replace with a no-op registry
// so feature-tests pass and library boot doesn't blow up. define() records
// the name but never upgrades: libraries that only gate on the registry
// being present work; ones that actually need WC upgrade behavior do not.

(function () {
  "use strict";
  try {
    const registry = {
      _defs: new Map(),
      define(name, ctor) {
        this._defs.set(String(name), ctor);
      },
      get(name) {
        return this._defs.get(String(name));
      },
      whenDefined(name) {
        return this._defs.has(String(name))
          ? Promise.resolve(this._defs.get(String(name)))
          : new Promise(() => {});
      },
      upgrade() {},
      getName(ctor) {
        for (const [k, v] of this._defs) if (v === ctor) return k;
        return null;
      },
    };
    Object.defineProperty(window, "customElements", {
      configurable: true,
      get() {
        return registry;
      },
    });
  } catch (_) {}
})();
