//! Reaching the Postgres instance `../compose.yaml` publishes.
//!
//! The check is a TCP connect rather than a real driver: this crate exists to give
//! the `tally` component a service dependency, and a driver would add a large
//! dependency tree that has nothing to do with what is being exercised.

use std::io;
use std::net::{TcpStream, ToSocketAddrs};
use std::time::Duration;

/// The address `compose.yaml` publishes on the host.
pub const DEFAULT_ADDR: &str = "127.0.0.1:5432";

/// Reports whether the database is accepting connections.
pub fn reachable(addr: &str, timeout: Duration) -> io::Result<bool> {
    let resolved = addr
        .to_socket_addrs()?
        .next()
        .ok_or_else(|| io::Error::new(io::ErrorKind::InvalidInput, "no address"))?;
    Ok(TcpStream::connect_timeout(&resolved, timeout).is_ok())
}

#[cfg(test)]
mod tests {
    use super::*;

    /// Requires the compose service to be up, so it is ignored by default and the
    /// suite passes on a machine with no container runtime.
    #[test]
    #[ignore = "requires the postgres service from rust/compose.yaml"]
    fn database_is_reachable() {
        assert!(reachable(DEFAULT_ADDR, Duration::from_secs(2)).unwrap());
    }
}
