---
title: PostgreSQL
---

```bash
miasma apps:create postgres -i postgres:13-alpine
miasma env:edit -a postgres
miasma apps:configure \
    --add-placement-constraint "node.labels.database == true" \
    --add-volume /dir/path/on/physical/machine:/var/lib/postgresql/data \
    --add-target-ports 5432 --add-published-ports 5432
```
