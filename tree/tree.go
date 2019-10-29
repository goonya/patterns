package tree

import (
	"strings"
)

// Branch ...
type Branch struct {
	ID       string
	ParentID string
	Type     string
}

// Tree ...
type Tree []Branch

func (tree Tree) String() string {
	var b strings.Builder
	var parentID string
	for _, branch := range tree {
		if parentID != branch.ParentID {
			b.WriteString("\n")
			b.WriteString(branch.ParentID)
			b.WriteString(" <- ")
			parentID = branch.ParentID
		} else {
			b.WriteString(",")
		}
		if branch.Type != "o" {
			b.WriteString(branch.Type)
			b.WriteString("_")
		}
		b.WriteString(branch.ID)
	}
	return b.String()
}

func f(childrenByParentID map[string]map[string]Branch, parentID string) {
	if children, ok := childrenByParentID[parentID]; ok {
		for _, branch := range children {
			if branch.Type == "o" {
				f(childrenByParentID, branch.ID)
				c, ok := childrenByParentID[branch.ID]
				if !ok || (len(c) == 0) {
					delete(children, branch.ID)
					childrenByParentID[parentID] = children
				}
			}
		}
		if len(children) == 0 {
			delete(childrenByParentID, parentID)
		}
	}
}

// RemoveEmptyBranches ...
func (tree Tree) RemoveEmptyBranches() {
	childrenByParentID := make(map[string]map[string]Branch)
	for _, branch := range tree {
		children, ok := childrenByParentID[branch.ParentID]
		if !ok {
			children = make(map[string]Branch)
		}
		children[branch.ID] = branch
		childrenByParentID[branch.ParentID] = children
	}
	//fmt.Println(childrenByParentID) // TODO: delete
	f(childrenByParentID, "#")
	//fmt.Println(childrenByParentID) // TODO: delete
	for _, branch := range tree {
		//
	}
}
