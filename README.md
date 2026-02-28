<div align="center">

# ☕ gotocafe

**CAFFEINE Macro Runner & Compiler for CAFFEE_Editor**

[![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Dependencies](https://img.shields.io/badge/Dependencies-CAFFEE__Editor%20%7C%20expect-orange)](#prerequisites)

ターミナルテキストエディタ「CAFFEE」の強力なマクロ言語 **CAFFEINE** を、<br>
より自由に、よりプログラマブルに扱うためのGo製CLIツール。

</div>

---

## 📖 概要 (About)

`gotocafe` は、[CAFFEE_Editor](https://github.com/iamthe000/CAFFEE_Editor) の自動化マクロ言語である `.caffeine` ファイルを、外部からシームレスに操作するためのツールです。

通常、CAFFEE_Editorを起動してから手動で実行する必要があるマクロをコマンドラインから一発で呼び出せるだけでなく、**マクロ自体をスタンドアロンな実行可能バイナリファイル（ELF/Mach-O等）にコンパイルする**という強力な機能を備えています。
これにより、自作のマクロスクリプトを単一のCLIツールとして配布・運用することが可能になります。

## 主な機能

- **Direct Run (直接実行モード)**: `.caffeine` ファイルを指定するだけで、自動的にCAFFEEを立ち上げ、マクロを即座に実行します。
- **Binary Compilation (バイナリ化モード)**: マクロコードをGoのバイナリに埋め込み、独立した実行可能ファイルを生成します。
- **TUI Automation**: バックエンドで `expect` を用いることで、複雑なターミナルUIのキーボード入力を完全にエミュレートします。

## 必須環境

本ツールをビルド・実行するには以下の環境が必要です。

- **Go Compiler** (v1.16+) : バイナリのビルド用
- **CAFFEE_Editor** : `caffee` コマンドとしてシステムのPATHが通っていること
- **expect** : TUI操作の自動化用（Linux / macOS）
  - Ubuntu/Debian: `sudo apt install expect`
  - macOS (Homebrew): `brew install expect`

## インストール

ソースコードをダウンロードし、ビルドしてパスの通ったディレクトリに配置します。

```bash
# 1. リポジトリをクローンまたはダウンロード
git clone [https://github.com/yourusername/gotocafe.git](https://github.com/yourusername/gotocafe.git)
cd gotocafe

# 2. ビルド
go build -o gotocafe gotocafe.go

# 3. 実行可能ファイルをPATHの通った場所に移動（例: /usr/local/bin/）
sudo mv gotocafe /usr/local/bin/
