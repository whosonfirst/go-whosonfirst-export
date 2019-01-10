package properties

func EnsureRequired(feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureName(feature)

	if err != nil {
		return nil, err
	}

	feature, err = EnsurePlacetype(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
