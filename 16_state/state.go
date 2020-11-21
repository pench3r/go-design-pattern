package state

import "fmt"

type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

type leaderApprovalState struct{}

func (leaderApprovalState) Approval(m *Machine) {
	fmt.Println("leader passed...")
	m.SetState(getFinanceApprovalState())
}

func (leaderApprovalState) Reject(m *Machine) {
	fmt.Println("leader reject...")
}

func (leaderApprovalState) GetName() string {
	return "LeaderApprovalState"
}

func getLeaderApprovalState() IState {
	return &leaderApprovalState{}
}

type financeApprovalState struct{}

func (financeApprovalState) Approval(m *Machine) {
	fmt.Println("finance passed...")
	fmt.Println("done...")
}

func (financeApprovalState) Reject(m *Machine) {
	fmt.Println("finance reject")
	m.SetState(getLeaderApprovalState())
}

func (financeApprovalState) GetName() string {
	return "FinanceApprovalState"
}

func getFinanceApprovalState() IState {
	return &financeApprovalState{}
}
