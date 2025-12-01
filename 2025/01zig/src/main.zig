const std = @import("std");
const scanner = @import("scanner.zig");

const Direction = enum {
    Right,
    Left,
};

const DIALSIZE: usize = 100;

const Dial = struct {
    index: usize,
    zero_index: usize,

    pub fn init(index: usize) Dial {
        return Dial{
            .index = index,
            .zero_index = 0,
        };
    }

    pub fn rotate(self: *Dial, direction: Direction, steps: usize) void {
        switch (direction) {
            .Right => {
                self.index = (self.index + steps) % DIALSIZE;
            },
            .Left => {
                self.index = (self.index + DIALSIZE - (steps % DIALSIZE)) % DIALSIZE;
            },
        }

        if (self.index == 0) {
            self.zero_index += 1;
        }
    }
};

const Dial0x434C49434B = struct {
    index: i32,
    zero_index: i32,

    pub fn init(index: i32) Dial0x434C49434B {
        return Dial0x434C49434B{
            .index = index,
            .zero_index = 0,
        };
    }

    pub fn rotate(self: *Dial0x434C49434B, direction: Direction, steps: i32) void {
        var i: i32 = 0;
        while (i < steps) : (i += 1) {
            switch (direction) {
                .Right => {
                    self.index += 1;
                    if (self.index == DIALSIZE) {
                        self.index = 0;
                    }
                },
                .Left => {
                    self.index -= 1;
                    if (self.index == -1) {
                        self.index = 99;
                    }
                },
            }

            if (self.index == 0) {
                self.zero_index += 1;
            }
        }
    }
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var scan_result = try scanner.scan(allocator, "input.txt");
    defer scan_result.close(allocator);

    var dial = Dial.init(50);
    while (scan_result.scanner.scan()) {
        const line = scan_result.scanner.text();
        std.debug.print("current Dial: {d}\n", .{dial.index});
        std.debug.print("Line is: {s}\n", .{line});

        const rotation_symbol = line[0];
        const steps_str = line[1..];
        const steps_number = try std.fmt.parseInt(usize, steps_str, 10);

        if (rotation_symbol == 'R') {
            dial.rotate(Direction.Right, steps_number);
        } else if (rotation_symbol == 'L') {
            dial.rotate(Direction.Left, steps_number);
        }
    }

    var scan_result2 = try scanner.scan(allocator, "input2.txt");
    defer scan_result2.close(allocator);

    var dial_method = Dial0x434C49434B.init(50);
    while (scan_result2.scanner.scan()) {
        const line = scan_result2.scanner.text();
        std.debug.print("current Dial: {d}\n", .{dial_method.index});
        std.debug.print("Line is: {s}\n", .{line});

        const rotation_symbol = line[0];
        const steps_str = line[1..];
        const steps_number = try std.fmt.parseInt(i32, steps_str, 10);

        if (rotation_symbol == 'R') {
            dial_method.rotate(Direction.Right, steps_number);
        } else if (rotation_symbol == 'L') {
            dial_method.rotate(Direction.Left, steps_number);
        }
    }

    std.debug.print("Final Dial Index: {d}\n", .{dial.index});
    std.debug.print("Answer is : {d}\n", .{dial.zero_index});
    std.debug.print("Final Dial0x434C49434B Index: {d}\n", .{dial_method.index});
    std.debug.print("Answer 2 is : {d}\n", .{dial_method.zero_index});
}
