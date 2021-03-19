import "strings"

func Rot90Clock(s string) string {
	subs := strings.Fields(s)
	l := len(subs[0])
	newsubs := make([]string, l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			newsubs[i] += string(subs[l-1-j][i])
		}
	}

	return strings.Join(newsubs, "\n")
}
func Diag1Sym(s string) string {
	subs := strings.Fields(s)
	l := len(subs[0])
	newsubs := make([]string, l)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			newsubs[i] += string(subs[j][i])
		}
	}

	return strings.Join(newsubs, "\n")
}
func SelfieAndDiag1(s string) string {
	subs := strings.Fields(s)
	subs2 := strings.Fields(Diag1Sym(s))
	l := len(subs[0])
	newsubs := make([]string, l)
	for i := 0; i < l; i++ {
		newsubs[i] = subs[i] + "|" + subs2[i]
	}
	return strings.Join(newsubs, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}