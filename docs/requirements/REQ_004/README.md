# REQ_004: 應用程式國際化 (i18n)

## 需求概述

為應用程式添加國際化（Internationalization, i18n）支援，使用配置文件切換語言，並處理應用內所有文字內容。

## 需求背景

目前應用程式所有文字都是硬編碼的繁體中文，無法切換語言。需要實現國際化機制，讓用戶可以在不同語言間切換。

## 功能需求

### 1. 支援語言

- **繁體中文 (zh-TW)**：預設語言
- **簡體中文 (zh-CN)**：簡體版本
- **英文 (en)**：國際通用語言

### 2. 語言切換機制

- 在「一般設定」頁面提供語言選擇器
- 語言設定儲存在配置文件中（與 REQ_003 整合）
- 切換語言後即時更新所有介面文字
- 應用重啟後保持上次選擇的語言

### 3. 涉及範圍

需要國際化的所有文字內容：

#### 側邊欄 (Sidebar.jsx)
- 選單標題：設置、關於
- 子選單：一般設定、進階功能設定

#### 一般設定頁面 (GeneralSettingsPage)
- 頁面標題：一般設定
- 頁面描述：應用的基本配置
- 區塊標題：外觀
- 欄位標籤：主題模式、語言
- 下拉選項：
  - 主題：淺色、深色、跟隨系統
  - 語言：繁體中文、简体中文、English
- 佔位符：選擇主題、選擇語言

#### 進階功能設定頁面 (AdvancedSettingsPage)
- 頁面標題：進階功能設定
- 頁面描述：進階功能配置
- 區塊標題：資料庫連線設定
- 區塊描述：配置應用程式的資料庫連線參數
- 欄位標籤：
  - 資料庫類型
  - 主機位址
  - 連接埠
  - 使用者名稱
  - 密碼
  - 資料庫名稱
- 佔位符：
  - localhost
  - 1433
  - username
  - password
  - database_name
- 按鈕文字：
  - 測試連線 / 測試中...
  - 儲存設定 / 儲存中...
  - 顯示 / 隱藏（密碼切換）
- 資料庫類型選項：
  - Microsoft SQL Server
  - MySQL
- 測試結果訊息：
  - 連線成功
  - 連線失敗
  - 設定已儲存
  - 儲存失敗
  - 連線延遲: Xms

#### 關於頁面 (AboutPage)
- 頁面標題和所有文字內容
- 版本資訊、作者、授權等

#### 其他元件
- 錯誤訊息
- 提示訊息
- 確認對話框

## 技術方案

### 1. 前端實現

#### 使用 i18next

```bash
npm install i18next react-i18next
```

#### 目錄結構

```
frontend/src/
├── locales/
│   ├── zh-TW/
│   │   └── translation.json
│   ├── zh-CN/
│   │   └── translation.json
│   └── en/
│       └── translation.json
├── i18n.js
└── ...
```

#### 配置檔案範例

**frontend/src/locales/zh-TW/translation.json**
```json
{
  "sidebar": {
    "settings": "設置",
    "about": "關於",
    "general": "一般設定",
    "advanced": "進階功能設定"
  },
  "settings": {
    "general": {
      "title": "一般設定",
      "description": "應用的基本配置",
      "appearance": "外觀",
      "theme": {
        "label": "主題模式",
        "placeholder": "選擇主題",
        "light": "淺色",
        "dark": "深色",
        "system": "跟隨系統"
      },
      "language": {
        "label": "語言",
        "placeholder": "選擇語言",
        "zhTW": "繁體中文",
        "zhCN": "简体中文",
        "en": "English"
      }
    },
    "advanced": {
      "title": "進階功能設定",
      "description": "進階功能配置",
      "database": {
        "title": "資料庫連線設定",
        "description": "配置應用程式的資料庫連線參數",
        "type": "資料庫類型",
        "host": "主機位址",
        "port": "連接埠",
        "username": "使用者名稱",
        "password": "密碼",
        "database": "資料庫名稱",
        "testConnection": "測試連線",
        "testing": "測試中...",
        "saveConfig": "儲存設定",
        "saving": "儲存中...",
        "show": "顯示",
        "hide": "隱藏",
        "mssql": "Microsoft SQL Server",
        "mysql": "MySQL",
        "connectionSuccess": "連線成功",
        "connectionFailed": "連線失敗",
        "configSaved": "設定已儲存",
        "saveFailed": "儲存失敗",
        "latency": "連線延遲: {{value}}ms"
      }
    }
  },
  "about": {
    "title": "關於",
    "description": "關於本應用程式"
  }
}
```

**frontend/src/locales/zh-CN/translation.json**
```json
{
  "sidebar": {
    "settings": "设置",
    "about": "关于",
    "general": "一般设定",
    "advanced": "高级功能设定"
  },
  "settings": {
    "general": {
      "title": "一般设定",
      "description": "应用的基本配置",
      "appearance": "外观",
      "theme": {
        "label": "主题模式",
        "placeholder": "选择主题",
        "light": "浅色",
        "dark": "深色",
        "system": "跟随系统"
      },
      "language": {
        "label": "语言",
        "placeholder": "选择语言",
        "zhTW": "繁体中文",
        "zhCN": "简体中文",
        "en": "English"
      }
    },
    "advanced": {
      "title": "高级功能设定",
      "description": "高级功能配置",
      "database": {
        "title": "数据库连接设定",
        "description": "配置应用程序的数据库连接参数",
        "type": "数据库类型",
        "host": "主机地址",
        "port": "端口",
        "username": "用户名",
        "password": "密码",
        "database": "数据库名称",
        "testConnection": "测试连接",
        "testing": "测试中...",
        "saveConfig": "保存设定",
        "saving": "保存中...",
        "show": "显示",
        "hide": "隐藏",
        "mssql": "Microsoft SQL Server",
        "mysql": "MySQL",
        "connectionSuccess": "连接成功",
        "connectionFailed": "连接失败",
        "configSaved": "设定已保存",
        "saveFailed": "保存失败",
        "latency": "连接延迟: {{value}}ms"
      }
    }
  },
  "about": {
    "title": "关于",
    "description": "关于本应用程序"
  }
}
```

