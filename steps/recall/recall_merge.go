package recall_group

import "recommend/data"

func RecallMerge(rc *data.RequestContext) {
	for _, groups := range rc.RecallReasons {
		rc.Groups = append(rc.Groups, groups...)
	}
}
