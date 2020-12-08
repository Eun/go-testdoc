// Code generated by 'yaegi extract path/filepath'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"go/constant"
	"go/token"
	"path/filepath"
	"reflect"
)

func init() {
	Symbols["path/filepath"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Abs":           reflect.ValueOf(filepath.Abs),
		"Base":          reflect.ValueOf(filepath.Base),
		"Clean":         reflect.ValueOf(filepath.Clean),
		"Dir":           reflect.ValueOf(filepath.Dir),
		"ErrBadPattern": reflect.ValueOf(&filepath.ErrBadPattern).Elem(),
		"EvalSymlinks":  reflect.ValueOf(filepath.EvalSymlinks),
		"Ext":           reflect.ValueOf(filepath.Ext),
		"FromSlash":     reflect.ValueOf(filepath.FromSlash),
		"Glob":          reflect.ValueOf(filepath.Glob),
		"HasPrefix":     reflect.ValueOf(filepath.HasPrefix),
		"IsAbs":         reflect.ValueOf(filepath.IsAbs),
		"Join":          reflect.ValueOf(filepath.Join),
		"ListSeparator": reflect.ValueOf(constant.MakeFromLiteral("58", token.INT, 0)),
		"Match":         reflect.ValueOf(filepath.Match),
		"Rel":           reflect.ValueOf(filepath.Rel),
		"Separator":     reflect.ValueOf(constant.MakeFromLiteral("47", token.INT, 0)),
		"SkipDir":       reflect.ValueOf(&filepath.SkipDir).Elem(),
		"Split":         reflect.ValueOf(filepath.Split),
		"SplitList":     reflect.ValueOf(filepath.SplitList),
		"ToSlash":       reflect.ValueOf(filepath.ToSlash),
		"VolumeName":    reflect.ValueOf(filepath.VolumeName),
		"Walk":          reflect.ValueOf(filepath.Walk),

		// type definitions
		"WalkFunc": reflect.ValueOf((*filepath.WalkFunc)(nil)),
	}
}
