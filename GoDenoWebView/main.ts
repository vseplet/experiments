import "./source/services/serve.ts";

const worker = new Worker(
  new URL(import.meta.dirname + "/source/services/serve.ts", import.meta.url)
    .href,
  {
    type: "module",
  },
);

const lib = Deno.dlopen(
  import.meta.dirname + "/static/libexample.dylib",
  {
    openWebView: {
      parameters: [],
      result: "void",
      nonblocking: false,
    },
  },
);

lib.symbols.openWebView();
