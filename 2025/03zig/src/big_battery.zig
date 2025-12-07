const std = @import("std");
const batteries = @import("batteries.zig");

pub const BigBank = struct {
    batteries_str: []const u8,

    pub fn init(batteries_str: []const u8) BigBank {
        return BigBank{
            .batteries_str = batteries_str,
        };
    }

    pub fn outputVoltage(self: *const BigBank, allocator: std.mem.Allocator, n: usize) !i32 {
        _ = allocator;
        _ = self;
        _ = n;
        // TODO: Implement logic
        return 1; // Dummy return value
    }

    fn voltageFromNumbers(self: *const BigBank, allocator: std.mem.Allocator, nums: []const i32) !i32 {
        _ = self;
        _ = allocator;
        _ = nums;


        return 0;
    }
};

test "BigBank output voltage" {
    const test_cases = [_]struct {
        batteries: []const u8,
        expected_voltage: i64,
    }{
        .{ .batteries = "811111111111119", .expected_voltage = 811111111119 },
        .{ .batteries = "234234234234278", .expected_voltage = 434234234278 },
        .{ .batteries = "818181911112111", .expected_voltage = 888911112111 },
    };

    for (test_cases) |tc| {
        const big_bank = BigBank.init(tc.batteries);
        const result = try big_bank.outputVoltage(std.testing.allocator, 12);
        try std.testing.expectEqual(tc.expected_voltage, result);
    }
}
