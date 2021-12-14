package dtypes

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArrayTypes(t *testing.T) {
	arrayTypes := arrayType()
	expected := 2

	fmt.Printf("types of arrayTypes is %s\n", reflect.TypeOf(arrayTypes))
	if expected != arrayTypes[1] {
		t.Errorf("expected %d but got %d\n", expected, arrayTypes[1])
	}

}

func TestStructType(t *testing.T) {
	structType()
}

func TestInterf(t *testing.T) {
	interf()
}
