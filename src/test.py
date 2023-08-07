import re
import requests
from hashlib import sha256

headers = {
    'authority': 'raw.githubusercontent.com',
    'accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7',
    'accept-language': 'en-US,en;q=0.9',
    'cache-control': 'max-age=0',
    'referer': 'https://github.com/KamiCYun/Internship-Discord-Monitor/blob/main/test.md',
    'sec-ch-ua': '"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"',
    'sec-ch-ua-mobile': '?0',
    'sec-ch-ua-platform': '"Windows"',
    'sec-fetch-dest': 'document',
    'sec-fetch-mode': 'navigate',
    'sec-fetch-site': 'cross-site',
    'sec-fetch-user': '?1',
    'upgrade-insecure-requests': '1',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36',
}

response = requests.get('https://raw.githubusercontent.com/KamiCYun/Internship-Discord-Monitor/main/test.md', headers=headers)



db = set()

jobs = re.findall("\|\s\*{2}.+\*{2}\s\|\s.+\s\|\n", response.text)

for job in jobs:
    fields = job.split(" | ")
    hash_str = fields[0] + fields[1] + fields[2] + fields[4]
    job_hash = sha256(hash_str.encode('utf-8')).hexdigest()
    if job_hash not in db:
        title = re.findall("\*\*\[(.+)\].*\*\*", fields[0])
        if len(title) == 0:
            title = re.findall("\*\*(.+)\*\*", fields[0])
        title = title[0]

        position = fields[1]
        location = fields[2]
        url = re.findall("ref=\".*\">", fields[3])
        if len(url) == 0:
            url = "Closed"
        else:
            url = url[0]
        date = fields[4].replace(" |", "")

        print(title, position, location, url, date)