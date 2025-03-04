let pt = class He extends Error {
  constructor(e, n, o) {
    super(
      typeof n == "object" ? n.message : n || e,
      {
        cause: typeof n == "object" ? n.cause : o
      }
    ), this.type = e, Object.setPrototypeOf(this, He.prototype);
  }
};
function Ge(t) {
  return t.replace(/[A-Z]/g, (e) => `_${e.toLowerCase()}`);
}
function Fo(t) {
  return t.replace(/_[a-z]/g, (e) => e[1].toUpperCase());
}
const Jo = "ERR_INVALID_VALUE", Ko = "ERR_UNEXPECTED_VALUE", Qo = "ERR_UNEXPECTED_TYPE", ze = "ERR_PARSE";
function Ye(t, e) {
  const n = {};
  for (const o in t) {
    const s = t[o];
    if (!s)
      continue;
    let r, a;
    typeof s == "function" ? (r = o, a = s) : [r, a] = s;
    try {
      const i = a(e(r));
      i !== void 0 && (n[o] = i);
    } catch (i) {
      throw new pt(
        ze,
        `Parser for "${o}" property failed${r === o ? "" : `. Source field: "${r}"`}`,
        i
      );
    }
  }
  return n;
}
function Fe(t) {
  let e = t;
  if (typeof e == "string")
    try {
      e = JSON.parse(e);
    } catch (n) {
      throw new pt(Jo, { cause: n });
    }
  if (typeof e != "object" || !e || Array.isArray(e))
    throw new pt(Ko);
  return e;
}
function k(t, e) {
  return (n) => {
    const o = (s) => {
      if (!(n && s === void 0))
        try {
          return e(s);
        } catch (r) {
          throw new pt(ze, {
            message: `"${t}" transformer failed to parse the value`,
            cause: r
          });
        }
    };
    return /* @__PURE__ */ Object.assign(
      o,
      {
        isValid(s) {
          try {
            return o(s), !0;
          } catch {
            return !1;
          }
        }
      }
    );
  };
}
function xt(t, e) {
  return k(e || "object", (n) => {
    const o = Fe(n);
    return Ye(t, (s) => o[s]);
  });
}
function St(t) {
  throw new pt(Qo, `Unexpected value received: ${JSON.stringify(t)}`);
}
const Xo = k("boolean", (t) => {
  if (typeof t == "boolean")
    return t;
  const e = String(t);
  if (e === "1" || e === "true")
    return !0;
  if (e === "0" || e === "false")
    return !1;
  St(t);
}), D = k("string", (t) => {
  if (typeof t == "string" || typeof t == "number")
    return t.toString();
  St(t);
}), kt = k("number", (t) => {
  if (typeof t == "number")
    return t;
  if (typeof t == "string") {
    const e = Number(t);
    if (!Number.isNaN(e))
      return e;
  }
  St(t);
}), Je = k("date", (t) => t instanceof Date ? t : new Date(kt()(t) * 1e3));
function Ke(t, e) {
  return k(e || "searchParams", (n) => {
    typeof n != "string" && !(n instanceof URLSearchParams) && St(n);
    const o = typeof n == "string" ? new URLSearchParams(n) : n;
    return Ye(t, (s) => {
      const r = o.get(s);
      return r === null ? void 0 : r;
    });
  });
}
function Gt(t) {
  for (const e in t)
    t[e] = [Ge(e), t[e]];
  return t;
}
const Zo = (t) => {
  const e = kt(), n = kt(!0), o = D(), s = D(!0), r = Xo(!0), a = xt(Gt({
    addedToAttachmentMenu: r,
    allowsWriteToPm: r,
    firstName: o,
    id: e,
    isBot: r,
    isPremium: r,
    languageCode: s,
    lastName: s,
    photoUrl: s,
    username: s
  }), "User")(!0);
  return Ke(
    Gt({
      authDate: Je(),
      canSendAfter: n,
      chat: xt(
        Gt({
          id: e,
          type: o,
          title: o,
          photoUrl: s,
          username: s
        }),
        "Chat"
      )(!0),
      chatInstance: s,
      chatType: s,
      hash: o,
      queryId: s,
      receiver: a,
      startParam: s,
      user: a
    }),
    "initData"
  )(t);
};
function Ct(t) {
  return /^#[\da-f]{6}$/i.test(t);
}
function ts(t) {
  return /^#[\da-f]{3}$/i.test(t);
}
function Qe(t) {
  const e = t.replace(/\s/g, "").toLowerCase();
  if (Ct(e))
    return e;
  if (ts(e)) {
    let o = "#";
    for (let s = 0; s < 3; s += 1)
      o += e[1 + s].repeat(2);
    return o;
  }
  const n = e.match(/^rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)$/) || e.match(/^rgba\((\d{1,3}),(\d{1,3}),(\d{1,3}),\d{1,3}\)$/);
  if (!n)
    throw new Error(`Value "${t}" does not satisfy any of known RGB formats.`);
  return n.slice(1).reduce((o, s) => {
    const r = parseInt(s, 10).toString(16);
    return o + (r.length === 1 ? "0" : "") + r;
  }, "#");
}
const es = k("rgb", (t) => Qe(D()(t))), ns = k(
  "themeParams",
  (t) => {
    const e = es(!0);
    return Object.entries(Fe(t)).reduce((n, [o, s]) => (n[Fo(o)] = e(s), n), {});
  }
);
// @__NO_SIDE_EFFECTS__
function os(t) {
  return JSON.stringify(
    Object.fromEntries(
      Object.entries(t).map(([e, n]) => [Ge(e), n])
    )
  );
}
// @__NO_SIDE_EFFECTS__
function gi(t) {
  const { initDataRaw: e, startParam: n, showSettings: o, botInline: s } = t, r = new URLSearchParams();
  return r.set("tgWebAppPlatform", t.platform), r.set("tgWebAppThemeParams", /* @__PURE__ */ os(t.themeParams)), r.set("tgWebAppVersion", t.version), e && r.set("tgWebAppData", e), n && r.set("tgWebAppStartParam", n), typeof o == "boolean" && r.set("tgWebAppShowSettings", o ? "1" : "0"), typeof s == "boolean" && r.set("tgWebAppBotInline", s ? "1" : "0"), r.toString();
}
function ss(t, e) {
  return k("array", (n) => {
    let o;
    if (Array.isArray(n))
      o = n;
    else if (typeof n == "string")
      try {
        const s = JSON.parse(n);
        Array.isArray(s) && (o = s);
      } catch {
      }
    return o || St(n), o.map(t);
  });
}
function Xe(t) {
  return !!t && typeof t == "object" && !Array.isArray(t);
}
function Xt(...t) {
  return t.map((e) => {
    if (typeof e == "string")
      return e;
    if (Xe(e))
      return Xt(Object.entries(e).map((n) => n[1] && n[0]));
    if (Array.isArray(e))
      return Xt(...e);
  }).filter(Boolean).join(" ");
}
function wi(...t) {
  return t.reduce((e, n) => (Xe(n) && Object.entries(n).forEach(([o, s]) => {
    const r = Xt(e[o], s);
    r && (e[o] = r);
  }), e), {});
}
let _t = class Ze extends Error {
  constructor(e, n, o) {
    super(
      typeof n == "object" ? n.message : n || e,
      {
        cause: typeof n == "object" ? n.cause : o
      }
    ), this.type = e, Object.setPrototypeOf(this, Ze.prototype);
  }
};
function tn(t) {
  return t.replace(/[A-Z]/g, (e) => `_${e.toLowerCase()}`);
}
function rs(t) {
  return t.replace(/_[a-z]/g, (e) => e[1].toUpperCase());
}
const as = "ERR_INVALID_VALUE", is = "ERR_UNEXPECTED_VALUE", cs = "ERR_UNEXPECTED_TYPE", en = "ERR_PARSE";
function nn(t, e) {
  const n = {};
  for (const o in t) {
    const s = t[o];
    if (!s)
      continue;
    let r, a;
    typeof s == "function" ? (r = o, a = s) : [r, a] = s;
    try {
      const i = a(e(r));
      i !== void 0 && (n[o] = i);
    } catch (i) {
      throw new _t(
        en,
        `Parser for "${o}" property failed${r === o ? "" : `. Source field: "${r}"`}`,
        i
      );
    }
  }
  return n;
}
function on(t) {
  let e = t;
  if (typeof e == "string")
    try {
      e = JSON.parse(e);
    } catch (n) {
      throw new _t(as, { cause: n });
    }
  if (typeof e != "object" || !e || Array.isArray(e))
    throw new _t(is);
  return e;
}
function B(t, e) {
  return (n) => {
    const o = (s) => {
      if (!(n && s === void 0))
        try {
          return e(s);
        } catch (r) {
          throw new _t(en, {
            message: `"${t}" transformer failed to parse the value`,
            cause: r
          });
        }
    };
    return /* @__PURE__ */ Object.assign(
      o,
      {
        isValid(s) {
          try {
            return o(s), !0;
          } catch {
            return !1;
          }
        }
      }
    );
  };
}
function A(t, e) {
  return B(e || "object", (n) => {
    const o = on(n);
    return nn(t, (s) => o[s]);
  });
}
function vt(t) {
  throw new _t(cs, `Unexpected value received: ${JSON.stringify(t)}`);
}
const Nt = B("boolean", (t) => {
  if (typeof t == "boolean")
    return t;
  const e = String(t);
  if (e === "1" || e === "true")
    return !0;
  if (e === "0" || e === "false")
    return !1;
  vt(t);
}), C = B("string", (t) => {
  if (typeof t == "string" || typeof t == "number")
    return t.toString();
  vt(t);
}), dt = B("number", (t) => {
  if (typeof t == "number")
    return t;
  if (typeof t == "string") {
    const e = Number(t);
    if (!Number.isNaN(e))
      return e;
  }
  vt(t);
}), us = B("date", (t) => t instanceof Date ? t : new Date(dt()(t) * 1e3));
function sn(t, e) {
  return B(e || "searchParams", (n) => {
    typeof n != "string" && !(n instanceof URLSearchParams) && vt(n);
    const o = typeof n == "string" ? new URLSearchParams(n) : n;
    return nn(t, (s) => {
      const r = o.get(s);
      return r === null ? void 0 : r;
    });
  });
}
function zt(t) {
  for (const e in t)
    t[e] = [tn(e), t[e]];
  return t;
}
const ls = (t) => {
  const e = dt(), n = dt(!0), o = C(), s = C(!0), r = Nt(!0), a = A(zt({
    addedToAttachmentMenu: r,
    allowsWriteToPm: r,
    firstName: o,
    id: e,
    isBot: r,
    isPremium: r,
    languageCode: s,
    lastName: s,
    photoUrl: s,
    username: s
  }), "User")(!0);
  return sn(
    zt({
      authDate: us(),
      canSendAfter: n,
      chat: A(
        zt({
          id: e,
          type: o,
          title: o,
          photoUrl: s,
          username: s
        }),
        "Chat"
      )(!0),
      chatInstance: s,
      chatType: s,
      hash: o,
      queryId: s,
      receiver: a,
      startParam: s,
      user: a
    }),
    "initData"
  )(t);
};
function ps(t) {
  return /^#[\da-f]{6}$/i.test(t);
}
function _s(t) {
  return /^#[\da-f]{3}$/i.test(t);
}
function ds(t) {
  const e = t.replace(/\s/g, "").toLowerCase();
  if (ps(e))
    return e;
  if (_s(e)) {
    let o = "#";
    for (let s = 0; s < 3; s += 1)
      o += e[1 + s].repeat(2);
    return o;
  }
  const n = e.match(/^rgb\((\d{1,3}),(\d{1,3}),(\d{1,3})\)$/) || e.match(/^rgba\((\d{1,3}),(\d{1,3}),(\d{1,3}),\d{1,3}\)$/);
  if (!n)
    throw new Error(`Value "${t}" does not satisfy any of known RGB formats.`);
  return n.slice(1).reduce((o, s) => {
    const r = parseInt(s, 10).toString(16);
    return o + (r.length === 1 ? "0" : "") + r;
  }, "#");
}
const fs = B("rgb", (t) => ds(C()(t))), hs = B(
  "themeParams",
  (t) => {
    const e = fs(!0);
    return Object.entries(on(t)).reduce((n, [o, s]) => (n[rs(o)] = e(s), n), {});
  }
);
// @__NO_SIDE_EFFECTS__
function rn(t) {
  return JSON.stringify(
    Object.fromEntries(
      Object.entries(t).map(([e, n]) => [tn(e), n])
    )
  );
}
const bs = (t) => {
  const e = C(), n = C(!0), o = Nt(!0);
  return sn({
    botInline: ["tgWebAppBotInline", o],
    initData: ["tgWebAppData", ls(!0)],
    initDataRaw: ["tgWebAppData", n],
    platform: ["tgWebAppPlatform", e],
    showSettings: ["tgWebAppShowSettings", o],
    startParam: ["tgWebAppStartParam", n],
    themeParams: ["tgWebAppThemeParams", hs()],
    version: ["tgWebAppVersion", e]
  }, "launchParams")(t);
};
// @__NO_SIDE_EFFECTS__
function ms(t) {
  const { initDataRaw: e, startParam: n, showSettings: o, botInline: s } = t, r = new URLSearchParams();
  return r.set("tgWebAppPlatform", t.platform), r.set("tgWebAppThemeParams", /* @__PURE__ */ rn(t.themeParams)), r.set("tgWebAppVersion", t.version), e && r.set("tgWebAppData", e), n && r.set("tgWebAppStartParam", n), typeof o == "boolean" && r.set("tgWebAppShowSettings", o ? "1" : "0"), typeof s == "boolean" && r.set("tgWebAppBotInline", s ? "1" : "0"), r.toString();
}
const an = A({
  eventType: C(),
  eventData: (t) => t
}, "miniAppsMessage"), cn = B("fn", (t) => {
  if (typeof t == "function")
    return t;
  vt(t);
});
function gs(t) {
  return !!t && typeof t == "object" && !Array.isArray(t);
}
const ws = A({
  TelegramWebviewProxy: A({ postEvent: cn() })()
});
function un(t) {
  return ws().isValid(t);
}
function Es() {
  try {
    return window.self !== window.top;
  } catch {
    return !0;
  }
}
var ys = Object.defineProperty, Ss = (t, e, n) => e in t ? ys(t, e, { enumerable: !0, configurable: !0, writable: !0, value: n }) : t[e] = n, ln = (t, e, n) => Ss(t, typeof e != "symbol" ? e + "" : e, n);
let l = class pn extends Error {
  constructor(e, n, o) {
    super(
      typeof n == "object" ? n.message : n || e,
      {
        cause: typeof n == "object" ? n.cause : o
      }
    ), this.type = e, Object.setPrototypeOf(this, pn.prototype);
  }
};
function Zt(t, e, n) {
  return t.addEventListener(e, n), () => t.removeEventListener(e, n);
}
function rt(...t) {
  const e = t.flat(1);
  return [
    e.push.bind(e),
    () => {
      e.forEach((n) => {
        n();
      });
    }
  ];
}
function Cs(t, e) {
  return t instanceof l && t.type === e;
}
function se(t) {
  return (e) => Cs(e, t);
}
const _n = "ERR_ABORTED", dn = "ERR_CANCELED", fn = "ERR_TIMED_OUT";
function Oe(t) {
  return new l(_n, { cause: t });
}
const Ei = se(fn), yi = se(_n), Si = se(dn);
function Ie(t, e) {
  return t.reject = e.reject, t;
}
class f extends Promise {
  constructor(e, n) {
    let o, s;
    typeof e == "function" ? (o = e, s = n) : s = e;
    let r, a;
    super((i, p) => {
      s || (s = {});
      const { abortSignal: u } = s;
      if (u && u.aborted)
        return p(Oe(u.reason));
      const [d, w] = rt(), y = (tt) => (...Yo) => (w(), tt(...Yo)), Z = new AbortController(), { signal: G } = Z;
      a = y((tt) => {
        Z.abort(tt), p(tt);
      }), r = y(i), u && d(
        Zt(u, "abort", () => {
          a(Oe(u.reason));
        })
      );
      const { timeout: Ht } = s;
      if (Ht) {
        const tt = setTimeout(() => {
          a(new l(fn, `Timeout reached: ${Ht}ms`));
        }, Ht);
        d(() => {
          clearTimeout(tt);
        });
      }
      o && o(r, a, G);
    }), ln(this, "reject"), this.reject = a;
  }
  /**
   * Creates a new BetterPromise instance using executor, resolving promise when a result
   * was returned.
   * @param fn - function returning promise result.
   * @param options - additional options.
   */
  static withFn(e, n) {
    return new f((o, s, r) => {
      try {
        const a = e(r);
        return a instanceof Promise ? a.then(o, s) : o(a);
      } catch (a) {
        s(a);
      }
    }, n);
  }
  /**
   * @see Promise.resolve
   */
  static resolve(e) {
    return new f((n) => {
      n(e);
    });
  }
  /**
   * @see Promise.reject
   */
  static reject(e) {
    return new f((n, o) => {
      o(e);
    });
  }
  /**
   * Cancels the promise execution.
   */
  cancel() {
    this.reject(new l(dn));
  }
  /**
   * @see Promise.catch
   */
  catch(e) {
    return this.then(void 0, e);
  }
  /**
   * @see Promise.finally
   */
  finally(e) {
    return Ie(super.finally(e), this);
  }
  /**
   * @see Promise.then
   */
  then(e, n) {
    return Ie(super.then(e, n), this);
  }
}
function Me(t, e) {
  return t.resolve = e.resolve, t;
}
class ut extends f {
  constructor(e, n) {
    let o, s;
    typeof e == "function" ? (o = e, s = n) : s = e;
    let r;
    super((a, i, p) => {
      r = a, o && o(a, i, p);
    }, s), ln(this, "resolve"), this.resolve = r;
  }
  /**
   * Creates a new EnhancedPromise instance using executor, resolving promise when a result
   * was returned.
   * @param fn - function returning promise result.
   * @param options - additional options.
   */
  static withFn(e, n) {
    return new ut(
      (o, s, r) => f.withFn(e, { abortSignal: r }).then(o, s),
      n
    );
  }
  /**
   * @see Promise.resolve
   */
  static resolve(e) {
    return new ut((n) => {
      n(e);
    });
  }
  /**
   * @see Promise.reject
   */
  static reject(e) {
    return new ut((n, o) => {
      o(e);
    });
  }
  /**
   * @see Promise.catch
   */
  catch(e) {
    return this.then(void 0, e);
  }
  /**
   * @see Promise.finally
   */
  finally(e) {
    return Me(super.finally(e), this);
  }
  /**
   * @see Promise.then
   */
  then(e, n) {
    return Me(super.then(e, n), this);
  }
}
function vs(t, e) {
  return new f((n) => {
    setTimeout(n, t);
  }, { abortSignal: e });
}
function hn(t) {
  return `tapps/${t}`;
}
function P(t, e) {
  sessionStorage.setItem(hn(t), JSON.stringify(e));
}
function $(t) {
  const e = sessionStorage.getItem(hn(t));
  try {
    return e ? JSON.parse(e) : void 0;
  } catch {
  }
}
function re(t) {
  return t.replace(/[A-Z]/g, (e) => `-${e.toLowerCase()}`);
}
// @__NO_SIDE_EFFECTS__
function Ps(t, e) {
  e || (e = {});
  const {
    textColor: n,
    bgColor: o,
    shouldLog: s = !0
  } = e;
  function r(a, ...i) {
    if (!s || typeof s == "function" && !s())
      return;
    const p = "font-weight:bold;padding:0 5px;border-radius:5px";
    console[a](
      `%c${Intl.DateTimeFormat("en-GB", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
        fractionalSecondDigits: 3,
        timeZone: "UTC"
      }).format(/* @__PURE__ */ new Date())}%c / %c${t}`,
      `${p};background-color: lightblue;color:black`,
      "",
      `${p};${n ? `color:${n};` : ""}${o ? `background-color:${o}` : ""}`,
      ...i
    );
  }
  return [
    function(...a) {
      r("log", ...a);
    },
    function(...a) {
      r("error", ...a);
    }
  ];
}
function ae(t, e) {
  document.documentElement.style.setProperty(t, e);
}
function ie(t) {
  document.documentElement.style.removeProperty(t);
}
function $s(t, e) {
  e();
}
// @__NO_SIDE_EFFECTS__
function at(t, e) {
  e || (e = {});
  const n = e.equals || Object.is;
  let o = [], s = t;
  const r = (u) => {
    if (!n(s, u)) {
      const d = s;
      s = u, $s(p, () => {
        [...o].forEach(([w, y]) => {
          w(u, d), y && i(w, !0);
        });
      });
    }
  };
  function a(u) {
    const d = typeof u != "object" ? { once: u } : u;
    return {
      once: d.once || !1,
      signal: d.signal || !1
    };
  }
  const i = (u, d) => {
    const w = a(d), y = o.findIndex(([Z, G]) => Z === u && G.once === w.once && G.signal === w.signal);
    y >= 0 && o.splice(y, 1);
  }, p = Object.assign(
    function() {
      return Ts(p), s;
    },
    {
      destroy() {
        o = [];
      },
      set: r,
      reset() {
        r(t);
      },
      sub(u, d) {
        return o.push([u, a(d)]), () => i(u, d);
      },
      unsub: i,
      unsubAll() {
        o = o.filter((u) => u[1].signal);
      }
    }
  );
  return p;
}
const Yt = [];
function Ts(t) {
  Yt.length && Yt[Yt.length - 1].add(t);
}
const As = /* @__PURE__ */ at(!1), [ce, Bs] = /* @__PURE__ */ Ps("Bridge", {
  bgColor: "#9147ff",
  textColor: "white",
  shouldLog: As
}), Rs = {
  clipboard_text_received: A({
    req_id: C(),
    data: (t) => t === null ? t : C(!0)(t)
  }, "clipboard_text_received"),
  custom_method_invoked: A({
    req_id: C(),
    result: (t) => t,
    error: C(!0)
  }, "custom_method_invoked"),
  popup_closed: B("popup_closed", (t) => t ? A({
    button_id: (e) => e == null ? void 0 : C()(e)
  })()(t) : {}),
  viewport_changed: A({
    height: dt(),
    width: (t) => t == null ? window.innerWidth : dt()(t),
    is_state_stable: Nt(),
    is_expanded: Nt()
  }, "viewport_changed")
};
function Os(t) {
  const e = window, [, n] = rt(
    // Add "resize" event listener to make sure, we always have fresh viewport information.
    // The desktop version of Telegram is sometimes not sending the "viewport_changed"
    // event. For example, when the Main Button is shown. That's why we should
    // add our own listener to make sure viewport information is always fresh.
    // Issue: https://github.com/Telegram-Mini-Apps/telegram-apps/issues/10
    Zt(e, "resize", () => {
      t(["viewport_changed", {
        width: window.innerWidth,
        height: window.innerHeight,
        is_state_stable: !0,
        is_expanded: !0
      }]);
    }),
    // Add listener, which handles events sent from the Telegram web application and also events
    // generated by the local emitEvent function.
    Zt(e, "message", (o) => {
      if (o.source !== e.parent)
        return;
      let s;
      try {
        s = an()(o.data);
      } catch {
        return;
      }
      const { eventType: r, eventData: a } = s, i = Rs[r];
      try {
        const p = i ? i()(a) : a;
        ce("Event received:", p ? { eventType: r, eventData: p } : { eventType: r }), t([r, p]);
      } catch (p) {
        Bs(
          [
            `An error occurred processing the "${r}" event from the Telegram application.`,
            "Please, file an issue here:",
            "https://github.com/Telegram-Mini-Apps/telegram-apps/issues/new/choose"
          ].join(`
`),
          s,
          p
        );
      }
    })
  );
  return n;
}
const te = /* @__PURE__ */ at(), De = /* @__PURE__ */ at();
function bn() {
  return De() || De.set(Os(te.set)), te;
}
const Ft = /* @__PURE__ */ at({});
function mn(t) {
  let e = Ft()[t];
  return e || (e = /* @__PURE__ */ at(void 0, {
    equals() {
      return !1;
    }
  }), bn().sub((n) => {
    n && n[0] === t && e.set(n[1]);
  }), Ft.set({ ...Ft(), [t]: e })), e;
}
function S(t, e, n) {
  return mn(t).sub(e, n);
}
const Is = "ERR_METHOD_UNSUPPORTED", Ms = "ERR_RETRIEVE_LP_FAILED", Ds = "ERR_METHOD_PARAMETER_UNSUPPORTED", xs = "ERR_UNKNOWN_ENV", ks = "ERR_INVOKE_CUSTOM_METHOD_RESPONSE", Ns = /* @__PURE__ */ at("https://web.telegram.org");
function ue(t, e) {
  ce("Posting event:", e ? { eventType: t, eventData: e } : { eventType: t });
  const n = window;
  if (un(n)) {
    n.TelegramWebviewProxy.postEvent(t, JSON.stringify(e));
    return;
  }
  const o = JSON.stringify({ eventType: t, eventData: e });
  if (Es())
    return n.parent.postMessage(o, Ns());
  const { external: s } = n;
  if (A({ notify: cn() })().isValid(s)) {
    s.notify(o);
    return;
  }
  throw new l(xs);
}
function le(t, e, n) {
  n || (n = {});
  const { capture: o } = n, [s, r] = rt();
  return new f((a) => {
    (Array.isArray(e) ? e : [e]).forEach((i) => {
      s(
        S(i, (p) => {
          (!o || (Array.isArray(e) ? o({
            event: i,
            payload: p
          }) : o(p))) && a(p);
        })
      );
    }), (n.postEvent || ue)(t, n.params);
  }, n).finally(r);
}
function pe(t) {
  return bs()(t);
}
function gn(t) {
  return pe(
    t.replace(/^[^?#]*[?#]/, "").replace(/[?#]/g, "&")
  );
}
function Vs() {
  return gn(window.location.href);
}
function js() {
  const t = performance.getEntriesByType("navigation")[0];
  if (!t)
    throw new Error("Unable to get first navigation entry.");
  return gn(t.name);
}
const Us = "launchParams";
function Ls() {
  return pe($(Us) || "");
}
function wn(t) {
  P("launchParams", /* @__PURE__ */ ms(t));
}
function En(t) {
  return t instanceof Error ? t.message + (t.cause ? `
  ${En(t.cause)}` : "") : JSON.stringify(t);
}
function it() {
  const t = [];
  for (const e of [
    // Try to retrieve launch parameters from the current location. This method can return
    // nothing in case, location was changed, and then the page was reloaded.
    Vs,
    // Then, try using the lower level API - window.performance.
    js,
    // Finally, try to extract launch parameters from the session storage.
    Ls
  ])
    try {
      const n = e();
      return wn(n), n;
    } catch (n) {
      t.push(n);
    }
  throw new l(Ms, [
    "Unable to retrieve launch parameters from any known source. Perhaps, you have opened your app outside Telegram?",
    "📖 Refer to docs for more information:",
    "https://docs.telegram-mini-apps.com/packages/telegram-apps-sdk/environment",
    "Collected errors:",
    ...t.map((e) => `— ${En(e)}`)
  ].join(`
`));
}
function Ci(t) {
  if (t === "simple")
    try {
      return it(), !0;
    } catch {
      return !1;
    }
  return f.withFn(async () => {
    if (un(window))
      return !0;
    try {
      return await le("web_app_request_theme", "theme_changed", { timeout: 100 }), !0;
    } catch {
      return !1;
    }
  }, t);
}
function ee(t, e) {
  window.dispatchEvent(new MessageEvent("message", {
    data: JSON.stringify({ eventType: t, eventData: e }),
    // We specify window.parent to imitate the case, the parent iframe sent us this event.
    source: window.parent
  }));
}
function Ws(t, e) {
  if (typeof e == "string")
    try {
      const { eventType: n } = an()(e);
      n === "web_app_request_theme" && ee("theme_changed", {
        theme_params: JSON.parse(/* @__PURE__ */ rn(t))
      }), n === "web_app_request_viewport" && ee("viewport_changed", {
        width: window.innerWidth,
        height: window.innerHeight,
        is_state_stable: !0,
        is_expanded: !0
      });
    } catch {
    }
}
function vi(t) {
  var e;
  const n = typeof t == "string" ? pe(t) : t;
  wn(n);
  const o = (e = window.TelegramWebviewProxy) == null ? void 0 : e.postEvent;
  window.TelegramWebviewProxy = {
    postEvent(s, r) {
      Ws(n.themeParams, JSON.stringify({ eventType: s, eventData: r })), o == null || o(s, r);
    }
  }, ce("Environment was mocked by the mockTelegramEnv function");
}
function qs() {
  [
    ["TelegramGameProxy_receiveEvent"],
    // Windows Phone.
    ["TelegramGameProxy", "receiveEvent"],
    // Desktop.
    ["Telegram", "WebView", "receiveEvent"]
    // Android and iOS.
  ].forEach((t) => {
    let e = window;
    t.forEach((n, o, s) => {
      if (o === s.length - 1) {
        e[n] = ee;
        return;
      }
      n in e || (e[n] = {}), e = e[n];
    });
  });
}
function Hs() {
  ["TelegramGameProxy_receiveEvent", "TelegramGameProxy", "Telegram"].forEach((t) => {
    delete window[t];
  });
}
function K(t, e, n) {
  mn(t).unsub(e, n);
}
function Pi(t, e) {
  return bn().sub(t, e);
}
function $i(t, e) {
  te.unsub(t, e);
}
function yn(t) {
  return ({ req_id: e }) => e === t;
}
function xe(t) {
  return t.split(".").map(Number);
}
function Gs(t, e) {
  const n = xe(t), o = xe(e), s = Math.max(n.length, o.length);
  for (let r = 0; r < s; r += 1) {
    const a = n[r] || 0, i = o[r] || 0;
    if (a !== i)
      return a > i ? 1 : -1;
  }
  return 0;
}
function b(t, e) {
  return Gs(t, e) <= 0;
}
function j(t, e, n) {
  if (typeof n == "string") {
    if (t === "web_app_open_link") {
      if (e === "try_instant_view")
        return b("6.4", n);
      if (e === "try_browser")
        return b("7.6", n);
    }
    if (t === "web_app_set_header_color" && e === "color")
      return b("6.9", n);
    if (t === "web_app_close" && e === "return_back")
      return b("7.6", n);
    if (t === "web_app_setup_main_button" && e === "has_shine_effect")
      return b("7.10", n);
  }
  switch (t) {
    case "web_app_open_tg_link":
    case "web_app_open_invoice":
    case "web_app_setup_back_button":
    case "web_app_set_background_color":
    case "web_app_set_header_color":
    case "web_app_trigger_haptic_feedback":
      return b("6.1", e);
    case "web_app_open_popup":
      return b("6.2", e);
    case "web_app_close_scan_qr_popup":
    case "web_app_open_scan_qr_popup":
    case "web_app_read_text_from_clipboard":
      return b("6.4", e);
    case "web_app_switch_inline_query":
      return b("6.7", e);
    case "web_app_invoke_custom_method":
    case "web_app_request_write_access":
    case "web_app_request_phone":
      return b("6.9", e);
    case "web_app_setup_settings_button":
      return b("6.10", e);
    case "web_app_biometry_get_info":
    case "web_app_biometry_open_settings":
    case "web_app_biometry_request_access":
    case "web_app_biometry_request_auth":
    case "web_app_biometry_update_token":
      return b("7.2", e);
    case "web_app_setup_swipe_behavior":
      return b("7.7", e);
    case "web_app_share_to_story":
      return b("7.8", e);
    case "web_app_setup_secondary_button":
    case "web_app_set_bottom_bar_color":
      return b("7.10", e);
    default:
      return [
        "iframe_ready",
        "iframe_will_reload",
        "web_app_close",
        "web_app_data_send",
        "web_app_expand",
        "web_app_open_link",
        "web_app_ready",
        "web_app_request_theme",
        "web_app_request_viewport",
        "web_app_setup_main_button",
        "web_app_setup_closing_behavior"
      ].includes(t);
  }
}
function zs(t, e) {
  e || (e = "strict");
  const n = typeof e == "function" ? e : (o) => {
    const { method: s, version: r } = o;
    let a, i;
    if ("param" in o ? (a = `Parameter "${o.param}" of "${s}" method is unsupported in Mini Apps version ${r}`, i = Ds) : (a = `Method "${s}" is unsupported in Mini Apps version ${r}`, i = Is), e === "strict")
      throw new l(i, a);
    return console.warn(a);
  };
  return (o, s) => j(o, t) ? gs(s) && o === "web_app_set_header_color" && "color" in s && !j(o, "color", t) ? n({ version: t, method: o, param: "color" }) : ue(o, s) : n({ version: t, method: o });
}
function Ys(t, e, n, o) {
  return le("web_app_invoke_custom_method", "custom_method_invoked", {
    ...o || {},
    params: { method: t, params: e, req_id: n },
    capture: yn(n)
  }).then(({ result: s, error: r }) => {
    if (r)
      throw new l(ks, r);
    return s;
  });
}
function Fs() {
  return performance.getEntriesByType("navigation")[0];
}
function R() {
  const t = Fs();
  return !!t && t.type === "reload";
}
function ke(t, e) {
  return t.startsWith(e) ? t : `${e}${t}`;
}
function Js(t) {
  return new URL(
    typeof t == "string" ? t : [
      t.pathname || "",
      ke(t.search || "", "?"),
      ke(t.hash || "", "#")
    ].join(""),
    "http://a"
  );
}
let et;
function Ks(t, e) {
  et && et.set(t, e) || e();
}
function Qs(t) {
  if (et)
    return t();
  et = /* @__PURE__ */ new Map();
  try {
    t();
  } finally {
    et.forEach((e) => e()), et = void 0;
  }
}
// @__NO_SIDE_EFFECTS__
function c(t, e) {
  e || (e = {});
  const n = e.equals || Object.is;
  let o = [], s = t;
  const r = (u) => {
    if (!n(s, u)) {
      const d = s;
      s = u, Ks(p, () => {
        [...o].forEach(([w, y]) => {
          w(u, d), y && i(w, !0);
        });
      });
    }
  };
  function a(u) {
    const d = typeof u != "object" ? { once: u } : u;
    return {
      once: d.once || !1,
      signal: d.signal || !1
    };
  }
  const i = (u, d) => {
    const w = a(d), y = o.findIndex(([Z, G]) => Z === u && G.once === w.once && G.signal === w.signal);
    y >= 0 && o.splice(y, 1);
  }, p = Object.assign(
    function() {
      return Xs(p), s;
    },
    {
      destroy() {
        o = [];
      },
      set: r,
      reset() {
        r(t);
      },
      sub(u, d) {
        return o.push([u, a(d)]), () => i(u, d);
      },
      unsub: i,
      unsubAll() {
        o = o.filter((u) => u[1].signal);
      }
    }
  );
  return p;
}
const lt = [];
function Xs(t) {
  lt.length && lt[lt.length - 1].add(t);
}
// @__NO_SIDE_EFFECTS__
function h(t, e) {
  let n = /* @__PURE__ */ new Set();
  const o = /* @__PURE__ */ c(r(), e);
  function s() {
    o.set(r());
  }
  function r() {
    n.forEach((p) => p.unsub(s, { signal: !0 }));
    const a = /* @__PURE__ */ new Set();
    let i;
    lt.push(a);
    try {
      i = t();
    } finally {
      lt.pop();
    }
    return a.forEach((p) => {
      p.sub(s, { signal: !0 });
    }), n = a, i;
  }
  return Object.assign(function() {
    return o();
  }, {
    destroy: o.destroy,
    sub: o.sub,
    unsub: o.unsub,
    unsubAll: o.unsubAll
  });
}
const Zs = /* @__PURE__ */ c(/* @__PURE__ */ (() => {
  let t = 0;
  return () => (t += 1).toString();
})()), Sn = /* @__PURE__ */ c(ue), Q = /* @__PURE__ */ c("0.0");
function tr(t) {
  t || (t = {});
  const { postEvent: e } = t, n = t.version || it().version;
  Q.set(n), Sn.set(
    typeof e == "function" ? e : zs(n)
  );
}
function Cn() {
  return Zs()();
}
function Pt(t, e, n) {
  return Ys(t, e, Cn(), {
    ...n || {},
    postEvent: _
  });
}
const O = (t, e, n) => (n || (n = {}), n.postEvent || (n.postEvent = _), le(t, e, n)), _ = (t, e) => Sn()(t, e), ct = "ERR_POPUP_INVALID_PARAMS", vn = "ERR_INVALID_HOSTNAME", er = "ERR_INVALID_SLUG", nr = "ERR_DATA_INVALID_SIZE", or = "ERR_ACCESS_DENIED", T = "ERR_ALREADY_CALLED", Pn = "ERR_NOT_AVAILABLE", $n = "ERR_NOT_SUPPORTED", sr = "ERR_NOT_MOUNTED";
// @__NO_SIDE_EFFECTS__
function g(t, e) {
  function n() {
    return typeof e == "string" ? j(e, Q()) : e();
  }
  return Object.assign((...o) => {
    if (!n())
      throw new l($n);
    return t(...o);
  }, {
    isSupported: n
  });
}
// @__NO_SIDE_EFFECTS__
function N(t) {
  return (e) => /* @__PURE__ */ g(e, t);
}
function v(t, e) {
  e(), t.sub(e);
}
function I(t) {
  return /* @__PURE__ */ h(() => j(t, Q()));
}
function _e(t, e) {
  return (...n) => {
    if (!e())
      throw new l(sr);
    return t(...n);
  };
}
// @__NO_SIDE_EFFECTS__
function q(t) {
  return (e) => _e(e, t);
}
const Tn = "web_app_setup_back_button", An = "back_button_pressed", Bn = "backButton", ft = /* @__PURE__ */ c(!1), Rn = I(Tn), de = /* @__PURE__ */ N(Rn), On = /* @__PURE__ */ q(ft), rr = On(() => {
  z.set(!1);
}), z = /* @__PURE__ */ c(!1), ar = de(() => {
  ft() || (z.set(R() && $(Bn) || !1), v(z, In), ft.set(!0));
});
function In() {
  const t = z();
  _(Tn, { is_visible: t }), P(Bn, t);
}
const ir = de(
  (t) => S(An, t)
), cr = de((t) => {
  K(An, t);
}), ur = On(() => {
  z.set(!0);
});
function lr() {
  z.unsub(In), ft.set(!1);
}
const Ti = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  hide: rr,
  isMounted: ft,
  isSupported: Rn,
  isVisible: z,
  mount: ar,
  offClick: cr,
  onClick: ir,
  show: ur,
  unmount: lr
}, Symbol.toStringTag, { value: "Module" }));
// @__NO_SIDE_EFFECTS__
function Mn(t, e, {
  isMounting: n,
  isMounted: o,
  mountError: s,
  isSupported: r
}) {
  return /* @__PURE__ */ g(
    (a) => {
      if (o())
        return f.resolve();
      if (n())
        throw new l(T);
      return n.set(!0), f.withFn((i) => t({ abortSignal: i }), a).then(
        (i) => [!0, i],
        (i) => [!1, i]
      ).then((i) => {
        Qs(() => {
          if (n.set(!1), o.set(!0), i[0])
            e(i[1]);
          else {
            const p = i[1];
            throw s.set(p), p;
          }
        });
      });
    },
    r || (() => !0)
  );
}
const x = /* @__PURE__ */ c(), At = /* @__PURE__ */ c(!1), Bt = /* @__PURE__ */ c(!1), fe = /* @__PURE__ */ c(!1), Dn = /* @__PURE__ */ c(!1), xn = /* @__PURE__ */ c(void 0);
function he(t) {
  return t.available ? {
    available: !0,
    tokenSaved: t.token_saved,
    deviceId: t.device_id,
    accessRequested: t.access_requested,
    type: t.type,
    accessGranted: t.access_granted
  } : {
    available: !1
  };
}
const Ne = "web_app_biometry_get_info", pr = /* @__PURE__ */ g(
  (t) => O(Ne, "biometry_info_received", t).then(he),
  Ne
), kn = "web_app_biometry_request_auth", _r = "web_app_biometry_request_access", dr = "web_app_biometry_open_settings", fr = "web_app_biometry_update_token", be = "biometry_info_received", Nn = "biometry", me = I(kn), hr = /* @__PURE__ */ N(me), ge = /* @__PURE__ */ q(fe), br = ge(
  (t) => {
    if (At())
      return f.reject(new l(T));
    const e = x();
    return !e || !e.available ? f.reject(new l(Pn)) : (At.set(!0), t || (t = {}), O(kn, "biometry_auth_requested", {
      ...t,
      params: { reason: (t.reason || "").trim() }
    }).then((n) => {
      const { token: o } = n;
      return typeof o == "string" && x.set({ ...e, token: o }), n;
    }).finally(() => {
      At.set(!1);
    }));
  }
), mr = hr(() => {
  _(dr);
}), gr = ge(
  (t) => Bt() ? f.reject(new l(T)) : (Bt.set(!0), t || (t = {}), O(_r, be, {
    ...t,
    params: { reason: t.reason || "" }
  }).then(he).then((e) => {
    if (!e.available)
      throw new l(Pn);
    return x.set(e), e.accessGranted;
  }).finally(() => {
    Bt.set(!1);
  }))
), wr = /* @__PURE__ */ Mn(
  (t) => {
    const e = R() && $(Nn);
    return e || pr(t);
  },
  (t) => {
    S(be, Vn), v(x, jn), x.set(t);
  },
  { isMounted: fe, mountError: xn, isMounting: Dn, isSupported: me }
), Vn = (t) => {
  x.set(he(t));
};
function jn() {
  const t = x();
  t && P(Nn, t);
}
function Er() {
  K(be, Vn), x.unsub(jn);
}
const yr = ge(
  (t) => (t || (t = {}), O(fr, "biometry_token_updated", {
    ...t,
    params: {
      token: t.token || "",
      reason: t.reason
    }
  }).then((e) => e.status))
), Ai = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  authenticate: br,
  isAuthenticating: At,
  isMounted: fe,
  isMounting: Dn,
  isRequestingAccess: Bt,
  isSupported: me,
  mount: wr,
  mountError: xn,
  openSettings: mr,
  requestAccess: gr,
  state: x,
  unmount: Er,
  updateToken: yr
}, Symbol.toStringTag, { value: "Module" })), Un = "closingConfirmation", ht = /* @__PURE__ */ c(!1), Ln = /* @__PURE__ */ q(ht), Sr = Ln(() => {
  Y.set(!1);
}), Y = /* @__PURE__ */ c(!1), Cr = Ln(() => {
  Y.set(!0);
});
function vr() {
  ht() || (Y.set(R() && $(Un) || !1), v(Y, Wn), ht.set(!0));
}
function Wn() {
  const t = Y();
  _("web_app_setup_closing_behavior", { need_confirmation: t }), P(Un, t);
}
function Pr() {
  Y.unsub(Wn), ht.set(!1);
}
const Bi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  disableConfirmation: Sr,
  enableConfirmation: Cr,
  isConfirmationEnabled: Y,
  isMounted: ht,
  mount: vr,
  unmount: Pr
}, Symbol.toStringTag, { value: "Module" })), $r = "web_app_invoke_custom_method", qn = I($r), Ut = /* @__PURE__ */ N(qn), Tr = Ut((t, e) => {
  const n = Array.isArray(t) ? t : [t];
  return n.length ? Pt("deleteStorageValues", { keys: n }, e).then() : f.resolve();
});
function Ar(t, e) {
  const n = Array.isArray(t) ? t : [t];
  return n.length ? Pt("getStorageValues", { keys: n }, e).then((o) => {
    const s = xt(
      Object.fromEntries(n.map((r) => [r, D()]))
    )()(o);
    return Array.isArray(t) ? s : s[t];
  }) : f.resolve(typeof t == "string" ? "" : {});
}
const Br = Ut(Ar), Rr = Ut(
  (t) => Pt("getStorageKeys", {}, t).then(ss(D())())
), Or = Ut(
  (t, e, n) => Pt("saveStorageValue", {
    key: t,
    value: e
  }, n).then()
), Ri = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  deleteItem: Tr,
  getItem: Br,
  getKeys: Rr,
  isSupported: qn,
  setItem: Or
}, Symbol.toStringTag, { value: "Module" })), Lt = "web_app_trigger_haptic_feedback", Hn = I(Lt), we = /* @__PURE__ */ N(Hn), Ir = we((t) => {
  _(Lt, { type: "impact", impact_style: t });
}), Mr = we((t) => {
  _(Lt, { type: "notification", notification_type: t });
}), Dr = we(() => {
  _(Lt, { type: "selection_change" });
}), Oi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  impactOccurred: Ir,
  isSupported: Hn,
  notificationOccurred: Mr,
  selectionChanged: Dr
}, Symbol.toStringTag, { value: "Module" })), Ee = /* @__PURE__ */ c(void 0);
function M(t) {
  return /* @__PURE__ */ h(() => {
    const e = Ee();
    return e ? e[t] : void 0;
  });
}
const Gn = M("authDate"), zn = M("canSendAfter"), xr = /* @__PURE__ */ h(() => {
  const t = Gn(), e = zn();
  return e && t ? new Date(t.getTime() + e * 1e3) : void 0;
}), kr = M("chat"), Nr = M("chatType"), Vr = M("chatInstance"), jr = M("hash"), Ur = M("queryId"), Yn = /* @__PURE__ */ c(), Lr = M("receiver");
function Wr() {
  const t = it();
  Ee.set(t.initData), Yn.set(t.initDataRaw);
}
const qr = M("startParam"), Hr = M("user"), Ii = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  authDate: Gn,
  canSendAfter: zn,
  canSendAfterDate: xr,
  chat: kr,
  chatInstance: Vr,
  chatType: Nr,
  hash: jr,
  queryId: Ur,
  raw: Yn,
  receiver: Lr,
  restore: Wr,
  startParam: qr,
  state: Ee,
  user: Hr
}, Symbol.toStringTag, { value: "Module" }));
function Mi(t) {
  return Zo()(t);
}
const Fn = "web_app_open_invoice", Rt = /* @__PURE__ */ c(!1), Jn = I(Fn);
async function Kn(t, e, n) {
  if (Rt())
    throw new l(T);
  let o;
  if (e === "url") {
    const { hostname: s, pathname: r } = new URL(t, window.location.href);
    if (s !== "t.me")
      throw new l(vn);
    const a = r.match(/^\/(\$|invoice\/)([A-Za-z0-9\-_=]+)$/);
    if (!a)
      throw new l(er);
    [, , o] = a;
  } else
    o = t, n = e;
  return Rt.set(!0), O(Fn, "invoice_closed", {
    ...n,
    params: { slug: o },
    capture: (s) => o === s.slug
  }).then((s) => s.status).finally(() => {
    Rt.set(!1);
  });
}
const Gr = /* @__PURE__ */ g(Kn, Jn), Di = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  _open: Kn,
  isOpened: Rt,
  isSupported: Jn,
  open: Gr
}, Symbol.toStringTag, { value: "Module" }));
function Qn(t) {
  const e = Qe(t);
  return Math.sqrt(
    [0.299, 0.587, 0.114].reduce((n, o, s) => {
      const r = parseInt(e.slice(1 + s * 2, 1 + (s + 1) * 2), 16);
      return n + r * r * o;
    }, 0)
  ) < 120;
}
const bt = /* @__PURE__ */ c(!1), Ot = /* @__PURE__ */ c(!1), V = /* @__PURE__ */ c({});
function m(t) {
  return /* @__PURE__ */ h(() => V()[t]);
}
const zr = m("accentTextColor"), ye = m("bgColor"), Se = m("buttonColor"), Xn = m("buttonTextColor"), Zn = m("bottomBarBgColor"), Yr = m("destructiveTextColor"), Fr = m("headerBgColor"), Jr = m("hintColor"), Kr = /* @__PURE__ */ h(() => {
  const { bgColor: t } = V();
  return !t || Qn(t);
}), Qr = m("linkColor"), Vt = m("secondaryBgColor"), Xr = m("sectionBgColor"), Zr = m("sectionHeaderTextColor"), ta = m("sectionSeparatorColor"), ea = m("subtitleTextColor"), na = m("textColor");
function X(t) {
  return /* @__PURE__ */ h(() => $t()[t]);
}
const ot = /* @__PURE__ */ c({
  hasShineEffect: !1,
  isEnabled: !0,
  isLoaderVisible: !1,
  isVisible: !1,
  text: "Continue"
}), $t = /* @__PURE__ */ h(() => {
  const t = ot();
  return {
    ...t,
    backgroundColor: t.backgroundColor || Se() || "#2481cc",
    textColor: t.textColor || Xn() || "#ffffff"
  };
}), mt = /* @__PURE__ */ c(!1), oa = X("backgroundColor"), sa = X("hasShineEffect"), ra = X("isEnabled"), aa = X("isLoaderVisible"), ia = X("isVisible"), ca = X("text"), ua = X("textColor"), la = "web_app_setup_main_button", to = "main_button_pressed", eo = "mainButton", pa = /* @__PURE__ */ q(mt);
function _a() {
  if (!mt()) {
    const t = R() && $(eo);
    t && ot.set(t), ot.sub(no), v($t, oo), mt.set(!0);
  }
}
function da(t) {
  return S(to, t);
}
function fa(t) {
  K(to, t);
}
function no(t) {
  P(eo, t);
}
function oo() {
  const t = $t();
  t.text && _(la, {
    color: t.backgroundColor,
    has_shine_effect: t.hasShineEffect,
    is_active: t.isEnabled,
    is_progress_visible: t.isLoaderVisible,
    is_visible: t.isVisible,
    text: t.text,
    text_color: t.textColor
  });
}
const ha = pa((t) => {
  ot.set({
    ...ot(),
    ...Object.fromEntries(
      Object.entries(t).filter(([, e]) => e !== void 0)
    )
  });
});
function ba() {
  ot.unsub(no), $t.unsub(oo), mt.set(!1);
}
const xi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  backgroundColor: oa,
  hasShineEffect: sa,
  isEnabled: ra,
  isLoaderVisible: aa,
  isMounted: mt,
  isVisible: ia,
  mount: _a,
  offClick: fa,
  onClick: da,
  setParams: ha,
  state: $t,
  text: ca,
  textColor: ua,
  unmount: ba
}, Symbol.toStringTag, { value: "Module" }));
function ma(t) {
  return ns()(t);
}
const so = "themeParams", ro = "theme_changed", ga = _e((t) => {
  if (Ot())
    throw new l(T);
  t || (t = (o) => `--tg-theme-${re(o)}`);
  function e(o) {
    Object.entries(V()).forEach(([s, r]) => {
      r && o(s, r);
    });
  }
  function n() {
    e((o, s) => {
      ae(t(o), s);
    });
  }
  return n(), V.sub(n), Ot.set(!0), () => {
    e(ie), V.unsub(n), Ot.set(!1);
  };
}, bt);
function ao() {
  bt() || (S(ro, io), V.set(R() && $(so) || it().themeParams), bt.set(!0));
}
const io = (t) => {
  const e = ma(t.theme_params);
  V.set(e), P(so, e);
};
function wa() {
  K(ro, io), bt.set(!1);
}
// @__NO_SIDE_EFFECTS__
function Ea(t, e) {
  function n(o) {
    const s = e[o];
    return j(s[0], s[1], Q());
  }
  return Object.assign(
    (...o) => {
      for (const s in e)
        if (e[s][2](...o) && !n(s))
          throw new l($n, `Parameter "${s}" is not supported`);
      return t(...o);
    },
    t,
    { supports: n }
  );
}
// @__NO_SIDE_EFFECTS__
function co(t) {
  return /* @__PURE__ */ h(() => {
    const e = t();
    return Ct(e) ? e : e === "bg_color" ? ye() : Vt();
  });
}
const U = /* @__PURE__ */ c("bg_color"), Ce = /* @__PURE__ */ co(U), L = /* @__PURE__ */ c("bottom_bar_bg_color"), ve = /* @__PURE__ */ h(() => {
  const t = L();
  return Ct(t) ? t : t === "bottom_bar_bg_color" ? Zn() || Vt() : t === "secondary_bg_color" ? Vt() : ye();
}), W = /* @__PURE__ */ c("bg_color"), uo = /* @__PURE__ */ co(W), gt = /* @__PURE__ */ c(!1), It = /* @__PURE__ */ c(!1), ya = /* @__PURE__ */ h(() => {
  const t = Ce();
  return t ? Qn(t) : !1;
}), lo = /* @__PURE__ */ h(() => ({
  backgroundColor: U(),
  bottomBarColor: L(),
  headerColor: W()
})), Pe = "web_app_set_background_color", $e = "web_app_set_bottom_bar_color", jt = "web_app_set_header_color", po = "miniApp", _o = /* @__PURE__ */ h(() => [
  Pe,
  $e,
  jt
].some((t) => j(t, Q()))), Sa = /* @__PURE__ */ N(_o), Wt = /* @__PURE__ */ q(gt), Ca = Wt((t) => {
  if (It())
    throw new l(T);
  const [e, n] = rt();
  function o(s, r) {
    function a() {
      ae(s, r() || null);
    }
    a(), e(r.sub(a), ie.bind(null, s));
  }
  return t || (t = (s) => `--tg-${re(s)}`), o(t("bgColor"), Ce), o(t("bottomBarColor"), ve), o(t("headerColor"), uo), e(() => {
    It.set(!1);
  }), It.set(!0), n;
});
function va(t) {
  _("web_app_close", { return_back: t });
}
const Pa = Sa(() => {
  if (!gt()) {
    const t = R() && $(po);
    ao(), U.set(t ? t.backgroundColor : "bg_color"), L.set(t ? t.bottomBarColor : "bottom_bar_bg_color"), W.set(t ? t.headerColor : "bg_color"), mo.isSupported() && v(U, fo), go.isSupported() && v(L, ho), wo.isSupported() && v(W, bo), gt.set(!0);
  }
});
function fo() {
  Te(), _(Pe, { color: U() });
}
function ho() {
  Te(), _($e, { color: L() });
}
function bo() {
  const t = W();
  Te(), _(jt, Ct(t) ? { color: t } : { color_key: t });
}
function $a() {
  _("web_app_ready");
}
function Te() {
  P(po, lo());
}
const mo = /* @__PURE__ */ g(
  Wt((t) => {
    U.set(t);
  }),
  Pe
), go = /* @__PURE__ */ g(
  Wt((t) => {
    L.set(t);
  }),
  $e
), wo = /* @__PURE__ */ Ea(
  /* @__PURE__ */ g(
    Wt((t) => {
      W.set(t);
    }),
    jt
  ),
  {
    color: [jt, "color", Ct]
  }
);
function Ta() {
  U.unsub(fo), L.unsub(ho), W.unsub(bo), gt.set(!1);
}
const ki = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  backgroundColor: U,
  backgroundColorRGB: Ce,
  bindCssVars: Ca,
  bottomBarColor: L,
  bottomBarColorRGB: ve,
  close: va,
  headerColor: W,
  headerColorRGB: uo,
  isCssVarsBound: It,
  isDark: ya,
  isMounted: gt,
  isSupported: _o,
  mount: Pa,
  ready: $a,
  setBackgroundColor: mo,
  setBottomBarColor: go,
  setHeaderColor: wo,
  state: lo,
  unmount: Ta
}, Symbol.toStringTag, { value: "Module" }));
function Aa(t) {
  const e = t.message.trim(), n = (t.title || "").trim(), o = t.buttons || [];
  if (n.length > 64)
    throw new l(ct, `Invalid title: ${n}`);
  if (!e || e.length > 256)
    throw new l(ct, `Invalid message: ${e}`);
  if (o.length > 3)
    throw new l(ct, `Invalid buttons count: ${o.length}`);
  return {
    title: n,
    message: e,
    buttons: o.length ? o.map((s, r) => {
      const a = s.id || "";
      if (a.length > 64)
        throw new l(ct, `Button with index ${r} has invalid id: ${a}`);
      if (!s.type || s.type === "default" || s.type === "destructive") {
        const i = s.text.trim();
        if (!i || i.length > 64)
          throw new l(ct, `Button with index ${r} has invalid text: ${i}`);
        return { type: s.type, text: i, id: a };
      }
      return { type: s.type, id: a };
    }) : [{ type: "close", id: "" }]
  };
}
const Eo = "web_app_open_popup", yo = I(Eo), Mt = /* @__PURE__ */ c(!1), Ba = /* @__PURE__ */ g(
  async (t) => {
    if (Mt())
      throw new l(T);
    Mt.set(!0);
    try {
      const { button_id: e = null } = await O(Eo, "popup_closed", {
        ...t,
        params: Aa(t)
      });
      return e;
    } finally {
      Mt.set(!1);
    }
  },
  yo
), Ni = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  isOpened: Mt,
  isSupported: yo,
  open: Ba
}, Symbol.toStringTag, { value: "Module" })), Ra = "web_app_close_scan_qr_popup", So = "web_app_open_scan_qr_popup", Oa = "scan_qr_popup_closed", Ia = "qr_text_received", Co = I(So), vo = /* @__PURE__ */ N(Co), ne = vo(() => {
  nt.set(!1), _(Ra);
}), nt = /* @__PURE__ */ c(!1);
function Ma(t) {
  return f.withFn((e) => {
    if (nt())
      throw new l(T);
    nt.set(!0), t || (t = {});
    const { onCaptured: n, text: o, capture: s } = t, [, r] = rt(
      // Whenever the scanner was closed for some reason (by a developer or a user), we should
      // resolve the promise with undefined.
      nt.sub(() => {
        a.resolve();
      }),
      // Whenever user closed the scanner, update the isOpened signal state.
      S(Oa, () => {
        nt.set(!1);
      }),
      // Whenever some QR was scanned, we should check if it must be captured.
      S(Ia, (i) => {
        n ? n(i.data) : (!s || s(i.data)) && (a.resolve(i.data), ne());
      })
    ), a = new ut({ abortSignal: e }).catch(ne).finally(r);
    return (t.postEvent || _)(So, { text: o }), a;
  }, t);
}
const Da = vo(Ma), Vi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  close: ne,
  isOpened: nt,
  isSupported: Co,
  open: Da
}, Symbol.toStringTag, { value: "Module" }));
function H(t) {
  return /* @__PURE__ */ h(() => Tt()[t]);
}
const st = /* @__PURE__ */ c({
  hasShineEffect: !1,
  isEnabled: !0,
  isLoaderVisible: !1,
  isVisible: !1,
  position: "left",
  text: "Cancel"
}), Tt = /* @__PURE__ */ h(() => {
  const t = st();
  return {
    ...t,
    backgroundColor: t.backgroundColor || ve() || "#000000",
    textColor: t.textColor || Se() || "#2481cc"
  };
}), wt = /* @__PURE__ */ c(!1), xa = H("backgroundColor"), ka = H("hasShineEffect"), Na = H("isEnabled"), Va = H("isLoaderVisible"), ja = H("isVisible"), Ua = H("position"), La = H("text"), Wa = H("textColor"), Po = "web_app_setup_secondary_button", $o = "secondary_button_pressed", To = "secondaryButton", Ao = I(Po), Ae = /* @__PURE__ */ N(Ao), qa = /* @__PURE__ */ q(wt), Ha = Ae(() => {
  if (!wt()) {
    const t = R() && $(To);
    t && st.set(t), st.sub(Bo), v(Tt, Ro), wt.set(!0);
  }
}), Ga = Ae(
  (t) => S($o, t)
), za = Ae(
  (t) => {
    K($o, t);
  }
);
function Bo(t) {
  P(To, t);
}
function Ro() {
  const t = Tt();
  t.text && _(Po, {
    color: t.backgroundColor,
    has_shine_effect: t.hasShineEffect,
    is_active: t.isEnabled,
    is_progress_visible: t.isLoaderVisible,
    is_visible: t.isVisible,
    position: t.position,
    text: t.text,
    text_color: t.textColor
  });
}
const Ya = qa((t) => {
  st.set({
    ...st(),
    ...Object.fromEntries(
      Object.entries(t).filter(([, e]) => e !== void 0)
    )
  });
});
function Fa() {
  st.unsub(Bo), Tt.unsub(Ro), wt.set(!1);
}
const ji = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  backgroundColor: xa,
  hasShineEffect: ka,
  isEnabled: Na,
  isLoaderVisible: Va,
  isMounted: wt,
  isSupported: Ao,
  isVisible: ja,
  mount: Ha,
  offClick: za,
  onClick: Ga,
  position: Ua,
  setParams: Ya,
  state: Tt,
  text: La,
  textColor: Wa,
  unmount: Fa
}, Symbol.toStringTag, { value: "Module" })), Oo = "web_app_setup_settings_button", Io = "settings_button_pressed", Mo = "settingsButton", Et = /* @__PURE__ */ c(!1), Do = I(Oo), Be = /* @__PURE__ */ N(Do), xo = /* @__PURE__ */ q(Et), Ja = xo(() => {
  F.set(!1);
}), F = /* @__PURE__ */ c(!1), Ka = Be(() => {
  Et() || (F.set(R() && $(Mo) || !1), v(F, ko), Et.set(!0));
});
function ko() {
  const t = F();
  _(Oo, { is_visible: t }), P(Mo, t);
}
const Qa = Be(
  (t) => S(Io, t)
), Xa = Be(
  (t) => {
    K(Io, t);
  }
), Za = xo(() => {
  F.set(!0);
});
function ti() {
  F.unsub(ko), Et.set(!1);
}
const Ui = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  hide: Ja,
  isMounted: Et,
  isSupported: Do,
  isVisible: F,
  mount: Ka,
  offClick: Xa,
  onClick: Qa,
  show: Za,
  unmount: ti
}, Symbol.toStringTag, { value: "Module" })), No = "web_app_setup_swipe_behavior", Vo = "swipeBehavior", yt = /* @__PURE__ */ c(!1), jo = I(No), ei = /* @__PURE__ */ N(jo), Uo = /* @__PURE__ */ q(yt), ni = Uo(() => {
  J.set(!1);
}), oi = Uo(() => {
  J.set(!0);
}), J = /* @__PURE__ */ c(!0), si = ei(() => {
  yt() || (J.set(R() && $(Vo) || !1), v(J, Lo), yt.set(!0));
});
function Lo() {
  const t = J();
  _(No, { allow_vertical_swipe: t }), P(Vo, t);
}
function ri() {
  J.unsub(Lo), yt.set(!1);
}
const Li = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  disableVertical: ni,
  enableVertical: oi,
  isMounted: yt,
  isSupported: jo,
  isVerticalEnabled: J,
  mount: si,
  unmount: ri
}, Symbol.toStringTag, { value: "Module" })), Wi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  accentTextColor: zr,
  backgroundColor: ye,
  bindCssVars: ga,
  bottomBarBgColor: Zn,
  buttonColor: Se,
  buttonTextColor: Xn,
  destructiveTextColor: Yr,
  headerBackgroundColor: Fr,
  hintColor: Jr,
  isCssVarsBound: Ot,
  isDark: Kr,
  isMounted: bt,
  linkColor: Qr,
  mount: ao,
  secondaryBackgroundColor: Vt,
  sectionBackgroundColor: Xr,
  sectionHeaderTextColor: Zr,
  sectionSeparatorColor: ta,
  state: V,
  subtitleTextColor: ea,
  textColor: na,
  unmount: wa
}, Symbol.toStringTag, { value: "Module" })), E = /* @__PURE__ */ c({
  height: 0,
  width: 0,
  isExpanded: !1,
  stableHeight: 0
}), Re = /* @__PURE__ */ c(!1), Dt = /* @__PURE__ */ c(!1), Wo = /* @__PURE__ */ c(!1), qo = /* @__PURE__ */ c(void 0);
function qt(t) {
  return /* @__PURE__ */ h(() => E()[t]);
}
const ai = qt("height"), ii = qt("isExpanded"), ci = /* @__PURE__ */ h(() => {
  const t = E();
  return t.height === t.stableHeight;
}), ui = qt("stableHeight"), li = qt("width");
function pi(t) {
  return O("web_app_request_viewport", "viewport_changed", t).then((e) => ({
    height: e.height,
    width: e.width,
    isExpanded: e.is_expanded,
    isStable: e.is_state_stable
  }));
}
const _i = _e((t) => {
  if (Dt())
    throw new l(T);
  t || (t = (o) => `--tg-viewport-${re(o)}`);
  const e = ["height", "width", "stableHeight"];
  function n() {
    e.forEach((o) => {
      ae(t(o), `${E()[o]}px`);
    });
  }
  return n(), E.sub(n), Dt.set(!0), () => {
    e.forEach(ie), E.unsub(n), Dt.set(!1);
  };
}, Re);
function di() {
  _("web_app_expand");
}
function Ho(t) {
  return {
    isExpanded: t.isExpanded,
    height: Jt(t.height),
    width: Jt(t.width),
    stableHeight: Jt(t.stableHeight)
  };
}
const fi = /* @__PURE__ */ Mn(
  (t) => {
    const e = R() && $("viewport");
    if (e)
      return e;
    if ([
      "macos",
      "tdesktop",
      "unigram",
      "webk",
      "weba",
      "web"
    ].includes(it().platform)) {
      const n = window;
      return {
        isExpanded: !0,
        height: n.innerHeight,
        width: n.innerWidth,
        stableHeight: n.innerHeight
      };
    }
    return t.timeout || (t.timeout = 1e3), pi(t).then((n) => ({
      height: n.height,
      isExpanded: n.isExpanded,
      stableHeight: n.isStable ? n.height : E().stableHeight,
      width: n.width
    }));
  },
  (t) => {
    S("viewport_changed", Go), v(E, zo), E.set(Ho(t));
  },
  { isMounted: Re, isMounting: Wo, mountError: qo }
), Go = (t) => {
  E.set(Ho({
    height: t.height,
    width: t.width,
    isExpanded: t.is_expanded,
    stableHeight: t.is_state_stable ? t.height : E().stableHeight
  }));
};
function zo() {
  P("viewport", E());
}
function Jt(t) {
  return Math.max(t, 0);
}
function hi() {
  K("viewport_changed", Go), E.unsub(zo);
}
const qi = /* @__PURE__ */ Object.freeze(/* @__PURE__ */ Object.defineProperty({
  __proto__: null,
  bindCssVars: _i,
  expand: di,
  height: ai,
  isCssVarsBound: Dt,
  isExpanded: ii,
  isMounted: Re,
  isMounting: Wo,
  isStable: ci,
  mount: fi,
  mountError: qo,
  stableHeight: ui,
  state: E,
  unmount: hi,
  width: li
}, Symbol.toStringTag, { value: "Module" })), Ve = "web_app_open_tg_link";
function Hi(t, e) {
  e || (e = {}), _("web_app_open_link", {
    url: Js(t).toString(),
    try_browser: e.tryBrowser,
    try_instant_view: e.tryInstantView
  });
}
function bi(t) {
  const { hostname: e, pathname: n, search: o } = new URL(t, "https://t.me");
  if (e !== "t.me")
    throw new l(vn);
  if (!j(Ve, Q())) {
    window.location.href = t;
    return;
  }
  _(Ve, { path_full: n + o });
}
function Gi(t, e) {
  bi(
    "https://t.me/share/url?" + new URLSearchParams({ url: t, text: e || "" }).toString().replace(/\+/g, "%20")
  );
}
const oe = "web_app_request_phone", je = "web_app_request_write_access", Kt = /* @__PURE__ */ c(!1), Qt = /* @__PURE__ */ c(!1);
function Ue(t) {
  return t || (t = {}), Pt("getRequestedContact", {}, {
    ...t,
    timeout: t.timeout || 5e3
  }).then(
    Ke({
      contact: xt({
        userId: ["user_id", kt()],
        phoneNumber: ["phone_number", D()],
        firstName: ["first_name", D()],
        lastName: ["last_name", D(!0)]
      })(),
      authDate: ["auth_date", Je()],
      hash: D()
    })()
  );
}
const zi = /* @__PURE__ */ g(
  (t) => new f(
    async (e, n, o) => {
      const s = {
        postEvent: (t || {}).postEvent,
        abortSignal: o
      };
      try {
        return e(await Ue(s));
      } catch {
      }
      if (await mi(s) !== "sent")
        throw new l(or);
      let a = 50;
      for (; !o.aborted; ) {
        try {
          return e(await Ue(s));
        } catch {
        }
        await vs(a), a += 50;
      }
    },
    t
  ),
  oe
), mi = /* @__PURE__ */ g(
  (t) => {
    if (Kt())
      throw new l(T);
    return Kt.set(!0), O(oe, "phone_requested", t).then((e) => e.status).finally(() => {
      Kt.set(!1);
    });
  },
  oe
), Yi = /* @__PURE__ */ g(
  (t) => {
    if (Qt())
      throw new l(T);
    return Qt.set(!0), O(je, "write_access_requested", t).then((e) => e.status).finally(() => {
      Qt.set(!1);
    });
  },
  je
), Le = "web_app_read_text_from_clipboard", We = "web_app_switch_inline_query", qe = "web_app_share_to_story", Fi = /* @__PURE__ */ g(
  (t) => {
    const e = Cn();
    return O(Le, "clipboard_text_received", {
      ...t,
      params: { req_id: e },
      capture: yn(e)
    }).then(({ data: n = null }) => n);
  },
  Le
);
function Ji(t) {
  const { size: e } = new Blob([t]);
  if (!e || e > 4096)
    throw new l(nr);
  _("web_app_data_send", { data: t });
}
const Ki = /* @__PURE__ */ g(
  (t, e) => {
    e || (e = {}), (e.postEvent || _)(qe, {
      text: e.text,
      media_url: t,
      widget_link: e.widgetLink
    });
  },
  qe
), Qi = /* @__PURE__ */ g(
  (t, e) => {
    _(We, {
      query: t,
      chat_types: e || []
    });
  },
  () => j(We, Q()) && !!it().botInline
);
function Xi() {
  return typeof window > "u";
}
function Zi(t) {
  tr(t), qs();
  const [e, n] = rt(
    S("reload_iframe", () => {
      _("iframe_will_reload"), window.location.reload();
    }),
    Hs
  ), { acceptCustomStyles: o = !0 } = t || {};
  if (o) {
    const s = document.createElement("style");
    s.id = "telegram-custom-styles", document.head.appendChild(s), e(
      S("set_custom_style", (r) => {
        s.innerHTML = r;
      }),
      () => {
        document.head.removeChild(s);
      }
    );
  }
  return _("iframe_ready", { reload_supported: !0 }), n;
}
export {
  Zs as $createRequestId,
  As as $debug,
  Sn as $postEvent,
  Ns as $targetOrigin,
  Q as $version,
  f as CancelablePromise,
  _n as ERR_ABORTED,
  or as ERR_ACCESS_DENIED,
  T as ERR_ALREADY_CALLED,
  dn as ERR_CANCELED,
  ks as ERR_CUSTOM_METHOD_ERR_RESPONSE,
  nr as ERR_DATA_INVALID_SIZE,
  vn as ERR_INVALID_HOSTNAME,
  er as ERR_INVALID_SLUG,
  Jo as ERR_INVALID_VALUE,
  Ds as ERR_METHOD_PARAMETER_UNSUPPORTED,
  Is as ERR_METHOD_UNSUPPORTED,
  Pn as ERR_NOT_AVAILABLE,
  sr as ERR_NOT_MOUNTED,
  $n as ERR_NOT_SUPPORTED,
  ze as ERR_PARSE,
  ct as ERR_POPUP_INVALID_PARAMS,
  Ms as ERR_RETRIEVE_LP_FAILED,
  fn as ERR_TIMED_OUT,
  Qo as ERR_UNEXPECTED_TYPE,
  Ko as ERR_UNEXPECTED_VALUE,
  xs as ERR_UNKNOWN_ENV,
  l as TypedError,
  Zt as addEventListener,
  br as authenticateBiometry,
  Ti as backButton,
  Ca as bindMiniAppCssVars,
  ga as bindThemeParamsCssVars,
  _i as bindViewportCssVars,
  Ai as biometry,
  xn as biometryMountError,
  x as biometryState,
  Xt as classNames,
  va as closeMiniApp,
  ne as closeQrScanner,
  Bi as closingBehavior,
  Ri as cloudStorage,
  Gs as compareVersions,
  zs as createPostEvent,
  qs as defineEventHandlers,
  Tr as deleteCloudStorageItem,
  ie as deleteCssVar,
  Sr as disableClosingConfirmation,
  ni as disableVerticalSwipes,
  ee as emitMiniAppsEvent,
  Cr as enableClosingConfirmation,
  oi as enableVerticalSwipes,
  di as expandViewport,
  Br as getCloudStorageItem,
  Rr as getCloudStorageKeys,
  Oi as hapticFeedback,
  Ir as hapticFeedbackImpactOccurred,
  Mr as hapticFeedbackNotificationOccurred,
  Dr as hapticFeedbackSelectionChanged,
  rr as hideBackButton,
  Ja as hideSettingsButton,
  Zi as init,
  Ii as initData,
  Gn as initDataAuthDate,
  zn as initDataCanSendAfter,
  xr as initDataCanSendAfterDate,
  kr as initDataChat,
  Vr as initDataChatInstance,
  Nr as initDataChatType,
  jr as initDataHash,
  Ur as initDataQueryId,
  Yn as initDataRaw,
  Lr as initDataReceiver,
  qr as initDataStartParam,
  Ee as initDataState,
  Hr as initDataUser,
  Di as invoice,
  Ys as invokeCustomMethod,
  yi as isAbortError,
  At as isAuthenticatingBiometry,
  ft as isBackButtonMounted,
  Rn as isBackButtonSupported,
  z as isBackButtonVisible,
  fe as isBiometryMounted,
  Dn as isBiometryMounting,
  me as isBiometrySupported,
  Si as isCanceledError,
  ht as isClosingBehaviorMounted,
  Y as isClosingConfirmationEnabled,
  qn as isCloudStorageSupported,
  Qn as isColorDark,
  Hn as isHapticFeedbackSupported,
  Es as isIframe,
  Rt as isInvoiceOpened,
  Jn as isInvoiceSupported,
  ra as isMainButtonEnabled,
  aa as isMainButtonLoaderVisible,
  mt as isMainButtonMounted,
  ia as isMainButtonVisible,
  It as isMiniAppCssVarsBound,
  ya as isMiniAppDark,
  gt as isMiniAppMounted,
  _o as isMiniAppSupported,
  Mt as isPopupOpened,
  yo as isPopupSupported,
  nt as isQrScannerOpened,
  Co as isQrScannerSupported,
  Ct as isRGB,
  ts as isRGBShort,
  Xe as isRecord,
  Bt as isRequestingBiometryAccess,
  Kt as isRequestingPhoneAccess,
  Qt as isRequestingWriteAccess,
  Xi as isSSR,
  Na as isSecondaryButtonEnabled,
  Va as isSecondaryButtonLoaderVisible,
  wt as isSecondaryButtonMounted,
  Ao as isSecondaryButtonSupported,
  ja as isSecondaryButtonVisible,
  Et as isSettingsButtonMounted,
  Do as isSettingsButtonSupported,
  F as isSettingsButtonVisible,
  yt as isSwipeBehaviorMounted,
  jo as isSwipeBehaviorSupported,
  Ci as isTMA,
  Ot as isThemeParamsCssVarsBound,
  Kr as isThemeParamsDark,
  bt as isThemeParamsMounted,
  Ei as isTimeoutError,
  J as isVerticalSwipesEnabled,
  Dt as isViewportCssVarsBound,
  ii as isViewportExpanded,
  Re as isViewportMounted,
  Wo as isViewportMounting,
  ci as isViewportStable,
  xi as mainButton,
  oa as mainButtonBackgroundColor,
  sa as mainButtonHasShineEffect,
  $t as mainButtonState,
  ca as mainButtonText,
  ua as mainButtonTextColor,
  wi as mergeClassNames,
  ki as miniApp,
  U as miniAppBackgroundColor,
  L as miniAppBottomBarColor,
  ve as miniAppBottomBarColorRGB,
  W as miniAppHeaderColor,
  uo as miniAppHeaderColorRGB,
  $a as miniAppReady,
  lo as miniAppState,
  vi as mockTelegramEnv,
  ar as mountBackButton,
  wr as mountBiometry,
  vr as mountClosingBehavior,
  _a as mountMainButton,
  Pa as mountMiniApp,
  Ha as mountSecondaryButton,
  Ka as mountSettingsButton,
  si as mountSwipeBehavior,
  ao as mountThemeParams,
  fi as mountViewport,
  K as off,
  cr as offBackButtonClick,
  fa as offMainButtonClick,
  za as offSecondaryButtonClick,
  Xa as offSettingsButtonClick,
  S as on,
  ir as onBackButtonClick,
  da as onMainButtonClick,
  Ga as onSecondaryButtonClick,
  Qa as onSettingsButtonClick,
  mr as openBiometrySettings,
  Gr as openInvoice,
  Hi as openLink,
  Ba as openPopup,
  Da as openQrScanner,
  bi as openTelegramLink,
  Mi as parseInitData,
  ma as parseThemeParams,
  Ni as popup,
  ue as postEvent,
  Vi as qrScanner,
  Fi as readTextFromClipboard,
  Hs as removeEventHandlers,
  le as request,
  pr as requestBiometry,
  gr as requestBiometryAccess,
  zi as requestContact,
  mi as requestPhoneAccess,
  Yi as requestWriteAccess,
  Wr as restoreInitData,
  it as retrieveLaunchParams,
  ji as secondaryButton,
  xa as secondaryButtonBackgroundColor,
  ka as secondaryButtonHasShineEffect,
  Ua as secondaryButtonPosition,
  Tt as secondaryButtonState,
  La as secondaryButtonText,
  Wa as secondaryButtonTextColor,
  Ji as sendData,
  gi as serializeLaunchParams,
  os as serializeThemeParams,
  Or as setCloudStorageItem,
  ae as setCssVar,
  ha as setMainButtonParams,
  mo as setMiniAppBackgroundColor,
  go as setMiniAppBottomBarColor,
  wo as setMiniAppHeaderColor,
  Ya as setSecondaryButtonParams,
  Ui as settingsButton,
  Ki as shareStory,
  Gi as shareURL,
  ur as showBackButton,
  Za as showSettingsButton,
  Pi as subscribe,
  j as supports,
  Li as swipeBehavior,
  Qi as switchInlineQuery,
  Wi as themeParams,
  zr as themeParamsAccentTextColor,
  ye as themeParamsBackgroundColor,
  Zn as themeParamsBottomBarBgColor,
  Se as themeParamsButtonColor,
  Xn as themeParamsButtonTextColor,
  Yr as themeParamsDestructiveTextColor,
  Fr as themeParamsHeaderBackgroundColor,
  Jr as themeParamsHintColor,
  Qr as themeParamsLinkColor,
  Vt as themeParamsSecondaryBackgroundColor,
  Xr as themeParamsSectionBackgroundColor,
  Zr as themeParamsSectionHeaderTextColor,
  ta as themeParamsSectionSeparatorColor,
  V as themeParamsState,
  ea as themeParamsSubtitleTextColor,
  na as themeParamsTextColor,
  Qe as toRGB,
  Fe as toRecord,
  lr as unmountBackButton,
  Er as unmountBiometry,
  Pr as unmountClosingBehavior,
  ba as unmountMainButton,
  Ta as unmountMiniApp,
  Fa as unmountSecondaryButton,
  ti as unmountSettingsButton,
  ri as unmountSwipeBehavior,
  wa as unmountThemeParams,
  hi as unmountViewport,
  $i as unsubscribe,
  yr as updateBiometryToken,
  qi as viewport,
  ai as viewportHeight,
  qo as viewportMountError,
  ui as viewportStableHeight,
  E as viewportState,
  li as viewportWidth
};
//# sourceMappingURL=index.js.map
