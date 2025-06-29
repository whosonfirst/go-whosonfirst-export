package properties

import (
	"context"
	"fmt"
	"sync"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"github.com/whosonfirst/go-whosonfirst-id"
)

const ID_PROVIDER string = "org.whosonfirst.id.provider"

var provider_once = sync.OnceValues(func() (id.Provider, error) {
	return id.NewProvider(context.Background())
})

func idProvider(ctx context.Context) (id.Provider, error) {

	v := ctx.Value(ID_PROVIDER)

	if v != nil {

		switch v.(type) {
		case id.Provider:
			return v.(id.Provider), nil
		default:
			return nil, fmt.Errorf("Invalid Id provider")
		}
	}

	return provider_once()
}

func EnsureWOFId(ctx context.Context, feature []byte) ([]byte, error) {

	provider, err := idProvider(ctx)

	if err != nil {
		return nil, err
	}

	var wof_id int64

	rsp := gjson.GetBytes(feature, "properties.wof:id")

	if rsp.Exists() {

		wof_id = rsp.Int()

	} else {

		i, err := provider.NewID(ctx)

		if err != nil {
			return nil, err
		}

		wof_id = i

		feature, err = sjson.SetBytes(feature, "properties.wof:id", wof_id)

		if err != nil {
			return nil, err
		}
	}

	id := gjson.GetBytes(feature, "id")

	if !id.Exists() {

		feature, err = sjson.SetBytes(feature, "id", wof_id)

		if err != nil {
			return nil, err
		}

	}

	return feature, nil
}
