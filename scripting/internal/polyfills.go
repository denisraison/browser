package internal

import (
	_ "embed"

	"github.com/gost-dom/browser/scripting/internal/js"
)

//go:embed polyfills/xpath/xpath.js
var xpath []byte

//go:embed polyfills/FastestSmallestTextEncoderDecoder/EncoderDecoderTogether.min.js
var textEncoderDecoder []byte

//go:embed polyfills/crypto/crypto.js
var webCrypto []byte

//go:embed polyfills/match_media/match_media.js
var matchMedia []byte

//go:embed polyfills/computed_style/computed_style.js
var computedStyle []byte

//go:embed polyfills/client_rects/client_rects.js
var clientRects []byte

//go:embed polyfills/observers/observers.js
var observers []byte

//go:embed polyfills/web_animations/web_animations.js
var webAnimations []byte

//go:embed polyfills/popover/popover.js
var popover []byte

//go:embed polyfills/dialog/dialog.js
var dialog []byte

//go:embed polyfills/view_transitions/view_transitions.js
var viewTransitions []byte

//go:embed polyfills/custom_elements/custom_elements.js
var customElements []byte

//go:embed polyfills/focus/focus.js
var focus []byte

func InstallPolyfills[T any](host js.ScriptEngine[T]) {
	host.InstallPolyfill(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(([k,v]) => { cb(v,k) })
		}
	`, "gost-dom/polyfills/formdata.js")
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
	host.InstallPolyfill(string(webCrypto), "gost-dom/polyfills/crypto.js")
	host.InstallPolyfill(string(matchMedia), "gost-dom/polyfills/match-media.js")
	host.InstallPolyfill(string(computedStyle), "gost-dom/polyfills/computed-style.js")
	host.InstallPolyfill(string(clientRects), "gost-dom/polyfills/client-rects.js")
	host.InstallPolyfill(string(observers), "gost-dom/polyfills/observers.js")
	host.InstallPolyfill(string(webAnimations), "gost-dom/polyfills/web-animations.js")
	host.InstallPolyfill(string(popover), "gost-dom/polyfills/popover.js")
	host.InstallPolyfill(string(dialog), "gost-dom/polyfills/dialog.js")
	host.InstallPolyfill(string(viewTransitions), "gost-dom/polyfills/view-transitions.js")
	host.InstallPolyfill(string(customElements), "gost-dom/polyfills/custom-elements.js")
	host.InstallPolyfill(string(focus), "gost-dom/polyfills/focus.js")
}
