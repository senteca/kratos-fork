package internal

import (
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/bxcodec/faker"

	"github.com/ory/x/randx"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/flow/registration"
	"github.com/ory/kratos/selfservice/form"
)

func RegisterFakes() {
	if err := faker.AddProvider("birthdate", func(v reflect.Value) (interface{}, error) {
		return time.Now().Add(time.Duration(rand.Int())).Round(time.Second).UTC(), nil
	}); err != nil {
		panic(err)
	}

	if err := faker.AddProvider("time_types", func(v reflect.Value) (interface{}, error) {
		es := make([]time.Time, rand.Intn(5))
		for k := range es {
			es[k] = time.Now().Add(time.Duration(rand.Int())).Round(time.Second).UTC()
		}
		return es, nil
	}); err != nil {
		panic(err)
	}

	if err := faker.AddProvider("http_header", func(v reflect.Value) (interface{}, error) {
		headers := http.Header{}
		for i := 0; i <= rand.Intn(5); i++ {
			values := make([]string, rand.Intn(4)+1)
			for k := range values {
				values[k] = randx.MustString(8, randx.AlphaNum)
			}
			headers[randx.MustString(8, randx.AlphaNum)] = values
		}

		return headers, nil
	}); err != nil {
		panic(err)
	}

	if err := faker.AddProvider("time_type", func(v reflect.Value) (interface{}, error) {
		return time.Now().Add(time.Duration(rand.Int())).Round(time.Second).UTC(), nil
	}); err != nil {
		panic(err)
	}

	if err := faker.AddProvider("login_request_methods", func(v reflect.Value) (interface{}, error) {
		methods := map[identity.CredentialsType]*login.RequestMethod{}
		for i := 0; i <= rand.Intn(3); i++ {
			var f form.HTMLForm
			if err := faker.FakeData(&f); err != nil {
				return nil, err
			}
			ct := identity.CredentialsType(randx.MustString(8, randx.AlphaLower))
			methods[ct] = &login.RequestMethod{
				Method: ct,
				Config: &f,
			}
		}

		return methods, nil
	}); err != nil {
		panic(err)
	}

	if err := faker.AddProvider("registration_request_methods", func(v reflect.Value) (interface{}, error) {
		methods := map[identity.CredentialsType]*registration.RequestMethod{}
		for i := 0; i <= rand.Intn(3); i++ {
			var f form.HTMLForm
			if err := faker.FakeData(&f); err != nil {
				return nil, err
			}
			ct := identity.CredentialsType(randx.MustString(8, randx.AlphaLower))
			methods[ct] = &registration.RequestMethod{
				Method: ct,
				Config: &f,
			}
		}

		return methods, nil
	}); err != nil {
		panic(err)
	}
}
