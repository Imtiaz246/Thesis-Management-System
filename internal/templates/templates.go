package templates

import _ "embed"

var (
	//go:embed email_verify.tmpl
	EmailVerifyTmpl string
)
