package dockerclient

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func NewMockClient() *MockClient {
	return &MockClient{}
}

func (client *MockClient) Info() (*Info, error) {
	args := client.Mock.Called()
	return args.Get(0).(*Info), args.Error(1)
}

func (client *MockClient) ListContainers(all bool) ([]Container, error) {
	args := client.Mock.Called(all)
	return args.Get(0).([]Container), args.Error(1)
}

func (client *MockClient) InspectContainer(id string) (*ContainerInfo, error) {
	args := client.Mock.Called(id)
	return args.Get(0).(*ContainerInfo), args.Error(1)
}

func (client *MockClient) CreateContainer(config *ContainerConfig, name string) (string, error) {
	args := client.Mock.Called(config, name)
	return args.String(0), args.Error(1)
}

func (client *MockClient) ContainerLogs(id string, stdout bool, stderr bool) (io.ReadCloser, error) {
	args := client.Mock.Called(id, stdout, stderr)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (client *MockClient) StartContainer(id string, config *HostConfig) error {
	args := client.Mock.Called(id, config)
	return args.Error(0)
}

func (client *MockClient) StopContainer(id string, timeout int) error {
	args := client.Mock.Called(id, timeout)
	return args.Error(0)
}

func (client *MockClient) RestartContainer(id string, timeout int) error {
	args := client.Mock.Called(id, timeout)
	return args.Error(0)
}

func (client *MockClient) KillContainer(id string) error {
	args := client.Mock.Called(id)
	return args.Error(0)
}

func (client *MockClient) StartMonitorEvents(cb Callback, args ...interface{}) {
	client.Mock.Called(cb, args)
}

func (client *MockClient) StopAllMonitorEvents() {
	client.Mock.Called()
}

func (client *MockClient) Version() (*Version, error) {
	args := client.Mock.Called()
	return args.Get(0).(*Version), args.Error(1)
}

func (client *MockClient) PullImage(name, tag string) error {
	args := client.Mock.Called(name, tag)
	return args.Error(0)
}

func (client *MockClient) RemoveContainer(id string, force bool) error {
	args := client.Mock.Called(id, force)
	return args.Error(0)
}

func (client *MockClient) ListImages() ([]*Image, error) {
	args := client.Mock.Called()
	return args.Get(0).([]*Image), args.Error(1)
}

func (client *MockClient) RemoveImage(name string) error {
	args := client.Mock.Called(name)
	return args.Error(0)
}
