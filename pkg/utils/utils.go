package utils

import v1 "k8s.io/api/core/v1"

//remove pod from  pod slice
func RemovePod(p []v1.Pod, i int) []v1.Pod {
	p[len(p)-1], p[i] = p[i], p[len(p)-1]
	return p[:len(p)-1]
}
