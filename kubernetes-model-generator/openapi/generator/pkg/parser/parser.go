/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package parser

import (
	"fmt"
	"github.com/fabric8io/kubernetes-client/kubernetes-model-generator/openapi/generator/pkg/kubernetes"
	"github.com/fabric8io/kubernetes-client/kubernetes-model-generator/openapi/generator/pkg/openapi"
	"k8s.io/gengo/v2/parser"
	"k8s.io/gengo/v2/types"
	"strings"
)

const genClient = "+genclient"
const genClientPrefix = genClient + ":"
const groupNamePrefix = "+groupName="
const versionNamePrefix = "+versionName="

type Module struct {
	patterns []string
	parser   *parser.Parser
	universe *types.Universe
}

type Fabric8Info struct {
	// https://github.com/kubernetes/community/blob/495011674de058660011593e6c6c842c83a1fd24/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Type    string
	Group   string
	Version string
	Kind    string
	Scope   string
}

func NewModule(patterns ...string) *Module {
	p := parser.New()
	err := p.LoadPackages(patterns...)
	if err != nil {
		panic(fmt.Sprintf("error loading packages: %v", err))
	}
	universe, err := p.NewUniverse()
	if err != nil {
		panic(fmt.Sprintf("error creating universe: %v", err))
	}
	return &Module{
		patterns: patterns,
		parser:   p,
		universe: &universe,
	}
}

func (oam *Module) ExtractInfo(definitionName string) *Fabric8Info {
	pkg := oam.resolvePackage(definitionName)
	typ := oam.universe.Type(types.ParseFullyQualifiedName(definitionName))
	fabric8Info := &Fabric8Info{}
	fabric8Info.Type = resolveType(typ)
	fabric8Info.Group = groupName(pkg)
	fabric8Info.Version = versionName(pkg)
	fabric8Info.Kind = typ.Name.Name
	fabric8Info.Scope = scope(typ)
	return fabric8Info
}

func (oam *Module) ApiName(definitionName string) string {
	// Don't treat k8s.io types, json is expected to contain the full Go definition name instead of the group/version
	if strings.HasPrefix(definitionName, "k8s.io/") {
		return openapi.FriendlyName(definitionName)
	}
	lastSeparator := strings.LastIndex(definitionName, ".")
	typeName := definitionName[lastSeparator+1:]
	pkg := oam.resolvePackage(definitionName)
	gn := groupName(pkg)
	if gn == "" {
		return openapi.FriendlyName(definitionName)
	}
	groupParts := strings.Split(gn, ".")
	for i, j := 0, len(groupParts)-1; i < j; i, j = i+1, j-1 {
		groupParts[i], groupParts[j] = groupParts[j], groupParts[i]
	}
	return strings.Join(groupParts, ".") + "." + versionName(pkg) + "." + typeName
}

func (oam *Module) resolvePackage(definitionName string) *types.Package {
	lastSeparator := strings.LastIndex(definitionName, ".")
	packageName := definitionName[:lastSeparator]
	pkg := oam.universe.Package(packageName)
	if pkg == nil || pkg.Name == "" {
		_, err := oam.parser.LoadPackagesTo(oam.universe, packageName)
		if err != nil {
			panic(fmt.Sprintf("error loading packages: %v", err))
		}
	}
	return pkg
}

func groupName(pkg *types.Package) string {
	return findTag(pkg, groupNamePrefix)
}

func versionName(pkg *types.Package) string {
	// Some packages have a versionName tag, but it's not the standard (usually package name is used instead)
	v := findTag(pkg, versionNamePrefix)
	if v != "" {
		return v
	}
	return pkg.Name
}

func findTag(pkg *types.Package, tag string) string {
	for _, c := range pkg.Comments {
		if strings.HasPrefix(c, tag) {
			t := strings.TrimPrefix(c, tag)
			// Incredibly 🤦, there are some operators with a badly defined groupName tag:
			// https://github.com/openshift/installer/blob/66e9daae9dae59c2ec167a6c31f2f2c127382357/pkg/types/doc.go#L1
			t = strings.ReplaceAll(t, "\"", "")
			return t
		}
	}
	return ""
}

func resolveType(typ *types.Type) string {
	// types with a genClient annotation are always objects
	// types with a listType are always lists
	// types with metaType are always lists or objects
	for _, c := range append(typ.CommentLines, typ.SecondClosestCommentLines...) {
		if strings.TrimSpace(c) == genClient {
			return "object"
		}
	}
	isList := false
	isObject := false
	for _, m := range typ.Members {
		if m.Embedded == true && m.Type.Name == kubernetes.ListMeta {
			isList = true
			// stop iterating, if it's a list then it's not an object
			break
		} else if m.Embedded == true && m.Type.Name == kubernetes.TypeMeta {
			isObject = true
			// keep iterating, maybe it's a list
		}
	}
	if isList {
		return "list"
	}
	if isObject {
		return "object"
	}
	return "nested"
}

func scope(typ *types.Type) string {
	scope := "Namespaced"
	for _, c := range append(typ.CommentLines, typ.SecondClosestCommentLines...) {
		if strings.Contains(c, genClientPrefix+"nonNamespaced") {
			scope = "Clustered"
			break
		}
	}
	return scope
}
