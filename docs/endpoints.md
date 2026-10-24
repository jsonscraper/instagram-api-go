# API endpoint reference

Generated from the public Postman collections. Duplicate paths are represented once. This is a route/parameter catalogue, not a guaranteed response schema. Some optional parameters are enabled in Postman; check the linked collection for validation rules. `license_key` is configured on the client.

## Instagram

| Method name | GET path | Query parameters | Purpose |
| --- | --- | --- | --- |
| `V1UserById` | `/v1/user/by/id` | `id`, `cache_timeout` | Get user info by ID |
| `V1UserByUsername` | `/v1/user/by/username` | `username`, `cache_timeout` | Get user info by username (by ID is faster) |
| `V1UserByUrl` | `/v1/user/by/url` | `url`, `cache_timeout` | Get user info by URL (by ID is faster) |
| `V1UserAbout` | `/v1/user/about` | `id`, `cache_timeout` | Get info about user by ID |
| `V1UserMediasChunk` | `/v1/user/medias/chunk` | `user_id`, `end_cursor`, `cache_timeout` | Get posts (chunked) |
| `V1UserMediasPinned` | `/v1/user/medias/pinned` | `user_id`, `amount`, `cache_timeout` | Get pinned posts |
| `V1UserClipsChunk` | `/v1/user/clips/chunk` | `user_id`, `end_cursor`, `cache_timeout` | Get reels (chunked) |
| `V1UserTagMediasChunk` | `/v1/user/tag/medias/chunk` | `user_id`, `max_id`, `cache_timeout` | Get tagged (mentions) posts |
| `V1UserStories` | `/v1/user/stories` | `user_id`, `amount`, `force`, `cache_timeout` | Get user stories by ID |
| `V1UserStoriesByUsername` | `/v1/user/stories/by/username` | `username`, `amount`, `force`, `cache_timeout` | Get user stories by username (by ID is faster) |
| `V1UserHighlights` | `/v1/user/highlights` | `user_id`, `amount`, `force`, `cache_timeout` | Get user highlights by ID |
| `V1UserHighlightsByUsername` | `/v1/user/highlights/by/username` | `username`, `amount`, `force`, `cache_timeout` | Get user highlights by username (by ID is faster) |
| `V1UserSearchFollowers` | `/v1/user/search/followers` | `user_id`, `query`, `force`, `cache_timeout` | Search in followers list by query (not working for legacy verified accounts) |
| `V1UserSearchFollowing` | `/v1/user/search/following` | `user_id`, `query`, `force`, `cache_timeout` | Search in followings list by query (not working for legacy verified accounts) |
| `V1UserFollowingChunk` | `/v1/user/following/chunk` | `user_id`, `max_id`, `cache_timeout` | Get user following (chunked) |
| `V1UserFollowersChunk` | `/v1/user/followers/chunk` | `user_id`, `max_id`, `cache_timeout` | Get user followers (chunked) |
| `V1MediaById` | `/v1/media/by/id` | `id`, `cache_timeout` | Get media info by ID |
| `V1MediaByCode` | `/v1/media/by/code` | `code`, `cache_timeout` | Get media info by shortcode |
| `V1MediaByUrl` | `/v1/media/by/url` | `url`, `cache_timeout` | Get media info by URL |
| `V1MediaInsight` | `/v1/media/insight` | `media_id`, `cache_timeout` | Get media insights |
| `V1MediaCommentsChunk` | `/v1/media/comments/chunk` | `id`, `min_id`, `max_id`, `can_support_threading`, `cache_timeout` | Get media comments (chunked) |
| `V1MediaLikers` | `/v1/media/likers` | `id`, `cache_timeout` | Get media likers |
| `V1MediaUser` | `/v1/media/user` | `media_id`, `cache_timeout` | Get author of the media |
| `V1MediaOembed` | `/v1/media/oembed` | `url`, `cache_timeout` | Get media oEmbed data |
| `V1MediaCodeFromPk` | `/v1/media/code/from/pk` | `pk`, `cache_timeout` | Get media code from media ID (PK) (unlimited usage) |
| `V1MediaPkFromCode` | `/v1/media/pk/from/code` | `code`, `cache_timeout` | Get media ID (PK) from shortcode (unlimited usage) |
| `V1MediaPkFromUrl` | `/v1/media/pk/from/url` | `url`, `cache_timeout` | Get media ID (PK) from URL (unlimited usage) |
| `V1StoryById` | `/v1/story/by/id` | `id`, `cache_timeout` | Get story by ID |
| `V1StoryByUrl` | `/v1/story/by/url` | `url`, `cache_timeout` | Get story by URL |
| `V1StoryDownload` | `/v1/story/download` | `id` | Download story by ID |
| `V1StoryDownloadByUrl` | `/v1/story/download/by/url` | `url` | Download story by direct media URL (unlimited usage) |
| `V1StoryDownloadByStoryUrl` | `/v1/story/download/by/story/url` | `url` | Download story file by story URL (IG story link) |
| `V1HighlightById` | `/v1/highlight/by/id` | `id`, `cache_timeout` | Get highlight by ID |
| `V1HighlightByUrl` | `/v1/highlight/by/url` | `url`, `cache_timeout` | Get highlight by URL |
| `V1LocationById` | `/v1/location/by/id` | `id`, `cache_timeout` | Get location by ID |
| `V1LocationMediasTop` | `/v1/location/medias/top` | `location_pk`, `amount`, `cache_timeout` | Get location top medias |
| `V1LocationMediasRecent` | `/v1/location/medias/recent` | `location_pk`, `amount`, `cache_timeout` | Get location recent medias |
| `V1LocationMediasTopChunk` | `/v1/location/medias/top/chunk` | `location_pk`, `max_id`, `cache_timeout` | Get location top medias (chunked) |
| `V1LocationMediasRecentChunk` | `/v1/location/medias/recent/chunk` | `location_pk`, `max_id`, `cache_timeout` | Get location recent medias (chunked) |
| `V1LocationSearch` | `/v1/location/search` | `lat`, `lng`, `cache_timeout` | Search locations |
| `V1HashtagByName` | `/v1/hashtag/by/name` | `name`, `cache_timeout` | Get hashtag by name |
| `V1HashtagMediasTop` | `/v1/hashtag/medias/top` | `name`, `amount`, `cache_timeout` | Get hashtag top medias |
| `V1HashtagMediasTopChunk` | `/v1/hashtag/medias/top/chunk` | `name`, `max_id`, `cache_timeout` | Get hashtag top medias (chunked) |
| `V1HashtagMediasTopRecentChunk` | `/v1/hashtag/medias/top/recent/chunk` | `name`, `max_id`, `cache_timeout` | Get hashtag recent medias (chunked) |
| `V1HashtagMediasClipsChunk` | `/v1/hashtag/medias/clips/chunk` | `name`, `max_id`, `cache_timeout` | Get hashtag clips (chunked) |
| `V1SearchHashtags` | `/v1/search/hashtags` | `query`, `cache_timeout` | Search hashtags |
| `V1SearchUsers` | `/v1/search/users` | `query`, `cache_timeout` | Search users |
| `V1SearchMusic` | `/v1/search/music` | `query`, `cache_timeout` | Search music |
| `V1FbsearchPlaces` | `/v1/fbsearch/places` | `query`, `lat`, `lng`, `cache_timeout` | Search places |
| `V1FbsearchTopsearch` | `/v1/fbsearch/topsearch` | `query`, `cache_timeout` | Top search |
| `V2UserById` | `/v2/user/by/id` | `id`, `cache_timeout` | Get user info by ID |
| `V2UserByUsername` | `/v2/user/by/username` | `username`, `cache_timeout` | Get user info by username |
| `V2UserstreamById` | `/v2/userstream/by/id` | `id`, `cache_timeout` | Get user info stream by ID |
| `V2UserstreamByUsername` | `/v2/userstream/by/username` | `username`, `cache_timeout` | Get user info stream by username |
| `V2UserStories` | `/v2/user/stories` | `user_id`, `force`, `safe_int`, `cache_timeout` | Get user stories by ID |
| `V2UserStoriesByUsername` | `/v2/user/stories/by/username` | `username`, `force`, `safe_int`, `cache_timeout` | Get user stories by username |
| `V2UserClips` | `/v2/user/clips` | `user_id`, `page_id`, `safe_int`, `cache_timeout` | Get user reels by ID |
| `V2UserFollowing` | `/v2/user/following` | `user_id`, `page_id`, `cache_timeout` | Get user following list |
| `V2UserFollowers` | `/v2/user/followers` | `user_id`, `page_id`, `cache_timeout` | Get user followers list |
| `V2UserTagMedias` | `/v2/user/tag/medias` | `user_id`, `page_id`, `cache_timeout` | Get media where user is tagged |
| `V2UserHighlights` | `/v2/user/highlights` | `user_id`, `amount`, `force`, `cache_timeout` | Get user highlights by ID |
| `V2UserHighlightsByUsername` | `/v2/user/highlights/by/username` | `username`, `amount`, `force`, `cache_timeout` | Get user highlights by username (by ID is faster) |
| `V2UserExploreBusinessesById` | `/v2/user/explore/businesses/by/id` | `user_id`, `cache_timeout` | Get business suggestions for user |
| `V2UserSuggestedProfiles` | `/v2/user/suggested/profiles` | `user_id`, `expand_suggestion`, `cache_timeout` | Get suggested profiles for user |
| `V2MediaInfoById` | `/v2/media/info/by/id` | `id`, `cache_timeout` | Get media info by ID |
| `V2MediaInfoByCode` | `/v2/media/info/by/code` | `code`, `cache_timeout` | Get media info by shortcode |
| `V2MediaInfoByUrl` | `/v2/media/info/by/url` | `url`, `cache_timeout` | Get media info by URL |
| `V2MediaComments` | `/v2/media/comments` | `id`, `page_id`, `can_support_threading`, `cache_timeout` | Get media comments |
| `V2MediaLikers` | `/v2/media/likers` | `id`, `cache_timeout` | Get media likers |
| `V2MediaCommentOffensive` | `/v2/media/comment/offensive` | `media_id`, `comment`, `cache_timeout` | Check comment for offensiveness |
| `V2MediaCommentsReplies` | `/v2/media/comments/replies` | `media_id`, `comment_id`, `min_id`, `cache_timeout` | Get comment replies |
| `V2StoryById` | `/v2/story/by/id` | `id`, `cache_timeout` | Get story by ID |
| `V2TrackByCanonicalId` | `/v2/track/by/canonical/id` | `canonical_id`, `page_id`, `cache_timeout` | Get music feed by canonical ID |
| `V2TrackById` | `/v2/track/by/id` | `track_id`, `page_id`, `cache_timeout` | Get music feed by ID |
| `V2HashtagByName` | `/v2/hashtag/by/name` | `name`, `cache_timeout` | Get hashtag by name |
| `V2HashtagMediasTop` | `/v2/hashtag/medias/top` | `name`, `page_id`, `cache_timeout` | Get hashtag top medias |
| `V2HashtagMediasRecent` | `/v2/hashtag/medias/recent` | `name`, `page_id`, `cache_timeout` | Get hashtag recent medias |
| `V2HighlightById` | `/v2/highlight/by/id` | `id`, `cache_timeout` | Get highlight by ID |
| `V2SearchHashtags` | `/v2/search/hashtags` | `query`, `page_token`, `cache_timeout` | Search hashtags |
| `V2SearchMusic` | `/v2/search/music` | `query`, `next_max_id`, `cache_timeout` | Search music |
| `V2FbsearchAccounts` | `/v2/fbsearch/accounts` | `query`, `page_token`, `cache_timeout` | Search accounts (users) |
| `V2FbsearchPlaces` | `/v2/fbsearch/places` | `query`, `cache_timeout` | Search places (locations) |
| `V2FbsearchTopsearch` | `/v2/fbsearch/topsearch` | `query`, `next_max_id`, `cache_timeout` | Top search |
| `V2FbsearchReels` | `/v2/fbsearch/reels` | `query`, `reels_max_id`, `cache_timeout` | Search reels (clips) |
| `GqlCommentsChunk` | `/gql/comments/chunk` | `media_id`, `sort_order`, `end_cursor`, `cache_timeout` | Get comments list by media ID (chunked) |
| `GqlCommentsThreadedChunk` | `/gql/comments/threaded/chunk` | `media_id`, `comment_id`, `end_cursor`, `cache_timeout` | Get threaded replies list by comment ID (chunked) |
| `GqlCommentLikersChunk` | `/gql/comment/likers/chunk` | `comment_id`, `end_cursor`, `cache_timeout` | Get comment likers list by comment ID (chunked) |
| `GqlTopsearch` | `/gql/topsearch` | `query`, `cache_timeout` | Search top content by keyword |
| `GqlMediaLikers` | `/gql/media/likers` | `media_id`, `cache_timeout` | Get media likers by media ID |
| `GqlMediaUsertags` | `/gql/media/usertags` | `media_ids`, `cache_timeout` | Get video usertags by media IDs |
| `GqlUserFollowersChunk` | `/gql/user/followers/chunk` | `user_id`, `force`, `end_cursor`, `cache_timeout` | Get user followers (chunked) |
| `GqlUserFollowingChunk` | `/gql/user/following/chunk` | `user_id`, `force`, `end_cursor`, `cache_timeout` | Get user following (chunked) |
| `GqlUserReposts` | `/gql/user/reposts` | `user_id`, `repost_next_max_id`, `flat`, `cache_timeout` | Get user reposts |
| `GqlUserClips` | `/gql/user/clips` | `user_id`, `max_id`, `sort_by_views`, `flat`, `cache_timeout` | Get user clips (reels) |
| `GqlUserMedias` | `/gql/user/medias` | `user_id`, `profile_grid_items_cursor`, `flat`, `cache_timeout` | Get user profile media (grid) |
| `V3FbsearchAccounts` | `/v3/fbsearch/accounts` | `query`, `page_token`, `cache_timeout` | Search users |
| `V3FbsearchPlaces` | `/v3/fbsearch/places` | `query`, `cache_timeout` | Search places (locations) |

