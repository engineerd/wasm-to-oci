extern crate libwasm2oci;

fn main() {
    let res = libwasm2oci::pull_wasm("cnabregistry.azurecr.io/wasm-to-oci:v2", "test.wasm").unwrap();
    println!("{}", res);
}
