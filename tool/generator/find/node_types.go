package find

import "github.com/JosephNaberhaus/agnostic/tool/generator/model"

func AllNodeTypes(specs []model.Spec) []string {
	nodeTypeSet := map[string]struct{}{}
	for _, spec := range specs {
		for _, typ := range spec.Types {
			nodeTypeSet[typ] = struct{}{}
		}
	}

	allNodeTypes := make([]string, 0, len(nodeTypeSet))
	for nodeType := range nodeTypeSet {
		allNodeTypes = append(allNodeTypes, nodeType)
	}

	return allNodeTypes
}
