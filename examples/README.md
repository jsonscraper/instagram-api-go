# Instagram go request examples

Every file below is a complete single-request CLI program with environment-based authentication, all documented query options, timeout configuration, error handling and full response output. Sample defaults come from Postman; replace account names, IDs and URLs for your use case. `--help` requires no API key and makes no network request.

Set `JSONSCRAPER_INSTAGRAM_LICENSE_KEY` and install the SDK as described in the main README. Each execution makes one live request and may consume quota. None of these programs loops through a feed automatically.

## Get user info by ID

[Open the complete example](endpoints/v1_user_by_id.go) · `GET /v1/user/by/id`

```sh
go run examples/endpoints/v1_user_by_id.go --help
```

## Get user info by username (by ID is faster)

[Open the complete example](endpoints/v1_user_by_username.go) · `GET /v1/user/by/username`

```sh
go run examples/endpoints/v1_user_by_username.go --help
```

## Get user info by URL (by ID is faster)

[Open the complete example](endpoints/v1_user_by_url.go) · `GET /v1/user/by/url`

```sh
go run examples/endpoints/v1_user_by_url.go --help
```

## Get info about user by ID

[Open the complete example](endpoints/v1_user_about.go) · `GET /v1/user/about`

```sh
go run examples/endpoints/v1_user_about.go --help
```

## Get posts (chunked)

[Open the complete example](endpoints/v1_user_medias_chunk.go) · `GET /v1/user/medias/chunk`

```sh
go run examples/endpoints/v1_user_medias_chunk.go --help
```

## Get pinned posts

[Open the complete example](endpoints/v1_user_medias_pinned.go) · `GET /v1/user/medias/pinned`

```sh
go run examples/endpoints/v1_user_medias_pinned.go --help
```

## Get reels (chunked)

[Open the complete example](endpoints/v1_user_clips_chunk.go) · `GET /v1/user/clips/chunk`

```sh
go run examples/endpoints/v1_user_clips_chunk.go --help
```

## Get tagged (mentions) posts

[Open the complete example](endpoints/v1_user_tag_medias_chunk.go) · `GET /v1/user/tag/medias/chunk`

```sh
go run examples/endpoints/v1_user_tag_medias_chunk.go --help
```

## Get user stories by ID

[Open the complete example](endpoints/v1_user_stories.go) · `GET /v1/user/stories`

```sh
go run examples/endpoints/v1_user_stories.go --help
```

## Get user stories by username (by ID is faster)

[Open the complete example](endpoints/v1_user_stories_by_username.go) · `GET /v1/user/stories/by/username`

```sh
go run examples/endpoints/v1_user_stories_by_username.go --help
```

## Get user highlights by ID

[Open the complete example](endpoints/v1_user_highlights.go) · `GET /v1/user/highlights`

```sh
go run examples/endpoints/v1_user_highlights.go --help
```

## Get user highlights by username (by ID is faster)

[Open the complete example](endpoints/v1_user_highlights_by_username.go) · `GET /v1/user/highlights/by/username`

```sh
go run examples/endpoints/v1_user_highlights_by_username.go --help
```

## Search in followers list by query (not working for legacy verified accounts)

[Open the complete example](endpoints/v1_user_search_followers.go) · `GET /v1/user/search/followers`

```sh
go run examples/endpoints/v1_user_search_followers.go --help
```

## Search in followings list by query (not working for legacy verified accounts)

[Open the complete example](endpoints/v1_user_search_following.go) · `GET /v1/user/search/following`

```sh
go run examples/endpoints/v1_user_search_following.go --help
```

## Get user following (chunked)

[Open the complete example](endpoints/v1_user_following_chunk.go) · `GET /v1/user/following/chunk`

```sh
go run examples/endpoints/v1_user_following_chunk.go --help
```

## Get user followers (chunked)

[Open the complete example](endpoints/v1_user_followers_chunk.go) · `GET /v1/user/followers/chunk`

```sh
go run examples/endpoints/v1_user_followers_chunk.go --help
```

## Get media info by ID

[Open the complete example](endpoints/v1_media_by_id.go) · `GET /v1/media/by/id`

```sh
go run examples/endpoints/v1_media_by_id.go --help
```

## Get media info by shortcode

