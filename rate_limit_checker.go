package main

type Rate struct {
	Burst     int
	Sustained int
	Count     int
}

type RateChecker struct {
	rates map[string]Rate
}

func (r *RateChecker) Increment(endpoint string) bool {
	rate := r.rates[endpoint]
	if r.CheckRate(endpoint) {
		rate.Count++
		return true
	} else {
		return false
	}
}

func (r RateChecker) CheckRate(endpoint string) bool {
	rate := r.rates[endpoint]
	if rate.Count+1 > rate.Burst {
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

func initializeRateChecker(config *[]Route, rateChecker *RateChecker) {
	for _, route := range *config {
		(*rateChecker).rates[route.Endpoint] = Rate{
			Burst:     route.Burst,
			Sustained: route.Sustained,
			Count:     0,
		}
	}
}
