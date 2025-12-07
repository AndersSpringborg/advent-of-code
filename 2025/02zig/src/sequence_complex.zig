const std = @import("std");
const common = @import("common.zig");
const expectEqual = std.testing.expectEqual;

const Set = struct {
    numbers: []i64,
    len: usize,
    allocator: std.mem.Allocator,

    pub fn init(allocator: std.mem.Allocator) !Set {
        const numbers = try allocator.alloc(i64, 10000);
        return Set{ 
            .numbers = numbers,
            .len = 0,
            .allocator = allocator,
        };
    }

    pub fn deinit(self: *Set) void {
        self.allocator.free(self.numbers);
    }

    pub fn addNumber(self: *Set, n: i64) !void {
        for (self.numbers[0..self.len]) |num| {
            if (num == n) return;
        }
        self.numbers[self.len] = n;
        self.len += 1;
    }

    pub fn slice(self: *Set) []i64 {
        return self.numbers[0..self.len];
    }
};

pub fn analyseSingleComplex(allocator: std.mem.Allocator, sequence: *const common.Sequence) ![]i64 {
    var invalid_ids = try Set.init(allocator);
    defer invalid_ids.deinit();

    const numbers = try common.enumerateSequence(allocator, sequence);
    defer allocator.free(numbers);

    for (numbers) |number| {
        const number_str = try common.numberToString(allocator, number);
        defer allocator.free(number_str);

        const splitsizesArr = try allocator.alloc(usize, number_str.len);
        const split_sizes = try splitSizes(splitsizesArr, number_str);
        defer allocator.free(split_sizes);

        const string_len = number_str.len;

        for (split_sizes) |size| {
            const substring_size = string_len / size;
            const pattern = number_str[0..substring_size];

            const parts = try splitStringIntoParts(allocator, number_str, size);
            defer allocator.free(parts);

            var all_same = true;
            for (parts) |part| {
                if (!std.mem.eql(u8, part, pattern)) {
                    all_same = false;
                    break;
                }
            }

            if (all_same) {
                try invalid_ids.addNumber(number);
                break;
            }
        }
    }

    const result = try allocator.alloc(i64, invalid_ids.len);
    @memcpy(result, invalid_ids.slice());
    return result;
}

fn splitStringIntoParts(allocator: std.mem.Allocator, s: []const u8, size: usize) ![][]const u8 {
    const parts = try allocator.alloc([]const u8, size);
    
    const string_len = s.len;
    const substring_size = string_len / size;

    var i: usize = 0;
    while (i < size) : (i += 1) {
        const start = i * substring_size;
        const end = start + substring_size;
        parts[i] = s[start..end];
    }

    return parts;
}

fn splitSizes(sizes: []usize, number_str: []const u8) ![]usize {
    const length = number_str.len;

    var count: usize = 0;

    var divider: usize = 2;
    while (divider <= length) : (divider += 1) {
        if (length % divider == 0) {
            sizes[count] = divider;
            count += 1;
        }
    }

    return sizes[0..count];
}

pub fn analyseComplex(allocator: std.mem.Allocator, sequences: []common.Sequence) !i64 {
    var all_count: usize = 0;
    
    // First pass: count total
    for (sequences) |*seq| {
        const ids = try analyseSingleComplex(allocator, seq);
        defer allocator.free(ids);
        all_count += ids.len;
    }

    // Second pass: collect
    const all_invalid_ids = try allocator.alloc(i64, all_count);
    defer allocator.free(all_invalid_ids);
    
    var offset: usize = 0;
    for (sequences) |*seq| {
        const ids = try analyseSingleComplex(allocator, seq);
        defer allocator.free(ids);
        
        @memcpy(all_invalid_ids[offset..offset+ids.len], ids);
        offset += ids.len;
    }

    return common.sumInvalidIds(all_invalid_ids);
}
const TestCase = struct {
    sequence: common.Sequence,
    expected_sum: i64,
    description: []const u8,
};



test "test analyse complex" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_cases = [_]struct {
        sequence: common.Sequence,
        expected: []const i64,
    }{
        .{ .sequence = .{ .start = 11, .end = 22 }, .expected = &[_]i64{ 11, 22 } },
        .{ .sequence = .{ .start = 95, .end = 115 }, .expected = &[_]i64{ 99, 111 } },
        .{ .sequence = .{ .start = 998, .end = 1012 }, .expected = &[_]i64{ 999, 1010 } },
        .{ .sequence = .{ .start = 1188511880, .end = 1188511890 }, .expected = &[_]i64{1188511885} },
        .{ .sequence = .{ .start = 222220, .end = 222224 }, .expected = &[_]i64{222222} },
        .{ .sequence = .{ .start = 1698522, .end = 1698528 }, .expected = &[_]i64{} },
        .{ .sequence = .{ .start = 446443, .end = 446449 }, .expected = &[_]i64{446446} },
        .{ .sequence = .{ .start = 38593856, .end = 38593862 }, .expected = &[_]i64{38593859} },
        .{ .sequence = .{ .start = 565653, .end = 565659 }, .expected = &[_]i64{565656} },
        .{ .sequence = .{ .start = 824824821, .end = 824824827 }, .expected = &[_]i64{824824824} },
        .{ .sequence = .{ .start = 2121212118, .end = 2121212124 }, .expected = &[_]i64{2121212121} },
    };

    var total: i64 = 0;

    for (test_cases) |tc| {
        const result = try analyseSingleComplex(allocator, &tc.sequence);
        defer allocator.free(result);

        for (result) |n| {
            total += n;
        }

        try std.testing.expectEqual( result, tc.expected);
    }

    try std.testing.expectEqual(@as(i64, 4174379265), total);
}

test "test split sizes" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_cases = [_]struct {
        number: []const u8,
        expected_split_sizes: []const usize,
    }{
        .{ .number = "12", .expected_split_sizes = &[_]usize{2} },
        .{ .number = "123", .expected_split_sizes = &[_]usize{3} },
        .{ .number = "1234", .expected_split_sizes = &[_]usize{ 2, 4 } },
        .{ .number = "12345", .expected_split_sizes = &[_]usize{5} },
        .{ .number = "123456", .expected_split_sizes = &[_]usize{ 2, 3, 6 } },
    };

    for (test_cases) |tc| {
        const sizesArr = try allocator.alloc(usize, tc.number.len);
        const result = try splitSizes(sizesArr, tc.number);
        defer allocator.free(result);

        try std.testing.expectEqual(result, tc.expected_split_sizes);
    }
}

test "test split into parts" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const test_cases = [_]struct {
        number: []const u8,
        size: usize,
        expected_parts: []const []const u8,
    }{
        .{ .number = "12", .size = 2, .expected_parts = &[_][]const u8{ "1", "2" } },
        .{ .number = "123", .size = 3, .expected_parts = &[_][]const u8{ "1", "2", "3" } },
        .{ .number = "112233", .size = 3, .expected_parts = &[_][]const u8{ "11", "22", "33" } },
    };

    for (test_cases) |tc| {
        const result = try splitStringIntoParts(allocator, tc.number, tc.size);
        defer allocator.free(result);

        try std.testing.expectEqual(result, tc.expected_parts);
    }
}