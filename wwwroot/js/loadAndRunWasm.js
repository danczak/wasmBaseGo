document.addEventListener('contextmenu', event => event.preventDefault());
(async function loadAndRunGoWasm() {
    const go = new Go();
    const buffer = await (await fetch("main.wasm")).arrayBuffer();
    const result = await WebAssembly.instantiate(buffer, go.importObject);
    go.run(result.instance)
})()

