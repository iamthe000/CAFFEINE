package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotocafe <file.caffeine> [cafename=<output_bin>]")
		os.Exit(1)
	}

	macroFile := os.Args[1]
	outputBin := ""

	for _, arg := range os.Args[2:] {
		if strings.HasPrefix(arg, "cafename=") {
			outputBin = strings.TrimPrefix(arg, "cafename=")
		}
	}

	if outputBin == "" {

		err := runMacro(macroFile)
		if err != nil {
			fmt.Printf("実行エラー: %v\n", err)
			os.Exit(1)
		}
	} else {

		err := compileMacro(macroFile, outputBin)
		if err != nil {
			fmt.Printf("コンパイルエラー: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("🎉 成功: %s をバイナリ '%s' としてコンパイルしました！\n", macroFile, outputBin)
	}
}

func runMacro(macroPath string) error {
	absPath, err := filepath.Abs(macroPath)
	if err != nil {
		return err
	}

	expectScript := fmt.Sprintf(`
set timeout -1
spawn caffee
sleep 0.5
send "\x10"
sleep 0.1
send "macro %s\r"
interact
`, absPath)

	cmd := exec.Command("expect", "-c", expectScript)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func compileMacro(macroPath, outputBin string) error {
	content, err := os.ReadFile(macroPath)
	if err != nil {
		return err
	}

	goCode := fmt.Sprintf(`package main

import (
	"fmt"
	"os"
	"os/exec"
)

const macroContent = %c%s%c

func main() {
	
	tmpFile, err := os.CreateTemp("", "*.caffeine")
	if err != nil {
		fmt.Println("エラー: 一時ファイルを作成できませんでした:", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(macroContent)); err != nil {
		fmt.Println("エラー: マクロを展開できませんでした:", err)
		os.Exit(1)
	}
	tmpFile.Close()

	
	expectScript := fmt.Sprintf(`+"`"+`
set timeout -1
spawn caffee
sleep 0.5
send "\x10"
sleep 0.1
send "macro %%s\r"
interact
`+"`"+`, tmpFile.Name())

	cmd := exec.Command("expect", "-c", expectScript)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("実行エラー:", err)
		os.Exit(1)
	}
}
`, '`', string(content), '`')

	tmpDir, err := os.MkdirTemp("", "gotocafe_build")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	goFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(goFile, []byte(goCode), 0644); err != nil {
		return err
	}

	cmd := exec.Command("go", "build", "-o", outputBin, goFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
