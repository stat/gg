package env

//
// Accessors
//

// Get

func Get[T any](key string) (T, error) {
	var value T

	aquired, err := Aquire[T](key)

	if err != nil {
		return value, err
	}

	return *aquired, nil
}

// Load

func Load[T any](key string, v *T) error {
	aquired, err := Aquire[T](key)

	if err != nil {
		return err
	}

	*v = *aquired

	return nil
}
