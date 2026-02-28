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

## 概要

`gotocafe` は、[CAFFEE_Editor](https://github.com/iamthe000/CAFFEE_Editor) の自動化マクロ言語である `.caffeine` ファイルを、外部からシームレスに操作するためのツールです。

通常、CAFFEE_Editorを起動してから手動で実行する必要があるマクロをコマンドラインから一発で呼び出せるだけでなく、**マクロ自体をスタンドアロンな実行可能バイナリファイル（ELF/Mach-O等）にコンパイルする**という強力な機能を備えています。
これにより、自作のマクロスクリプトを単一のCLIツールとして配布・運用することが可能になります。

## 主な機能

- **Direct Run (直接実行モード)**: `.caffeine` ファイルを指定するだけで、自動的にCAFFEEを立ち上げ、マクロを即座に実行します。
- **Binary Compilation (バイナリ化モード)**: マクロコードをGoのバイナリに埋め込み、独立した実行可能ファイルを生成します。
- **TUI Automation**: バックエンドで `expect` を用いることで、複雑なターミナルUIのキーボード入力を完全にエミュレートします。

## 必須環境 (Prerequisites)

本ツールをビルド・実行するには以下の環境が必要です。

- **Go Compiler** (v1.16+) : バイナリのビルド用
- **CAFFEE_Editor** : `caffee` コマンドとしてシステムのPATHが通っていること
- **expect** : TUI操作の自動化用（Linux / macOS）
  - Ubuntu/Debian: `sudo apt install expect`
  - macOS (Homebrew): `brew install expect`

## インストール (Installation)

ソースコードをダウンロードし、ビルドしてパスの通ったディレクトリに配置します。

```bash
# 1. リポジトリをクローンまたはダウンロード
git clone [https://github.com/yourusername/gotocafe.git](https://github.com/yourusername/gotocafe.git)
cd gotocafe

# 2. ビルド
go build -o gotocafe gotocafe.go

# 3. 実行可能ファイルをPATHの通った場所に移動（例: /usr/local/bin/）
sudo mv gotocafe /usr/local/bin/

```

## 使い方

`gotocafe` には大きく分けて2つのモードがあります。

### 1. マクロをそのまま実行する (Direct Run Mode)

作成した `sample.caffeine` を、コマンドラインから即座にCAFFEEエディタ上で実行します。

```bash
gotocafe sample.caffeine

```

> **動作**: `caffee` が起動し、自動的に `Ctrl+P` -> `:macro sample.caffeine` が入力されて実行されます。

### 2. マクロをバイナリ化する (Compile Mode)

マクロを他の環境でも単体で動かせるように（※実行環境にCAFFEEとexpectは必要です）、専用のコマンドツールとしてコンパイルします。`cafename=` 引数に出力ファイル名を指定してください。

```bash
# sample.caffeine を 'my_macro_tool' というバイナリに変換
gotocafe sample.caffeine cafename=my_macro_tool

# 生成されたツールを実行
./my_macro_tool

```

> **動作**: `sample.caffeine` の中身がGoコード内に文字列として埋め込まれ、`go build` が実行されます。生成されたバイナリは、実行時に一時ファイルとしてマクロを展開し、自動的にCAFFEEを立ち上げて処理を実行します。

## アーキテクチャ (How it works)

`gotocafe` は以下の仕組みで動作しています。

1. **疑似端末エミュレーション**: Goの `os/exec` から `expect` スクリプトを呼び出し、ターミナルエミュレータ上のプロセスとして `caffee` を `spawn` します。
2. **キー入力のインジェクト**: 起動直後のウェイト（約0.5秒）を挟んだ後、コマンドモードに入るための制御文字 `\x10` (Ctrl+P) を送信し、マクロ実行コマンドを流し込みます。
3. **制御の移譲**: マクロ実行開始後、`interact` コマンドによってプロセスの制御をユーザーに返し、ユーザーがそのままエディタの操作を継続できるようにしています。

## 🤝 コントリビューション

Issueの報告やPull Requestは大歓迎です。
バグ報告や新機能の提案は、GitHubのIssueトラッカーからお願いいたします。

## 📄 ライセンス

This project is licensed under the MIT License

```

---

### おすすめのカスタマイズポイント
* リポジトリ名やユーザー名のURL（`https://github.com/yourusername/...`）をご自身の環境に合わせて書き換えてください。
* 必要であれば、CAFFEINEマクロの簡単なサンプルコード（前回の `sample.caffeine` など）を `使い方` の項目に追記すると、初めてツールを見る人にとってさらに親切になります。

```
