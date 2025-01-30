package strutil

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type StringCaseKind string

const (
	CamelCase          StringCaseKind = "camel_case"
	PascalCase         StringCaseKind = "pascal_case"
	SnakeCase          StringCaseKind = "snake_case"
	KebabCase          StringCaseKind = "kebab_case"
	ScreamingSnakeCase StringCaseKind = "screaming_snake_case"
)

type StringCase struct {
	Kind  StringCaseKind
	Parts []string
}

func (sc StringCase) String() string {
	if len(sc.Parts) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(sc.Parts) * 8)

	switch sc.Kind {
	case CamelCase:
		builder.WriteString(sc.Parts[0])
		for _, part := range sc.Parts[1:] {
			builder.WriteString(capitalize(part))
		}
	case PascalCase:
		for _, part := range sc.Parts {
			builder.WriteString(capitalize(part))
		}
	case SnakeCase:
		builder.WriteString(sc.Parts[0])
		for _, part := range sc.Parts[1:] {
			builder.WriteByte('_')
			builder.WriteString(part)
		}
	case KebabCase:
		builder.WriteString(sc.Parts[0])
		for _, part := range sc.Parts[1:] {
			builder.WriteByte('-')
			builder.WriteString(part)
		}
	case ScreamingSnakeCase:
		builder.WriteString(strings.ToUpper(sc.Parts[0]))
		for _, part := range sc.Parts[1:] {
			builder.WriteByte('_')
			builder.WriteString(strings.ToUpper(part))
		}
	}

	return builder.String()
}

func camelCase(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for _, part := range parts[1:] {
		result += capitalize(part)
	}
	return result
}

func pascalCase(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	result := capitalize(parts[0])
	for _, part := range parts[1:] {
		result += capitalize(part)
	}
	return result
}

func snakeCase(parts []string) string {
	return strings.Join(parts, "_")
}

func kebabCase(parts []string) string {
	return strings.Join(parts, "-")
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func screamingSnakeCase(parts []string) string {
	return strings.ToUpper(snakeCase(parts))
}

func screamingKebabCase(parts []string) string {
	return strings.ToUpper(kebabCase(parts))
}

func DetectCase(s string) StringCaseKind {
	if strings.Contains(s, "_") {
		return SnakeCase
	} else if strings.Contains(s, "-") {
		return KebabCase
	} else if len(s) > 0 && s[0] >= 'A' && s[0] <= 'Z' {
		return PascalCase
	} else {
		return CamelCase
	}
}

func ParseCase(s string) StringCase {
	if s == "" {
		return StringCase{
			Kind:  CamelCase,
			Parts: nil,
		}
	}

	kind := DetectCase(s)
	var parts []string

	switch kind {
	case SnakeCase, KebabCase:
		primarySep := byte('_')
		secondarySep := byte('-')
		if kind == KebabCase {
			primarySep, secondarySep = secondarySep, primarySep
		}

		tempParts := strings.Split(s, string(primarySep))

		for _, part := range tempParts {
			if strings.Contains(part, string(secondarySep)) {
				subParts := strings.Split(part, string(secondarySep))
				for _, subPart := range subParts {
					if subPart != "" {
						parts = append(parts, strings.ToLower(subPart))
					}
				}
			} else if part != "" {
				parts = append(parts, strings.ToLower(part))
			}
		}
	default:
		parts = splitCamelCase(s)
		for i := range parts {
			parts[i] = strings.ToLower(parts[i])
		}
	}

	if len(parts) == 0 {
		return StringCase{
			Kind:  kind,
			Parts: nil,
		}
	}

	return StringCase{
		Kind:  kind,
		Parts: parts,
	}
}

func (sc StringCase) ToCamelCase() string {
	return camelCase(sc.Parts)
}

func (sc StringCase) ToPascalCase() string {
	return pascalCase(sc.Parts)
}

func (sc StringCase) ToSnakeCase() string {
	return snakeCase(sc.Parts)
}

func (sc StringCase) ToKebabCase() string {
	return kebabCase(sc.Parts)
}

func (sc StringCase) ToScreamingSnakeCase() string {
	return screamingSnakeCase(sc.Parts)
}

func splitCamelCase(s string) []string {
	if s == "" {
		return nil
	}

	var parts []string
	var current []rune
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if isUpper(runes[i]) && i > 0 && !isUpper(runes[i-1]) {
			if len(current) > 0 {
				parts = append(parts, string(current))
				current = nil
			}
			current = append(current, runes[i])
		} else if isUpper(runes[i]) && i+1 < len(runes) && !isUpper(runes[i+1]) && len(current) > 1 {
			parts = append(parts, string(current))
			current = []rune{runes[i]}
		} else {
			current = append(current, runes[i])
		}
	}

	if len(current) > 0 {
		parts = append(parts, string(current))
	}

	for i := range parts {
		if i == 0 && !isUpper(runes[0]) {
			parts[i] = strings.ToLower(parts[i])
		} else if isAllUpper(parts[i]) {
			continue
		} else if i > 0 || isUpper(runes[0]) {
			parts[i] = cases.Title(language.English).String(strings.ToLower(parts[i]))
		}
	}

	return parts
}

func isAllUpper(s string) bool {
	for _, r := range s {
		if !isUpper(r) {
			return false
		}
	}
	return len(s) > 0
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}
