// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package transport

import "strings"

// CanonicalizeHeaderKey canonicalizes the given header key for storage into
// Headers.
func CanonicalizeHeaderKey(k string) string {
	// TODO: Deal with unsupported header keys (anything that's not a valid HTTP
	// header key).
	return strings.ToLower(k)
}

// Headers is the transport-level representation of application headers.
//
// 	var headers transport.Headers
// 	headers = headers.With("foo", "bar")
// 	headers = headers.With("baz", "qux")
type Headers struct {
	// This representation allows us to make zero-value valid
	items map[string]string

	// Represents raw headers that are not canonicalized nor mutated by the transport mapper. rawItems differs from
	// original items in implementation, where originalItems is kept in sync with the canonicalized items and various
	// transports will swap them interchangably and rawItems is an explicit set of headers with a specific contract
	// around mutation.
	rawItems map[string]string

	// keep a merged map of items and rawItems, with items taking precedence on key collisions.
	allItems map[string]string

	// original non-canonical headers, foo-bar will be treated as different value than Foo-bar
	originalItems map[string]string
}

// NewHeaders builds a new Headers object.
func NewHeaders() Headers {
	return Headers{}
}

// NewHeadersWithCapacity allocates a new Headers object with the given
// capacity. A capacity of zero or less is ignored.
func NewHeadersWithCapacity(capacity int) Headers {
	if capacity <= 0 {
		return Headers{}
	}
	return Headers{
		items:         make(map[string]string, capacity),
		rawItems:      make(map[string]string, capacity),
		originalItems: make(map[string]string, capacity),
		// stores at most len(items) + len(newitems) iff there are no collisions
		allItems: make(map[string]string, capacity*2),
	}
}

// With returns a Headers object with the given key-value pair added to it.
//
// The returned object MAY not point to the same Headers underlying data store
// as the original Headers so the returned Headers MUST always be used instead
// of the original object.
//
// 	headers = headers.With("foo", "bar").With("baz", "qux")
func (h Headers) With(k, v string) Headers {
	if h.allItems == nil {
		h.allItems = make(map[string]string)
	}
	if h.items == nil {
		h.items = make(map[string]string)
		h.originalItems = make(map[string]string)
	}
	key := CanonicalizeHeaderKey(k)
	h.items[key] = v
	h.originalItems[k] = v
	h.allItems[key] = v
	return h
}

// WithRaw returns a Headers object with the given key-value pair added to it,
// with the guarantee that the underlying transports will not modify the key or value.
//
// The returned object MAY not point to the same Headers underlying data store
// as the original Headers so the returned Headers MUST always be used instead
// of the original object.
//
// 	headers = headers.WithRaw("foo", "bar").WithRaw("baz", "qux")
func (h Headers) WithRaw(k, v string) Headers {
	if h.rawItems == nil {
		h.rawItems = make(map[string]string)
	}

	h.rawItems[k] = v

	// set value in all items iff this key is not already set to keep canonicalized precedence invariant
	if _, ok := h.items[CanonicalizeHeaderKey(k)]; !ok {
		if h.allItems == nil {
			h.allItems = make(map[string]string)
		}
		h.allItems[k] = v
	}
	return h
}

// Del deletes the header with the given name from the Headers map.
//
// This is a no-op if the key does not exist.
func (h Headers) Del(k string) {
	key := CanonicalizeHeaderKey(k)
	delete(h.items, key)
	delete(h.originalItems, k)
	delete(h.rawItems, k)
	// clean up both the canonicalized and raw keys in the allItems map
	delete(h.allItems, k)
	delete(h.allItems, key)
}

// Get retrieves the value associated with the given header name. For non-raw headers,
// the key will be converted to a canonicalized form. For raw headers, the key must be an
// exact match.
//
// If both a canonicalized and raw header exist with the same key, the canonicalized
// value is returned.
//
//  headers = headers.With("Foo", "bar").WithRaw("foo", "baz")
//  headers.Get("foo") => "bar"
//
//  headers = headers.With("foo", "bar").WithRaw("FOO", "baz")
//  headers.Get("FOO") => "baz"
func (h Headers) Get(k string) (string, bool) {
	v, ok := h.items[CanonicalizeHeaderKey(k)]
	if !ok {
		// If canonicalized key is not found fall back to rawItems. This precedence keeps
		// existing behavior consistent and will avoid breaking older middleware when
		// using WithRaw in newer middleware.
		v, ok = h.rawItems[k]
	}
	return v, ok
}

// Len returns the number of canonicalized headers defined on this object.
func (h Headers) Len() int {
	return len(h.items)
}

// Items returns the underlying map for this Headers object. The returned map
// MUST NOT be changed. Doing so will result in undefined behavior.
//
// Keys in the map are normalized using CanonicalizeHeaderKey.
func (h Headers) Items() map[string]string {
	return h.items
}

// RawItems returns the underlying map of raw headers for this Headers object. The returned map
// MUST NOT be changed. Doing so will result in undefined behavior.
func (h Headers) RawItems() map[string]string {
	return h.rawItems
}

// AllItems returns the underlying map of ALL headers for this Headers object. The returned map
// MUST NOT be changed. Doing so will result in undefined behavior.
func (h Headers) AllItems() map[string]string {
	return h.allItems
}

// OriginalItems returns the non-canonicalized version of the underlying map
// for this Headers object. The returned map MUST NOT be changed.
// Doing so will result in undefined behavior.
func (h Headers) OriginalItems() map[string]string {
	return h.originalItems
}

// HeadersFromMap builds a new Headers object from the given map of header
// key-value pairs.
func HeadersFromMap(m map[string]string) Headers {
	if len(m) == 0 {
		return Headers{}
	}
	headers := NewHeadersWithCapacity(len(m))
	for k, v := range m {
		headers = headers.With(k, v)
	}
	return headers
}
