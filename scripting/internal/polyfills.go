package internal

import (
	_ "embed"

	"github.com/gost-dom/browser/scripting/internal/js"
)

//go:embed polyfills/xpath/xpath.js
var xpath []byte

//go:embed polyfills/FastestSmallestTextEncoderDecoder/EncoderDecoderTogether.min.js
var textEncoderDecoder []byte

//go:embed polyfills/iterable/iterable.js
var iterable []byte

//go:embed polyfills/storage/storage.js
var webStorage []byte

//go:embed polyfills/css/css_style_sheet.js
var cssStyleSheet []byte

func InstallPolyfills[T any](host js.ScriptEngine[T]) {
	host.InstallPolyfill(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(([k,v]) => { cb(v,k) })
		}
	`, "gost-dom/polyfills/formdata.js")
	host.InstallPolyfill(string(iterable), "gost-dom/polyfills/iterable.js")
	host.InstallPolyfill(`
		Node.ELEMENT_NODE = 1;
		Node.ATTRIBUTE_NODE = 2;
		Node.TEXT_NODE = 3;
		Node.CDATA_SECTION_NODE = 4;
		Node.ENTITY_REFERENCE_NODE = 5;
		Node.ENTITY_NODE = 6;
		Node.PROCESSING_INSTRUCTION_NODE = 7;
		Node.COMMENT_NODE = 8;
		Node.DOCUMENT_NODE = 9;
		Node.DOCUMENT_TYPE_NODE = 10;
		Node.DOCUMENT_FRAGMENT_NODE = 11;
		Node.NOTATION_NODE = 12;
		Node.DOCUMENT_POSITION_DISCONNECTED = 0x01;
		Node.DOCUMENT_POSITION_PRECEDING = 0x02;
		Node.DOCUMENT_POSITION_FOLLOWING = 0x04;
		Node.DOCUMENT_POSITION_CONTAINS = 0x08;
		Node.DOCUMENT_POSITION_CONTAINED_BY = 0x10;
		Node.DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20;
	`, "gost-dom/polyfills/node.js")

	host.InstallPolyfill(string(xpath), "gost-dom/polyfills/xpath-jsdom.js")
	host.InstallPolyfill(string(textEncoderDecoder), "gost-dom/polyfills/text-encoder-decoder.js")
	host.InstallPolyfill(`
			const { XPathExpression, XPathResult } = window;
			const evaluate = XPathExpression.prototype.evaluate;
			XPathExpression.prototype.evaluate = function (context, type, res) {
				return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
			};
			Element.prototype.scrollIntoView = function() {};

	`, "gost-dom/polyfills/xpath-custom.js")
	host.InstallPolyfill(`
		Object.setPrototypeOf(DOMException, Error)
		Object.setPrototypeOf(DOMException.prototype, Error.prototype)
	`, "gost-dom/polyfills/errors.js")
	host.InstallPolyfill(string(webStorage), "gost-dom/polyfills/storage.js")
	host.InstallPolyfill(string(cssStyleSheet), "gost-dom/polyfills/css-style-sheet.js")
	host.InstallPolyfill(`
		if (typeof globalThis.CSS === "undefined") {
			globalThis.CSS = {
				escape(value) {
					return String(value).replace(/[!"#$%&'()*+,.\/:;<=>?@\[\\\]^\x60{|}~]/g, "\\$&");
				},
				supports() { return false; },
			};
		}
	`, "gost-dom/polyfills/css-namespace.js")
}
