package commons

import (
	"net/http"
)

// TraceHeaderValidatorMiddleware allows to extract and save information from the request
/* func TraceHeaderValidatorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	// TODO, verificar si este middleware es necesario o solo ignoramos si no viene el header
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("x-request-id")) < 1 {
			HandleError(w, &BadRequestError{Message: "missing x-request-id header"})
			return
		}

		next.ServeHTTP(w, r)
	})
} */

// ReqHeadersMiddleware set http headers for all request
func ReqHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO, crear un builder o algo parecido para inyectar headers a pedido y no inyectarlos todos
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("x-request-id", r.Header.Get("x-request-id"))
		// jaeger tracer headers
		w.Header().Set("x-b3-traceid", r.Header.Get("x-b3-traceid"))
		w.Header().Set("x-b3-spanid", r.Header.Get("x-b3-spanid"))
		w.Header().Set("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
		w.Header().Set("x-b3-sampled", r.Header.Get("x-b3-sampled"))
		w.Header().Set("x-b3-flags", r.Header.Get("x-b3-flags"))
		w.Header().Set("x-ot-span-context", r.Header.Get("x-ot-span-context"))
		// datadog tracer headers
		w.Header().Set("x-datadog-trace-id", r.Header.Get("x-datadog-trace-id"))
		w.Header().Set("x-datadog-parent-id", r.Header.Get("x-datadog-parent-id"))
		w.Header().Set("x-datadog-sampling-priority", r.Header.Get("x-datadog-sampling-priority"))

		next.ServeHTTP(w, r)
	})
}
