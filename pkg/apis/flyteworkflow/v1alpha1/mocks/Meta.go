// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// Meta is an autogenerated mock type for the Meta type
type Meta struct {
	mock.Mock
}

type Meta_GetAnnotations struct {
	*mock.Call
}

func (_m Meta_GetAnnotations) Return(_a0 map[string]string) *Meta_GetAnnotations {
	return &Meta_GetAnnotations{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetAnnotations() *Meta_GetAnnotations {
	c_call := _m.On("GetAnnotations")
	return &Meta_GetAnnotations{Call: c_call}
}

func (_m *Meta) OnGetAnnotationsMatch(matchers ...interface{}) *Meta_GetAnnotations {
	c_call := _m.On("GetAnnotations", matchers...)
	return &Meta_GetAnnotations{Call: c_call}
}

// GetAnnotations provides a mock function with given fields:
func (_m *Meta) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type Meta_GetCreationTimestamp struct {
	*mock.Call
}

func (_m Meta_GetCreationTimestamp) Return(_a0 v1.Time) *Meta_GetCreationTimestamp {
	return &Meta_GetCreationTimestamp{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetCreationTimestamp() *Meta_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp")
	return &Meta_GetCreationTimestamp{Call: c_call}
}

func (_m *Meta) OnGetCreationTimestampMatch(matchers ...interface{}) *Meta_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp", matchers...)
	return &Meta_GetCreationTimestamp{Call: c_call}
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *Meta) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

type Meta_GetEventVersion struct {
	*mock.Call
}

func (_m Meta_GetEventVersion) Return(_a0 v1alpha1.EventVersion) *Meta_GetEventVersion {
	return &Meta_GetEventVersion{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetEventVersion() *Meta_GetEventVersion {
	c_call := _m.On("GetEventVersion")
	return &Meta_GetEventVersion{Call: c_call}
}

func (_m *Meta) OnGetEventVersionMatch(matchers ...interface{}) *Meta_GetEventVersion {
	c_call := _m.On("GetEventVersion", matchers...)
	return &Meta_GetEventVersion{Call: c_call}
}

// GetEventVersion provides a mock function with given fields:
func (_m *Meta) GetEventVersion() v1alpha1.EventVersion {
	ret := _m.Called()

	var r0 v1alpha1.EventVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.EventVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.EventVersion)
	}

	return r0
}

type Meta_GetExecutionID struct {
	*mock.Call
}

func (_m Meta_GetExecutionID) Return(_a0 v1alpha1.WorkflowExecutionIdentifier) *Meta_GetExecutionID {
	return &Meta_GetExecutionID{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetExecutionID() *Meta_GetExecutionID {
	c_call := _m.On("GetExecutionID")
	return &Meta_GetExecutionID{Call: c_call}
}

func (_m *Meta) OnGetExecutionIDMatch(matchers ...interface{}) *Meta_GetExecutionID {
	c_call := _m.On("GetExecutionID", matchers...)
	return &Meta_GetExecutionID{Call: c_call}
}

// GetExecutionID provides a mock function with given fields:
func (_m *Meta) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowExecutionIdentifier
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowExecutionIdentifier)
	}

	return r0
}

type Meta_GetK8sWorkflowID struct {
	*mock.Call
}

func (_m Meta_GetK8sWorkflowID) Return(_a0 types.NamespacedName) *Meta_GetK8sWorkflowID {
	return &Meta_GetK8sWorkflowID{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetK8sWorkflowID() *Meta_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID")
	return &Meta_GetK8sWorkflowID{Call: c_call}
}

func (_m *Meta) OnGetK8sWorkflowIDMatch(matchers ...interface{}) *Meta_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID", matchers...)
	return &Meta_GetK8sWorkflowID{Call: c_call}
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *Meta) GetK8sWorkflowID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

type Meta_GetLabels struct {
	*mock.Call
}

func (_m Meta_GetLabels) Return(_a0 map[string]string) *Meta_GetLabels {
	return &Meta_GetLabels{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetLabels() *Meta_GetLabels {
	c_call := _m.On("GetLabels")
	return &Meta_GetLabels{Call: c_call}
}

func (_m *Meta) OnGetLabelsMatch(matchers ...interface{}) *Meta_GetLabels {
	c_call := _m.On("GetLabels", matchers...)
	return &Meta_GetLabels{Call: c_call}
}

// GetLabels provides a mock function with given fields:
func (_m *Meta) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type Meta_GetName struct {
	*mock.Call
}

func (_m Meta_GetName) Return(_a0 string) *Meta_GetName {
	return &Meta_GetName{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetName() *Meta_GetName {
	c_call := _m.On("GetName")
	return &Meta_GetName{Call: c_call}
}

func (_m *Meta) OnGetNameMatch(matchers ...interface{}) *Meta_GetName {
	c_call := _m.On("GetName", matchers...)
	return &Meta_GetName{Call: c_call}
}

// GetName provides a mock function with given fields:
func (_m *Meta) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Meta_GetNamespace struct {
	*mock.Call
}

func (_m Meta_GetNamespace) Return(_a0 string) *Meta_GetNamespace {
	return &Meta_GetNamespace{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetNamespace() *Meta_GetNamespace {
	c_call := _m.On("GetNamespace")
	return &Meta_GetNamespace{Call: c_call}
}

func (_m *Meta) OnGetNamespaceMatch(matchers ...interface{}) *Meta_GetNamespace {
	c_call := _m.On("GetNamespace", matchers...)
	return &Meta_GetNamespace{Call: c_call}
}

// GetNamespace provides a mock function with given fields:
func (_m *Meta) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Meta_GetOwnerReference struct {
	*mock.Call
}

func (_m Meta_GetOwnerReference) Return(_a0 v1.OwnerReference) *Meta_GetOwnerReference {
	return &Meta_GetOwnerReference{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetOwnerReference() *Meta_GetOwnerReference {
	c_call := _m.On("GetOwnerReference")
	return &Meta_GetOwnerReference{Call: c_call}
}

func (_m *Meta) OnGetOwnerReferenceMatch(matchers ...interface{}) *Meta_GetOwnerReference {
	c_call := _m.On("GetOwnerReference", matchers...)
	return &Meta_GetOwnerReference{Call: c_call}
}

// GetOwnerReference provides a mock function with given fields:
func (_m *Meta) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

type Meta_GetRawOutputDataConfig struct {
	*mock.Call
}

func (_m Meta_GetRawOutputDataConfig) Return(_a0 v1alpha1.RawOutputDataConfig) *Meta_GetRawOutputDataConfig {
	return &Meta_GetRawOutputDataConfig{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetRawOutputDataConfig() *Meta_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig")
	return &Meta_GetRawOutputDataConfig{Call: c_call}
}

func (_m *Meta) OnGetRawOutputDataConfigMatch(matchers ...interface{}) *Meta_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig", matchers...)
	return &Meta_GetRawOutputDataConfig{Call: c_call}
}

// GetRawOutputDataConfig provides a mock function with given fields:
func (_m *Meta) GetRawOutputDataConfig() v1alpha1.RawOutputDataConfig {
	ret := _m.Called()

	var r0 v1alpha1.RawOutputDataConfig
	if rf, ok := ret.Get(0).(func() v1alpha1.RawOutputDataConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.RawOutputDataConfig)
	}

	return r0
}

type Meta_GetSecurityContext struct {
	*mock.Call
}

func (_m Meta_GetSecurityContext) Return(_a0 core.SecurityContext) *Meta_GetSecurityContext {
	return &Meta_GetSecurityContext{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetSecurityContext() *Meta_GetSecurityContext {
	c_call := _m.On("GetSecurityContext")
	return &Meta_GetSecurityContext{Call: c_call}
}

func (_m *Meta) OnGetSecurityContextMatch(matchers ...interface{}) *Meta_GetSecurityContext {
	c_call := _m.On("GetSecurityContext", matchers...)
	return &Meta_GetSecurityContext{Call: c_call}
}

// GetSecurityContext provides a mock function with given fields:
func (_m *Meta) GetSecurityContext() core.SecurityContext {
	ret := _m.Called()

	var r0 core.SecurityContext
	if rf, ok := ret.Get(0).(func() core.SecurityContext); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.SecurityContext)
	}

	return r0
}

type Meta_GetServiceAccountName struct {
	*mock.Call
}

func (_m Meta_GetServiceAccountName) Return(_a0 string) *Meta_GetServiceAccountName {
	return &Meta_GetServiceAccountName{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnGetServiceAccountName() *Meta_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName")
	return &Meta_GetServiceAccountName{Call: c_call}
}

func (_m *Meta) OnGetServiceAccountNameMatch(matchers ...interface{}) *Meta_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName", matchers...)
	return &Meta_GetServiceAccountName{Call: c_call}
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *Meta) GetServiceAccountName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type Meta_IsInterruptible struct {
	*mock.Call
}

func (_m Meta_IsInterruptible) Return(_a0 bool) *Meta_IsInterruptible {
	return &Meta_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *Meta) OnIsInterruptible() *Meta_IsInterruptible {
	c_call := _m.On("IsInterruptible")
	return &Meta_IsInterruptible{Call: c_call}
}

func (_m *Meta) OnIsInterruptibleMatch(matchers ...interface{}) *Meta_IsInterruptible {
	c_call := _m.On("IsInterruptible", matchers...)
	return &Meta_IsInterruptible{Call: c_call}
}

// IsInterruptible provides a mock function with given fields:
func (_m *Meta) IsInterruptible() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
