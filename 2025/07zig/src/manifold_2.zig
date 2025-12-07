const std = @import("std");
const print = std.debug.print;

test "basecase" {
    const y: i32 = 0;
    const max: i32 = 5;
    try std.testing.expectEqual(true, y == (max - 5));
}

pub const Manifold = struct {
    lines_processed: []u8,
    line_length: usize = 0,
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator) !Manifold {
        const lines_processed = try allocator.alloc(u8, 0);
        errdefer allocator.free(lines_processed);

        return Manifold{
            .lines_processed = lines_processed,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Manifold, allocator: std.mem.Allocator) void {
        allocator.free(self.lines_processed);
    }

    pub fn printInput(self: *Manifold) void {
        const linewith_usize: usize = @intCast(self.line_length);
        if (linewith_usize == 0) {
            print("No lines processed.\n", .{});
            return;
        }
        const total_lines: usize = self.lines_processed.len / linewith_usize;

        for (0..total_lines) |row| {
            const start_index = row * linewith_usize;
            const end_index = start_index + linewith_usize;
            const line = self.lines_processed[start_index..end_index];
            print("{s}\n", .{line});
        }
    }

    pub fn getTimelines(self: *Manifold) isize {
        if (self.line_length == 0) {
            return 1;
        }

        const maxDepth: i32 = @intCast(self.lines_processed.len / self.line_length);
        return depth_firstSearch(self, -1, 0, maxDepth);
    }

    fn depth_firstSearch(self: *Manifold, x: i32, y: i32, max_depth: i32) isize {
        const linewith_usize: usize = @intCast(self.line_length);
        const y_usize: usize = @intCast(y); // Cast y to usize

        // Base case: if we've reached the maximum depth, return 1 timeline found
        if (y == (max_depth - 1)) {
            return 1;
        }
        // If x is -1, we need to find the starting position 'S' in this row
        if (x == -1) {
            for (0..linewith_usize) |column| {
                const char_index = y_usize * linewith_usize + column;
                if (self.lines_processed[char_index] == 'S') {
                    const column_isize: i32 = @intCast(column);
                    return self.depth_firstSearch(column_isize, y + 1, max_depth);
                }
            }
        }
        const x_usize: usize = @intCast(x); // Cast x to usize

        var timelines: isize = 0;
        const index: usize = y_usize * linewith_usize + x_usize;
        const current_char = self.lines_processed[index];
        if (current_char == '^') {
            // Explore left
            if (x > 0) {
                timelines += self.depth_firstSearch(x - 1, y + 1, max_depth);
            }
            // Explore right
            if (x < self.line_length - 1) {
                timelines += self.depth_firstSearch(x + 1, y + 1, max_depth);
            }
        }
        if (current_char == '.') {
            // Continue straight down
            timelines += self.depth_firstSearch(x, y + 1, max_depth);
        }

        return timelines;
    }

    pub fn getTimelinesDP(self: *Manifold) !isize {
        if (self.line_length == 0) return 1;

        const width: usize = self.line_length;
        const height: usize = self.lines_processed.len / width;

        // Create DP table
        var dp = try self.allocator.alloc(isize, self.lines_processed.len);
        defer self.allocator.free(dp);
        @memset(dp, 0);

        // Base case: bottom row has 1 timeline
        const last_row = height - 1;
        for (0..width) |x| {
            dp[last_row * width + x] = 1;
        }

        // Work backwards from second-to-last row
        var row: isize = @intCast(height - 2);
        while (row >= 0) : (row -= 1) {
            const row_usize: usize = @intCast(row);
            for (0..width) |x| {
                const idx = row_usize * width + x;
                const char = self.lines_processed[idx];

                if (char == '^') {
                    if (x > 0) {
                        const timelinesOnBelowLeft = dp[(row_usize + 1) * width + (x - 1)];
                        dp[idx] += timelinesOnBelowLeft;
                    }
                    if (x < width - 1) {
                        const timelinesOnBelowRight = dp[(row_usize + 1) * width + (x + 1)];
                        dp[idx] += timelinesOnBelowRight;
                    }
                } else if (char == '.' or char == 'S') {
                    // contine amount of timelines from below
                    const timelinesOnBelow = dp[(row_usize + 1) * width + x];
                    dp[idx] = timelinesOnBelow;
                }
            }
        }

        // print timelines
        for (0..height) |y| {
            for (0..width) |x| {
                const idx = y * width + x;
                const d : usize= @intCast(dp[idx]);
                print("{d:2} ", .{d});
            }
            print("\n", .{});

        }


        // Find S and return its value
        for (0..width) |x| {
            if (self.lines_processed[x] == 'S') {
                return dp[x];
            }
        }
        return 0;
    }

    pub fn parseLine(self: *Manifold, allocator: std.mem.Allocator, line: []const u8) !void {
        if (self.line_length == 0) {
            self.line_length = @intCast(line.len);
        }

        // increase size of lines_processed
        const old_len = self.lines_processed.len;
        const new_len = old_len + line.len;
        const new_lines_processed = try allocator.realloc(self.lines_processed, new_len);
        self.lines_processed = new_lines_processed;
        // copy line into lines_processed
        @memcpy(self.lines_processed[old_len..new_len], line);
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

    var manifold = try Manifold.init(allocator);
    defer manifold.deinit(allocator);

    try manifold.parseLine(allocator, ".......S.......");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, ".......^.......");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, "......^.^......");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, ".....^.^.^.....");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, "....^.^...^....");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, "...^.^...^.^...");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, "..^...^.....^..");
    try manifold.parseLine(allocator, "...............");
    try manifold.parseLine(allocator, ".^.^.^.^.^...^.");
    try manifold.parseLine(allocator, "...............");

    try std.testing.expectEqual(40, manifold.getTimelinesDP());
}

test "Manifold initialization" {
    const allocator = std.testing.allocator;

    var manifold = try Manifold.init(allocator);
    errdefer manifold.deinit(allocator);

    try std.testing.expectEqual(1, manifold.getTimelines());
}
