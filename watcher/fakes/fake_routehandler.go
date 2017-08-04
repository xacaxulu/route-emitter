// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/route-emitter/routingtable"
	"code.cloudfoundry.org/route-emitter/watcher"
)

type FakeRouteHandler struct {
	HandleEventStub        func(logger lager.Logger, event models.Event)
	handleEventMutex       sync.RWMutex
	handleEventArgsForCall []struct {
		logger lager.Logger
		event  models.Event
	}
	SyncStub        func(logger lager.Logger, desired []*models.LRPDeploymentSchedulingInfo, runningActual []*routingtable.ActualLRPRoutingInfo, domains models.DomainSet, cachedEvents map[string]models.Event)
	syncMutex       sync.RWMutex
	syncArgsForCall []struct {
		logger        lager.Logger
		desired       []*models.LRPDeploymentSchedulingInfo
		runningActual []*routingtable.ActualLRPRoutingInfo
		domains       models.DomainSet
		cachedEvents  map[string]models.Event
	}
	EmitStub        func(logger lager.Logger)
	emitMutex       sync.RWMutex
	emitArgsForCall []struct {
		logger lager.Logger
	}
	ShouldRefreshDesiredStub        func(*routingtable.ActualLRPRoutingInfo) bool
	shouldRefreshDesiredMutex       sync.RWMutex
	shouldRefreshDesiredArgsForCall []struct {
		arg1 *routingtable.ActualLRPRoutingInfo
	}
	shouldRefreshDesiredReturns struct {
		result1 bool
	}
	shouldRefreshDesiredReturnsOnCall map[int]struct {
		result1 bool
	}
	RefreshDesiredStub        func(lager.Logger, []*models.LRPDeploymentSchedulingInfo)
	refreshDesiredMutex       sync.RWMutex
	refreshDesiredArgsForCall []struct {
		arg1 lager.Logger
		arg2 []*models.LRPDeploymentSchedulingInfo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouteHandler) HandleEvent(logger lager.Logger, event models.Event) {
	fake.handleEventMutex.Lock()
	fake.handleEventArgsForCall = append(fake.handleEventArgsForCall, struct {
		logger lager.Logger
		event  models.Event
	}{logger, event})
	fake.recordInvocation("HandleEvent", []interface{}{logger, event})
	fake.handleEventMutex.Unlock()
	if fake.HandleEventStub != nil {
		fake.HandleEventStub(logger, event)
	}
}

func (fake *FakeRouteHandler) HandleEventCallCount() int {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return len(fake.handleEventArgsForCall)
}

func (fake *FakeRouteHandler) HandleEventArgsForCall(i int) (lager.Logger, models.Event) {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return fake.handleEventArgsForCall[i].logger, fake.handleEventArgsForCall[i].event
}

func (fake *FakeRouteHandler) Sync(logger lager.Logger, desired []*models.LRPDeploymentSchedulingInfo, runningActual []*routingtable.ActualLRPRoutingInfo, domains models.DomainSet, cachedEvents map[string]models.Event) {
	var desiredCopy []*models.LRPDeploymentSchedulingInfo
	if desired != nil {
		desiredCopy = make([]*models.LRPDeploymentSchedulingInfo, len(desired))
		copy(desiredCopy, desired)
	}
	var runningActualCopy []*routingtable.ActualLRPRoutingInfo
	if runningActual != nil {
		runningActualCopy = make([]*routingtable.ActualLRPRoutingInfo, len(runningActual))
		copy(runningActualCopy, runningActual)
	}
	fake.syncMutex.Lock()
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
		logger        lager.Logger
		desired       []*models.LRPDeploymentSchedulingInfo
		runningActual []*routingtable.ActualLRPRoutingInfo
		domains       models.DomainSet
		cachedEvents  map[string]models.Event
	}{logger, desiredCopy, runningActualCopy, domains, cachedEvents})
	fake.recordInvocation("Sync", []interface{}{logger, desiredCopy, runningActualCopy, domains, cachedEvents})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		fake.SyncStub(logger, desired, runningActual, domains, cachedEvents)
	}
}

