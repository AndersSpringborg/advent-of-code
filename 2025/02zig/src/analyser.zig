const std = @import("std");

pub const Sequence = struct { start: i32, end: i32 };

pub fn analyseSimple(sequence: *Sequence) !i32 {
    const sum = 0;

    return sum;
}

test "simple analyser" {
    const sequence = Sequence{ .start = 11, .end = 22 };
    const data = [_]struct { Sequence, i32 }{ Sequence{ .start = 2, .end = 3 }, 32 };

    var total = 0;
    for (data) |entry| {

        const res = analyseSimple(entry.@"0");
        total+=res;
        try std.testing.expect(entry.@"1" == res);


    }

}
