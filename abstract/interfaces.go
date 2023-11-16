package interfaces

import "reflect"

type Owner interface {
	ID() uint
	Own(asset Asset) bool
	Owner() Owner
}

type Asset interface {
	SetOwner(owner Owner) Asset
	Owner() Owner
}

func Own(owner Owner, asset Asset) bool {
	if owner == nil {
		return false
	}
	if owner.ID() == asset.Owner().ID() && reflect.TypeOf(owner) == reflect.TypeOf(asset.Owner()) {
		return true
	}
	return Own(owner.Owner(), asset)
}
