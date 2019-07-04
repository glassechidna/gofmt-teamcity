package main

import (
	"bytes"
	"fmt"
	"github.com/glassechidna/teamcitymsg"
	"github.com/sourcegraph/go-diff/diff"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	data, err := gofmtData()
	if err != nil {
		panic(err)
	}
	process(bytes.NewReader(data), os.Stdout)
}

func process(reader io.Reader, writer io.Writer) {
	diffReader := diff.NewMultiFileDiffReader(reader)
	parsed, err := diffReader.ReadAllFiles()
	if err != nil {
		panic(err)
	}

	replacements := diffsToReplacements(parsed)

	typeId := "gofmt"
	typeMsg := teamcitymsg.NewMsgInspectionType(typeId, "gofmt", "`gofmt` style violation", "Code style")
	fmt.Fprintln(writer, typeMsg.String())

	for _, re := range replacements {
		msg := teamcitymsg.NewMsgInspection(typeId, re.Path, re.Hunk, re.Line)
		fmt.Fprintln(writer, msg.String())
	}
}

type replacement struct {
	Path string
	Line int
	Hunk string
}

func diffsToReplacements(diffs []*diff.FileDiff) []replacement {
	var replacements []replacement

	for _, fileDiff := range diffs {
		for _, hunk := range fileDiff.Hunks {
			body := string(hunk.Body)
			lines := strings.Split(body, "\n")
			for idx, line := range lines {
				if strings.HasPrefix(line, "-") {
					replacements = append(replacements, replacement{
						Path: fileDiff.NewName,
						Line: int(hunk.OrigStartLine) + idx,
						Hunk: body,
					})
				}
			}
		}
	}

	return replacements
}

func gofmtData() ([]byte, error) {
	if stdinHasData() {
		return ioutil.ReadAll(os.Stdin)
	} else {
		return subcmdData()
	}
}

func subcmdData() ([]byte, error) {
	cmd := exec.Command("gofmt", "-d", ".")
	return cmd.CombinedOutput()
}

func stdinHasData() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
