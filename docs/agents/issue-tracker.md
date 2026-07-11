# Issue tracker: Linear

Issues and PRDs for this repo live in **Linear**, accessed through the **Linear MCP tools** (`mcp__claude_ai_Linear__*`). There is no CLI — call the MCP tools directly.

## Configuration

- **Team**: **Docudex** — key `DEX`, id `733dcec1-824d-44a6-b023-0f46777b6a8a`. Workspace: `linear.app/docudex`. Every `save_issue` that creates an issue needs this `teamId`; issue identifiers look like `DEX-123`.
- **GitHub sync**: not yet enabled. Once Linear's GitHub Issue Sync is configured (Settings → Integrations → GitHub in Linear's dashboard), issues created here mirror to `aureliushq/docudex` automatically.

## Conventions

- **Create an issue**: `save_issue` with `title`, `description` (markdown — pass real newlines, not `\n`), and the team's `teamId`. Omit `id` to create.
- **Read an issue**: `get_issue` by id/identifier (e.g. `DOC-123`). Use `list_comments` for its discussion.
- **List issues**: `list_issues` with filters (`team`, `state`, `label`, `assignee`, `query`). Returns identifier, title, state, labels.
- **Comment on an issue**: `save_comment` with the issue id and markdown `body`.
- **Apply / remove labels**: `save_issue` with the updated `labelIds`. Discover label ids via `list_issue_labels`. Labels are created in Linear's UI or via `create_issue_label` if missing.
- **Change state / close**: `save_issue` with the target `stateId` (resolve via `list_issue_statuses` — e.g. a `Done`/`Canceled` workflow state). Linear has no "close"; move the issue to a completed or canceled state.
- **Assign**: `save_issue` with `assigneeId` (resolve via `list_users`).

## Triage labels

The five canonical triage roles map to **Linear labels** of the same names — see `triage-labels.md`. Create them in Linear (or via `create_issue_label`) if they don't exist, then apply via `labelIds` on `save_issue`.

## When a skill says "publish to the issue tracker"

Create a Linear issue with `save_issue` under the configured team.

## When a skill says "fetch the relevant ticket"

Call `get_issue` for the identifier, plus `list_comments` for its discussion.

## Wayfinding operations

Used by `/wayfinder`. The **map** is a parent Linear issue with **child** issues (sub-issues) as tickets.

- **Map**: a Linear issue labelled `wayfinder:map` holding the Notes / Decisions-so-far / Fog body. Create with `save_issue`.
- **Child ticket**: a Linear sub-issue — `save_issue` with `parentId` set to the map's id. Labels: `wayfinder:<type>` (`research`/`prototype`/`grilling`/`task`). Once claimed, assign it to the driving dev via `assigneeId`.
- **Blocking**: use Linear's native issue relations (`blocks` / `blocked by`). Where a relation can't be set, fall back to a `Blocked by: DOC-<n>, DOC-<n>` line at the top of the child body. A ticket is unblocked when every blocker is in a completed/canceled state.
- **Frontier query**: `list_issues` scoped to the map's children in an unstarted/started state; drop any with an open blocker or an assignee; first in map order wins.
- **Claim**: `save_issue` setting `assigneeId` to the current user — the session's first write.
- **Resolve**: `save_comment` with the answer, move the child to a completed state via `save_issue`, then append a context pointer to the map's Decisions-so-far.
