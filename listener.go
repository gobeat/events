package events

import . "gitlab.com/gobeer/errors"

type Listener func(event Event) Error