**frontend/src/locales/en/translation.json**
```json
{
  "sidebar": {
    "settings": "Settings",
    "about": "About",
    "general": "General",
    "advanced": "Advanced"
  },
  "settings": {
    "general": {
      "title": "General Settings",
      "description": "Basic application configuration",
      "appearance": "Appearance",
      "theme": {
        "label": "Theme",
        "placeholder": "Select theme",
        "light": "Light",
        "dark": "Dark",
        "system": "System"
      },
      "language": {
        "label": "Language",
        "placeholder": "Select language",
        "zhTW": "繁體中文",
        "zhCN": "简体中文",
        "en": "English"
      }
    },
    "advanced": {
      "title": "Advanced Settings",
      "description": "Advanced configuration",
      "database": {
        "title": "Database Connection",
        "description": "Configure database connection parameters",
        "type": "Database Type",
        "host": "Host",
        "port": "Port",
        "username": "Username",
        "password": "Password",
        "database": "Database",
        "testConnection": "Test Connection",
        "testing": "Testing...",
        "saveConfig": "Save Settings",
        "saving": "Saving...",
        "show": "Show",
        "hide": "Hide",
        "mssql": "Microsoft SQL Server",
        "mysql": "MySQL",
        "connectionSuccess": "Connection successful",
        "connectionFailed": "Connection failed",
        "configSaved": "Settings saved",
        "saveFailed": "Save failed",
        "latency": "Latency: {{value}}ms"
      }
    }
  },
  "about": {
    "title": "About",
    "description": "About this application"
  }
}
```

#### i18n 初始化

**frontend/src/i18n.js**
```javascript
import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

import zhTW from './locales/zh-TW/translation.json';
import zhCN from './locales/zh-CN/translation.json';
import en from './locales/en/translation.json';

i18n
  .use(initReactI18next)
  .init({
    resources: {
      'zh-TW': { translation: zhTW },
      'zh-CN': { translation: zhCN },
      'en': { translation: en }
    },
    lng: 'zh-TW', // 預設語言
    fallbackLng: 'zh-TW',
    interpolation: {
      escapeValue: false
    }
  });

export default i18n;
```

#### 在組件中使用

```javascript
import { useTranslation } from 'react-i18next';

function Sidebar({ onNavigate }) {
  const { t } = useTranslation();
  
  const menuItems = [
    {
      id: 'settings',
      label: t('sidebar.settings'),
      icon: '⚙️',
      children: [
        { id: 'general', label: t('sidebar.general'), page: 'settings-general' },
        { id: 'advanced', label: t('sidebar.advanced'), page: 'settings-advanced' }
      ]
    }
  ];
  
  // ...
}
```

### 2. 與配置系統整合

#### 修改 config/models.go

```go
type Application struct {
    Language string `json:"language"` // "zh-TW" | "zh-CN" | "en"
    Theme    string `json:"theme"`
}
```

#### 修改 GeneralSettingsPage

- 載入時從配置讀取語言設定
- 切換語言時：
  1. 更新 i18n 實例的語言
  2. 儲存到配置文件
  3. 重新渲染所有組件

```javascript
const handleLanguageChange = async (lang) => {
  // 更新 i18n
  i18n.changeLanguage(lang);
  
  // 儲存到配置
  const config = await GetConfig();
  config.application.language = lang;
  await SaveConfig(config);
};
```

## 實作步驟

### Phase 1: 基礎設置
1. 安裝 i18next 和 react-i18next
2. 創建語言配置文件目錄結構
3. 創建 i18n.js 初始化文件
4. 在 main.jsx 中導入 i18n

### Phase 2: 建立翻譯檔案
1. 分析所有組件，提取文字內容
2. 創建三個語言的 translation.json
3. 確保所有文字都有對應的翻譯鍵

### Phase 3: 組件改造
1. 修改 Sidebar.jsx 使用 t() 函數
2. 修改 GeneralSettingsPage.jsx
3. 修改 AdvancedSettingsPage.jsx
4. 修改 AboutPage.jsx
5. 處理其他含文字的組件

### Phase 4: 配置整合
1. 從配置文件讀取語言設定
2. 初始化時設定正確語言
3. 實現語言切換功能
4. 儲存語言選擇到配置文件

### Phase 5: 測試
1. 測試三種語言切換
2. 測試配置持久化
3. 測試所有頁面的文字顯示
4. 測試重啟後語言保持

## 驗收標準

- [ ] 應用支援繁體中文、簡體中文、英文三種語言
- [ ] 所有介面文字都可以切換語言
- [ ] 語言選擇儲存在配置文件中
- [ ] 應用重啟後保持上次選擇的語言
- [ ] 切換語言後即時更新所有介面
- [ ] 翻譯內容準確、用詞一致
- [ ] 無硬編碼文字殘留
- [ ] 不影響現有功能

## 預期效果

- 用戶可以根據偏好選擇介面語言
- 提升國際用戶體驗
- 為未來添加更多語言奠定基礎
- 符合國際化應用標準

## 相關文件

- [i18next 文檔](https://www.i18next.com/)
- [react-i18next 文檔](https://react.i18next.com/)
- REQ_003: 應用程式配置文件（語言設定整合）
