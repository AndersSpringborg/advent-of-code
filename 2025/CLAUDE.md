# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository contains Advent of Code 2025 solutions implemented in both **Go** and **Zig**. Each day's puzzle has parallel implementations to compare performance characteristics between the two languages.

## Repository Structure

Solutions are organized by day number:
- `01/`, `02/`, `03/`, etc. - Go implementations
- `01zig/`, `02zig/`, `03zig/`, etc. - Zig implementations
- `benchmark/` - Zig benchmarking framework
- `README.md` - Performance timings comparison table

### Go Project Structure

Go solutions exist for days 01-06. Each day follows this pattern:
- `main.go` - Entry point (days 01-02, in root directory)
- `cmd/{day}/main.go` - Entry point (days 03+, where {day} is without leading zero, e.g., `cmd/3/main.go`)
- `internal/` - Day-specific solution logic
- `input.txt` and `input2.txt` - Puzzle inputs for parts 1 and 2
- Go module named `advent_{day}` (e.g., `advent_01`, `advent_03`)

### Zig Project Structure

Each Zig day follows this pattern:
- `src/main.zig` - Entry point with solution logic
- `build.zig` - Standard Zig build configuration
- `build.zig.zon` - Package metadata (requires Zig 0.16.0-dev.1204+)
- `input.txt` and `input2.txt` - Puzzle inputs

## Development Commands

### Go Solutions

```bash
# Navigate to a specific day
cd 01/  # or 02/, 03/, 04/, 05/, 06/

# Run the solution
go run main.go        # For days 01-02
go run ./cmd/3        # For day 03 (no leading zero in cmd path)
go run ./cmd/4        # For day 04
go run ./cmd/5        # For day 05
go run ./cmd/6        # For day 06

# Build the solution
go build -o advent_01 .         # For days 01-02
go build -o advent_03 ./cmd/3   # For days 03-06

# Run tests
go test ./...

# Run specific test
go test -run TestName ./internal
```

### Zig Solutions

```bash
# Navigate to a specific Zig day
cd 01zig/  # or 02zig/, 03zig/, etc.

# Build and run (debug mode)
zig build run

# Build optimized release
zig build -Doptimize=ReleaseFast

# Run tests
zig build test

# Build specific step
zig build install  # Default: builds executable to zig-out/
```

### Benchmarking

**Automated Benchmarking with Hyperfine:**

```bash
# Run comprehensive benchmarks of all solutions
./benchmark.sh
```

This script will:
- Build all Go and Zig solutions with optimizations (ReleaseFast for Zig)
- Run hyperfine benchmarks with warmup runs
- Generate markdown reports in `benchmark-results/`
- Create a summary comparison between Go and Zig implementations

Requirements: Install hyperfine with `brew install hyperfine`

**Manual Benchmarking:**

The `benchmark/` directory contains a Zig benchmarking framework:

```bash
cd benchmark/
zig build run
zig build test
```

## Architecture Notes

### Dual Language Strategy

This repository maintains parallel implementations in Go and Zig for performance comparison:
- **Go implementations**: Emphasis on idiomatic Go patterns, stdlib usage
- **Zig implementations**: Lower-level control, manual memory management with allocators

### Performance Testing Pattern

As documented in `README.md`, solutions are timed using the `time` command to compare real/user/sys metrics between Go and Zig implementations.

### Common Patterns

Both languages share similar solution structures:
1. File scanning utilities for reading puzzle inputs line-by-line
2. Part 1 and Part 2 often use different algorithms or data structures
3. Each part typically processes separate input files (`input.txt` and `input2.txt`)

### Zig-Specific Notes

- All Zig projects use `std.heap.GeneralPurposeAllocator` for memory management
- Memory must be explicitly freed with `defer` statements
- Scanner utilities are common across solutions for file I/O
- Minimum Zig version: 0.16.0-dev.1204+

### Go-Specific Notes

- Later days (03+) use proper project structure with `cmd/` and `internal/` separation
- Most solutions use `github.com/stretchr/testify` for testing
- File scanner helpers are shared across days
