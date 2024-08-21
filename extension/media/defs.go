package media

const (
	ErrUnimplemented   Error = "unimplemented method"
	ErrForbidden       Error = "forbidden"
	ErrMediaNotFound   Error = "media not found"
	ErrEpisodeNotFound Error = "episode not found"
)

func (e Error) Error() string {
	return string(e)
}
