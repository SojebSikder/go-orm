package parser

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func StripComments(src string) string {
	var out []string
	sc := bufio.NewScanner(strings.NewReader(src))
	for sc.Scan() {
		line := sc.Text()
		if idx := strings.Index(line, "//"); idx >= 0 {
			line = line[:idx]
		}
		out = append(out, line)
	}
	joined := strings.Join(out, "\n")
	for {
		start := strings.Index(joined, "/*")
		if start == -1 {
			break
		}
		end := strings.Index(joined[start+2:], "*/")
		if end == -1 {
			joined = joined[:start]
			break
		}
		joined = joined[:start] + joined[start+2+end+2:]
	}
	return joined
}

func ReadAllFromFileOrStdin(args []string) (string, error) {
	if len(args) >= 2 {
		path := args[1]
		b, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	var sb strings.Builder
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		println("Paste schema then Ctrl+D (Unix) or Ctrl+Z (Windows):")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		sb.WriteString(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return sb.String(), nil
}
