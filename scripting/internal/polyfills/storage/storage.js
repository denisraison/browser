// Minimal in-memory implementation of the Web Storage API. Each window has its
// own localStorage and sessionStorage instance because polyfills run in every
// script context. Storage is not persisted across windows or test runs.
//
// See also: https://html.spec.whatwg.org/multipage/webstorage.html
(function () {
	const reserved = new Set([
		"length", "key", "getItem", "setItem", "removeItem", "clear",
	]);

	function createStorage() {
		const data = new Map();

		const api = {
			get length() { return data.size; },
			key(n) {
				let i = 0;
				for (const k of data.keys()) {
					if (i++ === n) return k;
				}
				return null;
			},
			getItem(k) {
				const s = String(k);
				return data.has(s) ? data.get(s) : null;
			},
			setItem(k, v) { data.set(String(k), String(v)); },
			removeItem(k) { data.delete(String(k)); },
			clear() { data.clear(); },
		};

		return new Proxy(api, {
			get(target, prop) {
				if (reserved.has(prop)) {
					const v = target[prop];
					return typeof v === "function" ? v.bind(target) : v;
				}
				const key = String(prop);
				return data.has(key) ? data.get(key) : undefined;
			},
			set(target, prop, value) {
				if (reserved.has(prop)) return false;
				data.set(String(prop), String(value));
				return true;
			},
			has(target, prop) {
				return reserved.has(prop) || data.has(String(prop));
			},
			deleteProperty(target, prop) {
				if (reserved.has(prop)) return false;
				return data.delete(String(prop));
			},
		});
	}

	Object.defineProperty(window, "localStorage", {
		value: createStorage(),
		configurable: true,
		enumerable: true,
	});
	Object.defineProperty(window, "sessionStorage", {
		value: createStorage(),
		configurable: true,
		enumerable: true,
	});
})();
