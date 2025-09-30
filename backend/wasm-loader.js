// WASM Worker Loader for Cloudflare Workers
// This file loads and runs the Go WASM binary

import wasmData from './dist/worker.wasm';

// Go WebAssembly runtime (simplified version)
class Go {
  constructor() {
    this.argv = [];
    this.env = {};
    this.exit = (code) => {
      if (code !== 0) {
        console.warn('exit code:', code);
      }
    };
    this._exitPromise = new Promise((resolve) => {
      this._resolveExitPromise = resolve;
    });
    this._pendingEvent = null;
    this._scheduledTimeouts = new Map();
    this._nextCallbackTimeoutID = 1;
    
    const mem = new WebAssembly.Memory({ initial: 256 });
    const memView = new DataView(mem.buffer);
    
    this.importObject = {
      go: {
        'runtime.wasmExit': (sp) => {
          const code = memView.getInt32(sp + 8, true);
          this.exit(code);
        },
        'runtime.wasmWrite': (fd, sp, n) => {
          // Simplified console output
          const buf = new Uint8Array(mem.buffer, sp, n);
          const str = new TextDecoder().decode(buf);
          if (fd === 1) {
            console.log(str);
          } else if (fd === 2) {
            console.error(str);
          }
        },
        'runtime.resetMemoryDataView': () => {
          // Memory view is already set up
        },
        'runtime.nanotime1': (sp) => {
          // Return current time in nanoseconds
          const now = BigInt(Date.now()) * 1000000n;
          memView.setBigUint64(sp + 8, now, true);
        },
        'runtime.walltime': (sp) => {
          // Return current time in seconds and nanoseconds
          const now = Date.now();
          const sec = BigInt(Math.floor(now / 1000));
          const nsec = BigInt((now % 1000) * 1000000);
          memView.setBigUint64(sp + 8, sec, true);
          memView.setBigUint64(sp + 16, nsec, true);
        },
      }
    };
  }

  async run(instance) {
    // Run the Go program
    if (instance.exports._start) {
      instance.exports._start();
    }
    return this._exitPromise;
  }
}

export default {
  async fetch(request, env, ctx) {
    try {
      // Load and instantiate the WASM module
      const go = new Go();
      const instance = await WebAssembly.instantiate(wasmData, go.importObject);
      
      // Run the Go program
      await go.run(instance);
      
      // The Go program should handle the request through syumai/workers
      return new Response('Hello from Go WASM Worker!', {
        status: 200,
        headers: { 'Content-Type': 'text/plain' }
      });
    } catch (error) {
      console.error('WASM execution error:', error);
      return new Response('Internal Server Error', { status: 500 });
    }
  },
};
