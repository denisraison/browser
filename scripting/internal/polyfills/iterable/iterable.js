// Web IDL value iterables (https://webidl.spec.whatwg.org/#es-iterable) must
// expose forEach/keys/values on their prototype. The Go-side InstallIterator
// only wires Symbol.iterator and entries, so patch the prototypes here for any
// iterable types that script code is expected to iterate.
(function () {
	const iterables = [
		typeof NodeList !== "undefined" ? NodeList : null,
		typeof HTMLCollection !== "undefined" ? HTMLCollection : null,
	];

	for (const C of iterables) {
		if (!C || !C.prototype) continue;

		if (!C.prototype.forEach) {
			C.prototype.forEach = function (cb, thisArg) {
				for (let i = 0; i < this.length; i++) {
					const value = this[i] !== undefined ? this[i] : this.item(i);
					cb.call(thisArg, value, i, this);
				}
			};
		}

		if (!C.prototype.keys) {
			C.prototype.keys = function* () {
				for (let i = 0; i < this.length; i++) yield i;
			};
		}

		if (!C.prototype.values) {
			C.prototype.values = function* () {
				for (let i = 0; i < this.length; i++) {
					yield this[i] !== undefined ? this[i] : this.item(i);
				}
			};
		}
	}
})();
