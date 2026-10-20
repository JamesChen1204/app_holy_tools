package main

import (
	"context"
	"powermoat_app_tools/config"
)

// App struct
type App struct {
	ctx           context.Context
	configService *config.ConfigService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		configService: config.NewConfigService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 載入設定檔
	if err := a.configService.LoadConfig(); err != nil {
		// 記錄錯誤但不中斷啟動
		println("Failed to load config:", err.Error())
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return "Hello " + name + ", It's show time!"
}

// GetConfig 取得應用程式設定
func (a *App) GetConfig() *config.Config {
	return a.configService.GetConfig()
}

// SaveConfig 儲存應用程式設定
func (a *App) SaveConfig(cfg *config.Config) error {
	return a.configService.UpdateConfig(cfg)
}

// TestDatabaseConnection 測試資料庫連線
func (a *App) TestDatabaseConnection(db *config.Database) *config.ConnectionTestResult {
	return config.TestConnection(db)
}
