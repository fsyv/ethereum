// Package reputation implements contract function
package reputation

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestReputations_GetReputation(t *testing.T) {
	type args struct {
		address common.Address
	}
	tests := []struct {
		name    string
		r       *Reputation
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetReputation(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reputation.GetReputation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reputation.GetReputation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReputations_GetRepThreshold(t *testing.T) {
	tests := []struct {
		name    string
		r       *Reputation
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetRepThreshold()
			if (err != nil) != tt.wantErr {
				t.Errorf("Reputation.GetRepThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reputation.GetRepThreshold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReputations_UpdateReputation(t *testing.T) {
	type args struct {
		address common.Address
		score   *big.Int
	}
	tests := []struct {
		name    string
		r       *Reputation
		args    args
		want    *types.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UpdateReputation(tt.args.address, tt.args.score)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reputation.UpdateReputation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reputation.UpdateReputation() = %v, want %v", got, tt.want)
			}
		})
	}
}
