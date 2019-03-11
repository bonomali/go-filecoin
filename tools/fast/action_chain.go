package fast

import (
	"context"
	"encoding/json"

	"github.com/filecoin-project/go-filecoin/types"
)

// ChainHead runs the chain head command against the filecoin process.
func (f *Filecoin) ChainHead(ctx context.Context) ([]types.Block, error) {
	var out []types.Block
	if err := f.RunCmdJSONWithStdin(ctx, nil, &out, "go-filecoin", "chain", "head"); err != nil {
		return nil, err
	}
	return out, nil

}

// ChainLs runs the chain ls command against the filecoin process.
func (f *Filecoin) ChainLs(ctx context.Context) (*json.Decoder, error) {
	return f.RunCmdLDJSONWithStdin(ctx, nil, "go-filecoin", "chain", "ls")
}
