package main

import (
	"go/importer"
	"go/types"
	"regexp"

	"github.com/paralin/go-dota2/protocol/all"
)

// TypeRegex matches struct types in Go
var TypeRegex = regexp.MustCompile("(type)(.+)(struct{)(.+)(})")

// ProtoType represents a protobuf message type with some heuristics.
type ProtoType struct {
	// PakStr is the containing package import string
	PakStr string
	// Pak is the containing package.
	Pak *types.Package
	// TypeName is the name of the type, like CSODOTALobby
	TypeName string
	// Obj is the object.
	Obj types.Object
	// TypeStr is the fully qualified type string.
	TypeStr string
}

// BuildProtoTypeMap attempts to compile the proto type list.
func BuildProtoTypeMap() (map[string]*types.Package, map[string]*ProtoType, error) {
	packages := make(map[string]*types.Package)
	imp := importer.Default()

	importPak := func(pak string) error {
		ipak, err := imp.Import(pak)
		if err != nil {
			return err
		}

		packages[pak] = ipak
		return nil
	}

	// import all packages
	for _, pak := range all.Packages {
		if err := importPak(pak); err != nil {
			return nil, nil, err
		}
	}

	if err := importPak("github.com/faceit/go-steam/steamid"); err != nil {
		return nil, nil, err
	}

	// build a map of proto types
	protoMap := make(map[string]*ProtoType)
	for pakStr, pak := range packages {
		scope := pak.Scope()
		for _, nam := range scope.Names() {
			obj := scope.Lookup(nam)
			objStr := obj.String()
			if !obj.Exported() {
				continue
			}

			if !TypeRegex.MatchString(objStr) {
				continue
			}

			protoTyp := &ProtoType{
				PakStr:   pakStr,
				Pak:      pak,
				TypeName: obj.Name(),
				Obj:      obj,
				TypeStr:  obj.String(),
			}

			protoMap[obj.Name()] = protoTyp
		}
	}

	return packages, protoMap, nil
}
