// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

const locationTemplate = `
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

<div class="cleaner"></div>

{{ if .Tags }}
<section class="tags">
	<header>
		Tags:
	</header>

	<ul class="tags">
	{{range .Tags}}
	<li class="tag">
		<a href="{{.Route}}" title="{{.Description}}">{{.Name}}</a>
	</li>
	{{end}}
	</ul>
</section>
{{end}}

{{ if .RelatedItems }}
<section class="related-items">
<ol class="list">
{{range .RelatedItems}}
<li class="related-item">
	<a href="{{.Route}}">{{.Title}}</a>
	<p>{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
{{end}}

{{ if .Childs }}
<section class="childs">
<ol class="list">
{{range .Childs}}
<li class="child">
	<a href="{{.Route}}" class="child-title child-link">{{.Title}}</a>
	<p class="child-description">{{.Description}}</p>
</li>
{{end}}
</ol>
</section>
{{end}}
`