[Open the complete example](endpoints/v1_media_by_code.go) · `GET /v1/media/by/code`

```sh
go run examples/endpoints/v1_media_by_code.go --help
```

## Get media info by URL

[Open the complete example](endpoints/v1_media_by_url.go) · `GET /v1/media/by/url`

```sh
go run examples/endpoints/v1_media_by_url.go --help
```

## Get media insights

[Open the complete example](endpoints/v1_media_insight.go) · `GET /v1/media/insight`

```sh
go run examples/endpoints/v1_media_insight.go --help
```

## Get media comments (chunked)

[Open the complete example](endpoints/v1_media_comments_chunk.go) · `GET /v1/media/comments/chunk`

```sh
go run examples/endpoints/v1_media_comments_chunk.go --help
```

## Get media likers

[Open the complete example](endpoints/v1_media_likers.go) · `GET /v1/media/likers`

```sh
go run examples/endpoints/v1_media_likers.go --help
```

## Get author of the media

[Open the complete example](endpoints/v1_media_user.go) · `GET /v1/media/user`

```sh
go run examples/endpoints/v1_media_user.go --help
```

## Get media oEmbed data

[Open the complete example](endpoints/v1_media_oembed.go) · `GET /v1/media/oembed`

```sh
go run examples/endpoints/v1_media_oembed.go --help
```

## Get media code from media ID (PK) (unlimited usage)

[Open the complete example](endpoints/v1_media_code_from_pk.go) · `GET /v1/media/code/from/pk`

```sh
go run examples/endpoints/v1_media_code_from_pk.go --help
```

## Get media ID (PK) from shortcode (unlimited usage)

[Open the complete example](endpoints/v1_media_pk_from_code.go) · `GET /v1/media/pk/from/code`

```sh
go run examples/endpoints/v1_media_pk_from_code.go --help
```

## Get media ID (PK) from URL (unlimited usage)

[Open the complete example](endpoints/v1_media_pk_from_url.go) · `GET /v1/media/pk/from/url`

```sh
go run examples/endpoints/v1_media_pk_from_url.go --help
```

## Get story by ID

[Open the complete example](endpoints/v1_story_by_id.go) · `GET /v1/story/by/id`

```sh
go run examples/endpoints/v1_story_by_id.go --help
```

## Get story by URL

[Open the complete example](endpoints/v1_story_by_url.go) · `GET /v1/story/by/url`

```sh
go run examples/endpoints/v1_story_by_url.go --help
```

## Download story by ID

[Open the complete example](endpoints/v1_story_download.go) · `GET /v1/story/download`

```sh
go run examples/endpoints/v1_story_download.go --help
```

## Download story by direct media URL (unlimited usage)

[Open the complete example](endpoints/v1_story_download_by_url.go) · `GET /v1/story/download/by/url`

```sh
go run examples/endpoints/v1_story_download_by_url.go --help
```

## Download story file by story URL (IG story link)

[Open the complete example](endpoints/v1_story_download_by_story_url.go) · `GET /v1/story/download/by/story/url`

```sh
go run examples/endpoints/v1_story_download_by_story_url.go --help
```

## Get highlight by ID

[Open the complete example](endpoints/v1_highlight_by_id.go) · `GET /v1/highlight/by/id`

```sh
go run examples/endpoints/v1_highlight_by_id.go --help
```

## Get highlight by URL

[Open the complete example](endpoints/v1_highlight_by_url.go) · `GET /v1/highlight/by/url`

```sh
go run examples/endpoints/v1_highlight_by_url.go --help
```

## Get location by ID

[Open the complete example](endpoints/v1_location_by_id.go) · `GET /v1/location/by/id`

```sh
go run examples/endpoints/v1_location_by_id.go --help
```

## Get location top medias

[Open the complete example](endpoints/v1_location_medias_top.go) · `GET /v1/location/medias/top`

```sh
go run examples/endpoints/v1_location_medias_top.go --help
```

## Get location recent medias

[Open the complete example](endpoints/v1_location_medias_recent.go) · `GET /v1/location/medias/recent`

```sh
go run examples/endpoints/v1_location_medias_recent.go --help
```

## Get location top medias (chunked)

