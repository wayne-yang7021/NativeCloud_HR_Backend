#!/bin/bash

echo "📦 執行資料庫 migration..."
set -e

# 進入專案根目錄（確保執行位置正確）
cd "$(dirname "$0")/.."

# 執行 Go migration script
go run scripts/migrate

echo "✅ Migration 完成"