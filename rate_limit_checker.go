package main

type Rate struct {
	Burst          int
	Sustained      int
	RemainingToken int
}

func (r *Rate) Consume() {
	r.RemainingToken--
}

func (r *Rate) Replenish() {
	if r.RemainingToken < r.Burst {
		r.RemainingToken++
	}
}

type RateChecker struct {
	rates map[string]*Rate
}

func (r *RateChecker) Consume(endpoint string) bool {
	rate := r.rates[endpoint]
	if r.CheckRate(endpoint) {
		rate.Consume()
		return true
	} else {
		return false
	}
}

func (r RateChecker) CheckRate(endpoint string) bool {
	rate := r.rates[endpoint]
	if rate.RemainingToken == 0 {
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

func (r RateChecker) GetRemainingToken(endpoint string) int {
	return r.rates[endpoint].RemainingToken
}

func initializeRateChecker(config *[]Route, rateChecker *RateChecker) {
	for _, route := range *config {
		r := &Rate{
			Burst:          route.Burst,
			Sustained:      route.Sustained,
			RemainingToken: route.Burst,
		}
		(*rateChecker).rates[route.Endpoint] = r
		go initializeTokenRefiller(route.Sustained, r)
	}
}
