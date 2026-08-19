package fieldarchive

func IsTerminal(s string) bool { return s == "rendered" || s == "rejected" }
func CanRender(s string) bool  { return s == "" || s == "queued" }
