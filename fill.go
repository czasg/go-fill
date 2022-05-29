package fill

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Opt int

const (
	OptEnv    Opt = 1 << iota
	OptSilent     // ignore err and iterate over all fields.
)

type Payload struct {
	Value       reflect.Value
	Prefix      string
	Opt         Opt
	Field       reflect.Value
	StructField reflect.StructField
}

var (
	NotPointerStructErr = errors.New("only the pointer to a struct is supported")
)

func FillEnv(v interface{}) error {
	return Fill(v, OptEnv)
}

func Fill(v interface{}, opts ...Opt) error {
	ind := reflect.Indirect(reflect.ValueOf(v))
	if reflect.ValueOf(v).Kind() != reflect.Ptr || ind.Kind() != reflect.Struct {
		return NotPointerStructErr
	}
	var opt Opt = 0
	for _, o := range opts {
		opt = opt | o
	}
	return fill(Payload{Value: ind, Opt: opt})
}

func fill(payload Payload) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(payload.StructField.Name, err)
		}
	}()
	for i := 0; i < payload.Value.NumField(); i++ {
		payload.Field = payload.Value.Field(i)
		payload.StructField = payload.Value.Type().Field(i)
		err := fill2(payload)
		if err != nil && payload.Opt&OptSilent != OptSilent {
			return err
		}
	}
	return nil
}

func fill2(payload Payload) error {
	switch payload.Field.Kind() {
	case reflect.Struct:
		return parseStruct(payload)
	case reflect.String:
		return parseString(payload)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return parseInt(payload)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return parseUint(payload)
	case reflect.Float32, reflect.Float64:
		return parseFloat(payload)
	case reflect.Map:
		return parseMap(payload)
	case reflect.Slice:
		return parseSlice(payload)
	case reflect.Bool:
		return parseBool(payload)
	default:
		return nil
	}
}

/*
value `fill:"field,sep=_,default=df,require,empty"`
*/
func parseValue(payload Payload) (string, error) {
	fillStr, _ := payload.StructField.Tag.Lookup("fill")
	var value string
	fillName := strings.ToUpper(payload.StructField.Name)
	fillRequire := false
	fillDefault := ""
	sep := "_"
	for index, str := range strings.Split(fillStr, ",") {
		if index == 0 && str != "" {
			fillName = str
		} else if strings.Contains(str, "require") {
			fillRequire = true
		} else if strings.Contains(str, "default=") {
			fillDefault = strings.TrimPrefix(str, "default=")
		} else if strings.Contains(str, "sep=") {
			sep = strings.TrimPrefix(str, "sep=")
		} else if strings.Contains(str, "empty") {
			fillName = ""
		}
	}
	if payload.Opt&OptEnv == OptEnv {
		if payload.Prefix != "" {
			if fillName == "" {
				fillName = payload.Prefix
			} else {
				fillName = fmt.Sprintf("%s%s%s", payload.Prefix, sep, fillName)
			}
		}
		envValue, exist := os.LookupEnv(fillName)
		if exist {
			value = envValue
		}
	}
	if value == "" && fillDefault != "" {
		value = fillDefault
	}
	if value == "" && fillRequire {
		return "", fmt.Errorf("%s require", fillName)
	}
	return value, nil
}

/*
struct `fill:"field,sep=_,default=df,require,empty"`
*/
func parseStruct(payload Payload) error {
	payload.Value = payload.Field
	fieldName := strings.ToUpper(payload.StructField.Name)
	sep := "_"
	fillStr, exist := payload.StructField.Tag.Lookup("fill")
	if exist {
		for index, str := range strings.Split(fillStr, ",") {
			if index == 0 && str != "" {
				fieldName = str
			} else if strings.Contains(str, "sep=") {
				sep = strings.TrimPrefix(str, "sep=")
			} else if strings.Contains(str, "empty") {
				fieldName = ""
			}
		}
	}
	if payload.Prefix == "" {
		payload.Prefix = fieldName
		return fill(payload)
	}
	if fieldName == "" {
		return fill(payload)
	}
	payload.Prefix = fmt.Sprintf("%s%s%s", payload.Prefix, sep, fieldName)
	return fill(payload)
}

func parseString(payload Payload) error {
	if payload.Field.String() != "" {
		return nil
	}
	value, err := parseValue(payload)
	if err != nil {
		return err
	}
	if value == "" {
		return nil
	}
	payload.Field.SetString(value)
	return nil
}

func parseInt(payload Payload) error {
	if payload.Field.Int() != 0 {
		return nil
	}
	value, err := parseValue(payload)
	if err != nil {
		return err
	}
	if value == "" {
		return nil
	}
	iv, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("%s invalid [%s]", payload.StructField.Name, value)
	}
	payload.Field.SetInt(int64(iv))
	return nil
}

func parseUint(payload Payload) error {
	if payload.Field.Uint() != 0 {
		return nil
	}
	value, err := parseValue(payload)
	if err != nil {
		return err
	}
	if value == "" {
		return nil
	}
	iv, err := strconv.ParseUint(value, 0, 64)
	if err != nil {
		return fmt.Errorf("%s invalid [%s]", payload.StructField.Name, value)
	}
	payload.Field.SetUint(iv)
	return nil
}

func parseFloat(payload Payload) error {
	if payload.Field.Float() != 0 {
		return nil
	}
	value, err := parseValue(payload)
	if err != nil {
		return err
	}
	if value == "" {
		return nil
	}
	iv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("%s invalid [%s]", payload.StructField.Name, value)
	}
	payload.Field.SetFloat(iv)
	return nil
}

func parseBool(payload Payload) error {
	if payload.Field.Bool() {
		return nil
	}
	value, err := parseValue(payload)
	if err != nil {
		return err
	}
	if value == "" {
		return nil
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("%s invalid [%s]", payload.StructField.Name, value)
	}
	payload.Field.SetBool(b)
	return nil
}

func parseMap(payload Payload) error {
	if payload.Field.IsNil() {
		payload.Field.Set(reflect.MakeMap(payload.Field.Type()))
	}
	return nil
}

func parseSlice(payload Payload) error {
	if payload.Field.IsNil() {
		payload.Field.Set(reflect.MakeSlice(payload.Field.Type(), 0, 0))
	}
	return nil
}
