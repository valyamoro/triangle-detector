# Triangle Detector

Small Go project for detecting triangle patterns on candlestick data and rendering charts.

This README documents the recent refactor performed on the loader and JSON helper utilities (done in this work session).

## What changed (today)

- Introduced a unified loader `loader.go` that centralises data loading logic for candles.
  - Decides whether to read from a local JSON file (default) or fetch from Binance.
  - Fetch policy:
    - Explicit fetch when both `-symbol` and `-interval` flags are provided.
    - If the local data file is missing or empty and no explicit flags are provided, loader fetches using defaults.
    - Defaults: `BTCUSDT` symbol and `15m` interval.
    - Fetch requests are bounded to 50 candles and the loader validates requested ranges (it suggests narrower ranges when a request exceeds 50 candles).
  - When fetching, the loader saves results only if the local file was absent or empty to avoid silent overwrites.

- Added `helpers.go` with generic JSON helpers and small utilities:
  - `readJSONFile[T any](path string) ([]T, error)` — generic reader for JSON arrays.
  - `saveJSONFile[T any](path string, items []T) error` — generic writer.
  - `isEmptyFile(path string) (bool, error)` — checks whether a file is missing or has zero length.
  - `intervalToMilliseconds`, `parseTime`, `parseInt64`, `parseFloat` helpers used by the loader.

## Key files

- `loader.go` — unified loader and `LoadCandles` entrypoint.
- `helpers.go` — generic JSON helpers and utility functions.
- `main.go` — CLI entrypoint; uses `LoadCandles` to obtain candles and then runs detection + rendering.

## CLI usage (examples)

- Read local `candles.json` (default behavior when file exists):
  ```sh
  go run .
  ```

- Explicitly fetch from Binance (do not overwrite existing non-empty file):
  ```sh
  go run . -symbol BTCUSDT -interval 15m -start 2026-04-14T00:00:00Z -end 2026-04-14T13:00:00Z
  ```

- If `candles.json` is missing or empty, running without flags will fetch default data:
  ```sh
  go run .
  # loader will fetch BTCUSDT @ 15m and save candles.json (only if file was missing/empty)
  ```

## Design notes

- The loader intentionally avoids overwriting an existing, non-empty `candles.json` to prevent data loss; if you need to force an overwrite, consider adding a `-force` or `-save` flag (future enhancement).
- Generics in `helpers.go` let us reuse JSON I/O for any typed slice (e.g., `[]Candle`).
- Range validation prevents requests exceeding 50 candles and provides suggestions for acceptable start/end windows.

## Next steps (possible enhancements)

- Add `-force` flag to allow intentional overwrite of `candles.json`.
- Add `-count`/`-last` to request a specific number of candles instead of start/end.
- Improve CLI messages and add unit/integration tests for `LoadCandles` and helpers.

## Contact

If you want the README expanded with examples, diagrams, or developer notes, tell me what to add and I'll update it.
