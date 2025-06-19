package jobs

type Job struct {
	Title    string
	Location string
	Company  string
	Link     string
	Time     string
}

var Blacklist = map[string]struct{}{
	"jobs via dice": {}, "lensa": {}, "jobot": {},
}

var TopTechCompanies = map[string]struct{}{
	"google": {}, "meta": {}, "apple": {}, "microsoft": {}, "amazon": {},
	"nvidia": {}, "openai": {}, "tesla": {}, "palantir": {}, "stripe": {},
	"databricks": {}, "snowflake": {}, "linkedin": {}, "samsung": {}, "tiktok": {},
	"bytedance": {}, "netflix": {}, "adobe": {}, "intel": {}, "amd": {},
	"oracle": {}, "salesforce": {}, "airbnb": {}, "uber": {}, "lyft": {},
	"dropbox": {}, "snap": {}, "pinterest": {}, "doordash": {}, "robinhood": {},
	"coinbase": {}, "square": {}, "block": {}, "zendesk": {}, "asana": {},
	"twilio": {}, "github": {}, "digitalocean": {}, "shopify": {}, "spotify": {},
	"qualcomm": {}, "huawei": {}, "tencent": {}, "baidu": {}, "alibaba": {},
	"ibm": {}, "dell": {}, "hp": {}, "cisco": {}, "red hat": {},
	"cloudflare": {}, "figma": {}, "notion": {}, "monday.com": {}, "atlassian": {},
	"zapier": {}, "intercom": {}, "splunk": {}, "elastic": {}, "fastly": {},
	"unity": {}, "epic games": {}, "riot games": {}, "blizzard": {}, "valve": {},
	"rovi": {}, "arm": {}, "marvell": {}, "synopsys": {},
	"keysight": {}, "ni": {}, "luminar": {}, "waymo": {}, "cruise": {},
	"zoox": {}, "niantic": {}, "replit": {}, "hugging face": {}, "scale ai": {},
	"anthropic": {}, "runway": {}, "mistral ai": {}, "perplexity": {}, "character.ai": {},
	"stability ai": {}, "skydio": {}, "anduril": {}, "spacex": {}, "blue origin": {},
	"rocket lab": {}, "relativity space": {}, "calm": {}, "headspace": {}, "duolingo": {},
	"coursera": {}, "khan academy": {}, "chegg": {}, "udemy": {}, "edx": {},
	"naver": {}, "line": {}, "kakao": {}, "grab": {}, "gojek": {},
	"booking.com": {}, "expedia": {}, "yelp": {}, "tripadvisor": {}, "zillow": {},
	"glassdoor": {}, "indeed": {}, "monster": {}, "okta": {}, "auth0": {},
	"dashlane": {}, "1password": {}, "lastpass": {}, "bitwarden": {}, "samsara": {},
	"cloudera": {}, "confluent": {}, "mongodb": {}, "couchbase": {}, "datastax": {},
	"c3.ai": {}, "verkada": {}, "rippling": {}, "gusto": {}, "brex": {},
	"plaid": {}, "robin": {}, "nuro": {}, "aurora": {}, "embark": {},
	"argo ai": {}, "argo": {}, "dataminr": {}, "pagerduty": {},
	"sendbird": {}, "segment": {}, "postman": {}, "new relic": {}, "sentry": {},
	"bugsnag": {}, "launchdarkly": {}, "circleci": {}, "travis ci": {}, "vercel": {},
	"netlify": {}, "heroku": {}, "render": {}, "fly.io": {}, "supabase": {},
	"firebase": {}, "backblaze": {}, "wasabi": {}, "digital ocean": {}, "linode": {},
	"vultr": {}, "alation": {}, "collibra": {}, "domo": {}, "tableau": {},
	"looker": {}, "mode": {}, "hex": {}, "preset": {}, "thoughtspot": {},
	"airbyte": {}, "fivetran": {}, "census": {}, "rudderstack": {}, "apache": {},
	"jetbrains": {}, "intellij": {}, "pycharm": {}, "goland": {}, "eclipse": {},
	"vs code": {}, "visual studio": {}, "docker": {}, "kubernetes": {}, "terraform": {},
	"ansible": {}, "chef": {}, "puppet": {}, "hashicorp": {}, "redpanda": {},
}

