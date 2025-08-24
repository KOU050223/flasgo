.PHONY: build install uninstall test clean

# バイナリ名とインストール先
BINARY_NAME=flasgo
INSTALL_PATH=/usr/local/bin

# ビルド
build:
	go build -o $(BINARY_NAME) ./cmd

# ローカルにインストール（sudo権限必要）
install: build
	sudo cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(BINARY_NAME) が $(INSTALL_PATH) にインストールされました"

# アンインストール
uninstall:
	sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "$(BINARY_NAME) がアンインストールされました"

# テスト
test:
	go test ./...

# クリーンアップ
clean:
	rm -f $(BINARY_NAME)

# 開発用（ローカルビルドして実行）
dev: build
	./$(BINARY_NAME)