package main

import (
	"regexp"
)

type Utility struct {

}

func (utility Utility) _removeExtraSpaces(mediaInfo string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
    re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
    final := re_leadclose_whtsp.ReplaceAllString(mediaInfo, "")
    final = re_inside_whtsp.ReplaceAllString(final, " ")
    return final
}