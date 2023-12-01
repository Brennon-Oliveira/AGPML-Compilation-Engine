package header_configs

var InlineCss bool = false
var Title string = ""

var acceptedHeaderConfigs = []string{
	"inlineCss",
	"title",
}

func IsAcceptedHeaderConfig(attribute string) bool {
	for _, acceptedHeaderConfig := range acceptedHeaderConfigs {
		if attribute == acceptedHeaderConfig {
			return true
		}
	}
	return false
}
