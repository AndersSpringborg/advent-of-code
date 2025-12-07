const std = @import("std");
const scanner = @import("scanner.zig");
const print = @import("std").debug.print;
const Manifold = @import("manifold.zig");
const Manifold2 = @import("manifold_2.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var scan_result1 = try scanner.scan(allocator, "input.txt");
    defer scan_result1.close(allocator);

    var manifold = try Manifold.Manifold.init(allocator, 141);
    defer manifold.deinit(allocator);

    var manifold2 = try Manifold2.Manifold.init(allocator);
    defer manifold2.deinit(allocator);

    while (scan_result1.scanner.scan()) {
        const line = scan_result1.scanner.text();
        try manifold.parseLine(line);
        try manifold2.parseLine(allocator, line);
    }

    const splits = manifold.getSplits();
    const timelines = manifold2.getTimelinesDP();
    print("Total splits: {d}\n", .{splits});
    print("Total timelines: {any}\n", .{timelines});
}
