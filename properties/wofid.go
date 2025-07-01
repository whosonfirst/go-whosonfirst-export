package properties

import (
	"context"
	"fmt"
	"sync"

	"log/slog"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	wof_properties "github.com/whosonfirst/go-whosonfirst-feature/properties"
	"github.com/whosonfirst/go-whosonfirst-id"
)

const ID_PROVIDER string = "org.whosonfirst.id.provider"

var provider_once = sync.OnceValues(func() (id.Provider, error) {
	return id.NewProvider(context.Background())
})

func idProvider(ctx context.Context) (id.Provider, error) {

	v := ctx.Value(ID_PROVIDER)

	if v != nil {

		if _, ok := v.(id.Provider); ok {
			return v.(id.Provider), nil
		} else {
			return nil, fmt.Errorf("Invalid Id provider %T", v)
		}
	}

	return provider_once()
}

func EnsureWOFIdAlt(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, wof_properties.PATH_WOF_ID)

	if !rsp.Exists() {
		return nil, wof_properties.MissingProperty(wof_properties.PATH_WOF_ID)
	}

	return feature, nil
}

func EnsureWOFId(ctx context.Context, feature []byte) ([]byte, error) {

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("Verbose logging enabled")

	slog.Info("PR 1")

	provider, err := idProvider(ctx)

	if err != nil {
		return nil, err
	}

	slog.Info("PR 2")

	var wof_id int64

	rsp := gjson.GetBytes(feature, wof_properties.PATH_WOF_ID)

	slog.Info("PR 3")
	if rsp.Exists() {

		slog.Info("PR 4")
		wof_id = rsp.Int()

	} else {

		slog.Info("PR 5")
		i, err := provider.NewID(ctx)

		slog.Info("PR 5a")
		if err != nil {
			return nil, err
		}

		slog.Info("PR 6")
		wof_id = i

		feature, err = sjson.SetBytes(feature, wof_properties.PATH_WOF_ID, wof_id)

		if err != nil {
			return nil, SetPropertyFailed(wof_properties.PATH_WOF_ID, err)
		}
	}

	slog.Info("PR 7")
	feature, err = sjson.SetBytes(feature, wof_properties.PATH_ID, wof_id)

	if err != nil {
		return nil, SetPropertyFailed(wof_properties.PATH_ID, err)
	}

	return feature, nil
}
