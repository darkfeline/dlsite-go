// Copyright (C) 2021 Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dlsite

// A FetcherOption can be passed to NewFetcher to configure Fetcher creation.
type FetcherOption interface {
	apply(*Fetcher)
}

type cacheOption struct {
	path string
}

func (o cacheOption) apply(f *Fetcher) {
	f.cachePath = o.path
}

// CachePath sets the cache path of the Fetcher.
// If path is empty, no cache file is used.
func CachePath(path string) FetcherOption {
	return cacheOption{path}
}

// A FetchWorkOption can be passed to FetchWork to configure fetching.
type FetchWorkOption interface {
	apply(fetchWorkOptions) fetchWorkOptions
}

type fetchWorkOptions struct {
	ignoreCache bool
}

func mergeOptions(o ...FetchWorkOption) fetchWorkOptions {
	var opts fetchWorkOptions
	for _, o := range o {
		opts = o.apply(opts)
	}
	return opts
}

type ignoreCacheOption struct{}

func (ignoreCacheOption) apply(o fetchWorkOptions) fetchWorkOptions {
	o.ignoreCache = true
	return o
}

// IgnoreCache returns an option that ignores the cache when fetching.
// Updated work information is still added to the cache.
func IgnoreCache() FetchWorkOption {
	return ignoreCacheOption{}
}
