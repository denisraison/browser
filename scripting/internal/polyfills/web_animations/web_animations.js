// Element.prototype.animate (Web Animations API)
// Returns a finished Animation object immediately. finished/ready resolve
// right away so awaiters proceed; play/pause/cancel/reverse are no-ops.

(function () {
  "use strict";
  if (typeof Element.prototype.animate === "function") return;
  Element.prototype.animate = function () {
    const anim = {
      play() {},
      pause() {},
      cancel() {},
      finish() {},
      reverse() {},
      addEventListener() {},
      removeEventListener() {},
      playState: "finished",
      currentTime: 0,
      startTime: 0,
      playbackRate: 1,
    };
    const done = Promise.resolve(anim);
    Object.defineProperty(anim, "finished", { get: () => done });
    Object.defineProperty(anim, "ready", { get: () => done });
    return anim;
  };
})();
