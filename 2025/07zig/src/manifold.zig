const std = @import("std");
const print = std.debug.print;

pub const Manifold = struct {
    beamsize: usize,
    beams: []bool,
    splits: i32 = 0,
    timelines: u32 = 0,

    pub fn getSplits(self: *Manifold) i32 {
        return self.splits;
    }
    pub fn getTimelines(self: *Manifold) u64 {
        return self.timelines;
    }

    pub fn init(allocator: std.mem.Allocator, beam_count: usize) !Manifold {
        const beams = try allocator.alloc(bool, beam_count);
        errdefer allocator.free(beams);

        for (beams) |*beam| {
            beam.* = false;
        }

        return Manifold{ .beamsize = beam_count, .beams = beams };
    }

    pub fn deinit(self: *Manifold, allocator: std.mem.Allocator) void {
        allocator.free(self.beams);
    }

    pub fn parseLine(self: *Manifold, line: []const u8) !void {
        const beams_before = countTrue(self.beams);
        for (self.beams, 0..) |value, i| {
            const char = line[i];
            if (char == '^') {
                if (value) {
                    // Beam was already present, so it splits
                    self.splits += 1;
                }
                self.beams[i - 1] = true;
                self.beams[i + 1] = true;
                self.beams[i] = false;
                const beams_after = countTrue(self.beams);
                const addedTimeline = beams_after - beams_before;
                if (addedTimeline > 0) {
                    self.timelines += @intCast(addedTimeline);
                }
            } else if (char == 'S') {
                self.timelines += 1;
                self.beams[i] = true;
            }
        }
    }
};

// parseline on manifold.zig

pub fn countTrue(beams: []bool) i32 {
    var count: i32 = 0;
    for (beams) |beam| {
        if (beam) {
            count += 1;
        }
    }
    return count;
}

test "Manifold get splits" {
    const allocator = std.testing.allocator;

    const beam_count: usize = 15;
    var manifold = try Manifold.init(allocator, beam_count);
    defer manifold.deinit(allocator);

    try manifold.parseLine(".......S.......");
    try manifold.parseLine("...............");
    try manifold.parseLine(".......^.......");
    try manifold.parseLine("...............");
    try manifold.parseLine("......^.^......");
    try manifold.parseLine("...............");
    try manifold.parseLine(".....^.^.^.....");
    try manifold.parseLine("...............");
    try manifold.parseLine("....^.^...^....");
    try manifold.parseLine("...............");
    try manifold.parseLine("...^.^...^.^...");
    try manifold.parseLine("...............");
    try manifold.parseLine("..^...^.....^..");
    try manifold.parseLine("...............");
    try manifold.parseLine(".^.^.^.^.^...^.");
    try manifold.parseLine("...............");

    try std.testing.expectEqual(21, manifold.getSplits());
}

test "Manifold initialization" {
    const allocator = std.testing.allocator;
    const beam_count: usize = 5;

    var manifold = try Manifold.init(allocator, beam_count);
    defer manifold.deinit(allocator);

    try std.testing.expectEqual(beam_count, manifold.beams.len);
    try std.testing.expectEqual(0, manifold.getSplits());

    for (manifold.beams) |beam| {
        try std.testing.expectEqual(false, beam);
    }
}
