# Changelog

## [2.0.0](https://github.com/swibrow/how/compare/v1.0.1...v2.0.0) (2026-07-26)


### ⚠ BREAKING CHANGES

* remove command memory

### Features

* validate suggested commands exist before presenting to user ([4676ba1](https://github.com/swibrow/how/commit/4676ba1fbb7eea2b0ffeac16208a15d29ef47269))


### Bug Fixes

* **deps:** update module github.com/anthropics/anthropic-sdk-go to v1.26.0 ([#4](https://github.com/swibrow/how/issues/4)) ([ef2af8c](https://github.com/swibrow/how/commit/ef2af8cdb92fba7ca9de1419d09ae567b476e86c))
* **deps:** update module modernc.org/sqlite to v1.52.0 ([#22](https://github.com/swibrow/how/issues/22)) ([9497d9f](https://github.com/swibrow/how/commit/9497d9f88c70c5d636b51b8c97dfbfbc2fbe45dc))
* **deps:** update module modernc.org/sqlite to v1.54.0 ([#24](https://github.com/swibrow/how/issues/24)) ([3ec4ce5](https://github.com/swibrow/how/commit/3ec4ce583dd1556edf71a463e162d52006ee3603))
* handle LLM responses missing COMMAND: prefix ([26a21c5](https://github.com/swibrow/how/commit/26a21c5773f2e4a1b410b91a6a359d14218c3bee))


### Code Refactoring

* remove command memory ([f015a8d](https://github.com/swibrow/how/commit/f015a8db53649643138421c0bbfe7fcd84a3adbf))

## [1.0.1](https://github.com/swibrow/how/compare/v1.0.0...v1.0.1) (2026-02-26)


### Bug Fixes

* use dedicated PAT for homebrew tap publishing ([ed69ed7](https://github.com/swibrow/how/commit/ed69ed71cf54d69100fa3f0dd4a25f76241d3876))

## 1.0.0 (2026-02-26)


### Features

* allow custom prompts ([471ad2f](https://github.com/swibrow/how/commit/471ad2ff97827f4c41872b579ce44ecdd50f5b3d))
* append executed commands to shell history on exit 0 ([066fc21](https://github.com/swibrow/how/commit/066fc21a7e82d17a7de2e5ac75714f05612c4df6))
* simple memory implementation ([5c1aba7](https://github.com/swibrow/how/commit/5c1aba7c2b3810c592889095e68cc2193a251495))


### Bug Fixes

* avoid placeholders for values resolvable from environment ([76723b2](https://github.com/swibrow/how/commit/76723b25876d086133437b780058cdf46dff5292))
* improve prompt to utilise mutiple tools ([a6be3d6](https://github.com/swibrow/how/commit/a6be3d6659536dda629cfa9a2d4b9e5eef1f647a))
* resolve golangci-lint errors (errcheck, gocritic, staticcheck) ([604f5e1](https://github.com/swibrow/how/commit/604f5e108bfa95193918c3b27814766837fe15ff))
* strip backticks from LLM command output before execution ([930c8bc](https://github.com/swibrow/how/commit/930c8bcfbac2bb21f82346109bfbabce710baa57))


### Performance Improvements

* make it pretty ([abfb6e5](https://github.com/swibrow/how/commit/abfb6e56cae237c674655584c4a346e7ee9465b3))
