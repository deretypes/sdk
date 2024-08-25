package media

import "time"

type (
	Provider interface {
		Search(query string) ([]*Media, error)
		GetMediaID(URL string) (string, error)
		FindMedia(mediaID string) (*Media, error)
		FindEpisodes(mediaID string) ([]*Episode, error)
		Extract(episode *Episode, translationID, server string) (*ExtractData, error)
		GetSettings() (*Settings, error)
		SetSettings(opts ...SettingOpt)
	}

	SettingOpt func(*Settings)

	Settings struct {
		VideoServers []string `json:"videoServers"`
		SessionToken *string  `json:"sessionToken"`
	}

	Media struct {
		// "ID" of the extension.
		Provider string `json:"provider"`
		// ID is the media slug.
		// It is used to fetch episodes.
		ID string `json:"id"`
		// Title is the media title.
		Title       string  `json:"title"`
		Description *string `json:"description"`
		// Cover is the media cover URL. (Optional)
		Cover *string `json:"cover"`
		// URL is the media page URL.
		URL string `json:"url"`
		// OtherProps used when additional parameters need to be passed.
		OtherProps *map[string]interface{} `json:"otherProps"`
	}

	Episode struct {
		// "ID" of the extension.
		Provider string `json:"provider"`
		// "ID" of the translation.
		// It is used to fetch EpisodeServer.
		ID string `json:"id"`
		// Episode number.
		// From 0 to n.
		Number int `json:"number"`
		// Episode title. (Optional)
		Title *string `json:"title"`
		// Episode translations.
		// Translation used for get extracted data of episode.
		Translations []*Translation `json:"translations"`
		// Episode page URL.
		URL string `json:"url"`
		// "ID" of the episode Media.
		MediaID string `json:"mediaId"`
	}

	Translation struct {
		// "ID" of the translation.
		// It is used to fetch the episode details.
		ID string `json:"id"`
		// Translation title.
		// e.g. "Crunchyroll", "Erai-Raws"
		Title string `json:"title"`
		// URL of the translation source.
		URL string `json:"url"`
	}

	ExtractData struct {
		// "ID" of the extension.
		Provider string `json:"provider"`
		// Episode server name.
		// e.g. "vidcloud".
		Server string `json:"server"`
		// Video streams for the episode.
		Streams []*Stream `json:"streams"`
		// Video subtitles for the episode.
		Subtitles []*Subtitle `json:"subtitles"`
		// Request headers for the video stream.
		Headers map[string]string `json:"headers"`
		// When streams are expired.
		// Leave it empty if the episode URL is permanent.
		ExpiredAt *time.Time `json:"expiredAt"`
	}

	VideoSourceType string

	Stream struct {
		URL string `json:"url"`
		// Type of the video source.
		Type VideoSourceType `json:"type"`
		// Quality of the video source.
		// e.g. "default", "auto", "1080p".
		Quality string `json:"quality"`
	}

	Subtitle struct {
		URL string `json:"url"`
		// Use country code for mark captions language.
		// e.g. "en", "fr"
		Language string `json:"language"`
	}

	Error string
)

const (
	VideoSourceMP4  VideoSourceType = "mp4"
	VideoSourceM3U8 VideoSourceType = "m3u8"
)
