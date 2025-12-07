const std = @import("std");

pub const Bank = struct {
    batteries: []const u8,

    pub fn init(batteries: []const u8) Bank {
        return Bank{
            .batteries = batteries,
        };
    }

    pub fn outputVoltage(self: *const Bank, allocator: std.mem.Allocator) !i32 {
        const voltages = parseBatteriesToIntArray(allocator, self.batteries) catch |err| {
            return err;
        };


        var left: i32 = 0;
        var right: i32 = 0;

        for (0.., voltages) |i, voltage| {
            if (voltage > left and (i < voltages.len - 1)) {
                left = voltage;
                right = voltages[i+1];
            } else if (voltage >= right and i == voltages.len - 1) {
                right = voltage;
            }
        }
        return joltFromNumbers(allocator, left, right);
    }




};

fn joltFromNumbers(allocator: std.mem.Allocator, left: i32, right: i32) !i32 {
    const leftStr = try itoa(allocator, left);
    defer allocator.free(leftStr);
    const rightStr = try itoa(allocator, right);
    defer allocator.free(rightStr);

    const concatedetNumbers = try std.mem.concat(allocator, u8, &[_][]const u8{leftStr, rightStr});
    defer allocator.free(concatedetNumbers);

    return atoi(concatedetNumbers);
}

fn atoi(value: []const u8) i32 {
    var result: i32 = 0;
    for (value) |char| {
        result = result * 10 + (@as(i32, char) - @as(i32, '0'));
    }
    return result;
}

fn itoa(allocator: std.mem.Allocator, value: i32) ![]u8 {
    return try std.fmt.allocPrint(allocator, "{d}", .{value});
}


pub fn parseBatteriesToIntArray(allocator: std.mem.Allocator, input: []const u8) ![]i32 {
    var voltages = try allocator.alloc(i32, input.len);

    for (input, 0..) |char, i| {
        const num = try std.fmt.parseInt(i32, &[_]u8{char}, 10);
        voltages[i] = num;
    }

    return voltages;
}

test "Bank output voltage" {
    const test_cases = [_]struct {
        batteries: []const u8,
        expected_voltage: i64,
    }{
        .{ .batteries = "99", .expected_voltage = 99 },
        .{ .batteries = "987654321111111", .expected_voltage = 98 },
        .{ .batteries = "811111111111119", .expected_voltage = 89 },
        .{ .batteries = "234234234234278", .expected_voltage = 78 },
        .{ .batteries = "818181911112111", .expected_voltage = 92 },
    };

    for (test_cases) |tc| {
        const bank = Bank.init(tc.batteries);
        const result = try bank.outputVoltage(std.testing.allocator);
        try std.testing.expectEqual(tc.expected_voltage, result);
    }
}
