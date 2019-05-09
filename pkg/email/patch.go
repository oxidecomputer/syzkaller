	"regexp"
	diffStarted := false
		if lineMatchesDiffStart(ln) {
			diffStarted = true
			diff += ln + "\n"
		if diffStarted {
			if ln == "" || ln == "--" || ln == "-- " || ln[0] == '>' {
				diffStarted = false
				continue
			}
			if strings.HasPrefix(ln, " ") || strings.HasPrefix(ln, "+") ||
				strings.HasPrefix(ln, "-") || strings.HasPrefix(ln, "@") ||
				strings.HasPrefix(ln, "================") {
				diff += ln + "\n"
				continue
			}

func lineMatchesDiffStart(ln string) bool {
	diffRegexps := []*regexp.Regexp{
		regexp.MustCompile(`^(---|\+\+\+) [^\s]`),
		regexp.MustCompile(`^diff --git`),
		regexp.MustCompile(`^index [0-9a-f]+\.\.[0-9a-f]+`),
		regexp.MustCompile(`^new file mode [0-9]+`),
		regexp.MustCompile(`^Index: [^\s]`),
	}
	for _, re := range diffRegexps {
		if re.MatchString(ln) {
			return true
		}
	}
	return false
}