let i;
function y(e, s) {
  i && i.set(e, s) || s();
}
function m(e) {
  if (i)
    return e();
  i = /* @__PURE__ */ new Map();
  try {
    e();
  } finally {
    i.forEach((s) => s()), i = void 0;
  }
}
// @__NO_SIDE_EFFECTS__
function S(e, s) {
  s || (s = {});
  const d = s.equals || Object.is;
  let n = [], c = e;
  const a = (t) => {
    if (!d(c, t)) {
      const u = c;
      c = t, y(l, () => {
        [...n].forEach(([f, g]) => {
          f(t, u), g && r(f, !0);
        });
      });
    }
  };
  function o(t) {
    const u = typeof t != "object" ? { once: t } : t;
    return {
      once: u.once || !1,
      signal: u.signal || !1
    };
  }
  const r = (t, u) => {
    const f = o(u), g = n.findIndex(([h, p]) => h === t && p.once === f.once && p.signal === f.signal);
    g >= 0 && n.splice(g, 1);
  }, l = Object.assign(
    function() {
      return j(l), c;
    },
    {
      destroy() {
        n = [];
      },
      set: a,
      reset() {
        a(e);
      },
      sub(t, u) {
        return n.push([t, o(u)]), () => r(t, u);
      },
      unsub: r,
      unsubAll() {
        n = n.filter((t) => t[1].signal);
      }
    }
  );
  return l;
}
const b = [];
function j(e) {
  b.length && b[b.length - 1].add(e);
}
// @__NO_SIDE_EFFECTS__
function x(e, s) {
  let d = /* @__PURE__ */ new Set();
  const n = /* @__PURE__ */ S(a(), s);
  function c() {
    n.set(a());
  }
  function a() {
    d.forEach((l) => l.unsub(c, { signal: !0 }));
    const o = /* @__PURE__ */ new Set();
    let r;
    b.push(o);
    try {
      r = e();
    } finally {
      b.pop();
    }
    return o.forEach((l) => {
      l.sub(c, { signal: !0 });
    }), d = o, r;
  }
  return Object.assign(function() {
    return n();
  }, {
    destroy: n.destroy,
    sub: n.sub,
    unsub: n.unsub,
    unsubAll: n.unsubAll
  });
}
export {
  m as batch,
  x as computed,
  S as signal
};
//# sourceMappingURL=index.js.map
