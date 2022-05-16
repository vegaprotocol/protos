# API Convention

- Optional fields should use the `optional` keyword
- ID fields should carry a qualifier prefix and an `_id` suffix, eg. `party_id`.
- All service endpoints must be prefixed with a quantifying verb prefix, e.g. `List`, `Get`, `Submit`, `Estimate`, `Observe`.
- `*Request`, `*Response` messages must appear in the order that the endpoints were defined, e.g. the last endpoint will have its messages defined at the end of the file.
- Nested messages should be defined before their first use.
- All REST endpoints must be kebab-case (lower case, words separated by `-`)
