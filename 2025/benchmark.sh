#!/usr/bin/env bash

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if hyperfine is installed
if ! command -v hyperfine &> /dev/null; then
    echo -e "${RED}Error: hyperfine is not installed${NC}"
    echo "Install it with: brew install hyperfine"
    exit 1
fi

# Configuration
WARMUP=3
RUNS=10
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUTPUT_DIR="${REPO_ROOT}/benchmark-results"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

mkdir -p "${OUTPUT_DIR}"

echo -e "${BLUE}=== Advent of Code 2025 Benchmark Suite ===${NC}"
echo -e "Warmup runs: ${WARMUP}"
echo -e "Benchmark runs: ${RUNS}"
echo -e "Results will be saved to: ${OUTPUT_DIR}"
echo ""

# Build all solutions first
echo -e "${YELLOW}Building all solutions...${NC}"

build_go_solution() {
    local day=$1
    local dir="${REPO_ROOT}/${day}"

    if [[ ! -d "${dir}" ]]; then
        return 1
    fi

    cd "${dir}"

    # Determine the numeric day (remove leading zero for directory names)
    local day_num="${day#0}"

    # Check for cmd structure (days 03+) or main.go in root (days 01-02)
    if [[ -d "cmd/${day_num}" ]] && [[ -f "cmd/${day_num}/main.go" ]]; then
        echo "  Building Go day ${day} (cmd structure)..."
        go build -o "advent_${day}" "./cmd/${day_num}"
    elif [[ -d "cmd/${day}" ]] && [[ -f "cmd/${day}/main.go" ]]; then
        echo "  Building Go day ${day} (cmd structure with leading zero)..."
        go build -o "advent_${day}" "./cmd/${day}"
    elif [[ -d "cmd" ]] && [[ -f "cmd/main.go" ]]; then
            echo "  Building Go day ${day} (cmd structure with leading zero)..."
            go build -o "advent_${day}" "./cmd/main.go"
    elif [[ -f "main.go" ]]; then
        echo "  Building Go day ${day} (root main.go)..."
        go build -o "advent_${day}" .
    else
        return 1
    fi

    return 0
}

build_zig_solution() {
    local day=$1
    local dir="${REPO_ROOT}/${day}zig"

    if [[ ! -d "${dir}" ]]; then
        return 1
    fi

    cd "${dir}"
    echo "  Building Zig day ${day}..."
    zig build -Doptimize=ReleaseFast > /dev/null 2>&1

    return 0
}

# Find all available days
GO_DAYS=()
ZIG_DAYS=()

# Check for Go solutions (days 01-06 based on user info)
for day_num in {1..6}; do
    day=$(printf "%02d" ${day_num})
    day_dir="${REPO_ROOT}/${day}"
    if [[ -d "${day_dir}" ]]; then
        if build_go_solution "${day}"; then
            GO_DAYS+=("${day}")
        fi
    fi
done

# Check for Zig solutions
for day_dir in "${REPO_ROOT}"/[0-9][0-9]zig; do
    if [[ -d "${day_dir}" ]]; then
        day=$(basename "${day_dir}" | sed 's/zig$//')
        if build_zig_solution "${day}"; then
            ZIG_DAYS+=("${day}")
        fi
    fi
done

echo -e "${GREEN}Build complete!${NC}"
echo -e "Go solutions: ${GO_DAYS[*]}"
echo -e "Zig solutions: ${ZIG_DAYS[*]}"
echo ""

# Function to get the correct binary path
get_go_binary() {
    local day=$1
    local dir="${REPO_ROOT}/${day}"

    # All binaries are now built as advent_{day} in the day directory
    if [[ -f "${dir}/advent_${day}" ]]; then
        echo "${dir}/advent_${day}"
    else
        return 1
    fi
}

get_zig_binary() {
    local day=$1
    echo "${REPO_ROOT}/${day}zig/zig-out/bin/_${day}zig"
}

# Run benchmarks
RESULTS_FILE="${OUTPUT_DIR}/results_${TIMESTAMP}.md"
TEMP_JSON="${OUTPUT_DIR}/temp_${TIMESTAMP}.json"

echo "# Advent of Code 2025 Benchmark Results" > "${RESULTS_FILE}"
echo "" >> "${RESULTS_FILE}"
echo "Generated: $(date)" >> "${RESULTS_FILE}"
echo "" >> "${RESULTS_FILE}"
echo "Configuration:" >> "${RESULTS_FILE}"
echo "- Warmup runs: ${WARMUP}" >> "${RESULTS_FILE}"
echo "- Benchmark runs: ${RUNS}" >> "${RESULTS_FILE}"
echo "- Optimization: Go (default), Zig (ReleaseFast)" >> "${RESULTS_FILE}"
echo "" >> "${RESULTS_FILE}"

