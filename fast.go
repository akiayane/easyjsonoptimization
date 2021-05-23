package assignment2

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"io"
	"os"
	"strings"
)

//----------------------------------------------------------------------------------------------------------------------
// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]bool, 150)
	foundUsers := new(bytes.Buffer)

	users := make([]User, 0)

	scanner := bufio.NewScanner(file)

	user_buffer := User{}

	for scanner.Scan() {
		user_buffer.UnmarshalJSON(scanner.Bytes())
		users = append(users, user_buffer)
		user_buffer = User{}
	}

	for i, user := range users {
		isAnd := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android"){
				isAnd = true

				seenBrowsers[browser] = true

			} else if strings.Contains(browser, "MSIE"){
				isMSIE = true

				seenBrowsers[browser] = true

			}

		}

		if isAnd && isMSIE{
			email := strings.Replace(user.Email, "@", " [at] ", -1)
			foundUsers.WriteString(fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email))
		}
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers.String())
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
//----------------------------------------------------------------------------------------------------------------------
// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

type User struct {
	Browsers []string `json:"browsers"`
	Email string `json:"email"`
	Name string `json:"name"`
}

func easyjson9e1087fdDecodeAssignment2Ezz(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeAssignment2Ezz(&r, v)
	return r.Error()
}

