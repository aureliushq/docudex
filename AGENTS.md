## Agent skills

### Issue tracker

Issues live in Linear, managed via the Linear MCP tools. See `docs/agents/issue-tracker.md`. _(Bidirectional GitHub sync to be set up later via Linear's GitHub integration.)_

### Triage labels

Default five canonical triage labels. See `docs/agents/triage-labels.md`.

### Domain docs

Single-context: `CONTEXT.md` + `docs/adr/` at the repo root. See `docs/agents/domain.md`.

## Git conventions

- **Branch names**: `<type>/engg-<n>-<short-slug>` — `<type>` is `feat`, `fix`, or `chore`; `<short-slug>` is a 2–4-word kebab summary. Derive `<n>` from the Linear issue (`ENGG-<n>`). Example: `feat/engg-6-scaffolding`.
