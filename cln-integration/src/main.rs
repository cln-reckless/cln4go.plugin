//! Core Lightning Integration testing
//!
//! Author: Vincenzo Palazzo <vincenzopalazzo@member.fsf.org>
use std::sync::Once;

mod utils;

#[allow(unused_imports)]
use utils::*;

static INIT: Once = Once::new();

fn init() {
    // ignore error
    INIT.call_once(|| {
       env_logger::init();
    });
}

#[tokio::test(flavor = "multi_thread")]
async fn test_init_plugin() -> anyhow::Result<()> {
    init();
    Ok(())
}