type GlassdoorResponse []struct {
	Data struct {
		JobListings struct {
			Typename             string `json:"__typename,omitempty"`
			CompanyFilterOptions []struct {
				Typename  string `json:"__typename,omitempty"`
				ID        int    `json:"id,omitempty"`
				ShortName string `json:"shortName,omitempty"`
			} `json:"companyFilterOptions,omitempty"`
			FilterOptions       string `json:"filterOptions,omitempty"`
			IndeedCtk           string `json:"indeedCtk,omitempty"`
			IndexablePageForSeo bool   `json:"indexablePageForSeo,omitempty"`
			JobListings         []struct {
				Typename string `json:"__typename,omitempty"`
				Jobview  struct {
					Typename string `json:"__typename,omitempty"`
					Header   struct {
						Typename  string `json:"__typename,omitempty"`
						AdOrderID int    `json:"adOrderId,omitempty"`
						AgeInDays int    `json:"ageInDays,omitempty"`
						EasyApply bool   `json:"easyApply,omitempty"`
						Employer  struct {
							Typename  string `json:"__typename,omitempty"`
							ID        int    `json:"id,omitempty"`
							Name      string `json:"name,omitempty"`
							ShortName string `json:"shortName,omitempty"`
						} `json:"employer,omitempty"`
						EmployerNameFromSearch string `json:"employerNameFromSearch,omitempty"`
						Expired                bool   `json:"expired,omitempty"`
						Goc                    string `json:"goc,omitempty"`
						GocID                  int    `json:"gocId,omitempty"`
						IndeedJobAttribute     any    `json:"indeedJobAttribute,omitempty"`
						IsSponsoredEmployer    bool   `json:"isSponsoredEmployer,omitempty"`
						IsSponsoredJob         bool   `json:"isSponsoredJob,omitempty"`
						JobCountryID           int    `json:"jobCountryId,omitempty"`
						JobLink                string `json:"jobLink,omitempty"`
						JobResultTrackingKey   string `json:"jobResultTrackingKey,omitempty"`
						JobTitleText           string `json:"jobTitleText,omitempty"`
						LocID                  int    `json:"locId,omitempty"`
						LocationName           string `json:"locationName,omitempty"`
						LocationType           string `json:"locationType,omitempty"`
						NormalizedJobTitle     string `json:"normalizedJobTitle,omitempty"`
						Occupations            []struct {
							Typename string `json:"__typename,omitempty"`
							Key      string `json:"key,omitempty"`
						} `json:"occupations,omitempty"`
						PayCurrency          string `json:"payCurrency,omitempty"`
						PayPeriod            string `json:"payPeriod,omitempty"`
						PayPeriodAdjustedPay struct {
							Typename string  `json:"__typename,omitempty"`
							P10      float64 `json:"p10,omitempty"`
							P50      float64 `json:"p50,omitempty"`
							P90      float64 `json:"p90,omitempty"`
						} `json:"payPeriodAdjustedPay,omitempty"`
						Rating       float64 `json:"rating,omitempty"`
						SalarySource string  `json:"salarySource,omitempty"`
						SavedJobID   any     `json:"savedJobId,omitempty"`
						SeoJobLink   string  `json:"seoJobLink,omitempty"`
					} `json:"header,omitempty"`
					Job struct {
						Typename                 string `json:"__typename,omitempty"`
						DescriptionFragmentsText any    `json:"descriptionFragmentsText,omitempty"`
						ImportConfigID           int    `json:"importConfigId,omitempty"`
						JobTitleID               int    `json:"jobTitleId,omitempty"`
						JobTitleText             string `json:"jobTitleText,omitempty"`
						ListingID                int64  `json:"listingId,omitempty"`
					} `json:"job,omitempty"`
					JobListingAdminDetails struct {
						Typename                       string `json:"__typename,omitempty"`
						UserEligibleForAdminJobDetails bool   `json:"userEligibleForAdminJobDetails,omitempty"`
					} `json:"jobListingAdminDetails,omitempty"`
					Overview struct {
						Typename      string `json:"__typename,omitempty"`
						ShortName     string `json:"shortName,omitempty"`
						SquareLogoURL string `json:"squareLogoUrl,omitempty"`
					} `json:"overview,omitempty"`
				} `json:"jobview,omitempty"`
			} `json:"jobListings,omitempty"`
			JobSearchTrackingKey string `json:"jobSearchTrackingKey,omitempty"`
			JobsPageSeoData      struct {
				Typename            string `json:"__typename,omitempty"`
				PageMetaDescription string `json:"pageMetaDescription,omitempty"`
				PageTitle           string `json:"pageTitle,omitempty"`
			} `json:"jobsPageSeoData,omitempty"`
			PaginationCursors []struct {
				Typename   string `json:"__typename,omitempty"`
				Cursor     string `json:"cursor,omitempty"`
				PageNumber int    `json:"pageNumber,omitempty"`
			} `json:"paginationCursors,omitempty"`
			SearchResultsMetadata struct {
				Typename string `json:"__typename,omitempty"`
				FooterVO struct {
					Typename    string `json:"__typename,omitempty"`
					CountryMenu struct {
						Typename             string `json:"__typename,omitempty"`
						ChildNavigationLinks []struct {
							Typename string `json:"__typename,omitempty"`
							ID       string `json:"id,omitempty"`
							Link     string `json:"link,omitempty"`
							TextKey  string `json:"textKey,omitempty"`
						} `json:"childNavigationLinks,omitempty"`
					} `json:"countryMenu,omitempty"`
				} `json:"footerVO,omitempty"`
				HelpCenterDomain string `json:"helpCenterDomain,omitempty"`
				HelpCenterLocale string `json:"helpCenterLocale,omitempty"`
				JobAlert         struct {
					Typename   string `json:"__typename,omitempty"`
					JobAlertID any    `json:"jobAlertId,omitempty"`
				} `json:"jobAlert,omitempty"`
				JobSerpFaq struct {
					Typename  string `json:"__typename,omitempty"`
					Questions []struct {
						Typename string `json:"__typename,omitempty"`
						Answer   string `json:"answer,omitempty"`
						Question string `json:"question,omitempty"`
					} `json:"questions,omitempty"`
				} `json:"jobSerpFaq,omitempty"`
				JobSerpJobOutlook struct {
					Typename   string `json:"__typename,omitempty"`
					Heading    string `json:"heading,omitempty"`
					Occupation string `json:"occupation,omitempty"`
					Paragraph  string `json:"paragraph,omitempty"`
				} `json:"jobSerpJobOutlook,omitempty"`
				SearchCriteria struct {
					Typename         string `json:"__typename,omitempty"`
					ImplicitLocation struct {
						Typename             string `json:"__typename,omitempty"`
						ID                   int    `json:"id,omitempty"`
						LocalizedDisplayName string `json:"localizedDisplayName,omitempty"`
						Type                 string `json:"type,omitempty"`
					} `json:"implicitLocation,omitempty"`
					Keyword  string `json:"keyword,omitempty"`
					Location struct {
						Typename             string `json:"__typename,omitempty"`
						ID                   int    `json:"id,omitempty"`
						LocalizedDisplayName string `json:"localizedDisplayName,omitempty"`
						LocalizedShortName   string `json:"localizedShortName,omitempty"`
						ShortName            string `json:"shortName,omitempty"`
						Type                 string `json:"type,omitempty"`
					} `json:"location,omitempty"`
				} `json:"searchCriteria,omitempty"`
				ShowMachineReadableJobs bool `json:"showMachineReadableJobs,omitempty"`
			} `json:"searchResultsMetadata,omitempty"`
			SerpSeoLinksVO struct {
				Typename                   string   `json:"__typename,omitempty"`
				RelatedJobTitlesResults    []string `json:"relatedJobTitlesResults,omitempty"`
				SearchedJobTitle           string   `json:"searchedJobTitle,omitempty"`
				SearchedKeyword            any      `json:"searchedKeyword,omitempty"`
				SearchedLocationIDAsString string   `json:"searchedLocationIdAsString,omitempty"`
				SearchedLocationSeoName    string   `json:"searchedLocationSeoName,omitempty"`
				SearchedLocationType       string   `json:"searchedLocationType,omitempty"`
				TopCityIdsToNameResults    []struct {
					Typename string `json:"__typename,omitempty"`
					Key      string `json:"key,omitempty"`
					Value    string `json:"value,omitempty"`
				} `json:"topCityIdsToNameResults,omitempty"`
				TopEmployerIdsToNameResults []struct {
					Typename string `json:"__typename,omitempty"`
					Key      string `json:"key,omitempty"`
					Value    string `json:"value,omitempty"`
				} `json:"topEmployerIdsToNameResults,omitempty"`
				TopOccupationResults []string `json:"topOccupationResults,omitempty"`
			} `json:"serpSeoLinksVO,omitempty"`
			TotalJobsCount int `json:"totalJobsCount,omitempty"`
		} `json:"jobListings,omitempty"`
	} `json:"data,omitempty"`
}
