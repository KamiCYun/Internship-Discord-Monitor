package jobs

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/kamicyun/internship-discord-monitor/internal/utils"
)

func MonitorGlassdoor(delay uint32, keywords string) ([]Job, error) {
	var jobs []Job
	pageNumber := 1
	pageCursor := ""

	csrf, err := getGlassdoorCSRF()
	if err != nil {
		return nil, fmt.Errorf("error fetching Glassdoor CSRF token: %w", err)
	}

	headers := generateHeaders(*csrf)

	for {
		payload := generatePayload(keywords, pageNumber, pageCursor)
		resp, err := utils.CloudflareRequest("POST", "https://www.glassdoor.com/graph", headers, payload)
		if err != nil {
			return nil, fmt.Errorf("cloudflare request failed: %w", err)
		}

		var gdResp GlassdoorResponse
		if err := json.Unmarshal([]byte(resp.Body), &gdResp); err != nil {
			return nil, fmt.Errorf("malformed Glassdoor response: %w", err)
		}

		listings := gdResp[0].Data.JobListings.JobListings

		for _, job := range listings {
			jobs = append(jobs, Job{
				Title:    job.Jobview.Job.JobTitleText,
				Location: job.Jobview.Header.LocationName,
				Company:  job.Jobview.Overview.ShortName,
				Link:     job.Jobview.Header.SeoJobLink,
				Time:     "24 Hrs Ago",
			})
		}

		if len(listings) != 30 {
			log.Printf("Scraped %d total jobs.", len(jobs))
			break
		}

		log.Printf("Scraped %d jobs so far.", len(jobs))
		time.Sleep(time.Duration(delay) * time.Millisecond)
		paginationCursors := gdResp[0].Data.JobListings.PaginationCursors[0]
		pageNumber = paginationCursors.PageNumber
		pageCursor = paginationCursors.Cursor
	}

	return jobs, nil
}

func getGlassdoorCSRF() (*string, error) {
	headers := `{
		"accept-language": "en-US,en;q=0.9",
		"cache-control": "max-age=0",
		"priority": "u=0, i",
		"referer": "https://www.glassdoor.com/",
		"sec-ch-ua": "\"Google Chrome\";v=\"137\", \"Chromium\";v=\"137\", \"Not/A)Brand\";v=\"24\"",
		"sec-ch-ua-mobile": "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest": "document",
		"sec-fetch-mode": "navigate",
		"sec-fetch-site": "same-origin",
		"sec-fetch-user": "?1",
		"upgrade-insecure-requests": "1"
	}`

	resp, err := utils.CloudflareRequest("GET", "https://www.glassdoor.com/Job/united-states-software-engineer-jobs.htm", headers, "")
	if err != nil {
		return nil, fmt.Errorf("cloudflare request failed: %w", err)
	}

	re := regexp.MustCompile(`"token":\s*"([^"]+)"`)
	matches := re.FindStringSubmatch(string(resp.Body))
	if len(matches) < 2 {
		return nil, fmt.Errorf("unable to parse Glassdoor CSRF token")
	}

	csrf := matches[1]
	return &csrf, nil
}

func generateHeaders(csrf string) string {
	return fmt.Sprintf(`{
		"accept": "*/*",
		"accept-language": "en-US,en;q=0.9",
		"apollographql-client-name": "job-search-next",
		"apollographql-client-version": "7.189.3",
		"content-type": "application/json",
		"gd-csrf-token": "%s",
		"origin": "https://www.glassdoor.com",
		"priority": "u=1, i",
		"referer": "https://www.glassdoor.com/",
		"sec-ch-ua": "\"Google Chrome\";v=\"137\", \"Chromium\";v=\"137\", \"Not/A)Brand\";v=\"24\"",
		"sec-ch-ua-mobile": "?0",
		"sec-ch-ua-platform": "\"Windows\"",
		"sec-fetch-dest": "empty",
		"sec-fetch-mode": "cors",
		"sec-fetch-site": "same-origin",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
		"x-gd-job-page": "serp"
	}`, csrf)
}

