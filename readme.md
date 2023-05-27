My golang-web scraper. 

Below is an image that quickly says what I need to write for 2 paragraphs.

![Diagram of the PWA](https://github.com/georgekakarlis/go-scrape/blob/main/Diagram.png?raw=true)


Then application-wise the whole idea is to have a form where a user would input an URL of which he wishes to scrape the data from,
and then select on the radio buttons in which form of the data he wishes to get back. The options for now are CSV and json as default (from the browser parsing). Afterwards, when he clicks submit he sends a POST request to the Go API and after we deserialize the request we have on the /process endpoint a switch/case statement, which on Go executes only which-one is about to execute (on the form we pass the "generateCSV" id then we execute the "generateCSV" function and not the rest case statements). After that we send away Go-Colly to scrape the targeted URL and give back (for now only Ahrefs) the data. Then, we ensure directory folders (where the download files should be) eg /downloads/CSV/scraped-data.csv . After we write to that file and we send it to that directory. Then we share the filePath variable/headers to the download handler and from the client we use javascript to make the download button visible after we get back response (from the headers) from the API that the file is there and we can download the file. After that, finally, we delete the file from our directory and it relies trustfully on the client/browser. 

a few things to point out : 
1. OFC im working on an API limiter so that nobody can request more than 3 times in 5 minutes.
2. im working on saving to a private json/sqlite3 file how many urls have been scraped (sth like analytics)
3. maybe pricing model


future todos : 
Async scraper with ML(?) to know/understand which headers to use to facilitate the right combination of IPs and headers, CAPCHAs, and any other anti-scraping systems on each website call.




the most "difficult" thing to do was to ensureDirectory and the whole download and os.Remove thing.

More to come soon

Cheers.