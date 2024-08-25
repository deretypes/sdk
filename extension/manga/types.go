package manga

type (
	Provider interface {
		Search(query string) ([]*Manga, error)
		GetMangaID(URL string) (string, error)
		FindManga(mangaID string) (*Manga, error)
		FindChapters(mangaID string) ([]*Chapter, error)
		FindChapter(chapterID string) (*Chapter, error)
		Extract(chapter *Chapter, server string) ([]*ChapterPage, error)
		GetSettings() (*Settings, error)
		SetSettings(opts ...SettingOpt)
	}

	SettingOpt func(*Settings)

	Settings struct {
		ImageServers []string `json:"imageServers"`
		SessionToken *string  `json:"sessionToken"`
	}

	Manga struct {
		// "ID" of the provider.
		Provider string `json:"provider"`
		// ID is the media slug.
		// It is used to fetch episodes.
		ID string `json:"id"`
		// Title is the media title.
		Title string `json:"title"`
		// Cover is the media cover URL. (Optional)
		Cover *string `json:"cover"`
		// URL is the media page URL.
		URL string `json:"url"`
		// OtherProps used when additional parameters need to be passed.
		OtherProps *map[string]interface{} `json:"otherProps"`
	}

	Chapter struct {
		// "ID" of the provider.
		Provider string `json:"provider"`
		// ID of the chapter.
		// It is used to fetch pages.
		ID string `json:"id"`
		// e.g., "1", "1.5", "2", "3"
		Number string `json:"number"`
		// The chapter title.
		// It should be in this format: "Chapter X.Y - {title}" where X is the chapter number and Y is the subchapter number.
		Title string `json:"title"`
		// Chapter index.
		// From 0 to n.
		Index int `json:"index"`
		// Chapter page URL.
		URL string `json:"url"`
		// "ID" of the chapter Manga.
		MangaID string `json:"mangaId"`
	}

	DescrambleFn func(bytes []byte) ([]byte, error)

	ChapterPage struct {
		// "ID" of the provider.
		Provider string `json:"provider"`
		// Index of the page in the chapter.
		// From 0 to n.
		Index int `json:"index"`
		// URL of the chapter page.
		URL string `json:"url"`
		// Filename of the page.
		Filename string `json:"filename"`
		// Function for the page if image decoding is required.
		DescrambleFn *DescrambleFn
	}

	Error string
)
