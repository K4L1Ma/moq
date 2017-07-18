// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package dotimport_test

import (
	"dotimport"
	"sync"
)

var (
	lockServiceMockUser sync.RWMutex
)

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             UserFunc: func(ID string) (dotimport.User, error) {
// 	               panic("TODO: mock out the User method")
//             },
//         }
//
//         // TODO: use mockedService in code that requires Service
//         //       and then make assertions.
//
//     }
type ServiceMock struct {
	// UserFunc mocks the User method.
	UserFunc func(ID string) (dotimport.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// User holds details about calls to the User method.
		User []struct {
			// ID is the ID argument value.
			ID string
		}
	}
}

// User calls UserFunc.
func (mock *ServiceMock) User(ID string) (dotimport.User, error) {
	if mock.UserFunc == nil {
		panic("moq: ServiceMock.UserFunc is nil but Service.User was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	lockServiceMockUser.Lock()
	mock.calls.User = append(mock.calls.User, callInfo)
	lockServiceMockUser.Unlock()
	return mock.UserFunc(ID)
}

// UserCalls gets all the calls that were made to User.
// Check the length with:
//     len(mockedService.UserCalls())
func (mock *ServiceMock) UserCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockServiceMockUser.RLock()
	calls = mock.calls.User
	lockServiceMockUser.RUnlock()
	return calls
}