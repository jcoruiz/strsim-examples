# strsim-examples

Runnable programs that use [strsim](https://github.com/jcoruiz/strsim) for real tasks. Each folder is a `main` package; the README quotes part of what each one prints.

## Rules

- Every example is gofmt-clean, passes `go vet` and runs against the strsim version pinned in `go.mod`. CI runs them all.
- The outputs quoted in the README are copied from `go run` with that version. When strsim or an example changes, run it again and update the README.
- Examples import only strsim and the standard library.

## Checks

```bash
gofmt -l .
go vet ./...
for e in */; do go run "./$e"; done
```

# Git conventions

All git text is written in **English**: commit messages, PR titles and descriptions, branch names.

## Commits: Conventional Commits

Format: `<type>(<scope>): <subject>`

- **type** (required): `feat`, `fix`, `refactor`, `perf`, `test`, `docs`, `style`, `build`, `ci`, `chore`, `revert`.
- **scope** (optional): the area touched, in lowercase, such as `fuzzy-search`, `dedup`, `autocomplete`, `spell-checker`, `name-matcher`, `deps`, `readme`, `ci`.
- **subject**: imperative mood, lowercase start, no trailing period, ≤ 72 chars for the whole header.
- **body** (optional): blank line after the header, then 1–3 short lines or bullets explaining *what/why*. No multi-paragraph prose.
- **Breaking changes**: `feat(scope)!: ...` plus a `BREAKING CHANGE: ...` footer.
- One logical change per commit; don't mix refactors with features.
- No `Co-Authored-By` or AI attribution lines.
- Enforced locally by `scripts/hooks/commit-msg` (enable once per clone with `git config core.hooksPath scripts/hooks`) and in CI by the `commits` job, which runs the same hook on every pushed commit and on pull request titles. (A GitHub ruleset can't do it: this account doesn't support commit message rules.)

Examples:

```
feat(name-matcher): match names with accents across systems
fix: drop the redundant newlines that go vet rejects
build(deps): update strsim to v0.2.1
docs(readme): refresh the quoted outputs
```

## Pull requests

- Title follows the same format: it becomes the squash commit, and the `commits` job checks it.
- Description: short summary of the change and how it was verified. No AI attribution.

## Branches

`<type>/<short-kebab-description>`, such as `feat/address-matcher`, `fix/dedup-threshold`.
