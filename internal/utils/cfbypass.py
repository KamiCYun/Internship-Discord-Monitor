import cloudscraper
import sys
import json

def main():
    if len(sys.argv) < 3:
        print(json.dumps({"error": "Usage: fetcher.py <GET|POST> <url> [headers_json] [payload_json]"}))
        sys.exit(1)

    method = sys.argv[1].upper()
    url = sys.argv[2]

    headers = json.loads(sys.argv[3]) if len(sys.argv) > 3 else {}
    payload = json.loads(sys.argv[4]) if len(sys.argv) > 4 else None

    scraper = cloudscraper.create_scraper()

    try:
        if method == "GET":
            resp = scraper.get(url, headers=headers)
        elif method == "POST":
            resp = scraper.post(url, headers=headers, json=payload)
        else:
            print(json.dumps({"error": f"Unsupported method: {method}"}))
            sys.exit(1)

        result = {
            "status": resp.status_code,
            "body": resp.text
        }

        print(json.dumps(result))
    except Exception as e:
        print(json.dumps({"error": str(e)}))
        sys.exit(1)

if __name__ == "__main__":
    main()
