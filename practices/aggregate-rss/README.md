# Functional requirement
Ability to collect rss feeds from multiple newspaper sites, which api like /feeds, /rss,
Pull feed rss each 3-6 hours, in feed rss contains articles.
Make sure articles are synchronized
- Thumnail should be resize to device size and store on object storage such as s3, google cloud storage
- View and scroll new sorted articles by timestamp from thousand of sources 
- Click on article to redirect to the content article site
- Allow text search to find articles in among multiple article sites

# Non-functional requirement
Low latency < 200ms