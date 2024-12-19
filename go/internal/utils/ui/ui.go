package ui

const (
	IdPaginationTop    = "pagination-top"
	IdPaginationBottom = "pagination-bottom"
	IdAnchorHome       = "a-home"
	IdAnchorAbout      = "a-about"
	IdAnchorRegister   = "a-register"
	IdAnchorLogin      = "a-login"
)

const (
	TextAnchorHome     = "Home"
	TextAnchorAbout    = "About"
	TextAnchorLogin    = "Login"
	TextAnchorRegister = "Register"
)

const (
	TitleHome     = "Library"
	TitleAbout    = "Library - About"
	TitleLogin    = "Library - Login"
	TitleRegister = "Library - Register"
	TitleNotFound = "Not Found"
)

func TitleDefault(text string) string {
	return "Library - " + text
}

type Anchor struct {
	Id   string
	Text string
}

var PathToAnchor = map[string]Anchor{
	"/":         {Id: IdAnchorHome, Text: TextAnchorHome},
	"/about":    {Id: IdAnchorAbout, Text: TextAnchorAbout},
	"/login":    {Id: IdAnchorLogin, Text: TextAnchorLogin},
	"/register": {Id: IdAnchorRegister, Text: TextAnchorRegister},
}
