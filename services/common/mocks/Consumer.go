// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	amqp "github.com/streadway/amqp"
	mock "github.com/stretchr/testify/mock"

	msgbus "github.com/ukama/ukama/services/common/msgbus"

	testing "testing"
)

// Consumer is an autogenerated mock type for the Consumer type
type Consumer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Consumer) Close() {
	_m.Called()
}

// IsClosed provides a mock function with given fields:
func (_m *Consumer) IsClosed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Subscribe provides a mock function with given fields: queueName, exchangeName, exchangeType, routingKeys, consumerName, handlerFunc
func (_m *Consumer) Subscribe(queueName string, exchangeName string, exchangeType string, routingKeys []msgbus.RoutingKey, consumerName string, handlerFunc func(amqp.Delivery, chan<- bool)) error {
	ret := _m.Called(queueName, exchangeName, exchangeType, routingKeys, consumerName, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []msgbus.RoutingKey, string, func(amqp.Delivery, chan<- bool)) error); ok {
		r0 = rf(queueName, exchangeName, exchangeType, routingKeys, consumerName, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeToQueue provides a mock function with given fields: queueName, consumerName, handlerFunc
func (_m *Consumer) SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery, chan<- bool)) error {
	ret := _m.Called(queueName, consumerName, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, func(amqp.Delivery, chan<- bool)) error); ok {
		r0 = rf(queueName, consumerName, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeToServiceQueue provides a mock function with given fields: serviceName, exchangeName, routingKeys, consumerId, handlerFunc
func (_m *Consumer) SubscribeToServiceQueue(serviceName string, exchangeName string, routingKeys []msgbus.RoutingKey, consumerId string, handlerFunc func(amqp.Delivery, chan<- bool)) error {
	ret := _m.Called(serviceName, exchangeName, routingKeys, consumerId, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []msgbus.RoutingKey, string, func(amqp.Delivery, chan<- bool)) error); ok {
		r0 = rf(serviceName, exchangeName, routingKeys, consumerId, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeWithArgs provides a mock function with given fields: queueName, exchangeName, exchangeType, routingKeys, consumerName, queueArgs, handlerFunc
func (_m *Consumer) SubscribeWithArgs(queueName string, exchangeName string, exchangeType string, routingKeys []msgbus.RoutingKey, consumerName string, queueArgs map[string]interface{}, handlerFunc func(amqp.Delivery, chan<- bool)) error {
	ret := _m.Called(queueName, exchangeName, exchangeType, routingKeys, consumerName, queueArgs, handlerFunc)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []msgbus.RoutingKey, string, map[string]interface{}, func(amqp.Delivery, chan<- bool)) error); ok {
		r0 = rf(queueName, exchangeName, exchangeType, routingKeys, consumerName, queueArgs, handlerFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewConsumer creates a new instance of Consumer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumer(t testing.TB) *Consumer {
	mock := &Consumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}