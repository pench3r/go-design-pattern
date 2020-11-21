package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMacineApproval(t *testing.T) {
	m := &Machine{state: getLeaderApprovalState()}
	m.Approval()
	assert.Equal(t, "FinanceApprovalState", m.GetStateName())
	m.Reject()
	assert.Equal(t, "LeaderApprovalState", m.GetStateName())
}
