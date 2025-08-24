package filemaker

import (
	"fmt"
	"github.com/KOU050223/flasgo/internal/ui"
	"github.com/KOU050223/flasgo/types"
)

func Generator() {
	fmt.Println("✨ Flaskプロジェクトを作成します\n")
	
	// プロジェクト設定を収集
	config := collectProjectConfig()
	
	fmt.Printf("\n📁 プロジェクトを作成中...\n")
	
	// プロジェクト作成（実装予定）
	err := createProject(config)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		return
	}
	
	fmt.Printf("✅ %s プロジェクトが作成されました！\n", config.Name)
	printNextSteps(config)
}

// プロジェクト設定を対話的に収集
func collectProjectConfig() *types.ProjectConfig {
	config := &types.ProjectConfig{}
	
	// プロジェクト名
	config.Name = ui.PromptText("プロジェクト名", "myflaskapp")
	
	// アプリタイプ選択
	appOptions := make([]ui.Option, len(types.AppTypes))
	for i, appType := range types.AppTypes {
		appOptions[i] = ui.Option{Label: appType.Label, Value: appType.Value}
	}
	config.Type = ui.PromptSelect("どのタイプのFlaskアプリを作成しますか？", appOptions)
	
	// プロジェクト構造選択
	structOptions := make([]ui.Option, len(types.ProjectStructures))
	for i, structure := range types.ProjectStructures {
		structOptions[i] = ui.Option{Label: structure.Label, Value: structure.Value}
	}
	config.Structure = ui.PromptSelect("プロジェクト構造を選択してください", structOptions)
	
	// 追加機能選択
	featureOptions := make([]ui.Option, len(types.AdditionalFeatures))
	for i, feature := range types.AdditionalFeatures {
		featureOptions[i] = ui.Option{Label: feature.Label, Value: feature.Value}
	}
	config.Features = ui.PromptMultiSelect("追加機能を選択してください", featureOptions)
	
	return config
}

// デフォルト設定でプロジェクト生成（非対話モード）
func GenerateWithDefaults(projectName string) {
	fmt.Printf("✨ Flaskプロジェクト '%s' をデフォルト設定で作成します\n\n", projectName)
	
	// デフォルト設定
	config := &types.ProjectConfig{
		Name:      projectName,
		Type:      "webapp",  // デフォルトはWebアプリ
		Structure: "standard", // デフォルトは標準構造
		Features:  []string{"env"}, // デフォルトは環境変数管理のみ
	}
	
	fmt.Printf("設定:\n")
	fmt.Printf("  タイプ: %s (Webアプリ)\n", config.Type)
	fmt.Printf("  構造: %s (標準構造)\n", config.Structure)
	fmt.Printf("  機能: %v\n", config.Features)
	
	fmt.Printf("\n📁 プロジェクトを作成中...\n")
	
	// プロジェクト作成
	err := createProject(config)
	if err != nil {
		fmt.Printf("❌ エラー: %v\n", err)
		return
	}
	
	fmt.Printf("✅ %s プロジェクトが作成されました！\n", config.Name)
	printNextSteps(config)
}

// 次のステップを表示
func printNextSteps(config *types.ProjectConfig) {
	fmt.Printf("\n次のステップ:\n")
	fmt.Printf("  cd %s\n", config.Name)
	fmt.Printf("  pip install -r requirements.txt\n")
	fmt.Printf("  flask run\n")
}
