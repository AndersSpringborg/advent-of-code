const std = @import("std");
const common = @import("common.zig");

pub fn analyseSimple(allocator: std.mem.Allocator, sequences: []common.Sequence) !i64 {
    var sum: i64 = 0;

    for (sequences) |seq| {
        var num = seq.start;
        while (num <= seq.end) : (num += 1) {
            const s = try std.fmt.allocPrint(allocator, "{d}", .{num});
            defer allocator.free(s);

            const length = s.len;

            // Only even-length numbers can split into equal halves
            if (length % 2 == 0) {
                const mid = length / 2;
                const first_half = s[0..mid];
                const second_half = s[mid..];

                if (std.mem.eql(u8, first_half, second_half)) {
                    sum += num;
                }
            }
        }
    }

    return sum;
}


const TestCase = struct {
    sequence: common.Sequence,
    expected_sum: i64,
    description: []const u8,
};

test "analyseSimple test cases" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_cases = [_]TestCase{
        .{ .sequence = .{ .start = 11, .end = 22 }, .expected_sum = 33, .description = "11-22 should find 11 and 22" },
        .{ .sequence = .{ .start = 95, .end = 115 }, .expected_sum = 99, .description = "95-115 should find 99" },
        .{ .sequence = .{ .start = 998, .end = 1012 }, .expected_sum = 1010, .description = "998-1012 should find 1010" },
        .{ .sequence = .{ .start = 1188511880, .end = 1188511890 }, .expected_sum = 1188511885, .description = "should find 1188511885" },
        .{ .sequence = .{ .start = 222220, .end = 222224 }, .expected_sum = 222222, .description = "should find 222222" },
        .{ .sequence = .{ .start = 1698522, .end = 1698528 }, .expected_sum = 0, .description = "should find nothing" },
        .{ .sequence = .{ .start = 446443, .end = 446449 }, .expected_sum = 446446, .description = "should find 446446" },
        .{ .sequence = .{ .start = 38593856, .end = 38593862 }, .expected_sum = 38593859, .description = "should find 38593859" },
        .{ .sequence = .{ .start = 565653, .end = 565659 }, .expected_sum = 0, .description = "should find nothing" },
        .{ .sequence = .{ .start = 824824821, .end = 824824827 }, .expected_sum = 0, .description = "should find nothing" },
        .{ .sequence = .{ .start = 2121212118, .end = 2121212124 }, .expected_sum = 0, .description = "should find nothing" },
    };

    var total: i64 = 0;
    var passed: usize = 0;
    var failed: usize = 0;

    for (test_cases) |tc| {
        var sequences = [_]common.Sequence{ tc.sequence };
        const result = try analyseSimple(allocator, &sequences);
        total += result;

        if (result == tc.expected_sum) {
            std.debug.print("✓ PASS: {s}\n", .{tc.description});
            passed += 1;
        } else {
            std.debug.print("✗ FAIL: {s}\n", .{tc.description});
            std.debug.print("  Expected: {d}, Got: {d}\n", .{ tc.expected_sum, result });
            failed += 1;
        }
    }

    // Final assertions
    if (total == 1227775554) {
        std.debug.print("✓ PASS: Total sum matches expected value\n", .{});
        passed += 1;
    } else {
        std.debug.print("✗ FAIL: Total sum does not match\n", .{});
        std.debug.print("  Expected: 1227775554, Got: {d}\n", .{total});
        failed += 1;
    }

    std.debug.print("\n{d} passed, {d} failed\n", .{ passed, failed });
}