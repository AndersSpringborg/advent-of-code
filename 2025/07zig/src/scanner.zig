const std = @import("std");

pub const Scanner = struct {
    lines: std.mem.SplitIterator(u8, .scalar),
    current_line: ?[]const u8,

    pub fn scan(self: *Scanner) bool {
        self.current_line = self.lines.next();
        return self.current_line != null;
    }

    pub fn text(self: *Scanner) []const u8 {
        return self.current_line orelse "";
    }
};
pub const ScanResult = struct {
    scanner: Scanner,
    content: []u8,
    file: std.fs.File,

    pub fn close(self: *ScanResult, allocator: std.mem.Allocator) void {
        allocator.free(self.content);
        self.file.close();
    }
};

pub fn scan(allocator: std.mem.Allocator, path: []const u8) !ScanResult {
    const file = try std.fs.cwd().openFile(path, .{});
    errdefer file.close();

    const stat = try file.stat();
    const content = try allocator.alloc(u8, stat.size);
    errdefer allocator.free(content);

    const bytes_read = try file.read(content);
    const file_content = content[0..bytes_read];

    const lines = std.mem.splitScalar(u8, file_content, '\n');

    return ScanResult{
        .scanner = Scanner{
            .lines = lines,
            .current_line = null,
        },
        .content = content,
        .file = file,
    };
}
