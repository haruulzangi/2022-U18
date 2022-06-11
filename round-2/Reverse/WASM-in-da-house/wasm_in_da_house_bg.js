import * as wasm from './wasm_warmup_bg.wasm';

const lTextDecoder = typeof TextDecoder === 'undefined' ? (0, module.require)('util').TextDecoder : TextDecoder;

let cachedTextDecoder = new lTextDecoder('utf-8', { ignoreBOM: true, fatal: true });

cachedTextDecoder.decode();

let cachegetUint8Memory0 = null;
function getUint8Memory0() {
    if (cachegetUint8Memory0 === null || cachegetUint8Memory0.buffer !== wasm.memory.buffer) {
        cachegetUint8Memory0 = new Uint8Array(wasm.memory.buffer);
    }
    return cachegetUint8Memory0;
}

function getStringFromWasm0(ptr, len) {
    return cachedTextDecoder.decode(getUint8Memory0().subarray(ptr, ptr + len));
}
/**
*/
export function flag5() {
    wasm.flag5();
}

/**
*/
export function flag1() {
    wasm.flag1();
}

/**
*/
export function flag3() {
    wasm.flag3();
}

/**
*/
export function flag2() {
    wasm.flag2();
}

/**
*/
export function flag4() {
    wasm.flag4();
}

export function __wbg_alert_af3e395a6293c0d5(arg0, arg1) {
    alert(getStringFromWasm0(arg0, arg1));
};

