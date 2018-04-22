package linearity

import (
    "sort"
)

type Calculator struct {}

func (c Calculator) Linearity(conversation Conversation) float32 {
    branchLengths := conversation.BranchLengths()

    var totalLength = 0
    for _, branchLength := range branchLengths {
        totalLength += branchLength
    }

    numBranches := len(branchLengths)

    sort.Sort(sort.Reverse(sort.IntSlice(branchLengths)))
    longestBranch := branchLengths[0]

    return 1 - (float32(totalLength) / float32(numBranches)) / float32(longestBranch)
}