## Tiktok

| Method name | GET path | Query parameters | Purpose |
| --- | --- | --- | --- |
| `GetNoWatermarkUrl` | `/getNoWatermarkUrl` | `video_url`, `region`, `cache_timeout` | Get non watermarked video url |
| `GetNoWatermarkUrlById` | `/getNoWatermarkUrlByID` | `video_id`, `region`, `cache_timeout` | Get non watermarked video url by video ID |
| `GetVideoById` | `/getVideoByID` | `video_id`, `region`, `cache_timeout` | Get video info by ID |
| `GetVideoByUrl` | `/getVideoByUrl` | `video_url`, `region`, `cache_timeout` | Get video info by URL |
| `GetUnshortenedUrl` | `/getUnshortenedUrl` | `video_url`, `cache_timeout` | Convert short TikTok video URL to full TikTok video URL |
| `GetCommentsByUrl` | `/getCommentsByUrl` | `video_url`, `region`, `cursor`, `count`, `cache_timeout` | Get comments list by video URL |
| `GetCommentsById` | `/getCommentsByID` | `video_id`, `region`, `cursor`, `count`, `cache_timeout` | Get comments list by video ID |
| `GetCommentsReplies` | `/getCommentsReplies` | `video_id`, `comment_id`, `cursor`, `count`, `cache_timeout` | Get comments replies |
| `SearchUser` | `/searchUser` | `keyword`, `cursor`, `count`, `region`, `follower_count`, `profile_type`, `other_pref`, `cache_timeout` | Search for user by keyword |
| `SearchVideo` | `/searchVideo` | `keyword`, `cursor`, `count`, `cache_timeout`, `date_posted`, `sort_by`, `region` | Search for video by keyword |
| `SearchMusic` | `/searchMusic` | `keyword`, `cursor`, `count`, `cache_timeout`, `filter_by`, `sort_type`, `region` | Search for music by keyword |
| `SearchLive` | `/searchLive` | `keyword`, `offset`, `count`, `cache_timeout`, `region` | Search for live by keyword |
| `SearchHashtag` | `/searchHashtag` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search for hashtag by keyword |
| `SearchLocation` | `/searchLocation` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search for location by keyword |
| `SearchTop` | `/searchTop` | `keyword`, `cursor`, `count`, `region`, `sort_type`, `publish_time`, `cache_timeout` | Top search by text query |
| `SearchInsights` | `/searchInsights` | `keyword`, `cursor`, `count`, `region`, `cache_timeout` | Search insights by text query |
| `SearchCommentSuggestedUsers` | `/searchCommentSuggestedUsers` | `keyword`, `count`, `region`, `cache_timeout` | Search comment suggested users |
| `GetUserInfo` | `/getUserInfo` | `query`, `region`, `cache_timeout` | Get user info by profile URL or nickname |
| `GetUserInfoById` | `/getUserInfoByID` | `user_id`, `region`, `cache_timeout` | Get user info by profile ID |
| `GetBasicUserInfo` | `/getBasicUserInfo` | `query`, `region`, `cache_timeout` | Get user info by username |
| `GetUserFeed` | `/getUserFeed` | `query`, `max_cursor`, `count`, `region`, `with_pinned_posts`, `sort_type`, `cache_timeout` | Get user feed by profile URL or nickname |
| `GetUserFeedById` | `/getUserFeedByID` | `user_id`, `max_cursor`, `count`, `region`, `with_pinned_posts`, `sort_type`, `cache_timeout` | Get user feed by profile ID |
| `GetUserId` | `/getUserID` | `query`, `region`, `cache_timeout` | Get user ID by profile URL or nickname |
| `GetUsernameById` | `/getUsernameByID` | `user_id`, `region`, `cache_timeout` | Get username by TikTok user ID |
| `GetTwitterUsername` | `/getTwitterUsername` | `query`, `region`, `cache_timeout` | Get Twitter username by profile URL, nickname or TikTok user ID |
| `GetInstagramUsername` | `/getInstagramUsername` | `query`, `region`, `cache_timeout` | Get Instagram username by profile URL, nickname or TikTok user ID |
| `GetYoutubeChannel` | `/getYoutubeChannel` | `query`, `region`, `cache_timeout` | Get YouTube channel info by profile URL, nickname or TikTok user ID |
| `GetFollowers` | `/getFollowers` | `query`, `count`, `min_time`, `cache_timeout` | Get user followers by profile URL or nickname |
| `GetFollowersById` | `/getFollowersByID` | `user_id`, `count`, `min_time`, `cache_timeout` | Get user followers by profile ID |
| `GetFollowing` | `/getFollowing` | `query`, `count`, `min_time`, `cache_timeout` | Get user following by profile URL or nickname |
| `GetFollowingById` | `/getFollowingByID` | `user_id`, `count`, `min_time`, `cache_timeout` | Get user following by profile ID |
| `GetLiveEventInfoById` | `/getLiveEventInfoByID` | `event_id`, `cache_timeout` | Get live event info by ID |
| `GetHashtagInfo` | `/getHashtagInfo` | `query`, `region`, `cache_timeout` | Get hashtag info by URL or tag name |
| `GetHashtagFeed` | `/getHashtagFeed` | `query`, `cursor`, `count`, `region`, `cache_timeout` | Get hashtag feed by URL or tag name |
| `GetHashtagFeedById` | `/getHashtagFeedByID` | `id`, `cursor`, `count`, `region`, `cache_timeout` | Get hashtag feed by ID |
| `GetLocationInfoById` | `/getLocationInfoByID` | `location_id`, `cache_timeout` | Get location info by ID |
| `GetLocationFeedById` | `/getLocationFeedByID` | `location_id`, `count`, `cursor`, `region`, `cache_timeout` | Get videos by location ID |
| `GetMusicInfo` | `/getMusicInfo` | `query`, `region`, `cache_timeout` | Get music info by URL or ID |
| `GetMusicFeed` | `/getMusicFeed` | `query`, `cursor`, `count`, `cache_timeout`, `region` | Get music feed by URL or ID |
| `GetMusicCharts` | `/getMusicCharts` | `chart_type`, `cache_timeout`, `region` | Get music charts |
| `GetTrendingCategoriesList` | `/getTrendingCategoriesList` | `region`, `offset`, `count`, `cache_timeout` | Get trending categories list |
| `GetTrendingFeed` | `/getTrendingFeed` | `region`, `pull_type`, `cache_timeout` | Get trending feed |
| `GetStickersById` | `/getStickersByID` | `sticker_ids`, `cache_timeout` | Get effect information by IDs |
| `GetStickerById` | `/getStickerByID` | `sticker_id`, `cache_timeout` | Get effect information by ID |
| `GetStickerFeedById` | `/getStickerFeedByID` | `sticker_id`, `count`, `cursor`, `region`, `cache_timeout` | Get videos by effect ID |
| `GetUserStories` | `/getUserStories` | `user_id`, `count`, `min_cursor`, `region`, `cache_timeout` | Get user stories |
| `GetStoryInfoById` | `/getStoryInfoByID` | `story_id`, `region`, `cache_timeout` | Get story info by ID |
| `GetUserPlaylists` | `/getUserPlaylists` | `user_id`, `cursor`, `cache_timeout` | Get user playlists |
| `GetUserPlaylistById` | `/getUserPlaylistByID` | `user_id`, `playlist_id`, `cache_timeout` | Get user playlist by ID |
| `GetVideoByPlaylistId` | `/getVideoByPlaylistID` | `user_id`, `playlist_id`, `max_cursor`, `cache_timeout` | Get video by playlist ID |
| `GetAvailableRankingList` | `/getAvailableRankingList` | `room_id`, `user_id`, `cache_timeout` | Get available ranking list |
| `GetRankingInfo` | `/getRankingInfo` | `rank_type`, `room_id`, `user_id`, `cache_timeout` | Get ranking info |
| `GetLiveInfo` | `/getLiveInfo` | `room_id`, `user_id`, `cache_timeout` | Get live info |
| `GetLiveFeed` | `/getLiveFeed` | `draw_room_id`, `max_time`, `region`, `cache_timeout` | Get live feed |

