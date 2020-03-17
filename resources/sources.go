package resources

import (
	"time"

	"github.com/tipej/CoronaApp-Api/models"
)

var NewsSources = map[string]models.NewsSource{
	"NYPost": models.NewsSource{
		RssUrl:     "https://nypost.com/coronavirus/feed/",
		DateFormat: time.RFC1123Z,
	},
	"NYTimes": models.NewsSource{
		RssUrl:     "https://rss.nytimes.com/services/xml/rss/nyt/Health.xml",
		Keyword:    "Coronavirus (2019-nCoV)",
		DateFormat: time.RFC1123Z,
	},
	"Politico (EU)": models.NewsSource{
		RssUrl:     "https://www.politico.eu/tag/coronavirus/feed/",
		DateFormat: time.RFC1123Z,
	},
	"BBC (UK)": models.NewsSource{
		RssUrl:     "http://feeds.bbci.co.uk/news/health/rss.xml",
		Keyword:    "Coronavirus",
		DateFormat: time.RFC1123,
	},
	"Independent (UK)": models.NewsSource{
		RssUrl:     "http://www.independent.co.uk/news/health/rss",
		Keyword:    "Coronavirus",
		DateFormat: time.RFC1123,
	},
	"Sky News": models.NewsSource{
		RssUrl:     "http://feeds.skynews.com/feeds/rss/world.xml",
		Keyword:    "Coronavirus",
		DateFormat: time.RFC1123Z,
	},
}
