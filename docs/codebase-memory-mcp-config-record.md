# codebase-memory-mcp 安裝配置指令記錄

日期：2026-07-07

目標：安裝並配置 DeusData/codebase-memory-mcp，讓 Codex CLI 與 VS Code 可透過 MCP 使用 codebase knowledge graph。

## 安裝資訊

- 套件：DeusData/codebase-memory-mcp
- 安裝版本：`codebase-memory-mcp 0.8.1`
- 平台：Linux x86_64
- 安裝類型：UI variant
- Binary 路徑：`/home/franco/.local/bin/codebase-memory-mcp`
- UI 預設連接埠：`9749`

## 已執行指令

```bash
curl -fsSL https://raw.githubusercontent.com/DeusData/codebase-memory-mcp/main/install.sh | bash -s -- --ui
```

安裝腳本完成事項：

- 下載 `codebase-memory-mcp-ui-linux-amd64-portable.tar.gz`
- 驗證 SHA-256 checksum
- 安裝 binary 至 `/home/franco/.local/bin/codebase-memory-mcp`
- 執行 agent auto-configuration
- 偵測並配置 Codex CLI 與 VS Code
- 將 `/home/franco/.local/bin` 加入 `/home/franco/.bashrc`

## Codex CLI 配置

設定檔：`/home/franco/.codex/config.toml`

```toml
[mcp_servers.codebase-memory-mcp]
command = "/home/franco/.local/bin/codebase-memory-mcp"

[[hooks.SessionStart]]
matcher = "startup|resume|clear|compact"

[[hooks.SessionStart.hooks]]
type = "command"
command = 'echo "Code discovery: prefer codebase-memory-mcp (search_graph, trace_path, get_code_snippet, query_graph, search_code) over grep/file-read; run index_repository first if the project is not indexed."'
```

指令提示檔：`/home/franco/.codex/AGENTS.md`

該檔案已加入 codebase-memory-mcp 使用規則，優先順序為：

1. `search_graph`
2. `trace_path`
3. `get_code_snippet`
4. `query_graph`
5. `get_architecture`

## VS Code MCP 配置

設定檔：`/home/franco/.config/Code/User/mcp.json`

```json
{
  "servers": {
    "codebase-memory-mcp": {
      "type": "stdio",
      "command": "/home/franco/.local/bin/codebase-memory-mcp"
    }
  }
}
```

## 全域 codebase-memory-mcp 配置

已執行：

```bash
/home/franco/.local/bin/codebase-memory-mcp config set auto_index true
/home/franco/.local/bin/codebase-memory-mcp config set auto_index_limit 50000
```

目前配置：

```text
Configuration:
  auto_index                = true
  auto_index_limit          = 50000
```

## 驗證指令

```bash
/home/franco/.local/bin/codebase-memory-mcp --version
/home/franco/.local/bin/codebase-memory-mcp --help
/home/franco/.local/bin/codebase-memory-mcp config list
/home/franco/.local/bin/codebase-memory-mcp cli index_repository '{"repo_path":"/home/franco/Documents/Pojects/powermoat_installer"}'
```

驗證結果：

```text
codebase-memory-mcp 0.8.1
```

專案索引驗證結果：

```text
project: home-franco-Documents-Pojects-powermoat_installer
status: indexed
nodes: 3787
edges: 4791
excluded dirs: .git, netframework, build, frontend/dist, frontend/node_modules
```

## 使用方式

重新啟動 Codex CLI 或 VS Code 後，MCP server 會以 stdio 啟動。

第一次使用專案時，可要求 agent 執行 `index_repository`。目前已啟用 `auto_index=true`，因此新專案首次連線時也會自動索引。

若要啟動圖形化 UI，可使用：

```bash
/home/franco/.local/bin/codebase-memory-mcp --ui=true --port=9749
```

瀏覽器開啟：

```text
http://localhost:9749
```

## 維護命令

更新：

```bash
/home/franco/.local/bin/codebase-memory-mcp update
```

解除 agent 配置：

```bash
/home/franco/.local/bin/codebase-memory-mcp uninstall
```

注意：`uninstall` 會移除 agent configs、skills、hooks、instructions，但不會移除 binary 或 SQLite database。
