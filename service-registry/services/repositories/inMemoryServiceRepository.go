package repositories

import (
	"2018_2_iu7.corp/common/errors"
	"2018_2_iu7.corp/service-registry/services/models"
	"sync"
	"time"
)

const (
	DefaultExpireTime = 10

	ServiceStatusUP   = "UP"
	ServiceStatusDOWN = "DOWN"
)

type serviceInfo struct {
	models.ServiceInfo
	expireTime time.Duration
}

type InMemoryServiceRepository struct {
	rwMutex    *sync.RWMutex
	services   map[string][]serviceInfo
	expireTime time.Duration
}

func NewInMemoryServiceRepository(expireTime time.Duration) *InMemoryServiceRepository {
	return &InMemoryServiceRepository{
		rwMutex:    new(sync.RWMutex),
		services:   make(map[string][]serviceInfo, 0),
		expireTime: expireTime,
	}
}

func (r *InMemoryServiceRepository) StartMonitor() {
	const monitoringInterval = 5 * time.Second

	go func() {
		for {
			time.Sleep(monitoringInterval)

			r.rwMutex.Lock()
			for name, info := range r.services {
				for i := 0; i < len(info); i++ {
					if info[i].expireTime <= 0 {
						info[i].Status = ServiceStatusDOWN
					} else {
						info[i].expireTime -= monitoringInterval
					}
				}
				r.services[name] = info
			}
			r.rwMutex.Unlock()
		}
	}()
}

func (r *InMemoryServiceRepository) GetAllServicesInfo() ([]models.Service, error) {
	r.rwMutex.RLock()
	defer r.rwMutex.Unlock()

	services := make([]models.Service, 0)
	for name, info := range r.services {
		tmp := make([]models.ServiceInfo, 0)
		for _, v := range info {
			tmp = append(tmp, v.ServiceInfo)
		}

		service := models.Service{
			Name:     name,
			Replicas: tmp,
		}

		services = append(services, service)
	}

	return services, nil
}

func (r *InMemoryServiceRepository) GetServiceInfo(name string) (*models.Service, error) {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	info, ok := r.services[name]
	if !ok {
		return nil, errors.NewNotFoundError("service not found")
	}

	tmp := make([]models.ServiceInfo, 0)
	for _, v := range info {
		tmp = append(tmp, v.ServiceInfo)
	}

	return &models.Service{
		Name:     name,
		Replicas: tmp,
	}, nil
}

func (r *InMemoryServiceRepository) RegisterService(name string, addr string) error {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	replica := serviceInfo{expireTime: r.expireTime}
	replica.Address = addr
	replica.Status = ServiceStatusUP

	info, ok := r.services[name]
	if !ok {
		info = []serviceInfo{replica}
		r.services[name] = info
		return nil
	}

	for _, v := range info {
		if v.Address == addr {
			return errors.NewDuplicateError("service instance already registered")
		}
	}

	info = append(info, replica)
	r.services[name] = info

	return nil
}

func (r *InMemoryServiceRepository) UpdateService(name string, addr string) error {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	info, ok := r.services[name]
	if !ok {
		return errors.NewNotFoundError("service instance not found")
	}

	for i, v := range info {
		if v.Address == addr {
			v.Status = ServiceStatusUP
			v.expireTime = r.expireTime
			info[i] = v
			return nil
		}
	}

	return errors.NewNotFoundError("service instance not found")
}

func (r *InMemoryServiceRepository) UnregisterService(name string, addr string) error {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	info, ok := r.services[name]
	if !ok {
		return errors.NewNotFoundError("service instance not found")
	}

	n := len(info)
	for i := 0; i < n; i++ {
		if info[i].Address == addr {
			if i+1 == n {
				info = info[:i]
			} else {
				info = append(info[:i], info[i+1:]...)
			}
			r.services[name] = info
		}
	}

	return errors.NewNotFoundError("service instance not found")
}