[Open the complete example](endpoints/v1_location_medias_top_chunk.go) · `GET /v1/location/medias/top/chunk`

```sh
go run examples/endpoints/v1_location_medias_top_chunk.go --help
```

## Get location recent medias (chunked)

[Open the complete example](endpoints/v1_location_medias_recent_chunk.go) · `GET /v1/location/medias/recent/chunk`

```sh
go run examples/endpoints/v1_location_medias_recent_chunk.go --help
```

## Search locations

[Open the complete example](endpoints/v1_location_search.go) · `GET /v1/location/search`

```sh
go run examples/endpoints/v1_location_search.go --help
```

## Get hashtag by name

[Open the complete example](endpoints/v1_hashtag_by_name.go) · `GET /v1/hashtag/by/name`

```sh
go run examples/endpoints/v1_hashtag_by_name.go --help
```

## Get hashtag top medias

[Open the complete example](endpoints/v1_hashtag_medias_top.go) · `GET /v1/hashtag/medias/top`

```sh
go run examples/endpoints/v1_hashtag_medias_top.go --help
```

## Get hashtag top medias (chunked)

[Open the complete example](endpoints/v1_hashtag_medias_top_chunk.go) · `GET /v1/hashtag/medias/top/chunk`

```sh
go run examples/endpoints/v1_hashtag_medias_top_chunk.go --help
```

## Get hashtag recent medias (chunked)

[Open the complete example](endpoints/v1_hashtag_medias_top_recent_chunk.go) · `GET /v1/hashtag/medias/top/recent/chunk`

```sh
go run examples/endpoints/v1_hashtag_medias_top_recent_chunk.go --help
```

## Get hashtag clips (chunked)

[Open the complete example](endpoints/v1_hashtag_medias_clips_chunk.go) · `GET /v1/hashtag/medias/clips/chunk`

```sh
go run examples/endpoints/v1_hashtag_medias_clips_chunk.go --help
```

## Search hashtags

[Open the complete example](endpoints/v1_search_hashtags.go) · `GET /v1/search/hashtags`

```sh
go run examples/endpoints/v1_search_hashtags.go --help
```

## Search users

[Open the complete example](endpoints/v1_search_users.go) · `GET /v1/search/users`

```sh
go run examples/endpoints/v1_search_users.go --help
```

## Search music

[Open the complete example](endpoints/v1_search_music.go) · `GET /v1/search/music`

```sh
go run examples/endpoints/v1_search_music.go --help
```

## Search places

[Open the complete example](endpoints/v1_fbsearch_places.go) · `GET /v1/fbsearch/places`

```sh
go run examples/endpoints/v1_fbsearch_places.go --help
```

## Top search

[Open the complete example](endpoints/v1_fbsearch_topsearch.go) · `GET /v1/fbsearch/topsearch`

```sh
go run examples/endpoints/v1_fbsearch_topsearch.go --help
```

## Get user info by ID

[Open the complete example](endpoints/v2_user_by_id.go) · `GET /v2/user/by/id`

```sh
go run examples/endpoints/v2_user_by_id.go --help
```

## Get user info by username

[Open the complete example](endpoints/v2_user_by_username.go) · `GET /v2/user/by/username`

```sh
go run examples/endpoints/v2_user_by_username.go --help
```

## Get user info stream by ID

[Open the complete example](endpoints/v2_userstream_by_id.go) · `GET /v2/userstream/by/id`

```sh
go run examples/endpoints/v2_userstream_by_id.go --help
```

## Get user info stream by username

[Open the complete example](endpoints/v2_userstream_by_username.go) · `GET /v2/userstream/by/username`

```sh
go run examples/endpoints/v2_userstream_by_username.go --help
```

## Get user stories by ID

[Open the complete example](endpoints/v2_user_stories.go) · `GET /v2/user/stories`

```sh
go run examples/endpoints/v2_user_stories.go --help
```

## Get user stories by username

[Open the complete example](endpoints/v2_user_stories_by_username.go) · `GET /v2/user/stories/by/username`

```sh
go run examples/endpoints/v2_user_stories_by_username.go --help
```

## Get user reels by ID

[Open the complete example](endpoints/v2_user_clips.go) · `GET /v2/user/clips`

```sh
go run examples/endpoints/v2_user_clips.go --help
```

