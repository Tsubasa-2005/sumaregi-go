package sumaregi

import "strings"

// JoinScopes joins slices of scope strings into a single string separated by spaces.
func joinScopes(scopes []string) string {
	return strings.Join(scopes, " ")
}
