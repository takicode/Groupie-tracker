package artist

import "errors"

var ErrArtistNotFound = errors.New(
    "artist not found",
)