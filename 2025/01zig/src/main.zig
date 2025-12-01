const std = @import("std");

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

fn processInput(allocator: std.mem.Allocator, filename: []const u8, comptime DialType: type, initial_index: anytype) !DialType {
    const file = try std.fs.cwd().openFile(filename, .{});
    defer file.close();

    const stat = try file.stat();
    const content = try allocator.alloc(u8, stat.size);
    defer allocator.free(content);

    const bytes_read = try file.read(content);
    const file_content = content[0..bytes_read];

    var dial = DialType.init(initial_index);
    var lines = std.mem.splitScalar(u8, file_content, '\n');

    while (lines.next()) |line| {
        if (line.len == 0) continue;

        std.debug.print("current Dial: {d}\n", .{dial.index});
        std.debug.print("Line is: {s}\n", .{line});

        const rotation_symbol = line[0];
        const steps_str = line[1..];
        const steps = try std.fmt.parseInt(@TypeOf(dial.index), steps_str, 10);

        const direction: Direction = if (rotation_symbol == 'R') .Right else .Left;
        dial.rotate(direction, steps);
    }

    return dial;
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Process first input with Dial
    const dial = try processInput(allocator, "input.txt", Dial, @as(usize, 50));

    // Process second input with Dial0x434C49434B
    const dial2 = try processInput(allocator, "input2.txt", Dial0x434C49434B, @as(i32, 50));
    std.debug.print("Final Dial Index: {d}\n", .{dial.index});
    std.debug.print("Answer is : {d}\n", .{dial.zero_index});
    std.debug.print("Final Dial0x434C49434B Index: {d}\n", .{dial2.index});
    std.debug.print("Answer 2 is : {d}\n", .{dial2.zero_index});
}
