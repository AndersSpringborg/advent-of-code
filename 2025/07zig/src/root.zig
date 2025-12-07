//! By convention, root.zig is the root source file when making a library.
const std = @import("std");

test {
    @import("std").testing.refAllDecls(@This());
    _ = @import("manifold.zig");
    _ = @import("manifold_2.zig");
}
