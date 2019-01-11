package properties

import ()

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

	feature, err = EnsureGeom(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func EnsureGeom(feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureSrcGeom(feature)

	if err != nil {
		return nil, err
	}

	feature, err = EnsureGeomHash(feature)

	if err != nil {
		return nil, err
	}

	// TO DO
	// geom:latitude
	// geom:longitude
	// geom:bbox

	return feature, nil
}

func EnsureTimestamps(feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureCreated(feature)

	if err != nil {
		return nil, err
	}

	feature, err = EnsureLastModified(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
