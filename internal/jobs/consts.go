package jobs

type Job struct {
	Title    string
	Location string
	Company  string
	Link     string
	Time     string
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
