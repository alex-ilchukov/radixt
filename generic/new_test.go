package generic_test

import (
	"testing"

	"github.com/alex-ilchukov/radixt/evident"
	"github.com/alex-ilchukov/radixt/generic"
)

var newTests = []struct {
	strings []string
	e       evident.Tree
}{
	{strings: nil, e: nil},
	{strings: []string{}, e: nil},
	{
		strings: []string{
			"authority",
			"authorization",
			"author",
			"authentication",
			"auth",
			"content-type",
			"content-length",
			"content-disposition",
		},
		e: evident.Tree{
			"|": {
				"auth|4": {
					"entication|3": nil,
					"or|2": {
						"i|": {
							"ty|0":     nil,
							"zation|1": nil,
						},
					},
				},
				"content-|": {
					"disposition|7": nil,
					"length|6":      nil,
					"type|5":        nil,
				},
			},
		},
	},
}

const testNewError = "New Test %d: got that New(%v...) is\n\n%v\n\nwhich is " +
	"not equal to \n\n%v\n\n (but should be equal)"

func TestNew(t *testing.T) {
	for i, tt := range newTests {
		result := generic.New(tt.strings...)
		if !tt.e.Eq(result) {
			t.Errorf(testNewError, i, tt.strings, result, tt.e)
		}
	}
}

var newFromSVTests = []struct {
	sv []generic.SV
	e  evident.Tree
}{
	{sv: nil, e: nil},
	{sv: []generic.SV{}, e: nil},
	{
		sv: []generic.SV{
			{S: "authority", V: 123},
			{S: "authorization", V: 456},
			{S: "author", V: 789},
			{S: "authentication", V: 135},
			{S: "auth", V: 680},
			{S: "content-type", V: 234},
			{S: "content-length", V: 453},
			{S: "content-disposition", V: 656534},
			{S: "", V: 9000},
		},
		e: evident.Tree{
			"|9000": {
				"auth|680": {
					"entication|135": nil,
					"or|789": {
						"i|": {
							"ty|123":     nil,
							"zation|456": nil,
						},
					},
				},
				"content-|": {
					"disposition|656534": nil,
					"length|453":         nil,
					"type|234":           nil,
				},
			},
		},
	},
}

const testNewFromSVError = "NewFromSV Test %d: got that NewFromSV(%v...) is" +
	"\n\n%v\n\nwhich is not equal to \n\n%v\n\n (but should be equal)"

func TestNewFromSV(t *testing.T) {
	for i, tt := range newFromSVTests {
		result := generic.NewFromSV(tt.sv...)
		if !tt.e.Eq(result) {
			t.Errorf(testNewFromSVError, i, tt.sv, result, tt.e)
		}
	}
}