func (fake *FakeRouteHandler) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakeRouteHandler) SyncArgsForCall(i int) (lager.Logger, []*models.LRPDeploymentSchedulingInfo, []*routingtable.ActualLRPRoutingInfo, models.DomainSet, map[string]models.Event) {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return fake.syncArgsForCall[i].logger, fake.syncArgsForCall[i].desired, fake.syncArgsForCall[i].runningActual, fake.syncArgsForCall[i].domains, fake.syncArgsForCall[i].cachedEvents
}

func (fake *FakeRouteHandler) Emit(logger lager.Logger) {
	fake.emitMutex.Lock()
	fake.emitArgsForCall = append(fake.emitArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("Emit", []interface{}{logger})
	fake.emitMutex.Unlock()
	if fake.EmitStub != nil {
		fake.EmitStub(logger)
	}
}

func (fake *FakeRouteHandler) EmitCallCount() int {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return len(fake.emitArgsForCall)
}

func (fake *FakeRouteHandler) EmitArgsForCall(i int) lager.Logger {
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	return fake.emitArgsForCall[i].logger
}

func (fake *FakeRouteHandler) ShouldRefreshDesired(arg1 *routingtable.ActualLRPRoutingInfo) bool {
	fake.shouldRefreshDesiredMutex.Lock()
	ret, specificReturn := fake.shouldRefreshDesiredReturnsOnCall[len(fake.shouldRefreshDesiredArgsForCall)]
	fake.shouldRefreshDesiredArgsForCall = append(fake.shouldRefreshDesiredArgsForCall, struct {
		arg1 *routingtable.ActualLRPRoutingInfo
	}{arg1})
	fake.recordInvocation("ShouldRefreshDesired", []interface{}{arg1})
	fake.shouldRefreshDesiredMutex.Unlock()
	if fake.ShouldRefreshDesiredStub != nil {
		return fake.ShouldRefreshDesiredStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shouldRefreshDesiredReturns.result1
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredCallCount() int {
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	return len(fake.shouldRefreshDesiredArgsForCall)
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredArgsForCall(i int) *routingtable.ActualLRPRoutingInfo {
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	return fake.shouldRefreshDesiredArgsForCall[i].arg1
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredReturns(result1 bool) {
	fake.ShouldRefreshDesiredStub = nil
	fake.shouldRefreshDesiredReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredReturnsOnCall(i int, result1 bool) {
	fake.ShouldRefreshDesiredStub = nil
	if fake.shouldRefreshDesiredReturnsOnCall == nil {
		fake.shouldRefreshDesiredReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldRefreshDesiredReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRouteHandler) RefreshDesired(arg1 lager.Logger, arg2 []*models.LRPDeploymentSchedulingInfo) {
	var arg2Copy []*models.LRPDeploymentSchedulingInfo
	if arg2 != nil {
		arg2Copy = make([]*models.LRPDeploymentSchedulingInfo, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.refreshDesiredMutex.Lock()
	fake.refreshDesiredArgsForCall = append(fake.refreshDesiredArgsForCall, struct {
		arg1 lager.Logger
		arg2 []*models.LRPDeploymentSchedulingInfo
	}{arg1, arg2Copy})
	fake.recordInvocation("RefreshDesired", []interface{}{arg1, arg2Copy})
	fake.refreshDesiredMutex.Unlock()
	if fake.RefreshDesiredStub != nil {
		fake.RefreshDesiredStub(arg1, arg2)
	}
}

func (fake *FakeRouteHandler) RefreshDesiredCallCount() int {
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	return len(fake.refreshDesiredArgsForCall)
}

func (fake *FakeRouteHandler) RefreshDesiredArgsForCall(i int) (lager.Logger, []*models.LRPDeploymentSchedulingInfo) {
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	return fake.refreshDesiredArgsForCall[i].arg1, fake.refreshDesiredArgsForCall[i].arg2
}

func (fake *FakeRouteHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	fake.emitMutex.RLock()
	defer fake.emitMutex.RUnlock()
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouteHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ watcher.RouteHandler = new(FakeRouteHandler)
