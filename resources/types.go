package resources

type (
	// ResourcePredicateData is a container used to query predicate information
	// codebeat:disable[TOO_MANY_IVARS]
	ResourcePredicateData struct {
		Schema []struct {
			Name      string `json:"predicate"`
			Type      string
			Index     bool
			Reverse   bool
			Tokenizer []string
			List      bool
			Count     bool
			Upsert    bool
			Lang      bool
		}
	}
	// codebeat:enable[TOO_MANY_IVARS]

	// ResourceTypeData is a container returned from Dgraph about a type
	ResourceTypeData struct {
		Types []struct {
			Name   string
			Fields []struct {
				Name string
			}
		}
	}
)
