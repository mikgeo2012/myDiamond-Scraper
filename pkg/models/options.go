package models

import (
    "strings"
    "reflect"
    "fmt"
    "log"
)

// Name of the struct tag used in examples.
const (
    tagName = "validate"
    OptionTypes = "shape, cut, color, clarity, carat, price, lab"
)

type Options struct {
    Shape []string `validate:"shape,min=1,max=10"`
    Cut []string `validate:"cut,min=1,max=4"`
    Color []string `validate:"color,min=1,max=10"`
    Clarity []string `validate:"clarity,min=1,max=9"`
    Carat []float32 `validate:"carat,min=1,max=2"`
    Price []float32 `validate:"price,min=1,max=2"`
    Lab []string `validate:"lab,min=1,max=3"`
}

func GetDefault() Options {
    return Options{[]string{"round" ,"cushion" }, []string{"Very Good", "Ideal"},
        []string{"M" , "L" , "K"}, []string{"I1" ,"SI2"}, []float32{2.5, 29.0}, []float32{200.00, 5000000.00}, []string{"GIA"}}
}

func (o *Options) Validate() bool {

    ok, err := validateStruct(*o)
    if len(err) != 0 {
        for _,e := range err {
            log.Printf("Error: %v", e.Error())
        }
    }
    return ok
}

// Option Value Validator

// Generic data validator.
type Validator interface {
    // Validate method performs validation and returns result and optional error.
    Validate(val interface{}) (bool, error)
}

// DefaultValidator does not perform any validations.
type DefaultValidator struct {
    OptionsValidator
    SliceValidator
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {

    if ok,e := v.SliceValidator.Validate("", val); !ok {
        return false, fmt.Errorf("slice length error: %s", e.Error())
    }

    if ok,e := v.OptionsValidator.Validate(val); !ok {
        return false, fmt.Errorf("invalid option error: %s", e.Error())
    }

    return true, nil

}

// StringValidator validates string presence and/or its length.
type SliceValidator struct {
    Min int
    Max int
}

func (v SliceValidator) Validate(option string, val interface{}) (bool, error) {
    switch va := val.(type) {
    case []string:
        l := len(va)
        if l == 0 {
            return false, fmt.Errorf("cannot be empty")
        }

        if l < v.Min {
            return false, fmt.Errorf("should be at least %v strings long", v.Min)
        }

        if v.Max >= v.Min && l > v.Max {
            return false,fmt.Errorf("should be less than %v strings long", v.Max)
        }
    case []float32:
        l := len(va)
        if l == 0 {
            return false, fmt.Errorf("cannot be empty")
        }

        if l < v.Min {
            return false, fmt.Errorf("should be at least %v int long", v.Min)
        }

        if v.Max >= v.Min && l > v.Max {
            return false,fmt.Errorf("should be less than %v int long", v.Max)
        }
    default:
        return false, fmt.Errorf("Not valid type")
    }

    return true, nil
}

// NumberValidator performs numerical value validation.
// Its limited to int type for simplicity.
type OptionsValidator struct {
    options interface{}
}

func (v OptionsValidator) Validate(val interface{}) (bool, error) {
    opts := v.options

    switch va := opts.(type) {
    case []string:
        m := make(map[string]bool)
        for i := 0; i < len(va); i++ {
            m[va[i]] = true
        }

        ov, ok := val.([]string)
        if !ok {
            fmt.Errorf("type error: expected %s, got %s", "string", reflect.ValueOf(val).Type().String())
        }

        for _, o := range ov {
            if _, ok := m[o]; !ok {
                return false, fmt.Errorf("should not contain %s", o)
            }
        }
    case []float32:
        ov, ok := val.([]float32)
        if !ok {
            fmt.Errorf("type error: expected %s, got %s", "float", reflect.ValueOf(val).Type().String())
        }

        if ov[0] < va[0] || ov[0] > va[1] {
            return false, fmt.Errorf("lower bound %v is wrong", ov[0])
        } else if ov[1] < va[0] || ov[1] > va[1] {
            return false, fmt.Errorf("upper bound %v is wrong", ov[1])
        }


    }

    return true, nil
}

func getValidatorFromTag(tag string) Validator {
    args := strings.Split(tag, ",")

    var df DefaultValidator

    switch args[0] {
    case "shape":
        optval := OptionsValidator{[]string{"round" ,"cushion" ,"emerald" ,"oval" ,"radiant" ,"asscher" ,"marquise" ,"princess"}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "cut":
        optval := OptionsValidator{[]string{"Good", "Very Good", "Ideal", "TrueHearts"}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "color":
        optval := OptionsValidator{[]string{"M" , "L" , "K" , "J" , "I" , "H" , "G" , "F" , "E" , "D"}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "clarity":
        optval := OptionsValidator{[]string{"I1" ,"SI2" ,"SI1" ,"VS2" ,"VS1" ,"VVS2" ,"VVS1" ,"IF" ,"FL"}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "price":
        optval := OptionsValidator{[]float32{200.00, 5000000.00}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "carat":
        optval := OptionsValidator{[]float32{0.05, 30.0}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    case "lab":
        optval := OptionsValidator{[]string{"GIA", "AGS", "IGI"}}
        slival := SliceValidator{}
        fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &slival.Min, &slival.Max)
        df = DefaultValidator{optval, slival}
    default:
        df = DefaultValidator{}
    }

    return df
}

// Performs actual data validation using validator definitions on the struct
func validateStruct(s interface{}) (bool, []error) {
    errs := []error{}



    // ValueOf returns a Value representing the run-time data
    v := reflect.ValueOf(s)

    for i := 0; i < v.NumField(); i++ {
        // Get the field tag value
        tag := v.Type().Field(i).Tag.Get(tagName)

        // Skip if tag is not defined or ignored
        if tag == "" || tag == "-" {
            continue
        }

        // Get a validator that corresponds to a tag
        validator := getValidatorFromTag(tag)

        // Perform validation
        valid, err := validator.Validate(v.Field(i).Interface())

        // Append error to results
        if !valid && err != nil {
            errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
        }
    }

    if len(errs) != 0 {
        return false, errs
    } else {
        return true, nil
    }
}