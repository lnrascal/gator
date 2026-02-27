Gator

Installation:
- Install go
- Install postgres
- Run "go install" in directory using cli

Usage:
- Registration: gator register %username
- Authroization: gator login %username
- Clearing database: gator reset
- Users List: gator users
- Add Feed: gator addfeed %name %url
- Feeds List: gator feeds
- Follow: gator follow %url
- Unfollow: gator unfollow %url
- Subscribed feeds list: gator following
- Start scraping feeds: gator agg %time_interval (example: 10s, 1m, etc.)
- Browse feeds: gator browse %count

