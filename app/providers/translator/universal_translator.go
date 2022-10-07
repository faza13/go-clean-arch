package translator

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"log"
)

type ITranslator interface {
	locales.Translator

	// creates the translations for the locale given the 'key' and params passed in.
	// wraps ut.Translator.T to handle errors
	T(key interface{}, params ...string) string

	// creates the cardinal translations for the locale given the 'key', 'num' and 'digit' arguments
	//  and param passed in.
	// wraps ut.Translator.C to handle errors
	C(key interface{}, num float64, digits uint64, param string) string

	// creates the ordinal translations for the locale given the 'key', 'num' and 'digit' arguments
	// and param passed in.
	// wraps ut.Translator.O to handle errors
	O(key interface{}, num float64, digits uint64, param string) string

	//  creates the range translations for the locale given the 'key', 'num1', 'digit1', 'num2' and
	//  'digit2' arguments and 'param1' and 'param2' passed in
	// wraps ut.Translator.R to handle errors
	R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string

	// Currency returns the type used by the given locale.
	Currency() currency.Type

	GetTrans() ut.Translator
}

// implements Translator interface definition above.
type Translator struct {
	locales.Translator
	trans ut.Translator
}

var utrans *ut.UniversalTranslator

var transKey = struct {
	name string
}{
	name: "transKey",
}

func NewTranslatorProvider(locale string) ITranslator {
	en := en.New()
	utrans = ut.New(en, en, id.New())

	err := utrans.Import(ut.FormatJSON, "translations")
	if err != nil {
		log.Fatal(err)
	}

	err = utrans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

	var trans ut.Translator
	switch locale {
	case "id":
		trans, _ = utrans.GetTranslator("id")
		break
	default:
		trans, _ = utrans.GetTranslator("en")
		break
	}

	return &Translator{
		trans: trans,
	}
}

func (t *Translator) T(key interface{}, params ...string) string {

	s, err := t.trans.T(key, params...)
	if err != nil {
		log.Printf("issue translating key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *Translator) C(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating cardinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *Translator) O(key interface{}, num float64, digits uint64, param string) string {

	s, err := t.trans.C(key, num, digits, param)
	if err != nil {
		log.Printf("issue translating ordinal key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *Translator) R(key interface{}, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) string {

	s, err := t.trans.R(key, num1, digits1, num2, digits2, param1, param2)
	if err != nil {
		log.Printf("issue translating range key: '%v' error: '%s'", key, err)
	}

	return s
}

func (t *Translator) Currency() currency.Type {

	// choose your own locale. The reason it isn't mapped for you is because many
	// countries have multiple currencies; it's up to you and you're application how
	// and which currencies to use. I recommend adding a function it to to your custon translator
	// interface like defined above.
	switch t.Locale() {
	case "en":
		return currency.USD
	case "id":
		return currency.IDR
	default:
		return currency.USD
	}
}

func (t *Translator) GetTrans() ut.Translator {
	return t.trans
}
