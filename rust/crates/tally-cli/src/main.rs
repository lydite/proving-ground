//! The deployable. It is not the component: `rust/` is one Cargo workspace and
//! therefore one build unit, which is what a component tracks.

use tally_core::{floor_zero, tally};

/// The root `VERSION` file, written by `make version`. Embedding it is what makes
/// the `watch` entry on the `tally` component real: change `VERSION` and this
/// binary genuinely changes.
const VERSION: &str = include_str!("../../../../VERSION");

fn main() {
    let counts: Vec<i64> = std::env::args()
        .skip(1)
        .filter_map(|arg| arg.parse().ok())
        .map(floor_zero)
        .collect();

    println!("tally {}", VERSION.trim());
    println!("{}", tally(&counts));
}
