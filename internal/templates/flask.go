package templates

// シンプルなHello World Flask アプリ
var HelloWorldApp = `from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return '<h1>Hello, World!</h1>'

@app.route('/about')
def about():
    return '<h1>About Page</h1>'

if __name__ == '__main__':
    app.run(debug=True)
`

// Webアプリ用のapp.py (標準構造向け)
var WebAppMain = `from flask import Flask, render_template, request, flash, redirect, url_for
{{if .HasForms}}from flask_wtf import FlaskForm
from wtforms import StringField, SubmitField
from wtforms.validators import DataRequired{{end}}
{{if .HasDatabase}}from flask_sqlalchemy import SQLAlchemy{{end}}
{{if .HasEnv}}import os
from dotenv import load_dotenv

load_dotenv(){{end}}

app = Flask(__name__)
{{if .HasEnv}}app.config['SECRET_KEY'] = os.environ.get('SECRET_KEY') or 'dev-secret-key'
{{if .HasDatabase}}app.config['SQLALCHEMY_DATABASE_URI'] = os.environ.get('DATABASE_URL') or 'sqlite:///app.db'{{end}}
{{else}}app.config['SECRET_KEY'] = 'your-secret-key-here'
{{if .HasDatabase}}app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///app.db'{{end}}
{{end}}

{{if .HasDatabase}}db = SQLAlchemy(app)

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(80), unique=True, nullable=False)

    def __repr__(self):
        return f'<User {self.name}>'
{{end}}

{{if .HasForms}}class NameForm(FlaskForm):
    name = StringField('Name', validators=[DataRequired()])
    submit = SubmitField('Submit')
{{end}}

@app.route('/')
def index():
    return render_template('index.html')

{{if .HasForms}}@app.route('/form', methods=['GET', 'POST'])
def form():
    form = NameForm()
    if form.validate_on_submit():
        flash(f'Hello {form.name.data}!')
        return redirect(url_for('form'))
    return render_template('form.html', form=form)
{{end}}

{{if .HasDatabase}}@app.route('/users')
def users():
    users = User.query.all()
    return render_template('users.html', users=users)
{{end}}

if __name__ == '__main__':
    {{if .HasDatabase}}with app.app_context():
        db.create_all()
    {{end}}app.run(debug=True)
`

// REST API用のapp.py
var APIMain = `from flask import Flask, jsonify, request
{{if .HasDatabase}}from flask_sqlalchemy import SQLAlchemy{{end}}
{{if .HasEnv}}import os
from dotenv import load_dotenv

load_dotenv(){{end}}

app = Flask(__name__)
{{if .HasEnv}}{{if .HasDatabase}}app.config['SQLALCHEMY_DATABASE_URI'] = os.environ.get('DATABASE_URL') or 'sqlite:///api.db'{{end}}
{{else}}{{if .HasDatabase}}app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///api.db'{{end}}
{{end}}

{{if .HasDatabase}}db = SQLAlchemy(app)

class Item(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(80), nullable=False)
    description = db.Column(db.Text)

    def to_dict(self):
        return {
            'id': self.id,
            'name': self.name,
            'description': self.description
        }
{{end}}

@app.route('/api/health')
def health():
    return jsonify({'status': 'ok', 'message': 'API is running'})

{{if .HasDatabase}}@app.route('/api/items', methods=['GET'])
def get_items():
    items = Item.query.all()
    return jsonify([item.to_dict() for item in items])

@app.route('/api/items', methods=['POST'])
def create_item():
    data = request.get_json()
    item = Item(name=data['name'], description=data.get('description'))
    db.session.add(item)
    db.session.commit()
    return jsonify(item.to_dict()), 201

@app.route('/api/items/<int:item_id>', methods=['GET'])
def get_item(item_id):
    item = Item.query.get_or_404(item_id)
    return jsonify(item.to_dict())
{{else}}
items = [
    {'id': 1, 'name': 'Sample Item', 'description': 'This is a sample item'}
]

@app.route('/api/items', methods=['GET'])
def get_items():
    return jsonify(items)

@app.route('/api/items', methods=['POST'])
def create_item():
    data = request.get_json()
    new_item = {
        'id': len(items) + 1,
        'name': data['name'],
        'description': data.get('description')
    }
    items.append(new_item)
    return jsonify(new_item), 201
{{end}}

if __name__ == '__main__':
    {{if .HasDatabase}}with app.app_context():
        db.create_all()
    {{end}}app.run(debug=True)
`

