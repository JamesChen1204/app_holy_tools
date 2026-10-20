# codebase-memory-mcp Local Artifact 導入指南

本文定義新專案導入 `codebase-memory-mcp` 的固定模式：

- 每個專案都使用 repo-local `.codebase-memory/`。
- `.codebase-memory/graph.db.zst` 要提交到 Git，作為團隊共享的 code graph artifact。
- agent 的即時查詢仍使用本機 MCP SQLite cache，但新環境可從 repo artifact 加速初始化。

## 1. 前置條件

確認本機已安裝 `codebase-memory-mcp`：

```bash
codebase-memory-mcp --version
```

建議開啟自動索引：

```bash
codebase-memory-mcp config set auto_index true
codebase-memory-mcp config set auto_index_limit 50000
```

確認設定：

```bash
codebase-memory-mcp config list
```

## 2. 新專案 AGENTS.md

在新專案根目錄建立或更新 `AGENTS.md`：

```md
<!-- codebase-memory-mcp:start -->
# Codebase Knowledge Graph (codebase-memory-mcp)

This project uses codebase-memory-mcp to maintain a persistent local knowledge graph of the codebase.
Always prefer MCP graph tools over grep/glob/file-search for code discovery.

## Priority Order

1. `search_graph` - find functions, classes, routes, variables, and other indexed symbols by pattern or natural-language query.
2. `trace_path` - trace who calls a function, what it calls, or related call/data-flow paths.
3. `get_code_snippet` - read source for a specific function/class after finding its exact qualified name with `search_graph`.
4. `query_graph` - run Cypher queries for complex relationships, aggregations, or multi-hop analysis.
5. `get_architecture` - get a high-level project summary, packages, dependencies, entry points, routes, hotspots, and clusters.

## When to fall back to grep/glob

- Searching for string literals, error messages, config values, scripts, or non-code files.
- Searching files or content that are not represented in the codebase-memory-mcp index.
- When MCP graph tools return insufficient or stale results.

## Examples

- Find a handler: `search_graph(name_pattern=".*OrderHandler.*")`
- Search by intent: `search_graph(query="update installer settings")`
- Who calls it: `trace_path(function_name="OrderHandler", direction="inbound")`
- Read source: `get_code_snippet(qualified_name="pkg/orders.OrderHandler")`
- Architecture overview: `get_architecture(aspects=["summary", "dependencies", "clusters"])`
<!-- codebase-memory-mcp:end -->
```

## 3. Git ignore policy

本團隊模式是提交 repo-local artifact，所以不要在 `.gitignore` 加入：

```gitignore
.codebase-memory/
```

如果既有專案已經忽略 `.codebase-memory/`，請移除該規則。

建議提交這些檔案：

```text
.codebase-memory/graph.db.zst
.codebase-memory/artifact.json
.codebase-memory/.gitattributes
```

`.codebase-memory/.gitattributes` 會設定：

```gitattributes
graph.db.zst merge=ours binary
```

這能降低壓縮 artifact 發生 merge conflict 時的處理成本。

## 4. 第一次產生 local artifact

在新專案根目錄執行：

```bash
codebase-memory-mcp cli index_repository '{"repo_path":"'"$(pwd)"'","mode":"full","persistence":true}'
```

成功時輸出會包含：

```text
"artifact_present":true
"artifact_hint":"Persistent artifact written to .codebase-memory/graph.db.zst..."
```

確認檔案存在：

```bash
find .codebase-memory -maxdepth 2 -type f -print
```

預期至少看到：

```text
.codebase-memory/graph.db.zst
.codebase-memory/artifact.json
.codebase-memory/.gitattributes
```

## 5. 提交到 Git

確認 `.codebase-memory/graph.db.zst` 沒被 ignore：

```bash
git check-ignore -v .codebase-memory/graph.db.zst
```

若沒有輸出，代表 Git 會追蹤它。

提交：

```bash
git add AGENTS.md .codebase-memory/
git commit -m "Add codebase-memory local artifact"
```

## 6. 更新 local artifact

當專案有較大結構變更，或準備讓團隊共享最新 code graph 時，重新執行：

```bash
codebase-memory-mcp cli index_repository '{"repo_path":"'"$(pwd)"'","mode":"full","persistence":true}'
```

再提交更新後的 artifact：

```bash
git add .codebase-memory/
git commit -m "Refresh codebase-memory artifact"
```

## 7. 新成員使用方式

新成員 clone 專案後：

1. 安裝 `codebase-memory-mcp`。
2. 重新啟動 Codex / VS Code / 其他 agent。
3. 開啟專案。
4. 要求 agent 使用 codebase-memory-mcp 查詢，或執行一次索引：
5. run codebase-memory-mcp --ui=true --port=9749
```bash
codebase-memory-mcp cli index_repository '{"repo_path":"'"$(pwd)"'","mode":"full","persistence":true}'
```

若 repo 內已有 `.codebase-memory/graph.db.zst`，工具可用該 artifact bootstrap，避免每個人都從零開始完整重建 graph。

## 8. 疑難排解

### artifact_present 一直是 false

確認指令有帶：

```json
"persistence": true
```

也要確認 `.codebase-memory/` 沒有被 `.gitignore` 或 global gitignore 擋住。

### Git 沒有顯示 .codebase-memory

檢查 ignore 規則：

```bash
git check-ignore -v .codebase-memory/graph.db.zst
```

若有輸出，移除對應 ignore 規則。

### 專案索引狀態怪異

先列出目前專案：

```bash
codebase-memory-mcp cli list_projects
```

必要時刪除本機索引後重建：

```bash
codebase-memory-mcp cli delete_project '{"project":"PROJECT_NAME"}'
codebase-memory-mcp cli index_repository '{"repo_path":"'"$(pwd)"'","mode":"full","persistence":true}'
```

注意：`delete_project` 只刪 codebase-memory-mcp 的本機索引，不會刪原始碼，也不會刪 repo 內的 `.codebase-memory/`。

## 9. 團隊約定

- 新專案一律加入 `AGENTS.md` 的 codebase-memory-mcp 區塊。
- 新專案一律產生並提交 `.codebase-memory/graph.db.zst`。
- 不要把 `.codebase-memory/` 加入 `.gitignore`。
- 大型重構、模組搬移、API 大改後，要刷新 `.codebase-memory/` artifact。
- 一般小改動可依需求刷新；若 agent 查詢結果明顯 stale，立即刷新。



