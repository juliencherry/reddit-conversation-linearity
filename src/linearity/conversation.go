package linearity

type Conversation interface {
    BranchLengths() []int
}

type CommentWithChildren struct {
    utterance string
    Children []Conversation
}

func (c CommentWithChildren) BranchLengths() []int {
    lengths := []int{}
    for _, child := range c.Children {

        childLengths := child.BranchLengths()
        for _, childLength := range childLengths {
            lengths = append(lengths, childLength + 1)
        }
    }

    return lengths
}

type FinalComment struct {
    utterance string
}

func (FinalComment) BranchLengths() []int {
    return []int{1}
}