## Get user following list

[Open the complete example](endpoints/v2_user_following.go) · `GET /v2/user/following`

```sh
go run examples/endpoints/v2_user_following.go --help
```

## Get user followers list

[Open the complete example](endpoints/v2_user_followers.go) · `GET /v2/user/followers`

```sh
go run examples/endpoints/v2_user_followers.go --help
```

## Get media where user is tagged

[Open the complete example](endpoints/v2_user_tag_medias.go) · `GET /v2/user/tag/medias`

```sh
go run examples/endpoints/v2_user_tag_medias.go --help
```

## Get user highlights by ID

[Open the complete example](endpoints/v2_user_highlights.go) · `GET /v2/user/highlights`

```sh
go run examples/endpoints/v2_user_highlights.go --help
```

## Get user highlights by username (by ID is faster)

[Open the complete example](endpoints/v2_user_highlights_by_username.go) · `GET /v2/user/highlights/by/username`

```sh
go run examples/endpoints/v2_user_highlights_by_username.go --help
```

## Get business suggestions for user

[Open the complete example](endpoints/v2_user_explore_businesses_by_id.go) · `GET /v2/user/explore/businesses/by/id`

```sh
go run examples/endpoints/v2_user_explore_businesses_by_id.go --help
```

## Get suggested profiles for user

[Open the complete example](endpoints/v2_user_suggested_profiles.go) · `GET /v2/user/suggested/profiles`

```sh
go run examples/endpoints/v2_user_suggested_profiles.go --help
```

## Get media info by ID

[Open the complete example](endpoints/v2_media_info_by_id.go) · `GET /v2/media/info/by/id`

```sh
go run examples/endpoints/v2_media_info_by_id.go --help
```

## Get media info by shortcode

[Open the complete example](endpoints/v2_media_info_by_code.go) · `GET /v2/media/info/by/code`

```sh
go run examples/endpoints/v2_media_info_by_code.go --help
```

## Get media info by URL

[Open the complete example](endpoints/v2_media_info_by_url.go) · `GET /v2/media/info/by/url`

```sh
go run examples/endpoints/v2_media_info_by_url.go --help
```

## Get media comments

[Open the complete example](endpoints/v2_media_comments.go) · `GET /v2/media/comments`

```sh
go run examples/endpoints/v2_media_comments.go --help
```

## Get media likers

[Open the complete example](endpoints/v2_media_likers.go) · `GET /v2/media/likers`

```sh
go run examples/endpoints/v2_media_likers.go --help
```

## Check comment for offensiveness

[Open the complete example](endpoints/v2_media_comment_offensive.go) · `GET /v2/media/comment/offensive`

```sh
go run examples/endpoints/v2_media_comment_offensive.go --help
```

## Get comment replies

[Open the complete example](endpoints/v2_media_comments_replies.go) · `GET /v2/media/comments/replies`

```sh
go run examples/endpoints/v2_media_comments_replies.go --help
```

## Get story by ID

[Open the complete example](endpoints/v2_story_by_id.go) · `GET /v2/story/by/id`

```sh
go run examples/endpoints/v2_story_by_id.go --help
```

## Get music feed by canonical ID

[Open the complete example](endpoints/v2_track_by_canonical_id.go) · `GET /v2/track/by/canonical/id`

```sh
go run examples/endpoints/v2_track_by_canonical_id.go --help
```

## Get music feed by ID

[Open the complete example](endpoints/v2_track_by_id.go) · `GET /v2/track/by/id`

```sh
go run examples/endpoints/v2_track_by_id.go --help
```

## Get hashtag by name

[Open the complete example](endpoints/v2_hashtag_by_name.go) · `GET /v2/hashtag/by/name`

```sh
go run examples/endpoints/v2_hashtag_by_name.go --help
```

## Get hashtag top medias

[Open the complete example](endpoints/v2_hashtag_medias_top.go) · `GET /v2/hashtag/medias/top`

```sh
go run examples/endpoints/v2_hashtag_medias_top.go --help
```

## Get hashtag recent medias

[Open the complete example](endpoints/v2_hashtag_medias_recent.go) · `GET /v2/hashtag/medias/recent`

```sh
go run examples/endpoints/v2_hashtag_medias_recent.go --help
```

