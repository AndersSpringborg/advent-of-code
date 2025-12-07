const std = @import("std");
const scanner = @import("scanner.zig");
const batteries = @import("batteries.zig");
const big_battery = @import("big_battery.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var scan_result1 = try scanner.scan(allocator, "input.txt");
    defer scan_result1.close(allocator);

    var total_voltage: i32 = 0;

    while (scan_result1.scanner.scan()) {
        const line = scan_result1.scanner.text();
        const bank = batteries.Bank.init(line);
        const voltage = try bank.outputVoltage(allocator);
        total_voltage += voltage;
    }

    var scan_result2 = try scanner.scan(allocator, "input2.txt");
    defer scan_result2.close(allocator);

    var total_voltage2: i32 = 0;

    while (scan_result2.scanner.scan()) {
        const line = scan_result2.scanner.text();
        const big_bank = big_battery.BigBank.init(line);
        const voltage = try big_bank.outputVoltage(allocator, 12);
        total_voltage2 += voltage;
    }

    std.debug.print("Total Output Voltage: {d}\n", .{total_voltage});
    std.debug.print("Total Output Voltage 2: {d}\n", .{total_voltage2});
}
