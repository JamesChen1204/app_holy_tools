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

## Mermaid Diagrams

The `mermaid/` folder contains Mermaid diagrams derived from or related to the codebase. Agents may inspect these files when they need to understand program logic, installation flow, state transitions, or cross-module behavior before changing code.

Use the diagrams as a companion to the MCP graph:

- Start with MCP graph tools for indexed code discovery.
- Check `mermaid/` when visual flow or state-machine context would clarify the logic.
- Verify important conclusions against source code before making behavioral changes, because diagrams can lag behind implementation.
<!-- codebase-memory-mcp:end -->
