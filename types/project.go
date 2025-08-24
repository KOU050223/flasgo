package types

// プロジェクト設定
type ProjectConfig struct {
	Name      string   // プロジェクト名
	Type      string   // アプリタイプ (hello, webapp, api, fullstack)
	Structure string   // プロジェクト構造 (simple, standard, blueprint)
	Features  []string // 追加機能 (database, auth, forms, env)
	Path      string   // 作成先パス
}

// アプリタイプの定義
var AppTypes = []struct {
	Value string
	Label string
}{
	{"hello", "Hello World (シンプルな1ファイル)"},
	{"webapp", "Webアプリ (HTML forms + templates)"},
	{"api", "REST API (JSONレスポンス)"},
	{"fullstack", "フルスタック (Web + API)"},
}

// プロジェクト構造の定義
var ProjectStructures = []struct {
	Value string
	Label string
}{
	{"simple", "シンプル (app.py 1ファイル)"},
	{"standard", "標準構造 (app/, templates/, static/)"},
	{"blueprint", "Blueprint構造 (大規模プロジェクト向け)"},
}

// 追加機能の定義
var AdditionalFeatures = []struct {
	Value string
	Label string
}{
	{"database", "データベース (SQLAlchemy)"},
	{"auth", "認証機能 (Flask-Login)"},
	{"forms", "フォーム処理 (Flask-WTF)"},
	{"env", "環境変数管理 (.env)"},
}