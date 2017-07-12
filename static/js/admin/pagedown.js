var Markdown;
if (typeof exports === "object" && typeof require === "function") {
    Markdown = exports
} else {
    Markdown = {}
}
(function () {
    function c(e) {
        return e
    }

    function d(e) {
        return false
    }

    function b() {
    }

    b.prototype = {
        chain: function (g, f) {
            var e = this[g];
            if (!e) {
                throw new Error("unknown hook " + g)
            }
            if (e === c) {
                this[g] = f
            } else {
                this[g] = function (i) {
                    var h = Array.prototype.slice.call(arguments, 0);
                    h[0] = e.apply(null, h);
                    return f.apply(null, h)
                }
            }
        }, set: function (f, e) {
            if (!this[f]) {
                throw new Error("unknown hook " + f)
            }
            this[f] = e
        }, addNoop: function (e) {
            this[e] = c
        }, addFalse: function (e) {
            this[e] = d
        }
    };
    Markdown.HookCollection = b;
    function a() {
    }

    a.prototype = {
        set: function (e, f) {
            this["s_" + e] = f
        }, get: function (e) {
            return this["s_" + e]
        }
    };
    Markdown.Converter = function (Y) {
        var l = this.hooks = new b();
        l.addNoop("plainLinkText");
        l.addNoop("preConversion");
        l.addNoop("postNormalization");
        l.addNoop("preBlockGamut");
        l.addNoop("postBlockGamut");
        l.addNoop("preSpanGamut");
        l.addNoop("postSpanGamut");
        l.addNoop("postConversion");
        var B;
        var q;
        var e;
        var F;
        Y = Y || {};
        var W = c, r = c;
        if (Y.nonAsciiLetters) {
            (function () {
                var ac = /[Q\u00aa\u00b5\u00ba\u00c0-\u00d6\u00d8-\u00f6\u00f8-\u02c1\u02c6-\u02d1\u02e0-\u02e4\u02ec\u02ee\u0370-\u0374\u0376-\u0377\u037a-\u037d\u0386\u0388-\u038a\u038c\u038e-\u03a1\u03a3-\u03f5\u03f7-\u0481\u048a-\u0523\u0531-\u0556\u0559\u0561-\u0587\u05d0-\u05ea\u05f0-\u05f2\u0621-\u064a\u0660-\u0669\u066e-\u066f\u0671-\u06d3\u06d5\u06e5-\u06e6\u06ee-\u06fc\u06ff\u0710\u0712-\u072f\u074d-\u07a5\u07b1\u07c0-\u07ea\u07f4-\u07f5\u07fa\u0904-\u0939\u093d\u0950\u0958-\u0961\u0966-\u096f\u0971-\u0972\u097b-\u097f\u0985-\u098c\u098f-\u0990\u0993-\u09a8\u09aa-\u09b0\u09b2\u09b6-\u09b9\u09bd\u09ce\u09dc-\u09dd\u09df-\u09e1\u09e6-\u09f1\u0a05-\u0a0a\u0a0f-\u0a10\u0a13-\u0a28\u0a2a-\u0a30\u0a32-\u0a33\u0a35-\u0a36\u0a38-\u0a39\u0a59-\u0a5c\u0a5e\u0a66-\u0a6f\u0a72-\u0a74\u0a85-\u0a8d\u0a8f-\u0a91\u0a93-\u0aa8\u0aaa-\u0ab0\u0ab2-\u0ab3\u0ab5-\u0ab9\u0abd\u0ad0\u0ae0-\u0ae1\u0ae6-\u0aef\u0b05-\u0b0c\u0b0f-\u0b10\u0b13-\u0b28\u0b2a-\u0b30\u0b32-\u0b33\u0b35-\u0b39\u0b3d\u0b5c-\u0b5d\u0b5f-\u0b61\u0b66-\u0b6f\u0b71\u0b83\u0b85-\u0b8a\u0b8e-\u0b90\u0b92-\u0b95\u0b99-\u0b9a\u0b9c\u0b9e-\u0b9f\u0ba3-\u0ba4\u0ba8-\u0baa\u0bae-\u0bb9\u0bd0\u0be6-\u0bef\u0c05-\u0c0c\u0c0e-\u0c10\u0c12-\u0c28\u0c2a-\u0c33\u0c35-\u0c39\u0c3d\u0c58-\u0c59\u0c60-\u0c61\u0c66-\u0c6f\u0c85-\u0c8c\u0c8e-\u0c90\u0c92-\u0ca8\u0caa-\u0cb3\u0cb5-\u0cb9\u0cbd\u0cde\u0ce0-\u0ce1\u0ce6-\u0cef\u0d05-\u0d0c\u0d0e-\u0d10\u0d12-\u0d28\u0d2a-\u0d39\u0d3d\u0d60-\u0d61\u0d66-\u0d6f\u0d7a-\u0d7f\u0d85-\u0d96\u0d9a-\u0db1\u0db3-\u0dbb\u0dbd\u0dc0-\u0dc6\u0e01-\u0e30\u0e32-\u0e33\u0e40-\u0e46\u0e50-\u0e59\u0e81-\u0e82\u0e84\u0e87-\u0e88\u0e8a\u0e8d\u0e94-\u0e97\u0e99-\u0e9f\u0ea1-\u0ea3\u0ea5\u0ea7\u0eaa-\u0eab\u0ead-\u0eb0\u0eb2-\u0eb3\u0ebd\u0ec0-\u0ec4\u0ec6\u0ed0-\u0ed9\u0edc-\u0edd\u0f00\u0f20-\u0f29\u0f40-\u0f47\u0f49-\u0f6c\u0f88-\u0f8b\u1000-\u102a\u103f-\u1049\u1050-\u1055\u105a-\u105d\u1061\u1065-\u1066\u106e-\u1070\u1075-\u1081\u108e\u1090-\u1099\u10a0-\u10c5\u10d0-\u10fa\u10fc\u1100-\u1159\u115f-\u11a2\u11a8-\u11f9\u1200-\u1248\u124a-\u124d\u1250-\u1256\u1258\u125a-\u125d\u1260-\u1288\u128a-\u128d\u1290-\u12b0\u12b2-\u12b5\u12b8-\u12be\u12c0\u12c2-\u12c5\u12c8-\u12d6\u12d8-\u1310\u1312-\u1315\u1318-\u135a\u1380-\u138f\u13a0-\u13f4\u1401-\u166c\u166f-\u1676\u1681-\u169a\u16a0-\u16ea\u1700-\u170c\u170e-\u1711\u1720-\u1731\u1740-\u1751\u1760-\u176c\u176e-\u1770\u1780-\u17b3\u17d7\u17dc\u17e0-\u17e9\u1810-\u1819\u1820-\u1877\u1880-\u18a8\u18aa\u1900-\u191c\u1946-\u196d\u1970-\u1974\u1980-\u19a9\u19c1-\u19c7\u19d0-\u19d9\u1a00-\u1a16\u1b05-\u1b33\u1b45-\u1b4b\u1b50-\u1b59\u1b83-\u1ba0\u1bae-\u1bb9\u1c00-\u1c23\u1c40-\u1c49\u1c4d-\u1c7d\u1d00-\u1dbf\u1e00-\u1f15\u1f18-\u1f1d\u1f20-\u1f45\u1f48-\u1f4d\u1f50-\u1f57\u1f59\u1f5b\u1f5d\u1f5f-\u1f7d\u1f80-\u1fb4\u1fb6-\u1fbc\u1fbe\u1fc2-\u1fc4\u1fc6-\u1fcc\u1fd0-\u1fd3\u1fd6-\u1fdb\u1fe0-\u1fec\u1ff2-\u1ff4\u1ff6-\u1ffc\u203f-\u2040\u2054\u2071\u207f\u2090-\u2094\u2102\u2107\u210a-\u2113\u2115\u2119-\u211d\u2124\u2126\u2128\u212a-\u212d\u212f-\u2139\u213c-\u213f\u2145-\u2149\u214e\u2183-\u2184\u2c00-\u2c2e\u2c30-\u2c5e\u2c60-\u2c6f\u2c71-\u2c7d\u2c80-\u2ce4\u2d00-\u2d25\u2d30-\u2d65\u2d6f\u2d80-\u2d96\u2da0-\u2da6\u2da8-\u2dae\u2db0-\u2db6\u2db8-\u2dbe\u2dc0-\u2dc6\u2dc8-\u2dce\u2dd0-\u2dd6\u2dd8-\u2dde\u2e2f\u3005-\u3006\u3031-\u3035\u303b-\u303c\u3041-\u3096\u309d-\u309f\u30a1-\u30fa\u30fc-\u30ff\u3105-\u312d\u3131-\u318e\u31a0-\u31b7\u31f0-\u31ff\u3400-\u4db5\u4e00-\u9fc3\ua000-\ua48c\ua500-\ua60c\ua610-\ua62b\ua640-\ua65f\ua662-\ua66e\ua67f-\ua697\ua717-\ua71f\ua722-\ua788\ua78b-\ua78c\ua7fb-\ua801\ua803-\ua805\ua807-\ua80a\ua80c-\ua822\ua840-\ua873\ua882-\ua8b3\ua8d0-\ua8d9\ua900-\ua925\ua930-\ua946\uaa00-\uaa28\uaa40-\uaa42\uaa44-\uaa4b\uaa50-\uaa59\uac00-\ud7a3\uf900-\ufa2d\ufa30-\ufa6a\ufa70-\ufad9\ufb00-\ufb06\ufb13-\ufb17\ufb1d\ufb1f-\ufb28\ufb2a-\ufb36\ufb38-\ufb3c\ufb3e\ufb40-\ufb41\ufb43-\ufb44\ufb46-\ufbb1\ufbd3-\ufd3d\ufd50-\ufd8f\ufd92-\ufdc7\ufdf0-\ufdfb\ufe33-\ufe34\ufe4d-\ufe4f\ufe70-\ufe74\ufe76-\ufefc\uff10-\uff19\uff21-\uff3a\uff3f\uff41-\uff5a\uff66-\uffbe\uffc2-\uffc7\uffca-\uffcf\uffd2-\uffd7\uffda-\uffdc]/g;
                var Z = "Q".charCodeAt(0);
                var ad = "A".charCodeAt(0);
                var ab = "Z".charCodeAt(0);
                var aa = "a".charCodeAt(0) - ab - 1;
                W = function (ae) {
                    return ae.replace(ac, function (af) {
                        var ai = af.charCodeAt(0);
                        var ah = "";
                        var ag;
                        while (ai > 0) {
                            ag = (ai % 51) + ad;
                            if (ag >= Z) {
                                ag++
                            }
                            if (ag > ab) {
                                ag += aa
                            }
                            ah = String.fromCharCode(ag) + ah;
                            ai = ai / 51 | 0
                        }
                        return "Q" + ah + "Q"
                    })
                };
                r = function (ae) {
                    return ae.replace(/Q([A-PR-Za-z]{1,3})Q/g, function (af, ai) {
                        var aj = 0;
                        var ag;
                        for (var ah = 0; ah < ai.length; ah++) {
                            ag = ai.charCodeAt(ah);
                            if (ag > ab) {
                                ag -= aa
                            }
                            if (ag > Z) {
                                ag--
                            }
                            ag -= ad;
                            aj = (aj * 51) + ag
                        }
                        return String.fromCharCode(aj)
                    })
                }
            })()
        }
        var D = Y.asteriskIntraWordEmphasis ? v : t;
        this.makeHtml = function (Z) {
            if (B) {
                throw new Error("Recursive call to converter.makeHtml")
            }
            B = new a();
            q = new a();
            e = [];
            F = 0;
            Z = l.preConversion(Z);
            Z = Z.replace(/~/g, "~T");
            Z = Z.replace(/\$/g, "~D");
            Z = Z.replace(/\r\n/g, "\n");
            Z = Z.replace(/\r/g, "\n");
            Z = "\n\n" + Z + "\n\n";
            Z = P(Z);
            Z = Z.replace(/^[ \t]+$/mg, "");
            Z = l.postNormalization(Z);
            Z = s(Z);
            Z = p(Z);
            Z = h(Z);
            Z = T(Z);
            Z = Z.replace(/~D/g, "$$");
            Z = Z.replace(/~T/g, "~");
            Z = l.postConversion(Z);
            e = q = B = null;
            return Z
        };
        function p(Z) {
            Z = Z.replace(/^[ ]{0,3}\[([^\[\]]+)\]:[ \t]*\n?[ \t]*<?(\S+?)>?(?=\s|$)[ \t]*\n?[ \t]*((\n*)["(](.+?)[")][ \t]*)?(?:\n+)/gm, function (ac, ae, ad, ab, aa, af) {
                ae = ae.toLowerCase();
                B.set(ae, H(ad));
                if (aa) {
                    return ab
                } else {
                    if (af) {
                        q.set(ae, af.replace(/"/g, "&quot;"))
                    }
                }
                return ""
            });
            return Z
        }

        function s(ab) {
            var aa = "p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|script|noscript|form|fieldset|iframe|math|ins|del";
            var Z = "p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|script|noscript|form|fieldset|iframe|math";
            ab = ab.replace(/^(<(p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|script|noscript|form|fieldset|iframe|math|ins|del)\b[^\r]*?\n<\/\2>[ \t]*(?=\n+))/gm, Q);
            ab = ab.replace(/^(<(p|div|h[1-6]|blockquote|pre|table|dl|ol|ul|script|noscript|form|fieldset|iframe|math)\b[^\r]*?.*<\/\2>[ \t]*(?=\n+)\n)/gm, Q);
            ab = ab.replace(/\n[ ]{0,3}((<(hr)\b([^<>])*?\/?>)[ \t]*(?=\n{2,}))/g, Q);
            ab = ab.replace(/\n\n[ ]{0,3}(<!(--(?:|(?:[^>-]|-[^>])(?:[^-]|-[^-])*)--)>[ \t]*(?=\n{2,}))/g, Q);
            ab = ab.replace(/(?:\n\n)([ ]{0,3}(?:<([?%])[^\r]*?\2>)[ \t]*(?=\n{2,}))/g, Q);
            return ab
        }

        function V(Z) {
            Z = Z.replace(/(^\n+|\n+$)/g, "");
            return "\n\n~K" + (e.push(Z) - 1) + "K\n\n"
        }

        function Q(Z, aa) {
            return V(aa)
        }

        var g = function (Z) {
            return h(Z)
        };

        function h(ab, aa) {
            ab = l.preBlockGamut(ab, g);
            ab = k(ab);
            var Z = "<hr />\n";
            ab = ab.replace(/^[ ]{0,2}([ ]?\*[ ]?){3,}[ \t]*$/gm, Z);
            ab = ab.replace(/^[ ]{0,2}([ ]?-[ ]?){3,}[ \t]*$/gm, Z);
            ab = ab.replace(/^[ ]{0,2}([ ]?_[ ]?){3,}[ \t]*$/gm, Z);
            ab = S(ab);
            ab = w(ab);
            ab = i(ab);
            ab = l.postBlockGamut(ab, g);
            ab = s(ab);
            ab = O(ab, aa);
            return ab
        }

        function m(Z) {
            Z = l.preSpanGamut(Z);
            Z = x(Z);
            Z = A(Z);
            Z = M(Z);
            Z = I(Z);
            Z = J(Z);
            Z = R(Z);
            Z = Z.replace(/~P/g, "://");
            Z = H(Z);
            Z = D(Z);
            Z = Z.replace(/  +\n/g, " <br>\n");
            Z = l.postSpanGamut(Z);
            return Z
        }

        function A(aa) {
            var Z = /(<[a-z\/!$]("[^"]*"|'[^']*'|[^'">])*>|<!(--(?:|(?:[^>-]|-[^>])(?:[^-]|-[^-])*)--)>)/gi;
            aa = aa.replace(Z, function (ac) {
                var ab = ac.replace(/(.)<\/?code>(?=.)/g, "$1`");
                ab = E(ab, ac.charAt(1) == "!" ? "\\`*_/" : "\\`*_");
                return ab
            });
            return aa
        }

        function J(Z) {
            if (Z.indexOf("[") === -1) {
                return Z
            }
            Z = Z.replace(/(\[((?:\[[^\]]*\]|[^\[\]])*)\][ ]?(?:\n[ ]*)?\[(.*?)\])()()()()/g, j);
            Z = Z.replace(/(\[((?:\[[^\]]*\]|[^\[\]])*)\]\([ \t]*()<?((?:\([^)]*\)|[^()\s])*?)>?[ \t]*((['"])(.*?)\6[ \t]*)?\))/g, j);
            Z = Z.replace(/(\[([^\[\]]+)\])()()()()()/g, j);
            return Z
        }

        function j(af, al, ak, aj, ai, ah, ae, ad) {
            if (ad == undefined) {
                ad = ""
            }
            var ac = al;
            var aa = ak.replace(/:\/\//g, "~P");
            var ab = aj.toLowerCase();
            var Z = ai;
            var ag = ad;
            if (Z == "") {
                if (ab == "") {
                    ab = aa.toLowerCase().replace(/ ?\n/g, " ")
                }
                Z = "#" + ab;
                if (B.get(ab) != undefined) {
                    Z = B.get(ab);
                    if (q.get(ab) != undefined) {
                        ag = q.get(ab)
                    }
                } else {
                    if (ac.search(/\(\s*\)$/m) > -1) {
                        Z = ""
                    } else {
                        return ac
                    }
                }
            }
            Z = o(Z);
            var am = '<a href="' + Z + '"';
            if (ag != "") {
                ag = L(ag);
                ag = E(ag, "*_");
                am += ' title="' + ag + '"'
            }
            am += ">" + aa + "</a>";
            return am
        }

        function I(Z) {
            if (Z.indexOf("![") === -1) {
                return Z
            }
            Z = Z.replace(/(!\[(.*?)\][ ]?(?:\n[ ]*)?\[(.*?)\])()()()()/g, K);
            Z = Z.replace(/(!\[(.*?)\]\s?\([ \t]*()<?(\S+?)>?[ \t]*((['"])(.*?)\6[ \t]*)?\))/g, K);
            return Z
        }

        function L(Z) {
            return Z.replace(/>/g, "&gt;").replace(/</g, "&lt;").replace(/"/g, "&quot;").replace(/'/g, "&#39;")
        }

        function K(af, al, ak, aj, ai, ah, ae, ad) {
            var ac = al;
            var ab = ak;
            var aa = aj.toLowerCase();
            var Z = ai;
            var ag = ad;
            if (!ag) {
                ag = ""
            }
            if (Z == "") {
                if (aa == "") {
                    aa = ab.toLowerCase().replace(/ ?\n/g, " ")
                }
                Z = "#" + aa;
                if (B.get(aa) != undefined) {
                    Z = B.get(aa);
                    if (q.get(aa) != undefined) {
                        ag = q.get(aa)
                    }
                } else {
                    return ac
                }
            }
            ab = E(L(ab), "*_[]()");
            Z = E(Z, "*_");
            var am = '<img src="' + Z + '" alt="' + ab + '"';
            ag = L(ag);
            ag = E(ag, "*_");
            am += ' title="' + ag + '"';
            am += " />";
            return am
        }

        function k(Z) {
            Z = Z.replace(/^(.+)[ \t]*\n=+[ \t]*\n+/gm, function (aa, ab) {
                return "<h1>" + m(ab) + "</h1>\n\n"
            });
            Z = Z.replace(/^(.+)[ \t]*\n-+[ \t]*\n+/gm, function (ab, aa) {
                return "<h2>" + m(aa) + "</h2>\n\n"
            });
            Z = Z.replace(/^(\#{1,6})[ \t]*(.+?)[ \t]*\#*\n+/gm, function (aa, ad, ac) {
                var ab = ad.length;
                return "<h" + ab + ">" + m(ac) + "</h" + ab + ">\n\n"
            });
            return Z
        }

        function S(ab, Z) {
            ab += "~0";
            var aa = /^(([ ]{0,3}([*+-]|\d+[.])[ \t]+)[^\r]+?(~0|\n{2,}(?=\S)(?![ \t]*(?:[*+-]|\d+[.])[ \t]+)))/gm;
            if (F) {
                ab = ab.replace(aa, function (ae, ai, ah) {
                    var aj = ai;
                    var ag = (ah.search(/[*+-]/g) > -1) ? "ul" : "ol";
                    var af;
                    if (ag === "ol") {
                        af = parseInt(ah, 10)
                    }
                    var ad = n(aj, ag, Z);
                    ad = ad.replace(/\s+$/, "");
                    var ac = "<" + ag;
                    if (af && af !== 1) {
                        ac += ' start="' + af + '"'
                    }
                    ad = ac + ">" + ad + "</" + ag + ">\n";
                    return ad
                })
            } else {
                aa = /(\n\n|^\n?)(([ ]{0,3}([*+-]|\d+[.])[ \t]+)[^\r]+?(~0|\n{2,}(?=\S)(?![ \t]*(?:[*+-]|\d+[.])[ \t]+)))/g;
                ab = ab.replace(aa, function (ag, ak, aj, ah) {
                    var ai = ak;
                    var af = aj;
                    var ad = (ah.search(/[*+-]/g) > -1) ? "ul" : "ol";
                    var ae;
                    if (ad === "ol") {
                        ae = parseInt(ah, 10)
                    }
                    var al = n(af, ad);
                    var ac = "<" + ad;
                    if (ae && ae !== 1) {
                        ac += ' start="' + ae + '"'
                    }
                    al = ai + ac + ">\n" + al + "</" + ad + ">\n";
                    return al
                })
            }
            ab = ab.replace(/~0/, "");
            return ab
        }

        var u = {ol: "\\d+[.]", ul: "[*+-]"};

        function n(ab, aa, ae) {
            F++;
            ab = ab.replace(/\n{2,}$/, "\n");
            ab += "~0";
            var Z = u[aa];
            var ac = new RegExp("(^[ \\t]*)(" + Z + ")[ \\t]+([^\\r]+?(\\n+))(?=(~0|\\1(" + Z + ")[ \\t]+))", "gm");
            var ad = false;
            ab = ab.replace(ac, function (ag, ai, ah, af) {
                var al = af;
                var am = ai;
                var ak = /\n\n$/.test(al);
                var aj = ak || al.search(/\n{2,}/) > -1;
                if (aj || ad) {
                    al = h(z(al), true)
                } else {
                    al = S(z(al), true);
                    al = al.replace(/\n$/, "");
                    if (!ae) {
                        al = m(al)
                    }
                }
                ad = ak;
                return "<li>" + al + "</li>\n"
            });
            ab = ab.replace(/~0/g, "");
            F--;
            return ab
        }

        function w(Z) {
            Z += "~0";
            Z = Z.replace(/(?:\n\n|^\n?)((?:(?:[ ]{4}|\t).*\n+)+)(\n*[ ]{0,3}[^ \t\n]|(?=~0))/g, function (aa, ac, ab) {
                var ad = ac;
                var ae = ab;
                ad = G(z(ad));
                ad = P(ad);
                ad = ad.replace(/^\n+/g, "");
                ad = ad.replace(/\n+$/g, "");
                ad = "<pre><code>" + ad + "\n</code></pre>";
                return "\n\n" + ad + "\n\n" + ae
            });
            Z = Z.replace(/~0/, "");
            return Z
        }

        function x(Z) {
            Z = Z.replace(/(^|[^\\`])(`+)(?!`)([^\r]*?[^`])\2(?!`)/gm, function (ac, ae, ad, ab, aa) {
                var af = ab;
                af = af.replace(/^([ \t]*)/g, "");
                af = af.replace(/[ \t]*$/g, "");
                af = G(af);
                af = af.replace(/:\/\//g, "~P");
                return ae + "<code>" + af + "</code>"
            });
            return Z
        }

        function G(Z) {
            Z = Z.replace(/&/g, "&amp;");
            Z = Z.replace(/</g, "&lt;");
            Z = Z.replace(/>/g, "&gt;");
            Z = E(Z, "*_{}[]\\", false);
            return Z
        }

        function t(Z) {
            if (Z.indexOf("*") === -1 && Z.indexOf("_") === -1) {
                return Z
            }
            Z = W(Z);
            Z = Z.replace(/(^|[\W_])(?:(?!\1)|(?=^))(\*|_)\2(?=\S)([^\r]*?\S)\2\2(?!\2)(?=[\W_]|$)/g, "$1<strong>$3</strong>");
            Z = Z.replace(/(^|[\W_])(?:(?!\1)|(?=^))(\*|_)(?=\S)((?:(?!\2)[^\r])*?\S)\2(?!\2)(?=[\W_]|$)/g, "$1<em>$3</em>");
            return r(Z)
        }

        function v(Z) {
            if (Z.indexOf("*") === -1 && Z.indexOf("_") === -1) {
                return Z
            }
            Z = W(Z);
            Z = Z.replace(/(?=[^\r][*_]|[*_])(^|(?=\W__|(?!\*)[\W_]\*\*|\w\*\*\w)[^\r])(\*\*|__)(?!\2)(?=\S)((?:|[^\r]*?(?!\2)[^\r])(?=\S_|\w|\S\*\*(?:[\W_]|$)).)(?=__(?:\W|$)|\*\*(?:[^*]|$))\2/g, "$1<strong>$3</strong>");
            Z = Z.replace(/(?=[^\r][*_]|[*_])(^|(?=\W_|(?!\*)(?:[\W_]\*|\D\*(?=\w)\D))[^\r])(\*|_)(?!\2\2\2)(?=\S)((?:(?!\2)[^\r])*?(?=[^\s_]_|(?=\w)\D\*\D|[^\s*]\*(?:[\W_]|$)).)(?=_(?:\W|$)|\*(?:[^*]|$))\2/g, "$1<em>$3</em>");
            return r(Z)
        }

        function i(Z) {
            Z = Z.replace(/((^[ \t]*>[ \t]?.+\n(.+\n)*\n*)+)/gm, function (aa, ab) {
                var ac = ab;
                ac = ac.replace(/^[ \t]*>[ \t]?/gm, "~0");
                ac = ac.replace(/~0/g, "");
                ac = ac.replace(/^[ \t]+$/gm, "");
                ac = h(ac);
                ac = ac.replace(/(^|\n)/g, "$1  ");
                ac = ac.replace(/(\s*<pre>[^\r]+?<\/pre>)/gm, function (ad, ae) {
                    var af = ae;
                    af = af.replace(/^  /mg, "~0");
                    af = af.replace(/~0/g, "");
                    return af
                });
                return V("<blockquote>\n" + ac + "\n</blockquote>")
            });
            return Z
        }

        function O(ag, Z) {
            ag = ag.replace(/^\n+/g, "");
            ag = ag.replace(/\n+$/g, "");
            var ah = ag.split(/\n{2,}/g);
            var ae = [];
            var aa = /~K(\d+)K/;
            var ab = ah.length;
            for (var ac = 0; ac < ab; ac++) {
                var ad = ah[ac];
                if (aa.test(ad)) {
                    ae.push(ad)
                } else {
                    if (/\S/.test(ad)) {
                        ad = m(ad);
                        ad = ad.replace(/^([ \t]*)/g, "<p>");
                        ad += "</p>";
                        ae.push(ad)
                    }
                }
            }
            if (!Z) {
                ab = ae.length;
                for (var ac = 0; ac < ab; ac++) {
                    var af = true;
                    while (af) {
                        af = false;
                        ae[ac] = ae[ac].replace(/~K(\d+)K/g, function (ai, aj) {
                            af = true;
                            return e[aj]
                        })
                    }
                }
            }
            return ae.join("\n\n")
        }

        function H(Z) {
            Z = Z.replace(/&(?!#?[xX]?(?:[0-9a-fA-F]+|\w+);)/g, "&amp;");
            Z = Z.replace(/<(?![a-z\/?!]|~D)/gi, "&lt;");
            return Z
        }

        function M(Z) {
            Z = Z.replace(/\\(\\)/g, y);
            Z = Z.replace(/\\([`*_{}\[\]()>#+-.!])/g, y);
            return Z
        }

        var C = "[-A-Z0-9+&@#/%?=~_|[\\]()!:,.;]", X = "[-A-Z0-9+&@#/%=~_|[\\])]",
            N = new RegExp('(="|<)?\\b(https?|ftp)(://' + C + "*" + X + ")(?=$|\\W)", "gi"), U = new RegExp(X, "i");

        function f(ad, ag, ai, ac) {
            if (ag) {
                return ad
            }
            if (ac.charAt(ac.length - 1) !== ")") {
                return "<" + ai + ac + ">"
            }
            var ae = ac.match(/[()]/g);
            var Z = 0;
            for (var aa = 0; aa < ae.length; aa++) {
                if (ae[aa] === "(") {
                    if (Z <= 0) {
                        Z = 1
                    } else {
                        Z++
                    }
                } else {
                    Z--
                }
            }
            var ab = "";
            if (Z < 0) {
                var ah = new RegExp("\\){1," + (-Z) + "}$");
                ac = ac.replace(ah, function (aj) {
                    ab = aj;
                    return ""
                })
            }
            if (ab) {
                var af = ac.charAt(ac.length - 1);
                if (!U.test(af)) {
                    ab = af + ab;
                    ac = ac.substr(0, ac.length - 1)
                }
            }
            return "<" + ai + ac + ">" + ab
        }

        function R(aa) {
            aa = aa.replace(N, f);
            var Z = function (ad, ac) {
                var ab = o(ac);
                return '<a href="' + ab + '">' + l.plainLinkText(ac) + "</a>"
            };
            aa = aa.replace(/<((https?|ftp):[^'">\s]+)>/gi, Z);
            return aa
        }

        function T(Z) {
            Z = Z.replace(/~E(\d+)E/g, function (aa, ac) {
                var ab = parseInt(ac);
                return String.fromCharCode(ab)
            });
            return Z
        }

        function z(Z) {
            Z = Z.replace(/^(\t|[ ]{1,4})/gm, "~0");
            Z = Z.replace(/~0/g, "");
            return Z
        }

        function P(ac) {
            if (!/\t/.test(ac)) {
                return ac
            }
            var ab = ["    ", "   ", "  ", " "], aa = 0, Z;
            return ac.replace(/[\n\t]/g, function (ad, ae) {
                if (ad === "\n") {
                    aa = ae + 1;
                    return ad
                }
                Z = (ae - aa) % 4;
                aa = ae + 1;
                return ab[Z]
            })
        }

        function o(Z) {
            Z = L(Z);
            Z = E(Z, "*_:()[]");
            return Z
        }

        function E(ad, aa, ab) {
            var Z = "([" + aa.replace(/([\[\]\\])/g, "\\$1") + "])";
            if (ab) {
                Z = "\\\\" + Z
            }
            var ac = new RegExp(Z, "g");
            ad = ad.replace(ac, y);
            return ad
        }

        function y(Z, ab) {
            var aa = ab.charCodeAt(0);
            return "~E" + aa + "E"
        }
    }
})();
(function () {
    var a = {}, u = {}, m = {}, w = window.document, n = window.RegExp, g = window.navigator, q = {lineLength: 72},
        v = {
            isIE: /msie/.test(g.userAgent.toLowerCase()),
            isIE_5or6: /msie 6/.test(g.userAgent.toLowerCase()) || /msie 5/.test(g.userAgent.toLowerCase()),
            isOpera: /opera/.test(g.userAgent.toLowerCase())
        };
    var i = {
        bold: "Strong <strong> Ctrl+B",
        boldexample: "strong text",
        italic: "Emphasis <em> Ctrl+I",
        italicexample: "emphasized text",
        link: "Hyperlink <a> Ctrl+L",
        linkdescription: "enter link description here",
        linkdialog: '<p><b>Insert Hyperlink</b></p><p>http://example.com/ "optional title"</p>',
        linkname: null,
        quote: "Blockquote <blockquote> Ctrl+Q",
        quoteexample: "Blockquote",
        code: "Code Sample <pre><code> Ctrl+K",
        codeexample: "enter code here",
        image: "Image <img> Ctrl+G",
        imagedescription: "enter image description here",
        imagedialog: "<p><b>Insert Image</b></p><p>http://example.com/images/diagram.jpg \"optional title\"<br>Need <a href='http://www.google.com/search?q=free+image+hosting' target='_blank'>free image hosting?</a></p>",
        imagename: null,
        olist: "Numbered List <ol> Ctrl+O",
        ulist: "Bulleted List <ul> Ctrl+U",
        litem: "List item",
        heading: "Heading <h1>/<h2> Ctrl+H",
        headingexample: "Heading",
        more: "More contents <!--more--> Ctrl+M",
        fullscreen: "FullScreen Ctrl+J",
        exitFullscreen: "Exit FullScreen Ctrl+E",
        fullscreenUnsupport: "Sorry, the browser dont support fullscreen api",
        hr: "Horizontal Rule <hr> Ctrl+R",
        undo: "Undo - Ctrl+Z",
        redo: "Redo - Ctrl+Y",
        redomac: "Redo - Ctrl+Shift+Z",
        ok: "OK",
        cancel: "Cancel",
        help: "Markdown Editing Help"
    };
    var c = "http://";
    var d = "http://";
    Markdown.Editor = function (B, D, A) {
        A = A || {};
        if (typeof A.handler === "function") {
            A = {helpButton: A}
        }
        A.strings = A.strings || {};
        if (A.helpButton) {
            A.strings.help = A.strings.help || A.helpButton.title
        }
        var y = function (F) {
            var E = A.strings[F] || i[F];
            if ("imagename" == F || "linkname" == F) {
                A.strings[F] = null
            }
            return E
        };
        D = D || "";
        var x = this.hooks = new Markdown.HookCollection();
        x.addNoop("onPreviewRefresh");
        x.addNoop("postBlockquoteCreation");
        x.addFalse("insertImageDialog");
        x.addFalse("insertLinkDialog");
        x.addNoop("makeButton");
        x.addNoop("enterFullScreen");
        x.addNoop("enterFakeFullScreen");
        x.addNoop("exitFullScreen");
        this.getConverter = function () {
            return B
        };
        var C = this, z;
        this.run = function () {
            if (z) {
                return
            }
            z = new k(D);
            var G = new l(x, y);
            var E = new h(B, z, function () {
                x.onPreviewRefresh()
            });
            var H, F;
            if (!/\?noundo/.test(w.location.href)) {
                H = new e(function () {
                    E.refresh();
                    if (F) {
                        F.setUndoRedoButtonStates()
                    }
                }, z);
                this.textOperation = function (J) {
                    H.setCommandMode();
                    J();
                    C.refreshPreview()
                }
            }
            fullScreenManager = new o(x, y);
            F = new t(D, z, x, H, E, G, fullScreenManager, A.helpButton, y);
            F.setUndoRedoButtonStates();
            var I = C.refreshPreview = function () {
                E.refresh(true)
            };
            I()
        }
    };
    function s() {
    }

    s.prototype.findTags = function (y, A) {
        var x = this;
        var z;
        if (y) {
            z = a.extendRegExp(y, "", "$");
            this.before = this.before.replace(z, function (B) {
                x.startTag = x.startTag + B;
                return ""
            });
            z = a.extendRegExp(y, "^", "");
            this.selection = this.selection.replace(z, function (B) {
                x.startTag = x.startTag + B;
                return ""
            })
        }
        if (A) {
            z = a.extendRegExp(A, "", "$");
            this.selection = this.selection.replace(z, function (B) {
                x.endTag = B + x.endTag;
                return ""
            });
            z = a.extendRegExp(A, "^", "");
            this.after = this.after.replace(z, function (B) {
                x.endTag = B + x.endTag;
                return ""
            })
        }
    };
    s.prototype.trimWhitespace = function (y) {
        var x, A, z = this;
        if (y) {
            x = A = ""
        } else {
            x = function (B) {
                z.before += B;
                return ""
            };
            A = function (B) {
                z.after = B + z.after;
                return ""
            }
        }
        this.selection = this.selection.replace(/^(\s*)/, x).replace(/(\s*)$/, A)
    };
    s.prototype.skipLines = function (z, y, x) {
        if (z === undefined) {
            z = 1
        }
        if (y === undefined) {
            y = 1
        }
        z++;
        y++;
        var A;
        var B;
        if (navigator.userAgent.match(/Chrome/)) {
            "X".match(/()./)
        }
        this.selection = this.selection.replace(/(^\n*)/, "");
        this.startTag = this.startTag + n.$1;
        this.selection = this.selection.replace(/(\n*$)/, "");
        this.endTag = this.endTag + n.$1;
        this.startTag = this.startTag.replace(/(^\n*)/, "");
        this.before = this.before + n.$1;
        this.endTag = this.endTag.replace(/(\n*$)/, "");
        this.after = this.after + n.$1;
        if (this.before) {
            A = B = "";
            while (z--) {
                A += "\\n?";
                B += "\n"
            }
            if (x) {
                A = "\\n*"
            }
            this.before = this.before.replace(new n(A + "$", ""), B)
        }
        if (this.after) {
            A = B = "";
            while (y--) {
                A += "\\n?";
                B += "\n"
            }
            if (x) {
                A = "\\n*"
            }
            this.after = this.after.replace(new n(A, ""), B)
        }
    };
    function k(x) {
        this.buttonBar = w.getElementById("wmd-button-bar" + x);
        this.preview = w.getElementById("wmd-preview" + x);
        this.input = w.getElementById("text")
    }

    a.isVisible = function (x) {
        if (window.getComputedStyle) {
            return window.getComputedStyle(x, null).getPropertyValue("display") !== "none"
        } else {
            if (x.currentStyle) {
                return x.currentStyle.display !== "none"
            }
        }
    };
    a.addEvent = function (y, x, z) {
        if (y.attachEvent) {
            y.attachEvent("on" + x, z)
        } else {
            y.addEventListener(x, z, false)
        }
    };
    a.removeEvent = function (y, x, z) {
        if (y.detachEvent) {
            y.detachEvent("on" + x, z)
        } else {
            y.removeEventListener(x, z, false)
        }
    };
    a.fixEolChars = function (x) {
        x = x.replace(/\r\n/g, "\n");
        x = x.replace(/\r/g, "\n");
        return x
    };
    a.extendRegExp = function (z, B, y) {
        if (B === null || B === undefined) {
            B = ""
        }
        if (y === null || y === undefined) {
            y = ""
        }
        var A = z.toString();
        var x;
        A = A.replace(/\/([gim]*)$/, function (C, D) {
            x = D;
            return ""
        });
        A = A.replace(/(^\/|\/$)/g, "");
        A = B + A + y;
        return new n(A, x)
    };
    u.getTop = function (z, y) {
        var x = z.offsetTop;
        if (!y) {
            while (z = z.offsetParent) {
                x += z.offsetTop
            }
        }
        return x
    };
    u.getHeight = function (x) {
        return x.offsetHeight || x.scrollHeight
    };
    u.getWidth = function (x) {
        return x.offsetWidth || x.scrollWidth
    };
    u.getPageSize = function () {
        var y, z;
        var x, C;
        if (self.innerHeight && self.scrollMaxY) {
            y = w.body.scrollWidth;
            z = self.innerHeight + self.scrollMaxY
        } else {
            if (w.body.scrollHeight > w.body.offsetHeight) {
                y = w.body.scrollWidth;
                z = w.body.scrollHeight
            } else {
                y = w.body.offsetWidth;
                z = w.body.offsetHeight
            }
        }
        if (self.innerHeight) {
            x = self.innerWidth;
            C = self.innerHeight
        } else {
            if (w.documentElement && w.documentElement.clientHeight) {
                x = w.documentElement.clientWidth;
                C = w.documentElement.clientHeight
            } else {
                if (w.body) {
                    x = w.body.clientWidth;
                    C = w.body.clientHeight
                }
            }
        }
        var B = Math.max(y, x);
        var A = Math.max(z, C);
        return [B, A, x, C]
    };
    function e(J, G) {
        var M = this;
        var H = [];
        var E = 0;
        var D = "none";
        var y;
        var z;
        var C;
        var x = function (O, N) {
            if (D != O) {
                D = O;
                if (!N) {
                    A()
                }
            }
            if (!v.isIE || D != "moving") {
                z = setTimeout(F, 1)
            } else {
                C = null
            }
        };
        var F = function (N) {
            C = new b(G, N);
            z = undefined
        };
        this.setCommandMode = function () {
            D = "command";
            A();
            z = setTimeout(F, 0)
        };
        this.canUndo = function () {
            return E > 1
        };
        this.canRedo = function () {
            if (H[E + 1]) {
                return true
            }
            return false
        };
        this.undo = function () {
            if (M.canUndo()) {
                if (y) {
                    y.restore();
                    y = null
                } else {
                    H[E] = new b(G);
                    H[--E].restore();
                    if (J) {
                        J()
                    }
                }
            }
            D = "none";
            G.input.focus();
            F()
        };
        this.redo = function () {
            if (M.canRedo()) {
                H[++E].restore();
                if (J) {
                    J()
                }
            }
            D = "none";
            G.input.focus();
            F()
        };
        var A = function () {
            var N = C || new b(G);
            if (!N) {
                return false
            }
            if (D == "moving") {
                if (!y) {
                    y = N
                }
                return
            }
            if (y) {
                if (H[E - 1].text != y.text) {
                    H[E++] = y
                }
                y = null
            }
            H[E++] = N;
            H[E + 1] = null;
            if (J) {
                J()
            }
        };
        var I = function (N) {
            var P = false;
            if ((N.ctrlKey || N.metaKey) && !N.altKey) {
                var O = N.charCode || N.keyCode;
                var Q = String.fromCharCode(O);
                switch (Q.toLowerCase()) {
                    case"y":
                        M.redo();
                        P = true;
                        break;
                    case"z":
                        if (!N.shiftKey) {
                            M.undo()
                        } else {
                            M.redo()
                        }
                        P = true;
                        break
                }
            }
            if (P) {
                if (N.preventDefault) {
                    N.preventDefault()
                }
                if (window.event) {
                    window.event.returnValue = false
                }
                return
            }
        };
        var L = function (N) {
            if (!N.ctrlKey && !N.metaKey) {
                var O = N.keyCode;
                if ((O >= 33 && O <= 40) || (O >= 63232 && O <= 63235)) {
                    x("moving")
                } else {
                    if (O == 8 || O == 46 || O == 127) {
                        x("deleting")
                    } else {
                        if (O == 13) {
                            x("newlines")
                        } else {
                            if (O == 27) {
                                x("escape")
                            } else {
                                if ((O < 16 || O > 20) && O != 91) {
                                    x("typing")
                                }
                            }
                        }
                    }
                }
            }
        };
        var B = function () {
            a.addEvent(G.input, "keypress", function (O) {
                if ((O.ctrlKey || O.metaKey) && !O.altKey && (O.keyCode == 89 || O.keyCode == 90)) {
                    O.preventDefault()
                }
            });
            var N = function () {
                if (v.isIE || (C && C.text != G.input.value)) {
                    if (z == undefined) {
                        D = "paste";
                        A();
                        F()
                    }
                }
            };
            a.addEvent(G.input, "keydown", I);
            a.addEvent(G.input, "keydown", L);
            a.addEvent(G.input, "mousedown", function () {
                x("moving")
            });
            G.input.onpaste = N;
            G.input.ondrop = N
        };
        var K = function () {
            B();
            F(true);
            A()
        };
        K()
    }

    function b(y, A) {
        var x = this;
        var z = y.input;
        this.init = function () {
            if (!a.isVisible(z)) {
                return
            }
            if (!A && w.activeElement && w.activeElement !== z) {
                return
            }
            this.setInputAreaSelectionStartEnd();
            this.scrollTop = z.scrollTop;
            if (!this.text && z.selectionStart || z.selectionStart === 0) {
                this.text = z.value
            }
        };
        this.setInputAreaSelection = function () {
            if (!a.isVisible(z)) {
                return
            }
            if (z.selectionStart !== undefined && !v.isOpera) {
                z.focus();
                z.selectionStart = x.start;
                z.selectionEnd = x.end;
                z.scrollTop = x.scrollTop
            } else {
                if (w.selection) {
                    if (w.activeElement && w.activeElement !== z) {
                        return
                    }
                    z.focus();
                    var B = z.createTextRange();
                    B.moveStart("character", -z.value.length);
                    B.moveEnd("character", -z.value.length);
                    B.moveEnd("character", x.end);
                    B.moveStart("character", x.start);
                    B.select()
                }
            }
        };
        this.setInputAreaSelectionStartEnd = function () {
            if (!y.ieCachedRange && (z.selectionStart || z.selectionStart === 0)) {
                x.start = z.selectionStart;
                x.end = z.selectionEnd
            } else {
                if (w.selection) {
                    x.text = a.fixEolChars(z.value);
                    var E = y.ieCachedRange || w.selection.createRange();
                    var F = a.fixEolChars(E.text);
                    var D = "\x07";
                    var C = D + F + D;
                    E.text = C;
                    var G = a.fixEolChars(z.value);
                    E.moveStart("character", -C.length);
                    E.text = F;
                    x.start = G.indexOf(D);
                    x.end = G.lastIndexOf(D) - D.length;
                    var B = x.text.length - a.fixEolChars(z.value).length;
                    if (B) {
                        E.moveStart("character", -F.length);
                        while (B--) {
                            F += "\n";
                            x.end += 1
                        }
                        E.text = F
                    }
                    if (y.ieCachedRange) {
                        x.scrollTop = y.ieCachedScrollTop
                    }
                    y.ieCachedRange = null;
                    this.setInputAreaSelection()
                }
            }
        };
        this.restore = function () {
            if (x.text != undefined && x.text != z.value) {
                z.value = x.text
            }
            this.setInputAreaSelection();
            z.scrollTop = x.scrollTop
        };
        this.getChunks = function () {
            var B = new s();
            B.before = a.fixEolChars(x.text.substring(0, x.start));
            B.startTag = "";
            B.selection = a.fixEolChars(x.text.substring(x.start, x.end));
            B.endTag = "";
            B.after = a.fixEolChars(x.text.substring(x.end));
            B.scrollTop = x.scrollTop;
            return B
        };
        this.setChunks = function (B) {
            B.before = B.before + B.startTag;
            B.after = B.endTag + B.after;
            this.start = B.before.length;
            this.end = B.before.length + B.selection.length;
            this.text = B.before + B.selection + B.after;
            this.scrollTop = B.scrollTop
        };
        this.init()
    }

    function h(R, A, L) {
        var y = this;
        var F;
        var E;
        var O;
        var z = 3000;
        var G = "delayed";
        var C = function (T, U) {
            a.addEvent(T, "input", U);
            T.onpaste = U;
            T.ondrop = U;
            a.addEvent(T, "keypress", U);
            a.addEvent(T, "keydown", U)
        };
        var K = function () {
            var T = 0;
            if (window.innerHeight) {
                T = window.pageYOffset
            } else {
                if (w.documentElement && w.documentElement.scrollTop) {
                    T = w.documentElement.scrollTop
                } else {
                    if (w.body) {
                        T = w.body.scrollTop
                    }
                }
            }
            return T
        };
        var D = function () {
            if (!A.preview) {
                return
            }
            var V = A.input.value;
            if (V && V == O) {
                return
            } else {
                O = V
            }
            var U = new Date().getTime();
            V = R.makeHtml(V);
            var T = new Date().getTime();
            E = T - U;
            x(V)
        };
        var Q = function () {
            if (F) {
                clearTimeout(F);
                F = undefined
            }
            if (G !== "manual") {
                var T = 0;
                if (G === "delayed") {
                    T = E
                }
                if (T > z) {
                    T = z
                }
                F = setTimeout(D, T)
            }
        };
        var B = function (T) {
            if (T.scrollHeight <= T.clientHeight) {
                return 1
            }
            return T.scrollTop / (T.scrollHeight - T.clientHeight)
        };
        var S = function () {
            if (A.preview) {
                A.preview.scrollTop = (A.preview.scrollHeight - A.preview.clientHeight) * B(A.preview)
            }
        };
        this.refresh = function (T) {
            if (T) {
                O = "";
                D()
            } else {
                Q()
            }
        };
        this.processingTime = function () {
            return E
        };
        var H = true;
        var I = function (W) {
            var V = A.preview;
            var U = V.parentNode;
            var T = V.nextSibling;
            U.removeChild(V);
            V.innerHTML = W;
            if (!T) {
                U.appendChild(V)
            } else {
                U.insertBefore(V, T)
            }
        };
        var N = function (T) {
            A.preview.innerHTML = T
        };
        var J;
        var M = function (U) {
            if (J) {
                return J(U)
            }
            try {
                N(U);
                J = N
            } catch (T) {
                J = I;
                J(U)
            }
        };
        var x = function (V) {
            var T = u.getTop(A.input) - K();
            if (A.preview) {
                M(V);
                L()
            }
            S();
            if (H) {
                H = false;
                return
            }
            var U = u.getTop(A.input) - K();
            if (v.isIE) {
                setTimeout(function () {
                    window.scrollBy(0, U - T)
                }, 0)
            } else {
                window.scrollBy(0, U - T)
            }
        };
        var P = function () {
            C(A.input, Q);
            D();
            if (A.preview) {
                A.preview.scrollTop = 0
            }
        };
        P()
    }

    m.createBackground = function () {
        var y = w.createElement("div"), z = y.style;
        y.className = "wmd-prompt-background";
        z.position = "absolute";
        z.top = "0";
        z.zIndex = "1000";
        if (v.isIE) {
            z.filter = "alpha(opacity=50)"
        } else {
            z.opacity = "0.5"
        }
        var x = u.getPageSize();
        z.height = x[1] + "px";
        if (v.isIE) {
            z.left = w.documentElement.scrollLeft;
            z.width = w.documentElement.clientWidth
        } else {
            z.left = "0";
            z.width = "100%"
        }
        w.body.appendChild(y);
        return y
    };
    m.dialog = function (A, E, z, B) {
        var y;
        var D = function (F) {
            var G = (F.charCode || F.keyCode);
            if (G === 27) {
                C(true)
            }
        };
        var C = function (F) {
            a.removeEvent(w.body, "keydown", D);
            y.parentNode.removeChild(y);
            E(F);
            return false
        };
        var x = function () {
            y = w.createElement("div");
            y.className = "wmd-prompt-dialog";
            y.setAttribute("role", "dialog");
            var F = w.createElement("div");
            var H = w.createElement("form"), G = H.style;
            H.onsubmit = function () {
                return C(false)
            };
            y.appendChild(H);
            H.appendChild(F);
            if ("function" == typeof(A)) {
                A.call(this, F)
            } else {
                F.innerHTML = A
            }
            var J = w.createElement("button");
            J.type = "button";
            J.className = "btn btn-s primary";
            J.onclick = function () {
                return C(false)
            };
            J.innerHTML = z;
            var I = w.createElement("button");
            I.type = "button";
            I.className = "btn btn-s";
            I.onclick = function () {
                return C(true)
            };
            I.innerHTML = B;
            H.appendChild(J);
            H.appendChild(I);
            a.addEvent(w.body, "keydown", D);
            w.body.appendChild(y)
        };
        setTimeout(function () {
            x()
        }, 0)
    };
    m.prompt = function (C, G, A, y, E) {
        var x;
        var z;
        if (G === undefined) {
            G = ""
        }
        var B = function (H) {
            var I = (H.charCode || H.keyCode);
            if (I === 27) {
                D(true)
            }
        };
        var D = function (H) {
            a.removeEvent(w.body, "keydown", B);
            var I = z.value;
            if (H) {
                I = null
            } else {
                I = I.replace(/^http:\/\/(https?|ftp):\/\//, "$1://");
                if (!/^(?:https?|ftp):\/\//.test(I)) {
                    I = "http://" + I
                }
            }
            x.parentNode.removeChild(x);
            A(I);
            return false
        };
        var F = function () {
            x = w.createElement("div");
            x.className = "wmd-prompt-dialog";
            x.setAttribute("role", "dialog");
            var H = w.createElement("div");
            H.innerHTML = C;
            x.appendChild(H);
            var J = w.createElement("form"), I = J.style;
            J.onsubmit = function () {
                return D(false)
            };
            x.appendChild(J);
            z = w.createElement("input");
            z.type = "text";
            z.value = G;
            J.appendChild(z);
            var L = w.createElement("button");
            L.type = "button";
            L.className = "btn btn-s primary";
            L.onclick = function () {
                return D(false)
            };
            L.innerHTML = y;
            var K = w.createElement("button");
            K.type = "button";
            K.className = "btn btn-s";
            K.onclick = function () {
                return D(true)
            };
            K.innerHTML = E;
            J.appendChild(L);
            J.appendChild(K);
            a.addEvent(w.body, "keydown", B);
            w.body.appendChild(x)
        };
        setTimeout(function () {
            F();
            var I = G.length;
            if (z.selectionStart !== undefined) {
                z.selectionStart = 0;
                z.selectionEnd = I
            } else {
                if (z.createTextRange) {
                    var H = z.createTextRange();
                    H.collapse(false);
                    H.moveStart("character", -I);
                    H.moveEnd("character", I);
                    H.select()
                }
            }
            z.focus()
        }, 0)
    };
    function t(K, E, L, N, A, J, I, D, y) {
        var C = E.input, G = {};
        B();
        var F = "keydown";
        if (v.isOpera) {
            F = "keypress"
        }
        a.addEvent(C, F, function (P) {
            if ((P.ctrlKey || P.metaKey) && !P.altKey && !P.shiftKey) {
                var R = P.charCode || P.keyCode;
                var O = String.fromCharCode(R).toLowerCase();
                switch (O) {
                    case"b":
                        M(G.bold);
                        break;
                    case"i":
                        M(G.italic);
                        break;
                    case"l":
                        M(G.link);
                        break;
                    case"q":
                        M(G.quote);
                        break;
                    case"k":
                        M(G.code);
                        break;
                    case"g":
                        M(G.image);
                        break;
                    case"o":
                        M(G.olist);
                        break;
                    case"u":
                        M(G.ulist);
                        break;
                    case"m":
                        M(G.more);
                        break;
                    case"j":
                        M(G.fullscreen);
                        break;
                    case"e":
                        M(G.exitFullscreen);
                        break;
                    case"h":
                        M(G.heading);
                        break;
                    case"r":
                        M(G.hr);
                        break;
                    case"y":
                        M(G.redo);
                        break;
                    case"z":
                        if (P.shiftKey) {
                            M(G.redo)
                        } else {
                            M(G.undo)
                        }
                        break;
                    default:
                        return
                }
                if (P.preventDefault) {
                    P.preventDefault()
                }
                if (window.event) {
                    window.event.returnValue = false
                }
            } else {
                if (P.keyCode == 9 && window.fullScreenEntered) {
                    var Q = {};
                    Q.textOp = z("doTab");
                    M(Q);
                    if (P.preventDefault) {
                        P.preventDefault()
                    }
                    if (window.event) {
                        window.event.returnValue = false
                    }
                }
            }
        });
        a.addEvent(C, "keyup", function (P) {
            if (P.shiftKey && !P.ctrlKey && !P.metaKey) {
                var Q = P.charCode || P.keyCode;
                if (Q === 13) {
                    var O = {};
                    O.textOp = z("doAutoindent");
                    M(O)
                }
            }
        });
        if (v.isIE) {
            a.addEvent(C, "keydown", function (O) {
                var P = O.keyCode;
                if (P === 27) {
                    return false
                }
            })
        }
        function M(P) {
            C.focus();
            if (P.textOp) {
                if (N) {
                    N.setCommandMode()
                }
                var R = new b(E);
                if (!R) {
                    return
                }
                var S = R.getChunks();
                var O = function () {
                    C.focus();
                    if (S) {
                        R.setChunks(S)
                    }
                    R.restore();
                    A.refresh()
                };
                var Q = P.textOp(S, O);
                if (!Q) {
                    O()
                }
            }
            if (P.execute) {
                P.execute(N)
            }
        }

        function x(O, Q) {
            var R = "0px";
            var T = "-20px";
            var P = "-40px";
            var S = O.getElementsByTagName("span")[0];
            if (Q) {
                S.style.backgroundPosition = O.XShift + " " + R;
                O.onmouseover = function () {
                    S.style.backgroundPosition = this.XShift + " " + P
                };
                O.onmouseout = function () {
                    S.style.backgroundPosition = this.XShift + " " + R
                };
                if (v.isIE) {
                    O.onmousedown = function () {
                        if (w.activeElement && w.activeElement !== E.input) {
                            return
                        }
                        E.ieCachedRange = document.selection.createRange();
                        E.ieCachedScrollTop = E.input.scrollTop
                    }
                }
                if (!O.isHelp) {
                    O.onclick = function () {
                        if (this.onmouseout) {
                            this.onmouseout()
                        }
                        M(this);
                        return false
                    }
                }
            } else {
                S.style.backgroundPosition = O.XShift + " " + T;
                O.onmouseover = O.onmouseout = O.onclick = function () {
                }
            }
        }

        function z(O) {
            if (typeof O === "string") {
                O = J[O]
            }
            return function () {
                O.apply(J, arguments)
            }
        }

        function B() {
            console.log("buttonRow");
            var S = E.buttonBar;
            var W = "0px";
            var R = "-20px";
            var U = "-40px";
            var V = document.createElement("ul");
            V.id = "wmd-button-row" + K;
            V.className = "wmd-button-row";
            V = S.appendChild(V);
            var O = 0;
            var X = function (ae, ac, ab, ad) {
                var aa = document.createElement("li");
                aa.className = "wmd-button";
                aa.style.left = O + "px";
                O += 25;
                var Z = document.createElement("span");
                aa.id = ae + K;
                aa.appendChild(Z);
                aa.title = ac;
                aa.XShift = ab;
                if (ad) {
                    aa.textOp = ad
                }
                x(aa, true);
                V.appendChild(aa);
                return aa
            };
            var Q = function (aa) {
                var Z = document.createElement("li");
                Z.className = "wmd-spacer wmd-spacer" + aa;
                Z.id = "wmd-spacer" + aa + K;
                V.appendChild(Z);
                O += 25
            };
            G.bold = X("wmd-bold-button", y("bold"), "0px", z("doBold"));
            G.italic = X("wmd-italic-button", y("italic"), "-20px", z("doItalic"));
            Q(1);
            G.link = X("wmd-link-button", y("link"), "-40px", z(function (Z, aa) {
                return this.doLinkOrImage(Z, aa, false)
            }));
            G.quote = X("wmd-quote-button", y("quote"), "-60px", z("doBlockquote"));
            G.code = X("wmd-code-button", y("code"), "-80px", z("doCode"));
            G.image = X("wmd-image-button", y("image"), "-100px", z(function (Z, aa) {
                return this.doLinkOrImage(Z, aa, true)
            }));
            Q(2);
            G.olist = X("wmd-olist-button", y("olist"), "-120px", z(function (Z, aa) {
                this.doList(Z, aa, true)
            }));
            G.ulist = X("wmd-ulist-button", y("ulist"), "-140px", z(function (Z, aa) {
                this.doList(Z, aa, false)
            }));
            G.heading = X("wmd-heading-button", y("heading"), "-160px", z("doHeading"));
            G.hr = X("wmd-hr-button", y("hr"), "-180px", z("doHorizontalRule"));
            G.more = X("wmd-more-button", y("more"), "-280px", z("doMore"));
            Q(3);
            G.undo = X("wmd-undo-button", y("undo"), "-200px", null);
            G.undo.execute = function (Z) {
                if (Z) {
                    Z.undo()
                }
            };
            var Y = /win/.test(g.platform.toLowerCase()) ? y("redo") : y("redomac");
            G.redo = X("wmd-redo-button", Y, "-220px", null);
            G.redo.execute = function (Z) {
                if (Z) {
                    Z.redo()
                }
            };
            Q(4);
            G.fullscreen = X("wmd-fullscreen-button", y("fullscreen"), "-240px", null);
            G.fullscreen.execute = function () {
                I.doFullScreen(G, true)
            };
            G.exitFullscreen = X("wmd-exit-fullscreen-button", y("exitFullscreen"), "-260px", null);
            G.exitFullscreen.style.display = "none";
            G.exitFullscreen.execute = function () {
                I.doFullScreen(G, false)
            };
            L.makeButton(G, X, z, m);
            if (D) {
                var T = document.createElement("li");
                var P = document.createElement("span");
                T.appendChild(P);
                T.className = "wmd-button wmd-help-button";
                T.id = "wmd-help-button" + K;
                T.XShift = "-300px";
                T.isHelp = true;
                T.style.right = "0px";
                T.title = y("help");
                T.onclick = D.handler;
                x(T, true);
                V.appendChild(T);
                G.help = T
            }
            H()
        }

        function H() {
            if (N) {
                x(G.undo, N.canUndo());
                x(G.redo, N.canRedo())
            }
        }

        this.setUndoRedoButtonStates = H
    }

    function l(y, x) {
        this.hooks = y;
        this.getString = x
    }

    var f = l.prototype;
    f.prefixes = "(?:\\s{4,}|\\s*>|\\s*-\\s+|\\s*\\d+\\.|=|\\+|-|_|\\*|#|\\s*\\[[^\n]]+\\]:)";
    f.unwrap = function (y) {
        var x = new n("([^\\n])\\n(?!(\\n|" + this.prefixes + "))", "g");
        y.selection = y.selection.replace(x, "$1 $2")
    };
    f.wrap = function (y, x) {
        this.unwrap(y);
        var A = new n("(.{1," + x + "})( +|$\\n?)", "gm"), z = this;
        y.selection = y.selection.replace(A, function (B, C) {
            if (new n("^" + z.prefixes, "").test(B)) {
                return B
            }
            return C + "\n"
        });
        y.selection = y.selection.replace(/\s+$/, "")
    };
    f.doBold = function (x, y) {
        return this.doBorI(x, y, 2, this.getString("boldexample"))
    };
    f.doItalic = function (x, y) {
        return this.doBorI(x, y, 1, this.getString("italicexample"))
    };
    f.doBorI = function (D, B, C, x) {
        D.trimWhitespace();
        D.selection = D.selection.replace(/\n{2,}/g, "\n");
        var A = /(\**$)/.exec(D.before)[0];
        var y = /(^\**)/.exec(D.after)[0];
        var E = Math.min(A.length, y.length);
        if ((E >= C) && (E != 2 || C != 1)) {
            D.before = D.before.replace(n("[*]{" + C + "}$", ""), "");
            D.after = D.after.replace(n("^[*]{" + C + "}", ""), "")
        } else {
            if (!D.selection && y) {
                D.after = D.after.replace(/^([*_]*)/, "");
                D.before = D.before.replace(/(\s?)$/, "");
                var z = n.$1;
                D.before = D.before + y + z
            } else {
                if (!D.selection && !y) {
                    D.selection = x
                }
                var F = C <= 1 ? "*" : "**";
                D.before = D.before + F;
                D.after = F + D.after
            }
        }
        return
    };
    f.stripLinkDefs = function (y, x) {
        y = y.replace(/^[ ]{0,3}\[(\d+)\]:[ \t]*\n?[ \t]*<?(\S+?)>?[ \t]*\n?[ \t]*(?:(\n*)["(](.+?)[")][ \t]*)?(?:\n+|$)/gm, function (C, D, z, A, B) {
            x[D] = C.replace(/\s*$/, "");
            if (A) {
                x[D] = C.replace(/["(](.+?)[")]$/, "");
                return A + B
            }
            return ""
        });
        return y
    };
    f.addLinkDef = function (E, A) {
        var x = 0;
        var z = {};
        E.before = this.stripLinkDefs(E.before, z);
        E.selection = this.stripLinkDefs(E.selection, z);
        E.after = this.stripLinkDefs(E.after, z);
        var y = "";
        var D = /(\[)((?:\[[^\]]*\]|[^\[\]])*)(\][ ]?(?:\n[ ]*)?\[)(\d+)(\])/g;
        var C = function (G) {
            x++;
            G = G.replace(/^[ ]{0,3}\[(\d+)\]:/, "  [" + x + "]:");
            y += "\n" + G
        };
        var B = function (H, K, I, J, L, G) {
            I = I.replace(D, B);
            if (z[L]) {
                C(z[L]);
                return K + I + J + x + G
            }
            return H
        };
        E.before = E.before.replace(D, B);
        if (A) {
            C(A)
        } else {
            E.selection = E.selection.replace(D, B)
        }
        var F = x;
        E.after = E.after.replace(D, B);
        if (E.after) {
            E.after = E.after.replace(/\n*$/, "")
        }
        if (!E.after) {
            E.selection = E.selection.replace(/\n*$/, "")
        }
        E.after += "\n\n" + y;
        return F
    };
    function p(x) {
        return x.replace(/^\s*(.*?)(?:\s+"(.+)")?\s*$/, function (z, y, B) {
            var A = false;
            y = y.replace(/%(?:[\da-fA-F]{2})|\?|\+|[^\w\d-./[\]]/g, function (C) {
                if (C.length === 3 && C.charAt(0) == "%") {
                    return C.toUpperCase()
                }
                switch (C) {
                    case"?":
                        A = true;
                        return "?";
                        break;
                    case"+":
                        if (A) {
                            return "%20"
                        }
                        break
                }
                return encodeURI(C)
            });
            if (B) {
                B = B.trim ? B.trim() : B.replace(/^\s*/, "").replace(/\s*$/, "");
                B = B.replace(/"/g, "quot;").replace(/\(/g, "&#40;").replace(/\)/g, "&#41;").replace(/</g, "&lt;").replace(/>/g, "&gt;")
            }
            return B ? y + ' "' + B + '"' : y
        })
    }

    f.doLinkOrImage = function (x, z, C) {
        x.trimWhitespace();
        x.findTags(/\s*!?\[/, /\][ ]?(?:\n[ ]*)?(\[.*?\])?/);
        var y;
        if (x.endTag.length > 1 && x.startTag.length > 0) {
            x.startTag = x.startTag.replace(/!?\[/, "");
            x.endTag = "";
            this.addLinkDef(x, null)
        } else {
            x.selection = x.startTag + x.selection + x.endTag;
            x.startTag = x.endTag = "";
            if (/\n\n/.test(x.selection)) {
                this.addLinkDef(x, null);
                return
            }
            var A = this;
            var B = function (G) {
                y.parentNode.removeChild(y);
                if (G !== null) {
                    x.selection = (" " + x.selection).replace(/([^\\](?:\\\\)*)(?=[[\]])/g, "$1\\").substr(1);
                    var F = " [999]: " + p(G);
                    var E = A.addLinkDef(x, F);
                    x.startTag = C ? "![" : "[";
                    x.endTag = "][" + E + "]";
                    if (!x.selection) {
                        if (C) {
                            var H = A.getString("imagename");
                            x.selection = H || A.getString("imagedescription")
                        } else {
                            var D = A.getString("linkname");
                            x.selection = D || A.getString("linkdescription")
                        }
                    }
                }
                z()
            };
            y = m.createBackground();
            if (C) {
                if (!this.hooks.insertImageDialog(B)) {
                    m.prompt(this.getString("imagedialog"), c, B, this.getString("ok"), this.getString("cancel"))
                }
            } else {
                if (!this.hooks.insertLinkDialog(B)) {
                    m.prompt(this.getString("linkdialog"), d, B, this.getString("ok"), this.getString("cancel"))
                }
            }
            return true
        }
    };
    f.doAutoindent = function (z, A) {
        var y = this, x = false;
        z.before = z.before.replace(/(\n|^)[ ]{0,3}([*+-]|\d+[.])[ \t]*\n$/, "\n\n");
        z.before = z.before.replace(/(\n|^)[ ]{0,3}>[ \t]*\n$/, "\n\n");
        z.before = z.before.replace(/(\n|^)[ \t]+\n$/, "\n\n");
        if (!z.selection && !/^[ \t]*(?:\n|$)/.test(z.after)) {
            z.after = z.after.replace(/^[^\n]*/, function (B) {
                z.selection = B;
                return ""
            });
            x = true
        }
        if (/(\n|^)[ ]{0,3}([*+-]|\d+[.])[ \t]+.*\n$/.test(z.before)) {
            if (y.doList) {
                y.doList(z)
            }
        }
        if (/(\n|^)[ ]{0,3}>[ \t]+.*\n$/.test(z.before)) {
            if (y.doBlockquote) {
                y.doBlockquote(z)
            }
        }
        if (/(\n|^)(\t|[ ]{4,}).*\n$/.test(z.before)) {
            if (y.doCode) {
                y.doCode(z)
            }
        }
        if (x) {
            z.after = z.selection + z.after;
            z.selection = ""
        }
    };
    f.doBlockquote = function (E, z) {
        E.selection = E.selection.replace(/^(\n*)([^\r]+?)(\n*)$/, function (K, J, I, H) {
            E.before += J;
            E.after = H + E.after;
            return I
        });
        E.before = E.before.replace(/(>[ \t]*)$/, function (I, H) {
            E.selection = H + E.selection;
            return ""
        });
        E.selection = E.selection.replace(/^(\s|>)+$/, "");
        E.selection = E.selection || this.getString("quoteexample");
        var D = "", C = "", F;
        if (E.before) {
            var G = E.before.replace(/\n$/, "").split("\n");
            var B = false;
            for (var A = 0; A < G.length; A++) {
                var y = false;
                F = G[A];
                B = B && F.length > 0;
                if (/^>/.test(F)) {
                    y = true;
                    if (!B && F.length > 1) {
                        B = true
                    }
                } else {
                    if (/^[ \t]*$/.test(F)) {
                        y = true
                    } else {
                        y = B
                    }
                }
                if (y) {
                    D += F + "\n"
                } else {
                    C += D + F;
                    D = "\n"
                }
            }
            if (!/(^|\n)>/.test(D)) {
                C += D;
                D = ""
            }
        }
        E.startTag = D;
        E.before = C;
        if (E.after) {
            E.after = E.after.replace(/^\n?/, "\n")
        }
        E.after = E.after.replace(/^(((\n|^)(\n[ \t]*)*>(.+\n)*.*)+(\n[ \t]*)*)/, function (H) {
            E.endTag = H;
            return ""
        });
        var x = function (I) {
            var H = I ? "> " : "";
            if (E.startTag) {
                E.startTag = E.startTag.replace(/\n((>|\s)*)\n$/, function (K, J) {
                    return "\n" + J.replace(/^[ ]{0,3}>?[ \t]*$/gm, H) + "\n"
                })
            }
            if (E.endTag) {
                E.endTag = E.endTag.replace(/^\n((>|\s)*)\n/, function (K, J) {
                    return "\n" + J.replace(/^[ ]{0,3}>?[ \t]*$/gm, H) + "\n"
                })
            }
        };
        if (/^(?![ ]{0,3}>)/m.test(E.selection)) {
            this.wrap(E, q.lineLength - 2);
            E.selection = E.selection.replace(/^/gm, "> ");
            x(true);
            E.skipLines()
        } else {
            E.selection = E.selection.replace(/^[ ]{0,3}> ?/gm, "");
            this.unwrap(E);
            x(false);
            if (!/^(\n|^)[ ]{0,3}>/.test(E.selection) && E.startTag) {
                E.startTag = E.startTag.replace(/\n{0,2}$/, "\n\n")
            }
            if (!/(\n|^)[ ]{0,3}>.*$/.test(E.selection) && E.endTag) {
                E.endTag = E.endTag.replace(/^\n{0,2}/, "\n\n")
            }
        }
        E.selection = this.hooks.postBlockquoteCreation(E.selection);
        if (!/\n/.test(E.selection)) {
            E.selection = E.selection.replace(/^(> *)/, function (H, I) {
                E.startTag += I;
                return ""
            })
        }
    };
    f.doCode = function (x, y) {
        var A = /\S[ ]*$/.test(x.before);
        var C = /^[ ]*\S/.test(x.after);
        if ((!C && !A) || /\n/.test(x.selection)) {
            x.before = x.before.replace(/[ ]{4}$/, function (D) {
                x.selection = D + x.selection;
                return ""
            });
            var B = 1;
            var z = 1;
            if (/(\n|^)(\t|[ ]{4,}).*\n$/.test(x.before)) {
                B = 0
            }
            if (/^\n(\t|[ ]{4,})/.test(x.after)) {
                z = 0
            }
            x.skipLines(B, z);
            if (!x.selection) {
                x.startTag = "    ";
                x.selection = this.getString("codeexample")
            } else {
                if (/^[ ]{0,3}\S/m.test(x.selection)) {
                    if (/\n/.test(x.selection)) {
                        x.selection = x.selection.replace(/^/gm, "    ")
                    } else {
                        x.before += "    "
                    }
                } else {
                    x.selection = x.selection.replace(/^(?:[ ]{4}|[ ]{0,3}\t)/gm, "")
                }
            }
        } else {
            x.trimWhitespace();
            x.findTags(/`/, /`/);
            if (!x.startTag && !x.endTag) {
                x.startTag = x.endTag = "`";
                if (!x.selection) {
                    x.selection = this.getString("codeexample")
                }
            } else {
                if (x.endTag && !x.startTag) {
                    x.before += x.endTag;
                    x.endTag = ""
                } else {
                    x.startTag = x.endTag = ""
                }
            }
        }
    };
    f.doList = function (I, B, A) {
        var K = /(\n|^)(([ ]{0,3}([*+-]|\d+[.])[ \t]+.*)(\n.+|\n{2,}([*+-].*|\d+[.])[ \t]+.*|\n{2,}[ \t]+\S.*)*)\n*$/;
        var J = /^\n*(([ ]{0,3}([*+-]|\d+[.])[ \t]+.*)(\n.+|\n{2,}([*+-].*|\d+[.])[ \t]+.*|\n{2,}[ \t]+\S.*)*)\n*/;
        var x = "-";
        var F = 1;
        var D = function () {
            var L;
            if (A) {
                L = " " + F + ". ";
                F++
            } else {
                L = " " + x + " "
            }
            return L
        };
        var E = function (L) {
            if (A === undefined) {
                A = /^\s*\d/.test(L)
            }
            L = L.replace(/^[ ]{0,3}([*+-]|\d+[.])\s/gm, function (M) {
                return D()
            });
            return L
        };
        I.findTags(/(\n|^)*[ ]{0,3}([*+-]|\d+[.])\s+/, null);
        if (I.before && !/\n$/.test(I.before) && !/^\n/.test(I.startTag)) {
            I.before += I.startTag;
            I.startTag = ""
        }
        if (I.startTag) {
            var z = /\d+[.]/.test(I.startTag);
            I.startTag = "";
            I.selection = I.selection.replace(/\n[ ]{4}/g, "\n");
            this.unwrap(I);
            I.skipLines();
            if (z) {
                I.after = I.after.replace(J, E)
            }
            if (A == z) {
                return
            }
        }
        var C = 1;
        I.before = I.before.replace(K, function (L) {
            if (/^\s*([*+-])/.test(L)) {
                x = n.$1
            }
            C = /[^\n]\n\n[^\n]/.test(L) ? 1 : 0;
            return E(L)
        });
        if (!I.selection) {
            I.selection = this.getString("litem")
        }
        var G = D();
        var y = 1;
        I.after = I.after.replace(J, function (L) {
            y = /[^\n]\n\n[^\n]/.test(L) ? 1 : 0;
            return E(L)
        });
        I.trimWhitespace(true);
        I.skipLines(C, y, true);
        I.startTag = G;
        var H = G.replace(/./g, " ");
        this.wrap(I, q.lineLength - H.length);
        I.selection = I.selection.replace(/\n/g, "\n" + H)
    };
    f.doHeading = function (z, A) {
        z.selection = z.selection.replace(/\s+/g, " ");
        z.selection = z.selection.replace(/(^\s+|\s+$)/g, "");
        if (!z.selection) {
            z.startTag = "## ";
            z.selection = this.getString("headingexample");
            z.endTag = " ##";
            return
        }
        var B = 0;
        z.findTags(/#+[ ]*/, /[ ]*#+/);
        if (/#+/.test(z.startTag)) {
            B = n.lastMatch.length
        }
        z.startTag = z.endTag = "";
        z.findTags(null, /\s?(-+|=+)/);
        if (/=+/.test(z.endTag)) {
            B = 1
        }
        if (/-+/.test(z.endTag)) {
            B = 2
        }
        z.startTag = z.endTag = "";
        z.skipLines(1, 1);
        var C = B == 0 ? 2 : B - 1;
        if (C > 0) {
            var y = C >= 2 ? "-" : "=";
            var x = z.selection.length;
            if (x > q.lineLength) {
                x = q.lineLength
            }
            z.endTag = "\n";
            while (x--) {
                z.endTag += y
            }
        }
    };
    f.doHorizontalRule = function (x, y) {
        x.startTag = "----------\n";
        x.selection = "";
        x.skipLines(2, 1, true)
    };
    f.doMore = function (x, y) {
        x.startTag = "<!--more-->\n\n";
        x.selection = "";
        x.skipLines(2, 0, true)
    };
    f.doTab = function (x, y) {
        x.startTag = "    ";
        x.selection = ""
    };
    function o(x, y) {
        this.fullScreenBind = false;
        this.hooks = x;
        this.getString = y;
        this.isFakeFullScreen = false
    }

    function j() {
        var y = {
            fullScreenChange: ["onfullscreenchange", "onwebkitfullscreenchange", "onmozfullscreenchange", "onmsfullscreenchange"],
            requestFullscreen: ["requestFullscreen", "webkitRequestFullScreen", "mozRequestFullScreen", "msRequestFullScreen"],
            cancelFullscreen: ["cancelFullscreen", "exitFullScreen", "webkitCancelFullScreen", "mozCancelFullScreen", "msCancelFullScreen"]
        }, z = {};
        for (var A in y) {
            var x = y[A].length, C = false;
            for (var B = 0; B < x; B++) {
                var D = y[A][B];
                if ("undefined" != typeof(document[D]) || "undefined" != typeof(document.body[D])) {
                    z[A] = D;
                    C = true;
                    break
                }
            }
            if (!C) {
                return false
            }
        }
        return z
    }

    function r() {
        return document.fullScreen || document.mozFullScreen || document.webkitIsFullScreen || document.msIsFullScreen
    }

    o.prototype.doFullScreen = function (z, A) {
        var x = j(), y = this;
        if (!x) {
            alert(y.getString("fullscreenUnsupport"));
            return false
        }
        if (!this.fullScreenBind) {
            a.addEvent(document, x.fullScreenChange.substring(2), function () {
                if (!r()) {
                    z.fullscreen.style.display = "";
                    z.exitFullscreen.style.display = "none";
                    y.hooks.exitFullScreen()
                } else {
                    z.fullscreen.style.display = "none";
                    z.exitFullscreen.style.display = "";
                    y.hooks.enterFullScreen()
                }
            });
            this.fullScreenBind = true
        }
        if (A) {
            if (y.isFakeFullScreen) {
                document.body[x.requestFullscreen]("webkitRequestFullScreen" == x.requestFullscreen ? Element.ALLOW_KEYBOARD_INPUT : null);
                y.isFakeFullScreen = false
            } else {
                if (!r()) {
                    z.exitFullscreen.style.display = "";
                    y.hooks.enterFakeFullScreen();
                    y.isFakeFullScreen = true
                }
            }
            window.fullScreenEntered = true
        } else {
            if (y.isFakeFullScreen) {
                z.exitFullscreen.style.display = "none";
                y.hooks.exitFullScreen()
            } else {
                if (r()) {
                    document[x.cancelFullscreen]()
                }
            }
            y.isFakeFullScreen = false;
            window.fullScreenEntered = false
        }
    }
})();