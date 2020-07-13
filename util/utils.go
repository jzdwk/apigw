/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package util

// Empty checks if a string referenced by s or s itself is empty.
func Empty(s *string) bool {
	return s == nil || *s == ""
}
