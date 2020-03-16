package resources

import (
	"github.com/tipej/CoronaApp-Api/models"
)

var NewsSources = map[string]models.NewsSource{
	"NYPost": models.NewsSource{
		RssUrl: "https://nypost.com/coronavirus/feed/",
	},
	"NYTimes": models.NewsSource{
		RssUrl:  "https://rss.nytimes.com/services/xml/rss/nyt/Health.xml",
		Keyword: "Coronavirus (2019-nCoV)",
	},
	"Politico (EU)": models.NewsSource{
		RssUrl: "https://www.politico.eu/tag/coronavirus/feed/",
	},
}