# Compare Go vs Zig for each day
echo -e "${BLUE}Running benchmarks...${NC}"
echo ""

# Array to store all benchmark data
declare -A BENCHMARK_DATA

for day in $(printf '%s\n' "${GO_DAYS[@]}" "${ZIG_DAYS[@]}" | sort -u); do
    go_binary=$(get_go_binary "${day}" 2>/dev/null || echo "")
    zig_binary=$(get_zig_binary "${day}" 2>/dev/null || echo "")

    # Skip if neither binary exists
    if [[ -z "${go_binary}" ]] && [[ ! -f "${zig_binary}" ]]; then
        continue
    fi

    echo -e "${YELLOW}Day ${day}${NC}"

    # Build hyperfine command
    cmds=()
    names=()

    if [[ -n "${go_binary}" ]] && [[ -f "${go_binary}" ]]; then
        cmds+=("cd ${REPO_ROOT}/${day} && ${go_binary}")
        names+=("Day${day}-Go")
    fi

    if [[ -f "${zig_binary}" ]]; then
        cmds+=("cd ${REPO_ROOT}/${day}zig && ${zig_binary}")
        names+=("Day${day}-Zig")
    fi

    # Only benchmark if we have at least one implementation
    if [[ ${#cmds[@]} -gt 0 ]]; then
        # Build the hyperfine command with proper quoting
        hyperfine_cmd=(hyperfine --warmup "${WARMUP}" --runs "${RUNS}" --export-json "${OUTPUT_DIR}/day${day}_${TIMESTAMP}.json")

        # Add each command with its name
        for i in "${!cmds[@]}"; do
            hyperfine_cmd+=(--command-name "${names[$i]}" "${cmds[$i]}")
        done

        # Run hyperfine
        "${hyperfine_cmd[@]}"
    fi

    echo ""
done

# Generate consolidated table using Python for JSON parsing
echo "## Consolidated Results" >> "${RESULTS_FILE}"
echo "" >> "${RESULTS_FILE}"
echo "| Day | Language | Mean [ms] | Min [ms] | Max [ms] | Std Dev [ms] |" >> "${RESULTS_FILE}"
echo "|-----|----------|-----------|----------|----------|--------------|" >> "${RESULTS_FILE}"

# Process all JSON files and extract data using Python
python3 -c "
import json
from pathlib import Path

output_dir = Path('${OUTPUT_DIR}')
timestamp = '${TIMESTAMP}'
days = '${GO_DAYS[*]} ${ZIG_DAYS[*]}'.split()

results = []

for day in sorted(set(days)):
    json_file = output_dir / f'day{day}_{timestamp}.json'

    if json_file.exists():
        with open(json_file) as f:
            data = json.load(f)

        for result in data.get('results', []):
            cmd_name = result.get('command', '')
            # Parse 'Day01-Go' format
            if '-' in cmd_name:
                day_num = cmd_name.split('-')[0].replace('Day', '')
                lang = cmd_name.split('-')[1]

                mean_ms = result.get('mean', 0) * 1000
                min_ms = result.get('min', 0) * 1000
                max_ms = result.get('max', 0) * 1000
                stddev_ms = result.get('stddev', 0) * 1000

                results.append({
                    'day': day_num,
                    'lang': lang,
                    'mean': mean_ms,
                    'min': min_ms,
                    'max': max_ms,
                    'stddev': stddev_ms
                })

# Sort by day (numeric) and language
results.sort(key=lambda x: (int(x['day']), x['lang']))

for r in results:
    print(f\"| {r['day']:>3} | {r['lang']:8} | {r['mean']:9.3f} | {r['min']:8.3f} | {r['max']:8.3f} | {r['stddev']:12.3f} |\")
" >> "${RESULTS_FILE}"

echo "" >> "${RESULTS_FILE}"

echo -e "${GREEN}=== Benchmark Complete ===${NC}"
echo -e "Results saved to: ${RESULTS_FILE}"
echo ""

# Generate summary
echo -e "${BLUE}Generating summary...${NC}"

SUMMARY_FILE="${OUTPUT_DIR}/summary_${TIMESTAMP}.txt"

{
    echo "=== BENCHMARK SUMMARY ==="
    echo ""
    echo "Total Go solutions benchmarked: ${#GO_DAYS[@]}"
    echo "Total Zig solutions benchmarked: ${#ZIG_DAYS[@]}"
    echo ""
    echo "Detailed results: ${RESULTS_FILE}"
} > "${SUMMARY_FILE}"

cat "${SUMMARY_FILE}"

# Create a symlink to the latest results
ln -sf "results_${TIMESTAMP}.md" "${OUTPUT_DIR}/latest.md"

echo ""
echo -e "${GREEN}View latest results: ${OUTPUT_DIR}/latest.md${NC}"
