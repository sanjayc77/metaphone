// Copyright (c) 2012 Sanjay Chouksey
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package metaphone

import "testing"

type metaphoneTest struct {
	fn func(string) string
	in, out, desc string
}

var metaphoneTests = []metaphoneTest {
	metaphoneTest{dedup, "dropping", "droping", "should drop duplicate adjacent letters, except C"},
	metaphoneTest{dedup, "accelerate", "accelerate", "should not drop duplicat C"},
	metaphoneTest{dropInitialLetters, "knuth", "nuth", "should drop some initial letters"},
	metaphoneTest{dropInitialLetters, "gnat", "nat", "should drop some initial letters"},
	metaphoneTest{dropInitialLetters, "aegis", "egis", "should drop some initial letters"},
	metaphoneTest{dropInitialLetters, "pneumatic", "neumatic", "should drop some initial letters"},
	metaphoneTest{dropInitialLetters, "wrack", "rack", "should drop some initial letters"},
	metaphoneTest{dropInitialLetters, "garbage", "garbage", "should not drop other initial letters"},
	metaphoneTest{dropBafterMAtEnd, "dumb", "dum", "should b if if words end with mb"},
	metaphoneTest{dropBafterMAtEnd, "dumbo", "dumbo", "should not drop b after m if not at end of word"},
	metaphoneTest{cTransform, "change", "xhange", "should replace CH to X"},
	metaphoneTest{cTransform, "discharger", "diskharger", "should not replace CH to X if part of SCH"},
	metaphoneTest{cTransform, "aesthetician", "aesthetixian", "should replace CIA to X"},
	metaphoneTest{cTransform, "cieling", "sieling", "C should become S if followed by I, E, or Y"},
	metaphoneTest{cTransform, "cuss", "kuss", "should transform other C's to K"},
	metaphoneTest{dTransform, "abridge", "abrijge", "should transform D to J if followed by GE, GY, GI"},
	metaphoneTest{dTransform, "bid", "bit", "should transform D to T if not followed by GE, GY, GI"},
	metaphoneTest{dropG, "alight", "aliht", "should drop G before H if not at the end or before vowell"},
	metaphoneTest{dropG, "fright", "friht", "should drop G before H if not at the end or before vowell"},
  metaphoneTest{dropG, "aligned", "alined", "should drop G if followed by N or NED at the end"},
  metaphoneTest{dropG, "align", "alin", "should drop G if followed by N or NED at the end"},
  metaphoneTest{transformG, "age", "aje", "should transform G to J if followed by I, E or Y and not preceeded by G"},
  metaphoneTest{transformG, "gin", "jin", "should transform G to J if followed by I, E or Y and not preceeded by G"},
  metaphoneTest{transformG, "august", "aukust", "should transform G to K"},
  metaphoneTest{transformG, "aggrade", "akrade", "should transform G to K"},
  metaphoneTest{dropH, "alriht", "alrit", "should drop H if after vowell and not before vowell"},
  metaphoneTest{dropH, "that", "that", "should not drop H if after vowell"},
  metaphoneTest{dropH, "chump", "chump", "should not drop H if not before vowell"},
  metaphoneTest{transformCK, "check", "chek", "should transform CK to K"},
  metaphoneTest{transformPH, "phone", "fone", "should transform PH to F"},
  metaphoneTest{transformQ, "quack", "kuack", "should transform Q to K"},
  metaphoneTest{transformS, "shack", "xhack", "should transform S to X if followed by H, IO, or IA"},
  metaphoneTest{transformS, "sialagogues", "xialagogues", "should transform S to X if followed by H, IO, or IA"},
  metaphoneTest{transformS, "asia", "axia", "should transform S to X if followed by H, IO, or IA"},
  metaphoneTest{transformS, "substance", "substance", "should not transform S to X if not followed by H, IO, or IA"},
  metaphoneTest{transformT, "dementia", "demenxia", "should transform T to X if followed by IA or IO"},
  metaphoneTest{transformT, "abbreviation", "abbreviaxion", "should transform T to X if followed by IA or IO"},
  metaphoneTest{transformT, "that", "0at", "should transform TH to 0"},
  metaphoneTest{dropT, "backstitch", "backstich", "should drop T if followed by CH"},
  metaphoneTest{transformV, "vestige", "festige", "should transform V to F"},
  metaphoneTest{dropW, "bowl", "bol", "should drop W if not followed by vowell"},
  metaphoneTest{dropW, "warsaw", "warsa", "should drop W if not followed by vowell"},
  metaphoneTest{transformX, "xenophile", "senophile", "should transform X to S if at beginning"},
  metaphoneTest{transformX, "admixed", "admiksed", "should transform X to KS if not at beginning"},
  metaphoneTest{dropY, "analyzer", "analzer", "should drop Y if not followed by a vowell"},
  metaphoneTest{dropY, "specify", "specif", "should drop Y if not followed by a vowell"},
  metaphoneTest{dropY, "allying", "allying", "should not drop Y if followed by a vowell"},
  metaphoneTest{transformZ, "blaze", "blase", "should transform Z to S"},
  metaphoneTest{dropVowels, "ablaze", "ablz", "should drop all vowels except initial"},
  metaphoneTest{dropVowels, "adamantium", "admntm", "should drop all vowels except initial"},
  metaphoneTest{Process, "ablaze", "ABLS", "should do all"},
  metaphoneTest{Process, "transition", "TRNSXN", "should do all"},
  metaphoneTest{Process, "astronomical", "ASTRNMKL", "should do all"},
  metaphoneTest{Process, "buzzard", "BSRT", "should do all"},
  metaphoneTest{Process, "wonderer", "WNTRR", "should do all"},
  metaphoneTest{Process, "district", "TSTRKT", "should do all"},
  metaphoneTest{Process, "hockey", "HK", "should do all"},
  metaphoneTest{Process, "capital", "KPTL", "should do all"},
  metaphoneTest{Process, "penguin", "PNKN", "should do all"},
  metaphoneTest{Process, "garbonzo", "KRBNS", "should do all"},
  metaphoneTest{Process, "lightning", "LTNNK", "should do all"},
  metaphoneTest{Process, "light", "LT", "should do all"},
}

func TestMetaphone(t *testing.T) {
	for _, dt := range metaphoneTests {
		v := dt.fn(dt.in)
		if v != dt.out {
			t.Errorf("%v(%s) = %s, expected %s.", dt.fn, dt.in, v, dt.out)
		}
	}
}

func TestMetaphoneWithMaxLength(t *testing.T) {
  // should truncate to length specified if code exceeds
  v := ProcessWithMaxLength("phonetics", 4)
  if v != "FNTK" {
    t.Errorf("ProcessWithMaxLength(\"phonetics\", 4) = %s, expected %s.", v, "FNTK");
  }
  
  // should not truncate to length specified if code does not exceed
  v = ProcessWithMaxLength("phonetics", 8)
  if v != "FNTKS" {
    t.Errorf("ProcessWithMaxLength(\"phonetics\", 8) = %s, expected %s.", v, "FNTKS");
  }
}
