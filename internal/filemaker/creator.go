package filemaker

import (
	"bytes"
	"fmt"
	"github.com/KOU050223/flasgo/internal/templates"
	"github.com/KOU050223/flasgo/types"
	"os"
	"path/filepath"
	"text/template"
)

// テンプレート用のデータ構造
type TemplateData struct {
	ProjectName string
	HasDatabase bool
	HasAuth     bool
	HasForms    bool
	HasEnv      bool
}

// プロジェクト作成のメイン関数
func createProject(config *types.ProjectConfig) error {
	// プロジェクトディレクトリが既に存在するかチェック
	if _, err := os.Stat(config.Name); !os.IsNotExist(err) {
		return fmt.Errorf("プロジェクトディレクトリ '%s' は既に存在します", config.Name)
	}

	// プロジェクトディレクトリを作成
	if err := os.MkdirAll(config.Name, 0755); err != nil {
		return fmt.Errorf("プロジェクトディレクトリの作成に失敗: %v", err)
	}

	// テンプレートデータを準備
	templateData := prepareTemplateData(config)

	// プロジェクト構造に応じてファイルを作成
	switch config.Structure {
	case "simple":
		return createSimpleStructure(config, templateData)
	case "standard":
		return createStandardStructure(config, templateData)
	case "blueprint":
		return createBlueprintStructure(config, templateData)
	default:
		return fmt.Errorf("不明なプロジェクト構造: %s", config.Structure)
	}
}

// テンプレートデータを準備
func prepareTemplateData(config *types.ProjectConfig) *TemplateData {
	data := &TemplateData{
		ProjectName: config.Name,
	}

	for _, feature := range config.Features {
		switch feature {
		case "database":
			data.HasDatabase = true
		case "auth":
			data.HasAuth = true
		case "forms":
			data.HasForms = true
		case "env":
			data.HasEnv = true
		}
	}

	return data
}

// シンプル構造（1ファイル）を作成
func createSimpleStructure(config *types.ProjectConfig, data *TemplateData) error {
	appPath := filepath.Join(config.Name, "app.py")

	// アプリタイプに応じたテンプレートを選択
	var content string
	switch config.Type {
	case "hello":
		content = templates.HelloWorldApp
	case "webapp":
		content = processTemplate(templates.WebAppMain, data)
	case "api":
		content = processTemplate(templates.APIMain, data)
	default:
		content = templates.HelloWorldApp
	}

	if err := writeFile(appPath, content); err != nil {
		return err
	}

	// requirements.txtを作成
	reqPath := filepath.Join(config.Name, "requirements.txt")
	reqContent := templates.GenerateRequirements(config.Features, config.Type)
	if err := writeFile(reqPath, reqContent); err != nil {
		return err
	}

	// .envファイルを作成（必要な場合）
	if data.HasEnv {
		envPath := filepath.Join(config.Name, ".env")
		if err := writeFile(envPath, templates.EnvTemplate); err != nil {
			return err
		}
	}

	// README.mdを作成
	readmePath := filepath.Join(config.Name, "README.md")
	readmeContent := templates.GenerateReadme(config.Name, config.Type, data.HasDatabase, data.HasForms)
	if err := writeFile(readmePath, readmeContent); err != nil {
		return err
	}

	// .gitignoreを作成
	gitignorePath := filepath.Join(config.Name, ".gitignore")
	if err := writeFile(gitignorePath, templates.GitignoreTemplate); err != nil {
		return err
	}

	return nil
}

// 標準構造を作成
func createStandardStructure(config *types.ProjectConfig, data *TemplateData) error {
	// ディレクトリ構造を作成
	dirs := []string{
		filepath.Join(config.Name, "templates"),
		filepath.Join(config.Name, "static", "css"),
		filepath.Join(config.Name, "static", "js"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("ディレクトリ作成エラー: %v", err)
		}
	}

	// app.pyを作成
	appPath := filepath.Join(config.Name, "app.py")
	var appContent string
	switch config.Type {
	case "webapp":
		appContent = processTemplate(templates.WebAppMain, data)
	case "api":
		appContent = processTemplate(templates.APIMain, data)
	default:
		appContent = processTemplate(templates.WebAppMain, data)
	}

	if err := writeFile(appPath, appContent); err != nil {
		return err
	}

	// HTMLテンプレートを作成（APIタイプでない場合）
	if config.Type != "api" {
		baseTemplatePath := filepath.Join(config.Name, "templates", "base.html")
		if err := writeFile(baseTemplatePath, templates.BaseTemplate); err != nil {
			return err
		}

		indexTemplatePath := filepath.Join(config.Name, "templates", "index.html")
		if err := writeFile(indexTemplatePath, templates.IndexTemplate); err != nil {
			return err
		}

		// フォーム機能がある場合
		if data.HasForms {
			formTemplatePath := filepath.Join(config.Name, "templates", "form.html")
			if err := writeFile(formTemplatePath, templates.FormTemplate); err != nil {
				return err
			}
		}
	}

	// requirements.txtを作成
	reqPath := filepath.Join(config.Name, "requirements.txt")
	reqContent := templates.GenerateRequirements(config.Features, config.Type)
	if err := writeFile(reqPath, reqContent); err != nil {
		return err
	}

	// .envファイルを作成（必要な場合）
	if data.HasEnv {
		envPath := filepath.Join(config.Name, ".env")
		if err := writeFile(envPath, templates.EnvTemplate); err != nil {
			return err
		}
	}

	// README.mdを作成
	readmePath := filepath.Join(config.Name, "README.md")
	readmeContent := templates.GenerateReadme(config.Name, config.Type, data.HasDatabase, data.HasForms)
	if err := writeFile(readmePath, readmeContent); err != nil {
		return err
	}

	// .gitignoreを作成
	gitignorePath := filepath.Join(config.Name, ".gitignore")
	if err := writeFile(gitignorePath, templates.GitignoreTemplate); err != nil {
		return err
	}

	return nil
}

// Blueprint構造を作成（今後実装）
func createBlueprintStructure(config *types.ProjectConfig, data *TemplateData) error {
	// とりあえず標準構造と同じにしておく
	return createStandardStructure(config, data)
}

// テンプレートを処理
func processTemplate(templateStr string, data *TemplateData) string {
	tmpl, err := template.New("flask").Parse(templateStr)
	if err != nil {
		return templateStr // エラーの場合は元のテンプレートを返す
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return templateStr // エラーの場合は元のテンプレートを返す
	}

	return buf.String()
}

// ファイルに内容を書き込む
func writeFile(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("ファイル作成エラー (%s): %v", path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("ファイル書き込みエラー (%s): %v", path, err)
	}

	return nil
}