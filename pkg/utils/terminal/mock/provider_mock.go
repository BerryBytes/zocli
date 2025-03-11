package mock

type MockCredentialProvider struct {
}

type MockProvider interface {
	ReadInput() (string, error)
	ReadPassword() (string, error)
	Clear()
}

func NewMockProvider() MockProvider {
	return &MockCredentialProvider{}
}

func (m *MockCredentialProvider) ReadInput() (string, error) {
	return "mock.email@gmail.com", nil
}

func (m *MockCredentialProvider) ReadPassword() (string, error) {
	return "mockPassword", nil
}
func (m *MockCredentialProvider) Clear() {
}
