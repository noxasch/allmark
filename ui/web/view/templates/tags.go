// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

var tagmapContentTemplate = `
{{ if .Tags }}
<ul class="tags">
{{range .Tags}}
<li class="tag">
	<a name="{{.Name}}" title="{{.Description}}">{{.Name}}</a>
	<ol class="childs">
		{{range .Childs}}
		<li class="child">
			<a href="{{.Route}}" class="child-title child-link">{{.Title}}</a>
			<p class="child-description">{{.Description}}</p>
		</li>
		{{end}}
	</ol>
</li>
{{end}}
</ul>
{{else}}
There are currently no tags.
{{end}}
`

const tagmapTemplate = `
<header>
<h1 class="title">
{{.Title}}
</h1>
</header>

<section class="description">
{{.Description}}
</section>

<section class="content">
{{.Content}}
</section>
`
