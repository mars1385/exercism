package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	var result Node

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := map[int]*Node{}

	for _, val := range records {

		if val.ID == 0 && val.ID < val.Parent {
			return nil, errors.New("tree root cant have parent")
		}

		if val.ID != 0 && val.ID <= val.Parent {
			return nil, errors.New("invalid records")
		}

		_, hasNode := nodes[val.ID]
		if hasNode {
			return nil, errors.New("already has node")
		}
		_, lastNode := nodes[val.ID-1]
		if !lastNode && val.ID != 0 {
			return nil, errors.New("invalid record id")
		}

		if val.ID == 0 {

			result = Node{
				ID: val.ID,
			}

			nodes[val.ID] = &result
		}

		if val.ID != 0 {
			temp := Node{
				ID: val.ID,
			}
			nodes[val.ID] = &temp
			parent, has := nodes[val.Parent]

			if !has {
				return nil, errors.New("no parent")
			}

			parent.Children = append(parent.Children, &temp)
		}
	}

	return &result, nil
}
