package connect

import "j322.ica/gumroad-sammi/gumroad"

func gumroadToVar(sale gumroad.Sale) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range sale {
		switch len(v) {
		case 0:
		case 1:
			res[k] = v[0]
		default:
			res[k] = v
		}
	}
	return res
}
