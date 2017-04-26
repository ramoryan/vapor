package vapor

import (
	"reflect"
	"testing"
)

func TestIsVariableInitializer(t *testing.T) {
	if !isVariableInitializer("$myvar=kuningaz") || !isVariableInitializer("$myvar    =   treacherous gods") {
		t.Error("With $ sign it must be a variable initializer!")
	}

	if isVariableInitializer("myvar = ups") {
		t.Error("Variable initializer must be start with $ sign!")
	}
}

func TestParseVariable(t *testing.T) {
	// $var = text
	name, value, err := parseVariable("$ensiferum = kardhordozó")
	if name != "ensiferum" || value != "kardhordozó" {
		if err != nil {
			t.Error(err)
		} else {
			t.Error("Unknown variable error!")
		}
	}

	_, _, err = parseVariable("iron = metal")
	if err == nil {
		t.Error("Variable initializer must be start with $ sign!")
	}

	// $var2 = $var1
	name, value, err = parseVariable("$swordBearer = $ensiferum")
	if name != "swordBearer" || value != "kardhordozó" || err != nil {
		t.Error("Variable to variable error!")
	}

	// $var3 = #{ $var }
	name, value, err = parseVariable("$kardhordozo = #{ $swordBearer }")
	if name != "kardhordozo" || value != "kardhordozó" || err != nil {
		t.Error("Error when try to add interpolated variable!")
	}

	// $var4 = :filter text
	name, value, err = parseVariable("$dieLikeKings = :upper majesty")
	if name != "dieLikeKings" || value != "MAJESTY" || err != nil {
		t.Error("Initializing with filter has been broken!")
	}

	// $var5 = :filter $var
	name, value, err = parseVariable("$peaceTime = :lower $dieLikeKings")
	if name != "peaceTime" || value != "majesty" || err != nil {
		t.Error("Init error with filter and variable!")
	}

	// $var6 = :notvalidfilter text
	name, value, err = parseVariable("$notValid = :notvalid text")
	if err == nil {
		t.Error("Initializer must returns with filter error!")
	}

	// $var7 = :filter $undefinedVar
	name, value, err = parseVariable("$argh = :upper $undefinedVar")
	if err == nil {
		t.Error("Initializer must returns with undefined variable error!")
	}

	// $var8 = $anotherUndefined
	name, value, err = parseVariable("$argh2 = $anotherUndefined")
	if err == nil {
		t.Error("Initializer must returns with undefined variable error!")
	}

	// --- SLICES
	AddStrSliceVar("strSlice", []string{"a", "b", "c"})
	iface, ok := findVariable("strSlice")
	if !ok {
		t.Error("strSlice variable not found!")
	}

	switch reflect.TypeOf(iface).Kind() {
	case reflect.Slice:
		// Everything is ok!
		/*s := reflect.ValueOf(iface)
		t.Log(s)*/
	default:
		t.Error("Variable is not slice of strings!")
	}
}
