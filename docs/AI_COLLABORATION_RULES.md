# AI 協作規則

## 重要：所有 AI 助手必須遵守此規則

### 專案管理結構

本專案採用 **SPRINT 敏捷開發管理方式**，所有 AI 助手在協助開發時必須遵守以下規則：

## 📁 文檔結構

```
docs/
├── AI_COLLABORATION_RULES.md    # 本文件 - AI 協作規則
├── requirements/                 # 需求管理
│   ├── BACKLOG.md               # 產品待辦清單
│   ├── CURRENT_SPRINT.md        # 當前 Sprint 內容
│   └── [feature-name].md        # 個別功能需求文件
├── design/                       # 設計文件
│   ├── ARCHITECTURE.md          # 系統架構設計
│   ├── UI_UX.md                 # UI/UX 設計
│   └── API_DESIGN.md            # API 設計（如需要）
└── sprints/                      # Sprint 記錄
    ├── sprint-001/              # Sprint 編號
    │   ├── PLANNING.md          # Sprint 規劃
    │   ├── DAILY.md             # 每日進度
    │   ├── REVIEW.md            # Sprint 回顧
    │   └── RETROSPECTIVE.md     # Sprint 反思
    └── sprint-002/
        └── ...
```

## 🤖 AI 助手工作流程

### 1. 接收新需求時

**必須執行以下步驟（不詢問，直接執行）：**

1. **記錄需求**
   - 在 `docs/requirements/BACKLOG.md` 中新增需求項目
   - 如果是重要功能，在 `docs/requirements/` 建立獨立需求文件

2. **評估並分配 Sprint**
   - 檢查 `docs/requirements/CURRENT_SPRINT.md`
   - 決定是否加入當前 Sprint 或下個 Sprint
   - 更新相應文件

3. **設計階段**
   - 如果涉及架構變更，更新 `docs/design/ARCHITECTURE.md`
   - 如果涉及 UI 變更，更新 `docs/design/UI_UX.md`

4. **實作階段**
   - 記錄每日進度到當前 Sprint 的 `DAILY.md`
   - 完成後更新 Sprint 狀態

### 2. 實作原則

- ✅ **直接行動**：不詢問確認，直接實作
- ✅ **文檔優先**：先更新文檔，再寫代碼
- ✅ **追蹤進度**：每次重要變更都記錄到 DAILY.md
- ✅ **保持同步**：代碼與文檔必須保持一致

### 3. Sprint 管理規則

#### Sprint 週期
- 每個 Sprint 持續 1-2 週
- Sprint 編號格式：`sprint-XXX`（從 001 開始）

#### Sprint 文件更新時機

**PLANNING.md** - Sprint 開始時
- 列出所有計劃任務
- 設定優先級
- 估算工作量

**DAILY.md** - 每次實作時
- 記錄完成的任務
- 記錄遇到的問題
- 記錄決策

**REVIEW.md** - Sprint 結束時
- 總結完成項目
- 列出未完成項目
- 展示成果

**RETROSPECTIVE.md** - Sprint 結束時
- 做得好的地方
- 需要改進的地方
- 下個 Sprint 的行動計劃

### 4. 需求變更處理

當用戶提出變更時：

1. 立即更新 `BACKLOG.md`
2. 評估影響範圍
3. 更新相關設計文件
4. 在 `DAILY.md` 記錄變更原因
5. 執行變更
6. 更新測試文件（如有）

### 5. 文檔撰寫規範

- 使用繁體中文
- 使用 Markdown 格式
- 包含日期時間戳記
- 清楚標註狀態（待辦/進行中/完成/擱置）
- 記錄決策理由

### 6. 優先級定義

- **P0 (Critical)**: 阻塞性問題，立即處理
- **P1 (High)**: 高優先級功能，本 Sprint 完成
- **P2 (Medium)**: 中優先級，可以延後
- **P3 (Low)**: 低優先級，有時間再做

### 7. 狀態標記

- 🔴 **Blocked**: 被阻塞
- 🟡 **In Progress**: 進行中
- 🟢 **Done**: 已完成
- ⚪ **Todo**: 待辦
- 🔵 **Review**: 待審查

## 🚀 快速範例

### 用戶說：「加入一個登入功能」

AI 應該：

1. ✅ 在 `BACKLOG.md` 新增「用戶登入功能」
2. ✅ 建立 `requirements/user-authentication.md` 詳細需求
3. ✅ 更新 `ARCHITECTURE.md` 說明認證流程
4. ✅ 更新 `CURRENT_SPRINT.md` 將任務加入
5. ✅ 在 `sprints/sprint-XXX/DAILY.md` 記錄開始時間
6. ✅ 實作程式碼
7. ✅ 更新 `DAILY.md` 記錄完成

❌ **不要問**：「你要我加入登入功能嗎？」
❌ **不要問**：「需要哪種認證方式？」（自己決定最佳實踐）

## 📌 注意事項

1. **自主決策**：在合理範圍內做出技術決策，並在文檔中說明理由
2. **保持簡潔**：文檔要清晰但不冗長
3. **實用至上**：不要過度設計，focus 在可用的功能
4. **持續更新**：每次協作都更新相關文檔
5. **版本追蹤**：重要變更在文檔中記錄版本號

## 🔄 版本歷史

- v1.0.0 (2026-06-03) - 初始版本，建立 AI 協作規則
