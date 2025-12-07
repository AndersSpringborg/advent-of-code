//! By convention, root.zig is the root source file when making a library.
const std = @import("std");

// Import test modules to ensure tests are run
test {
    @import("std").testing.refAllDecls(@This());
    _ = @import("batteries.zig");
    _ = @import("big_battery.zig");
    _ = @import("scanner.zig");
}
