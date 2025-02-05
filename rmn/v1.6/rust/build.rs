use std::process::Command;

fn main() {
    let git_hash = Command::new("git")
        .args(["rev-parse", "HEAD"])
        .output()
        .ok()
        .filter(|output| output.status.success())
        .and_then(|x| String::from_utf8(x.stdout).ok());

    println!("cargo:rustc-env=GIT_HASH={:?}", git_hash);

    // Remove the old generated files
    let _ = std::fs::remove_file("src/keystore.rs");
    let _ = std::fs::remove_file("src/rageproxy.rs");
    let _ = std::fs::remove_file("src/rmn_offchain.rs");
    let _ = std::fs::remove_file("src/offchainreporting3_config.rs");

    let out_dir = std::env::var("CARGO_MANIFEST_DIR").unwrap() + "/src";
    prost_build::Config::new()
        .out_dir(&out_dir)
        .compile_protos(
        &[
            "../proto/serialization/rageproxy.proto",
            "../proto/serialization/rmn_offchain.proto",
            "../proto/serialization/keystore.proto",
            "../proto/ocr3config/offchainreporting3_offchain_config.proto",
        ],
        &["../proto"],
    )
    .unwrap();
}
