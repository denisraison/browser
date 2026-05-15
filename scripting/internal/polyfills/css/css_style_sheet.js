// Minimal stub for the constructable CSSStyleSheet API. Modern client-side
// frameworks (e.g., htmx 4) reference the constructor for feature detection
// or to build adopted stylesheets. The DOM in gost-dom does not apply any
// CSS, so rule storage is opaque, but enough of the surface exists for code
// that constructs sheets to run without throwing.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleSheet
(function () {
	class CSSStyleSheet {
		constructor(options) {
			this._rules = [];
			this.disabled = false;
			this.media = options && options.media ? String(options.media) : "";
			this.title = options && options.title ? String(options.title) : "";
		}

		get cssRules() { return this._rules; }

		replaceSync(_text) {
			this._rules = [];
		}

		replace(_text) {
			this._rules = [];
			return Promise.resolve(this);
		}

		insertRule(rule, index) {
			const at = index === undefined ? this._rules.length : index;
			this._rules.splice(at, 0, { cssText: rule });
			return at;
		}

		deleteRule(index) {
			this._rules.splice(index, 1);
		}
	}

	globalThis.CSSStyleSheet = CSSStyleSheet;

	if (!("adoptedStyleSheets" in Document.prototype)) {
		Object.defineProperty(Document.prototype, "adoptedStyleSheets", {
			get() {
				if (!this._adoptedStyleSheets) this._adoptedStyleSheets = [];
				return this._adoptedStyleSheets;
			},
			set(v) { this._adoptedStyleSheets = v; },
			configurable: true,
		});
	}
})();
