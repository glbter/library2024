package requestHeaders

type HxRequestHeader = string

const (
	// HxBoosted value indicates that the request is via an element using hx-boost
	HxBoosted HxRequestHeader = "Hx-Boosted"
	// HxCurrentURL value is the current URL of the browser
	HxCurrentURL HxRequestHeader = "Hx-Current-Url"
	// HxHistoryRestoreRequest value is “true” if the request is for history restoration after a miss in the local history cache
	HxHistoryRestoreRequest HxRequestHeader = "Hx-History-Restore-Request"
	// HxPrompt value is the user response to an `hx-prompt``
	HxPrompt HxRequestHeader = "Hx-Prompt"
	// HxRequest value is always "true"
	HxRequest HxRequestHeader = "Hx-Request"
	// HxTarget value is the `id` of the target element if it exists
	HxTarget HxRequestHeader = "Hx-Target"
	// HxTriggerName value is the `name` of the triggered element if it exists
	HxTriggerName HxRequestHeader = "Hx-Trigger-Name"
	// HxTrigger value is the `id` of the triggered element if it exists
	HxTrigger HxRequestHeader = "Hx-Trigger"
)
