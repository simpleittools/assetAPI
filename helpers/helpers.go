package helpers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Slugify(name string) string {
	// Convert to lowercase
	slug := strings.ToLower(name)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non-alphanumeric characters
	reg := regexp.MustCompile("[^a-zA-Z0-9-]")
	slug = reg.ReplaceAllString(slug, "")

	// Remove multiple consecutive hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")

	// Remove any leading or trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

func UintConv(input string) uint {
	str := input
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Panic(err)
	}
	uintNum := uint(num)
	return uintNum
}
