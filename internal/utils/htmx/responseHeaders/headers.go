package responseHeaders

type HxResponseHeader = string

const (
	// HxLocation allows you to do a client-side redirect that does not do a full page reload
	//
	// https://htmx.org/headers/hx-location/
	HxLocation HxResponseHeader = "Hx-Location"
	// HxPushUrl pushes a new url into the history stack
	//
	// https://htmx.org/headers/hx-push-url/
	HxPushUrl HxResponseHeader = "Hx-Push-Url"
	// HxRedirect can be used to do a client-side redirect to a new location
	//
	// https://htmx.org/headers/hx-redirect/
	HxRedirect HxResponseHeader = "Hx-Redirect"
	// HxRefresh
	//
	// if set to “true” the client-side will do a full refresh of the page
	HxRefresh HxResponseHeader = "Hx-Refresh"
	// HxReplaceUrl replaces the current URL in the location bar
	//
	// https://htmx.org/headers/hx-replace-url/
	HxReplaceUrl HxResponseHeader = "Hx-Replace-Url"
	// HxReswap allows you to specify how the response will be swapped. See hx-swap for possible values
	HxReswap HxResponseHeader = "Hx-Reswap"
	// HxRetarget value is a CSS selector that updates the target of the content update to a different element on the page
	HxRetarget HxResponseHeader = "Hx-Retarget"
	// HxReselect value is a CSS selector that allows you to choose which part of the response is used to be swapped in.
	// Overrides an existing `hx-select` on the triggering element
	HxReselect HxResponseHeader = "Hx-Reselect"
	// HxTrigger allows you to trigger client-side events
	//
	// https://htmx.org/headers/hx-trigger/
	HxTrigger HxResponseHeader = "Hx-Trigger"
	// HxTriggerAfterSettle allows you to trigger client-side events after the settle step
	//
	// https://htmx.org/headers/hx-trigger/
	HxTriggerAfterSettle HxResponseHeader = "Hx-Trigger-After-Settle"
	// HxTriggerAfterSwap allows you to trigger client-side events after the swap step
	//
	// https://htmx.org/headers/hx-trigger/
	HxTriggerAfterSwap HxResponseHeader = "Hx-Trigger-After-Swap"
)
