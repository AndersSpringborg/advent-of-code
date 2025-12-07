const std = @import("std");

pub const Sequence = struct {
    start: i64,
    end: i64,
};

pub fn enumerateSequence(allocator: std.mem.Allocator, sequence: *const Sequence) ![]i64 {
    const count = @as(usize, @intCast(sequence.end - sequence.start + 1));
    const result = try allocator.alloc(i64, count);
    
    var i: usize = 0;
    var num = sequence.start;
    while (num <= sequence.end) : (num += 1) {
        result[i] = num;
        i += 1;
    }

    return result;
}

pub fn numberToString(allocator: std.mem.Allocator, number: i64) ![]u8 {
    return std.fmt.allocPrint(allocator, "{d}", .{number});
}

pub fn sumInvalidIds(invalid_ids: []const i64) i64 {
    var total: i64 = 0;
    for (invalid_ids) |id| {
        total += id;
    }
    return total;
}