// HTMLテンプレート
var BaseTemplate = `<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{% block title %}Flask App{% endblock %}</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <a class="navbar-brand" href="{{ url_for('index') }}">Flask App</a>
        </div>
    </nav>

    <div class="container mt-4">
        {% with messages = get_flashed_messages() %}
            {% if messages %}
                {% for message in messages %}
                    <div class="alert alert-success alert-dismissible fade show" role="alert">
                        {{ message }}
                        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
                    </div>
                {% endfor %}
            {% endif %}
        {% endwith %}

        {% block content %}{% endblock %}
    </div>

    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>
`

var IndexTemplate = `{% extends "base.html" %}

{% block title %}Home - Flask App{% endblock %}

{% block content %}
<div class="row">
    <div class="col-md-8 mx-auto">
        <div class="jumbotron bg-light p-5 rounded">
            <h1 class="display-4">Hello, Flask!</h1>
            <p class="lead">This is your new Flask application.</p>
            <hr class="my-4">
            <p>Get started by editing your templates and routes.</p>
        </div>
    </div>
</div>
{% endblock %}
`

var FormTemplate = `{% extends "base.html" %}

{% block title %}Form - Flask App{% endblock %}

{% block content %}
<div class="row">
    <div class="col-md-6 mx-auto">
        <h2>Sample Form</h2>
        <form method="POST">
            {{ form.hidden_tag() }}
            <div class="mb-3">
                {{ form.name.label(class="form-label") }}
                {{ form.name(class="form-control") }}
            </div>
            <div class="mb-3">
                {{ form.submit(class="btn btn-primary") }}
            </div>
        </form>
    </div>
</div>
{% endblock %}
`

// requirements.txtの生成
func GenerateRequirements(features []string, appType string) string {
    requirements := []string{"Flask>=2.3.0"}
    
    for _, feature := range features {
        switch feature {
        case "database":
            requirements = append(requirements, "Flask-SQLAlchemy>=3.0.0")
        case "auth":
            requirements = append(requirements, "Flask-Login>=0.6.0")
        case "forms":
            requirements = append(requirements, "Flask-WTF>=1.1.0", "WTForms>=3.0.0")
        case "env":
            requirements = append(requirements, "python-dotenv>=1.0.0")
        }
    }
    
    if appType == "api" {
        requirements = append(requirements, "Flask-CORS>=4.0.0")
    }
    
    result := ""
    for _, req := range requirements {
        result += req + "\n"
    }
    return result
}

// .env ファイルのテンプレート
var EnvTemplate = `# Flask Configuration
FLASK_APP=app.py
FLASK_ENV=development
SECRET_KEY=your-secret-key-here

# Database
DATABASE_URL=sqlite:///app.db

# Other configurations
DEBUG=True
`