func generatePayload(keywords string, pageNumber int, pageCursor string) string {
	slug := strings.ReplaceAll(keywords, " ", "-")
	ko := generateKO(keywords)

	return fmt.Sprintf(`[{
		"operationName": "JobSearchResultsQuery",
		"variables": {
			"excludeJobListingIds": [],
			"filterParams": [
				{
					"filterKey": "fromAge",
					"values": "1"
				},
				{
					"filterKey": "sortBy",
					"values": "date_desc"
				},
				{
					"filterKey": "industryNId",
					"values": "10013"
				}
			],
			"keyword": "%v",
			"locationId": 1,
			"locationType": "COUNTRY",
			"numJobsToShow": 30,
			"originalPageUrl": "https://www.glassdoor.com/Job/united-states-%v-jobs-SRCH_IL.0,13_IN1_%v.htm",
			"parameterUrlInput": "IL.0,13_IN1_%v",
			"pageType": "SERP",
			"queryString": "",
			"seoFriendlyUrlInput": "united-states-%v-jobs",
			"seoUrl": true,
			"includeIndeedJobAttributes": false,
			"pageNumber": %d,
			"pageCursor": "%v"
		},
		"query": "query JobSearchResultsQuery($excludeJobListingIds: [Long!], $filterParams: [FilterParams], $keyword: String, $locationId: Int, $locationType: LocationTypeEnum, $numJobsToShow: Int!, $originalPageUrl: String, $pageCursor: String, $pageNumber: Int, $pageType: PageTypeEnum, $parameterUrlInput: String, $queryString: String, $seoFriendlyUrlInput: String, $seoUrl: Boolean, $includeIndeedJobAttributes: Boolean) {\n  jobListings(\n    contextHolder: {queryString: $queryString, pageTypeEnum: $pageType, searchParams: {excludeJobListingIds: $excludeJobListingIds, filterParams: $filterParams, keyword: $keyword, locationId: $locationId, locationType: $locationType, numPerPage: $numJobsToShow, pageCursor: $pageCursor, pageNumber: $pageNumber, originalPageUrl: $originalPageUrl, seoFriendlyUrlInput: $seoFriendlyUrlInput, parameterUrlInput: $parameterUrlInput, seoUrl: $seoUrl, searchType: SR, includeIndeedJobAttributes: $includeIndeedJobAttributes}}\n  ) {\n    companyFilterOptions {\n      id\n      shortName\n      __typename\n    }\n    filterOptions\n    indeedCtk\n    jobListings {\n      ...JobListingJobView\n      __typename\n    }\n    jobSearchTrackingKey\n    jobsPageSeoData {\n      pageMetaDescription\n      pageTitle\n      __typename\n    }\n    paginationCursors {\n      cursor\n      pageNumber\n      __typename\n    }\n    indexablePageForSeo\n    searchResultsMetadata {\n      searchCriteria {\n        implicitLocation {\n          id\n          localizedDisplayName\n          type\n          __typename\n        }\n        keyword\n        location {\n          id\n          shortName\n          localizedShortName\n          localizedDisplayName\n          type\n          __typename\n        }\n        __typename\n      }\n      footerVO {\n        countryMenu {\n          childNavigationLinks {\n            id\n            link\n            textKey\n            __typename\n          }\n          __typename\n        }\n        __typename\n      }\n      helpCenterDomain\n      helpCenterLocale\n      jobAlert {\n        jobAlertId\n        __typename\n      }\n      jobSerpFaq {\n        questions {\n          answer\n          question\n          __typename\n        }\n        __typename\n      }\n      jobSerpJobOutlook {\n        occupation\n        paragraph\n        heading\n        __typename\n      }\n      showMachineReadableJobs\n      __typename\n    }\n    serpSeoLinksVO {\n      relatedJobTitlesResults\n      searchedJobTitle\n      searchedKeyword\n      searchedLocationIdAsString\n      searchedLocationSeoName\n      searchedLocationType\n      topCityIdsToNameResults {\n        key\n        value\n        __typename\n      }\n      topEmployerIdsToNameResults {\n        key\n        value\n        __typename\n      }\n      topOccupationResults\n      __typename\n    }\n    totalJobsCount\n    __typename\n  }\n}\n\nfragment JobListingJobView on JobListingSearchResult {\n  jobview {\n    header {\n      indeedJobAttribute {\n        skills\n        extractedJobAttributes {\n          key\n          value\n          __typename\n        }\n        __typename\n      }\n      adOrderId\n      ageInDays\n      easyApply\n      employer {\n        id\n        name\n        shortName\n        __typename\n      }\n      expired\n      occupations {\n        key\n        __typename\n      }\n      employerNameFromSearch\n      goc\n      gocId\n      isSponsoredJob\n      isSponsoredEmployer\n      jobCountryId\n      jobLink\n      jobResultTrackingKey\n      normalizedJobTitle\n      jobTitleText\n      locationName\n      locationType\n      locId\n      payCurrency\n      payPeriod\n      payPeriodAdjustedPay {\n        p10\n        p50\n        p90\n        __typename\n      }\n      rating\n      salarySource\n      savedJobId\n      seoJobLink\n      __typename\n    }\n    job {\n      descriptionFragmentsText\n      importConfigId\n      jobTitleId\n      jobTitleText\n      listingId\n      __typename\n    }\n    jobListingAdminDetails {\n      userEligibleForAdminJobDetails\n      __typename\n    }\n    overview {\n      shortName\n      squareLogoUrl\n      __typename\n    }\n    __typename\n  }\n  __typename\n}"
	}]`, keywords, slug, ko, ko, slug, pageNumber, pageCursor)
}

func generateKO(keyword string) string {
	base := "united-states-"
	start := len(base)
	slug := strings.ReplaceAll(keyword, " ", "-")
	end := start + len(slug)
	return fmt.Sprintf("KO%d,%d", start, end)
}
