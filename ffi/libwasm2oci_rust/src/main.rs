extern crate libwasm2oci;

fn main() {
    libwasm2oci::pull_wasm("cnabregistry.azurecr.io/wasm-to-oci:v2", "test.wasm").unwrap();
}
