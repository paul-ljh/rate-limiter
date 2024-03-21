package main

type Rate struct {
	Burst     int
	Sustained int
	Count     int
}

func (r *Rate) Increment() {
	r.Count++
}

type RateChecker struct {
	rates map[string]*Rate
}

func (r *RateChecker) Increment(endpoint string) bool {
	rate := r.rates[endpoint]
	if r.CheckRate(endpoint) {
		rate.Increment()
		return true
	} else {
		return false
	}
}

func (r RateChecker) CheckRate(endpoint string) bool {
	rate := r.rates[endpoint]
	if rate.Count == rate.Burst {
		return false
	} else {
		return true
	}
}

func (r RateChecker) IsEndpointValid(endpoint string) bool {
	_, ok := r.rates[endpoint]
	return ok
}

func (r RateChecker) GetBurst(endpoint string) int {
	return r.rates[endpoint].Burst
}

func (r RateChecker) GetSustained(endpoint string) int {
	return r.rates[endpoint].Sustained
}

func (r RateChecker) GetCount(endpoint string) int {
	return r.rates[endpoint].Count
}

func initializeRateChecker(config *[]Route, rateChecker *RateChecker) {
	for _, route := range *config {
		(*rateChecker).rates[route.Endpoint] = &Rate{
			Burst:     route.Burst,
			Sustained: route.Sustained,
			Count:     0,
		}
	}
}