// README.mdテンプレート生成関数
func GenerateReadme(projectName string, appType string, hasDatabase bool, hasForms bool) string {
	var typeDescription string
	switch appType {
	case "hello":
		typeDescription = "シンプルなHello Worldアプリケーション"
	case "webapp":
		typeDescription = "HTMLテンプレートとフォームを含むWebアプリケーション"
	case "api":
		typeDescription = "JSON APIを提供するRESTfulアプリケーション"
	case "fullstack":
		typeDescription = "WebUIとAPIの両方を提供するフルスタックアプリケーション"
	default:
		typeDescription = "Flaskアプリケーション"
	}

	readme := "# " + projectName + "\n\n" + typeDescription + "\n\n"
	
	readme += `## セットアップ

### 1. 仮想環境の作成・有効化

` + "```bash" + `
python3 -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate
` + "```" + `

### 2. 依存関係のインストール

` + "```bash" + `
pip install -r requirements.txt
` + "```" + `

*注意: 新しいパッケージを追加した場合は ` + "`pip freeze > requirements.txt`" + ` で更新してください*

### 3. 環境変数の設定

` + "`.env`" + ` ファイルを編集して、必要な設定を行ってください：

` + "```bash" + `
SECRET_KEY=your-production-secret-key-here
` + "```" + `
`

	if hasDatabase {
		readme += `
### 4. データベースの初期化

` + "```bash" + `
flask shell
>>> from app import app, db
>>> with app.app_context():
...     db.create_all()
>>> exit()
` + "```" + `
`
	}

	readme += `
## 実行

### 開発サーバーの起動

` + "```bash" + `
flask run
` + "```" + `

または

` + "```bash" + `
python app.py
` + "```" + `

アプリケーションは http://localhost:5000 でアクセスできます。

## 機能
`

	if appType == "webapp" || appType == "fullstack" {
		readme += `
- Web UI
- HTMLテンプレート（Jinja2）`
		if hasForms {
			readme += `
- フォーム処理（Flask-WTF）`
		}
	}

	if appType == "api" || appType == "fullstack" {
		readme += `
- REST API
- JSON レスポンス`
	}

	if hasDatabase {
		readme += `
- データベース連携（SQLAlchemy）`
	}

	readme += "\n\n## プロジェクト構造\n\n```\n" + projectName + "/\n"
	readme += `├── app.py              # メインアプリケーション
├── requirements.txt    # Python依存関係
├── .env               # 環境変数設定
├── .gitignore         # Git除外ファイル
├── README.md          # このファイル`

	if appType != "hello" {
		readme += `
├── templates/         # HTMLテンプレート
│   ├── base.html
│   └── index.html
└── static/           # 静的ファイル
    ├── css/
    └── js/`
	}

	readme += `
` + "```" + `

## 開発

### デバッグモード

開発中は ` + "`.env`" + ` ファイルで ` + "`DEBUG=True`" + ` に設定されています。

### 本番環境

本番環境では以下の設定を変更してください：

- ` + "`SECRET_KEY`" + ` を強固なものに変更
- ` + "`DEBUG=False`" + ` に設定
- データベースURLを本番環境用に変更
`

	return readme
}

// .gitignoreテンプレート
var GitignoreTemplate = `# Byte-compiled / optimized / DLL files
__pycache__/
*.py[cod]
*$py.class

# C extensions
*.so

# Distribution / packaging
.Python
build/
develop-eggs/
dist/
downloads/
eggs/
.eggs/
lib/
lib64/
parts/
sdist/
var/
wheels/
pip-wheel-metadata/
share/python-wheels/
*.egg-info/
.installed.cfg
*.egg
MANIFEST

# PyInstaller
*.manifest
*.spec

# Installer logs
pip-log.txt
pip-delete-this-directory.txt

# Unit test / coverage reports
htmlcov/
.tox/
.nox/
.coverage
.coverage.*
.cache
nosetests.xml
coverage.xml
*.cover
*.py,cover
.hypothesis/
.pytest_cache/

# Translations
*.mo
*.pot

# Django stuff:
*.log
local_settings.py
db.sqlite3
db.sqlite3-journal

# Flask stuff:
instance/
.webassets-cache

# Scrapy stuff:
.scrapy

# Sphinx documentation
docs/_build/

# PyBuilder
target/

# Jupyter Notebook
.ipynb_checkpoints

# IPython
profile_default/
ipython_config.py

# pyenv
.python-version

# pipenv
Pipfile.lock

# PEP 582
__pypackages__/

# Celery stuff
celerybeat-schedule
celerybeat.pid

# SageMath parsed files
*.sage.py

# Environments
.env
.venv
env/
venv/
ENV/
env.bak/
venv.bak/

# Spyder project settings
.spyderproject
.spyproject

# Rope project settings
.ropeproject

# mkdocs documentation
/site

# mypy
.mypy_cache/
.dmypy.json
dmypy.json

# Pyre type checker
.pyre/

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db
`