package ix_ci

import (
	"strings"
)

func SplitVersionStr(x string) []string {
	return strings.Split(x, ".")
}
/*
func main() {
	s := "2.10.1"

	v := SplitVersionStr(s)

	fmt.Println(strings.Join(v, " "))
}
*/
