package helper

import "github.com/gtxiqbal/sac24/model/web"

func SliceUniqueNmsRequest(nmsSlice []web.NmsRequest) []web.NmsRequest {
	keys := make(map[web.NmsRequest]bool)
	list := []web.NmsRequest{}
	for _, entry := range nmsSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveIndexSliceString(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
