const std = @import("std");

test {
    @import("std").testing.refAllDecls(@This());
    _ = @import("sequence_analyser.zig");
    _ = @import("sequence_complex.zig");
}