## Get highlight by ID

[Open the complete example](endpoints/v2_highlight_by_id.go) · `GET /v2/highlight/by/id`

```sh
go run examples/endpoints/v2_highlight_by_id.go --help
```

## Search hashtags

[Open the complete example](endpoints/v2_search_hashtags.go) · `GET /v2/search/hashtags`

```sh
go run examples/endpoints/v2_search_hashtags.go --help
```

## Search music

[Open the complete example](endpoints/v2_search_music.go) · `GET /v2/search/music`

```sh
go run examples/endpoints/v2_search_music.go --help
```

## Search accounts (users)

[Open the complete example](endpoints/v2_fbsearch_accounts.go) · `GET /v2/fbsearch/accounts`

```sh
go run examples/endpoints/v2_fbsearch_accounts.go --help
```

## Search places (locations)

[Open the complete example](endpoints/v2_fbsearch_places.go) · `GET /v2/fbsearch/places`

```sh
go run examples/endpoints/v2_fbsearch_places.go --help
```

## Top search

[Open the complete example](endpoints/v2_fbsearch_topsearch.go) · `GET /v2/fbsearch/topsearch`

```sh
go run examples/endpoints/v2_fbsearch_topsearch.go --help
```

## Search reels (clips)

[Open the complete example](endpoints/v2_fbsearch_reels.go) · `GET /v2/fbsearch/reels`

```sh
go run examples/endpoints/v2_fbsearch_reels.go --help
```

## Get comments list by media ID (chunked)

[Open the complete example](endpoints/gql_comments_chunk.go) · `GET /gql/comments/chunk`

```sh
go run examples/endpoints/gql_comments_chunk.go --help
```

## Get threaded replies list by comment ID (chunked)

[Open the complete example](endpoints/gql_comments_threaded_chunk.go) · `GET /gql/comments/threaded/chunk`

```sh
go run examples/endpoints/gql_comments_threaded_chunk.go --help
```

## Get comment likers list by comment ID (chunked)

[Open the complete example](endpoints/gql_comment_likers_chunk.go) · `GET /gql/comment/likers/chunk`

```sh
go run examples/endpoints/gql_comment_likers_chunk.go --help
```

## Search top content by keyword

[Open the complete example](endpoints/gql_topsearch.go) · `GET /gql/topsearch`

```sh
go run examples/endpoints/gql_topsearch.go --help
```

## Get media likers by media ID

[Open the complete example](endpoints/gql_media_likers.go) · `GET /gql/media/likers`

```sh
go run examples/endpoints/gql_media_likers.go --help
```

## Get video usertags by media IDs

[Open the complete example](endpoints/gql_media_usertags.go) · `GET /gql/media/usertags`

```sh
go run examples/endpoints/gql_media_usertags.go --help
```

## Get user followers (chunked)

[Open the complete example](endpoints/gql_user_followers_chunk.go) · `GET /gql/user/followers/chunk`

```sh
go run examples/endpoints/gql_user_followers_chunk.go --help
```

## Get user following (chunked)

[Open the complete example](endpoints/gql_user_following_chunk.go) · `GET /gql/user/following/chunk`

```sh
go run examples/endpoints/gql_user_following_chunk.go --help
```

## Get user reposts

[Open the complete example](endpoints/gql_user_reposts.go) · `GET /gql/user/reposts`

```sh
go run examples/endpoints/gql_user_reposts.go --help
```

## Get user clips (reels)

[Open the complete example](endpoints/gql_user_clips.go) · `GET /gql/user/clips`

```sh
go run examples/endpoints/gql_user_clips.go --help
```

## Get user profile media (grid)

[Open the complete example](endpoints/gql_user_medias.go) · `GET /gql/user/medias`

```sh
go run examples/endpoints/gql_user_medias.go --help
```

## Search users

[Open the complete example](endpoints/v3_fbsearch_accounts.go) · `GET /v3/fbsearch/accounts`

```sh
go run examples/endpoints/v3_fbsearch_accounts.go --help
```

## Search places (locations)

[Open the complete example](endpoints/v3_fbsearch_places.go) · `GET /v3/fbsearch/places`

```sh
go run examples/endpoints/v3_fbsearch_places.go --help
```

