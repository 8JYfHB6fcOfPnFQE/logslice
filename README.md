# logslice

A fast log file slicer that extracts time-range segments from large structured or unstructured log files.

---

## Installation

```bash
go install github.com/yourusername/logslice@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logslice.git
cd logslice
go build -o logslice .
```

---

## Usage

```bash
logslice [flags] <logfile>
```

### Examples

Extract logs between two timestamps:

```bash
logslice --from "2024-01-15 08:00:00" --to "2024-01-15 09:00:00" app.log
```

Write output to a file:

```bash
logslice --from "2024-01-15 08:00:00" --to "2024-01-15 09:00:00" -o slice.log app.log
```

Specify a custom timestamp format:

```bash
logslice --from "2024-01-15T08:00:00Z" --to "2024-01-15T09:00:00Z" --format RFC3339 app.log
```

### Flags

| Flag | Description |
|------|-------------|
| `--from` | Start of the time range (inclusive) |
| `--to` | End of the time range (inclusive) |
| `--format` | Timestamp format (default: auto-detect) |
| `-o, --output` | Output file (default: stdout) |

---

## License

This project is licensed under the [MIT License](LICENSE).