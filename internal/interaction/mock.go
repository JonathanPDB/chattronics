package interaction

type MockUser struct {
	projectPrompt string
	counter       int
}

func CreateMockUser(projectPrompt string) User {
	return &MockUser{
		projectPrompt: projectPrompt,
		counter:       0,
	}
}

func (m *MockUser) ReadConsole() string {
	if m.counter == 0 {
		m.counter++
		return m.projectPrompt
	}

	return "no"
}

func (m *MockUser) AskQuestions(_ []string) string {
	return "Sorry, I don't have any information regarding any of the questions you asked. " +
		"Please proceed as with what you believe to be the best options. "
}

func (m *MockUser) IsUserSatisfied() bool {
	return true
}
