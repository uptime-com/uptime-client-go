package upctl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gobeam/stringy"
)

type FlagSet interface {
	StringVarP(ptr *string, name, shorthand, value, usage string)
	IntVarP(ptr *int, name, shorthand string, value int, usage string)
	Float64VarP(ptr *float64, name, shorthand string, value float64, usage string)
	BoolVarP(ptr *bool, name, shorthand string, value bool, usage string)
	StringSliceVarP(ptr *[]string, name, shorthand string, value []string, usage string)
}

func Bind(fs FlagSet, obj any) error {
	v := reflect.ValueOf(obj)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		if _, skip := tf.Tag.Lookup("skip"); skip {
			continue
		}
		vf := v.FieldByName(tf.Name)
		if tf.Type.Kind() == reflect.Ptr && vf.IsNil() {
			continue
		}
		flag, short, usage := tf.Tag.Get("flag"), tf.Tag.Get("short"), tf.Tag.Get("usage")
		if flag == "" {
			if strings.HasSuffix(t.PkgPath(), "/pkg/upapi") {
				if tf.Name == "PK" {
					continue
				}
				if tf.Name == "URL" {
					continue
				}
			}
			flag = stringy.New(tf.Name).KebabCase().ToLower()
		}

		switch {
		default:
			return fmt.Errorf("unsupported field kind: %v", tf.Type.Kind())
		case t.Field(i).Type.Kind() == reflect.String:
			ptr, val := ptrVal[string](vf)
			fs.StringVarP(ptr, flag, short, val, usage)
		case t.Field(i).Type.Kind() == reflect.Int:
			ptr, val := ptrVal[int](vf)
			fs.IntVarP(ptr, flag, short, val, usage)
		case t.Field(i).Type.Kind() == reflect.Float64:
			ptr, val := ptrVal[float64](vf)
			fs.Float64VarP(ptr, flag, short, val, usage)
		case t.Field(i).Type.Kind() == reflect.Bool:
			ptr, val := ptrVal[bool](vf)
			fs.BoolVarP(ptr, flag, short, val, usage)
		case t.Field(i).Type.Kind() == reflect.Slice && t.Field(i).Type.Elem().Kind() == reflect.String:
			ptr, val := ptrVal[[]string](vf)
			fs.StringSliceVarP(ptr, flag, short, val, usage)
		// recurse into structs and pointers
		case t.Field(i).Type.Kind() == reflect.Ptr:
			x := vf.Interface()
			if err := Bind(fs, x); err != nil {
				return err
			}
		case t.Field(i).Type.Kind() == reflect.Struct:
			x := vf.Addr().Interface()
			if err := Bind(fs, x); err != nil {
				return err
			}
		}
	}
	return nil
}

func ptrVal[T any](v reflect.Value) (ptr *T, val T) {
	ptr = v.Addr().Interface().(*T)
	val = v.Interface().(T)
	return
}

func ptr[T any](v T) *T {
	return &v
}

func parsePK(s string) (int, error) {
	var pk int
	_, err := fmt.Sscanf(s, "%d", &pk)
	if err != nil {
		return 0, fmt.Errorf("invalid PK: %s", s)
	}
	return pk, nil
}
