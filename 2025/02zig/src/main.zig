const std = @import("std");
const scanner = @import("scanner.zig");
const common = @import("common.zig");
const sequence_analyser = @import("sequence_analyser.zig");
const sequence_complex = @import("sequence_complex.zig");
const print = std.debug.print;

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();

    //const sequences = try getSequences(allocator, "input.txt");
    //defer allocator.free(sequences);

    //const sequences2 = try getSequences(allocator, "input2.txt");
    //defer allocator.free(sequences2);

    //const res1 = try sequence_analyser.analyseSimple(allocator, sequences);
    //const res2 = try sequence_complex.analyseSingleComplex(allocator, sequences2);

    //print("Task 1 - Sum of invalid IDs: {d}\n", .{res1});
    //print("Task 2 - Sum of invalid IDs (complex analysis): {d}\n", .{res2});
}

fn getSequences(allocator: std.mem.Allocator, input_file: []const u8) ![]common.Sequence {
    var scan_result = try scanner.scan(allocator, input_file);
    defer scan_result.close(allocator);

    _ = scan_result.scanner.scan();
    const line = scan_result.scanner.text();

    return try parseSequences(allocator, line);
}

fn parseSequences(allocator: std.mem.Allocator, line: []const u8) ![]common.Sequence {
    // Count ranges first
    var count: usize = 1;
    for (line) |c| {
        if (c == ',') count += 1;
    }

    const sequences = try allocator.alloc(common.Sequence, count);
    var idx: usize = 0;

    var it = std.mem.splitScalar(u8, line, ',');
    while (it.next()) |range| {
        var parts = std.mem.splitScalar(u8, range, '-');
        const start_str = parts.next() orelse continue;
        const end_str = parts.next() orelse continue;

        const start = try std.fmt.parseInt(i64, start_str, 10);
        const end = try std.fmt.parseInt(i64, end_str, 10);

        sequences[idx] = common.Sequence{ .start = start, .end = end };
        idx += 1;
    }

    return sequences[0..idx];
}
