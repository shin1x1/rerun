package version

import "fmt"

var (
	version = "dev"
	commit  = "dev"
)

func Text() string {
	return fmt.Sprintf("rerun version %s (%s)", version, commit)
}
