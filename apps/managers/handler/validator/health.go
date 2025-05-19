package validator

import (
	"fmt"
	"net/http"
)

type HealthState struct {
	Active int
}

func (p *ValidatorHandler) CheckHealth() {
	for _, v := range p.validatorList {
		resp, err := http.Get(fmt.Sprintf("%s/health", v.Url))
		if err != nil {
			p.logger.Error("event", "CheckHealth", "err", err)
			v.Active = false
			continue
		}
		resp.Body.Close()

		v.Active = true
	}
}
