package token

type TokenType struct {
	TypeInfo string
	Value    string
	PAXU     int
}

const (
	TYNUM    = "number"
	TYSTR    = "string"
	BULLE    = "null"
	NANINFO  = "NaN"
	BOOL     = "boolean"
	FUNCTION = "function"

	TRUE  = "true"
	FALSE = "false"
)

const (
	NNNN   = ""
	ELSE   = "else"
	TYPEOF = "typeof"
	FUN    = "function"
	FOR    = "for"
	DENYU  = "="
	QUFAN  = "!"
	BUDY   = "!="
	BUDYDY = "!=="
	Str    = "\""
	Str2   = "'"

	XIAND      = "=="
	MAOHAO     = ":"
	QUYU       = "%"
	HUO        = "|"
	HUOHUO     = "||"
	YU         = "&"
	YUYU       = "&&"
	XIAOYH     = "<"
	XIAOYHYH   = "<<"
	JIADEN     = "+="
	JANDEN     = "-="
	XIAOYHYHYH = "<<<"
	XIAOYHDY   = "<="
	DAYH       = ">"
	DAYHDY     = ">="
	DAYHYH     = ">>"
	DAYHYHYU   = ">>>"
	DXIAND     = "==="
	IF         = "if"
	ZHUOK      = "("
	YOUOK      = ")"
	YOUZ       = "["
	ZUOZ       = "]"
	ADD        = "+"
	YIHUO      = "^"
	DH         = ","
	SDD        = "-"
	CHU        = "/"
	ZUOKH      = "}"
	DUOZF      = "`"
	YOUKH      = "{"
	CHEN       = "*"
	KAIFAN     = "**"
	UPADD      = "++"
	UPASD      = "--"
	DONT       = "DONT"
	END        = ""
	Kong       = ""
	Dian       = "."
)

const (
	ZNO   = 0
	ONE   = -1
	TWO   = 2
	SCRRR = 3
	FOTT  = 4
	FIVE  = 5
	SIX   = 6
)
const (
	Slice  = "slice"
	Length = "length"
)

const (
	GetLength   = "len"
	ParseInt    = "parseInt"
	Wait        = "wait"
	ParseFloat  = "parseFloat"
	Print       = "cyout"
	GetChar     = "cychar"
	CharToStr   = "cystr"
	Input       = "input"
	AppendArray = "cyappend"
	Delete      = "delete"
)

const (
	Math        = "Math"
	Math_random = "random"
	Math_Pow    = "pow"
	Math_Sqrt   = "sqrt"
)

const (
	Objecte                = "Object"
	Objecte_setPrototypeOf = "setPrototypeOf"
	Objecte_keys           = "keys"
)

const (
	Promise      = "Promise"
	Promise_then = "then"
	Cbb_a        = "cbb_a"
	Cbb_b        = "cbb_b"
)

const (
	JSON           = "JSON"
	JSON_stringify = "stringify"
	JSON_parse     = "parse"
)

const (
	Date       = "Date"
	Date_now   = "now"
	Date_sleep = "sleep"
)

const (
	CONSOLE     = "console"
	CONSOLE_log = "log"
)

const (
	String              = "String"
	String_fromCharCode = "fromCharCode"
	String_strip        = "strip"
	String_replace      = "replace"
	String_split        = "split"
	String_decode       = "decode"
	String_encode       = "encode"
	String_newbyte      = "newbyte"
)
const (
	Etree         = "etree"
	Etree_HTML    = "HTML"
	Etree_xpath   = "xpath"
	Etree_gethtml = "gethtml"
)

const (
	Fs          = "fs"
	File        = "file"
	Fs_open     = "open"
	Fs_read     = "read"
	Fs_readCont = "readcont"
	Fs_close    = "close"
	Fs_write    = "write"
	Fs_cmd      = "cmd"
	Fs_ms       = "ms"
	Fs_encoding = "encoding"
)

const (
	Re         = "re"
	Re_findall = "findall"
	Re_sub     = "sub"
)

const (
	PUSH = "push"
	POP  = "pop"
	JION = "join"
)

const (
	Cyhttp           = "cyhttp"
	Cyhttp_get       = "get"
	Cyhttp_post      = "post"
	Cyhttp_ReHeaders = "headers"
	Headers          = "headers"
	Timeout          = "timeout"
	Params           = "params"
	Allow_redirects  = "allow_redirects"
	Proxies          = "proxies"
	Content          = "content"

	Status       = "status"
	Iserror      = "iserror"
	Text         = "text"
	Data         = "data"
	Json         = "json"
	Jsontext     = "jsontext"
	Headerstext  = "jsontext"
	Headerstext2 = "jsontext2"
)

const (
	Prog      = "Prog"
	IDENT     = "IDENT"
	NULL      = "null"
	BYTE      = "byte"
	INT       = "INT"
	VAR       = "var"
	CALL      = "call"
	APPLY     = "apply"
	Bin       = "Bin"
	OVER      = ";"
	HUANH     = "\n"
	TB        = "\t"
	TR        = "\r"
	Ass       = "Ass"
	Call      = "Call"
	NOP       = " "
	IfStat    = "IfStat"
	Block     = "Block"
	Arguments = "arguments"
	Unary     = "Unary"
	FuncD     = "FuncD"
	FuncE     = "FuncE"
	Member    = "Member"
	Stri      = "Stri"
	THIS      = "this"
	ENV       = "env"
	IN        = "in"
	CONTINUE  = "continue"
	ForS      = "ForS"
	ForI      = "ForI"
	ArrayE    = "ArrayE"
	Object    = "Object"
	Prop      = "Prop"
	BREAK     = "break"
	TRY       = "try"
	CATCH     = "catch"
	RETURN    = "return"
	Debug     = "debugger"
	NEW       = "new"
	Eval      = "eval"
	Require   = "require"
)
