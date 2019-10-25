module SocialMediaSite

replace SocialMediaSite/handlers => ./handlers

go 1.13

require (
	SocialMediaSite/handlers v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.2.0
)
