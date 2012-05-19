// Copyright (c) 2012 Sanjay Chouksey
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package metaphone

import (
	"fmt"
	"regexp"
	"bytes"
	"strings"
)

/*
func dedup(token) {
  return token.replace(/([^c])\1/g, '${1}');
}
*/

func dedup(token string) string {
	if len(token) < 2 {
		return token;
	}
	
	buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, string(token[0]))
	for i := 1; i < len(token); i++ {
		if token[i] == "c"[0] || token[i] != token[i - 1] {
			fmt.Fprint(buffer, string(token[i]))
		}
	}
	return string(buffer.Bytes())
}

func dropInitialLetters(token string) string {
  if regexp.MustCompile("^(kn|gn|pn|ae|wr)").MatchString(token) {
  return token[1:]
  }
  return token;
}

func dropBafterMAtEnd(token string) string {
	return regexp.MustCompile("mb$").ReplaceAllLiteralString(token, "m")
}

func cTransform(token string) string {
  token = strings.TrimSpace(regexp.MustCompile("([^s]|^)(c)(h)").ReplaceAllString(token, "${1}x${3}"))
  token = regexp.MustCompile("cia").ReplaceAllLiteralString(token, "xia")
  token = regexp.MustCompile("c(i|e|y)").ReplaceAllString(token, "s${1}")
  token = regexp.MustCompile("c").ReplaceAllLiteralString(token, "k")
  
  return token;
}

func dTransform(token string) string {
  token = regexp.MustCompile("d(ge|gy|gi)").ReplaceAllString(token, "j${1}")
  token = regexp.MustCompile("d").ReplaceAllLiteralString(token, "t")
  
  return token;
}

func dropG(token string) string {
  token = regexp.MustCompile("gh(^$|[^aeiou])").ReplaceAllString(token, "h${1}")
  token = regexp.MustCompile("g(n|ned)$").ReplaceAllString(token, "${1}")  
  
  return token;
}

func transformG(token string) string {
  token = regexp.MustCompile("([^g]|^)(g)(i|e|y)").ReplaceAllString(token, "${1}j${3}")
  token = regexp.MustCompile("gg").ReplaceAllLiteralString(token, "g")
  token = regexp.MustCompile("g").ReplaceAllLiteralString(token, "k")
  
  return token;
}

func dropH(token string) string {
  return regexp.MustCompile("([aeiou])h([^aeiou])").ReplaceAllString(token, "${1}${2}")
}

func transformCK(token string) string {
  return regexp.MustCompile("ck").ReplaceAllString(token, "k")
}
func transformPH(token string) string {
  return regexp.MustCompile("ph").ReplaceAllString(token, "f")
}

func transformQ(token string) string {
  return regexp.MustCompile("q").ReplaceAllString(token, "k")
}

func transformS(token string) string {
  return regexp.MustCompile("s(h|io|ia)").ReplaceAllString(token, "x${1}")
}

func transformT(token string) string {
  token = regexp.MustCompile("t(ia|io)").ReplaceAllString(token, "x${1}")
  token = regexp.MustCompile("th").ReplaceAllLiteralString(token, "0")
  
  return token;
}

func dropT(token string) string {
  return regexp.MustCompile("tch").ReplaceAllString(token, "ch")
}

func transformV(token string) string {
  return regexp.MustCompile("v").ReplaceAllString(token, "f")
}

func transformWH(token string) string {
  return regexp.MustCompile("^wh").ReplaceAllLiteralString(token, "w")
}

func dropW(token string) string {
  return regexp.MustCompile("w([^aeiou]|$)").ReplaceAllString(token, "${1}")
}

func transformX(token string) string {
  token = regexp.MustCompile("^x").ReplaceAllLiteralString(token, "s")
  token = regexp.MustCompile("x").ReplaceAllString(token, "ks")
  return token;
}

func dropY(token string) string {
  return regexp.MustCompile("y([^aeiou]|$)").ReplaceAllString(token, "${1}")
}

func transformZ(token string) string {
  return regexp.MustCompile("z").ReplaceAllLiteralString(token, "s")
}

func dropVowels(token string) string {
  buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, string(token[0]))
	fmt.Fprint(buffer, regexp.MustCompile("[aeiou]").ReplaceAllLiteralString(token[1:], ""))
  return string(buffer.Bytes())
}

func ProcessWithMaxLength(token string, maxLength int) string {
  token = strings.ToLower(token)
  token = dedup(token)
  token = dropInitialLetters(token)
  token = dropBafterMAtEnd(token)
  token = transformCK(token)
  token = cTransform(token)
  token = dTransform(token)
  token = dropG(token)
  token = transformG(token)
  token = dropH(token)
  token = transformPH(token)
  token = transformQ(token)
  token = transformS(token)
  token = transformX(token)
  token = transformT(token)
  token = dropT(token)
  token = transformV(token)
  token = transformWH(token)
  token = dropW(token)
  token = dropY(token)
  token = transformZ(token)
  token = dropVowels(token)
  
  token = strings.ToUpper(token)
  if len(token) >= maxLength {
      token = token[0:maxLength]
  }
  return token
}

func Process(token string) string {
  return ProcessWithMaxLength(token, 32)
}