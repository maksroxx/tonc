package build

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/maksroxx/tonc/util"
	"github.com/urfave/cli/v2"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func BuildAction(c *cli.Context) error {
	verbose := c.Bool("verbose")
	logf := func(format string, args ...interface{}) {
		if verbose {
			log.Printf(format, args...)
		}
	}

	contractPath := c.String("contract")
	buildDir := c.String("out")
	saveBOC := c.Bool("boc")
	saveJSON := c.Bool("json")
	saveHEX := c.Bool("hex")

	if !util.DirExists(buildDir) {
		logf("Creating build directory: %s", buildDir)
		if err := os.MkdirAll(buildDir, 0755); err != nil {
			return fmt.Errorf("failed to create build directory: %w", err)
		}
	}

	if contractPath != "" {
		if !util.FileExists(contractPath) {
			return fmt.Errorf("contract file does not exist: %s", contractPath)
		}
		fmt.Printf("🛠️  Compiling single contract: %s\n", contractPath)
		if err := compileContract(contractPath, buildDir, saveBOC, saveJSON, saveHEX, logf); err != nil {
			return err
		}
		return nil
	}

	srcDir := c.String("src")
	if !util.DirExists(srcDir) {
		return fmt.Errorf("source directory does not exist: %s", srcDir)
	}

	pattern := filepath.Join(srcDir, "*.fc")
	fcFiles, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("failed to scan source directory: %w", err)
	}
	if len(fcFiles) == 0 {
		return fmt.Errorf("no .fc files found in %s", srcDir)
	}

	fmt.Printf("🛠️  Compiling %d contracts from %s\n", len(fcFiles), srcDir)
	for i, fc := range fcFiles {
		fmt.Printf(" [%d/%d] %s\n", i+1, len(fcFiles), filepath.Base(fc))
		if err := compileContract(fc, buildDir, saveBOC, saveJSON, saveHEX, logf); err != nil {
			return fmt.Errorf("failed to compile %s: %w", fc, err)
		}
	}

	fmt.Println("🎉 All contracts compiled successfully.")
	return nil
}

func compileContract(contractPath, buildDir string, saveBOC, saveJSON, saveHEX bool, logf func(string, ...interface{})) error {
	contractName := strings.TrimSuffix(filepath.Base(contractPath), filepath.Ext(contractPath))
	logf("Start compiling contract: %s", contractName)

	outputFift := filepath.Join(buildDir, contractName+".fif")
	outputBOC := filepath.Join(buildDir, contractName+".cell.boc")
	outputJSON := filepath.Join(buildDir, contractName+".compiled.json")

	cmdFunc := exec.Command("func", "-o", outputFift, "-W"+outputBOC, "-AP", contractPath)
	cmdFunc.Stdout = os.Stdout
	cmdFunc.Stderr = os.Stderr
	if err := cmdFunc.Run(); err != nil {
		return fmt.Errorf("func compilation failed: %w", err)
	}
	logf("func compilation done")

	cmdFift := exec.Command("fift", outputFift)
	cmdFift.Stdout = os.Stdout
	cmdFift.Stderr = os.Stderr
	if err := cmdFift.Run(); err != nil {
		return fmt.Errorf("fift execution failed: %w", err)
	}
	logf("fift execution done")

	bocData, err := os.ReadFile(outputBOC)
	if err != nil {
		return fmt.Errorf("failed to read BOC file: %w", err)
	}

	cells, err := cell.FromBOC(bocData)
	if err != nil {
		return fmt.Errorf("failed to parse BOC: %w", err)
	}

	if !saveBOC {
		if err := os.Remove(outputBOC); err != nil && !errors.Is(err, os.ErrNotExist) {
			logf("Warning: failed to remove BOC file: %v", err)
		}
	} else {
		fmt.Printf(" - BOC saved: %s\n", outputBOC)
	}

	if saveJSON || saveHEX {
		data := map[string]string{}
		if saveHEX {
			data["hex"] = hex.EncodeToString(cells.ToBOC())
		}
		if saveJSON {
			jsonBytes, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			if err := os.WriteFile(outputJSON, jsonBytes, 0644); err != nil {
				return fmt.Errorf("failed to write JSON: %w", err)
			}
			fmt.Printf(" - JSON saved: %s\n", outputJSON)
		}
	}

	fmt.Printf("✅ Contract '%s' compiled successfully\n\n", contractName)
	return nil
}
