package upctl

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/uptime-com/uptime-client-go/v2/pkg/upapi"
)

type flagSetMock struct {
	mock.Mock
}

func (f *flagSetMock) StringVarP(ptr *string, name, shorthand, value, usage string) {
	f.Called(ptr, name, shorthand, value, usage)
}

func (f *flagSetMock) IntVarP(ptr *int, name, shorthand string, value int, usage string) {
	f.Called(ptr, name, shorthand, value, usage)
}

func (f *flagSetMock) BoolVarP(ptr *bool, name, shorthand string, value bool, usage string) {
	f.Called(ptr, name, shorthand, value, usage)
}

func (f *flagSetMock) StringSliceVarP(ptr *[]string, name, shorthand string, value []string, usage string) {
	f.Called(ptr, name, shorthand, value, usage)
}

func (f *flagSetMock) Float64VarP(ptr *float64, name, shorthand string, value float64, usage string) {
	f.Called(ptr, name, shorthand, value, usage)
}

func TestBind(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Run("string", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("StringVarP", mock.Anything, "str", "s", "default", "Lenin lives!").Return().Once()
			obj := struct {
				Str string `flag:"str" short:"s" usage:"Lenin lives!"`
			}{
				Str: "default",
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
		t.Run("int", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("IntVarP", mock.Anything, "int", "i", 42, "Lenin lives!").Return().Once()
			obj := struct {
				Int int `flag:"int" short:"i" usage:"Lenin lives!"`
			}{
				Int: 42,
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
		t.Run("bool", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("BoolVarP", mock.Anything, "bool", "b", true, "Lenin lives!").Return().Once()
			obj := struct {
				Bool bool `flag:"bool" short:"b" usage:"Lenin lives!"`
			}{
				Bool: true,
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
		t.Run("float64", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("Float64VarP", mock.Anything, "float", "f", 3.14, "Lenin lives!").Return().Once()
			obj := struct {
				Float64 float64 `flag:"float" short:"f" usage:"Lenin lives!"`
			}{
				Float64: 3.14,
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
	})
	t.Run("skip", func(t *testing.T) {
		fs := flagSetMock{}
		fs.On("StringVarP", mock.Anything, "str", "s", "default", "Lenin lives!").Return().Once()
		obj := struct {
			Bool bool   `flag:"bool" short:"b" usage:"Lenin lives!" skip:"-"`
			Str  string `flag:"str" short:"s" usage:"Lenin lives!"`
		}{
			Bool: true,
			Str:  "default",
		}
		err := Bind(&fs, &obj)
		require.NoError(t, err)
		fs.AssertExpectations(t)
	})
	t.Run("kebab case", func(t *testing.T) {
		fs := flagSetMock{}
		fs.On("BoolVarP", mock.Anything, "bool-bool", "", true, "").Return().Once()
		obj := struct {
			BoolBool bool
		}{
			BoolBool: true,
		}
		err := Bind(&fs, &obj)
		require.NoError(t, err)
		fs.AssertExpectations(t)
	})
	t.Run("nested", func(t *testing.T) {
		t.Run("struct", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("StringVarP", mock.Anything, "tag", "", "", "").Return().Once()
			fs.On("StringVarP", mock.Anything, "color-hex", "", "#ff0000", "").Return().Once()
			obj := struct {
				upapi.Tag
			}{
				Tag: upapi.Tag{
					ColorHex: "#ff0000",
				},
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
		t.Run("pointer", func(t *testing.T) {
			fs := flagSetMock{}
			fs.On("StringVarP", mock.Anything, "tag", "", "", "").Return().Once()
			fs.On("StringVarP", mock.Anything, "color-hex", "", "#ff0000", "").Return().Once()
			obj := struct {
				*upapi.Tag
			}{
				Tag: &upapi.Tag{
					ColorHex: "#ff0000",
				},
			}
			err := Bind(&fs, &obj)
			require.NoError(t, err)
			fs.AssertExpectations(t)
		})
	})
}
