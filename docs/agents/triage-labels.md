# Triage Labels

The skills speak in terms of five canonical triage roles. This file maps those roles to the actual label strings used in this repo's issue tracker.

| Canonical role    | Linear label      | Linear label id                        | Meaning                                  |
| ----------------- | ----------------- | -------------------------------------- | ---------------------------------------- |
| `needs-triage`    | `needs-triage`    | `589bbc4c-27dc-44ce-b337-ed28ec5a187c` | Maintainer needs to evaluate this issue  |
| `needs-info`      | `needs-info`      | `0184f2dd-5c4f-46ef-af76-92cb507ec36e` | Waiting on reporter for more information |
| `ready-for-agent` | `ready-for-agent` | `131afc19-cdf3-4058-bebc-8fdf5bb73c33` | Fully specified, ready for an AFK agent  |
| `ready-for-human` | `ready-for-human` | `1a957635-5a32-42de-8161-8211792c1fbd` | Requires human implementation            |
| `wontfix`         | `wontfix`         | `ed287dfe-a0e8-43c7-b943-ffcf426d6188` | Will not be actioned                     |

When a skill mentions a role (e.g. "apply the AFK-ready triage label"), apply the corresponding label from this table. Pass the label id in `labelIds` on `save_issue`.

All five labels exist in the Engineering team (`ENGG`). Edit the middle/right columns if you remap to different vocabulary later.
