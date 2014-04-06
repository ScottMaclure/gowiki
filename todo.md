# TODO

## Sooner

* Work out how to use a "master" wrapping template.
* Serve static content, use a /public/ folder structure.
* Split wiki.go into a second page.go file, to contain struct and functions for pages.

## Later

* Implement inter-page linking by converting instances of [PageName] to <a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)
* Dynamically parse all templates in /templates folder, instead of pre-defined ones.
* Support custom 404/error responses (errorHandler).
* Support nested hierarchies for pages.
* Write a function to convert Title to a file-friendly name. "Hello, World" -> hello-world.
* Support "move" operation for pages.

## Changelog

### 2014-04-06

* Redirect to index.txt for default route.
* Bug: redirect issue with /save/ POST. (issue was /save//index, caused 301)
* Finish tutorial steps (was up to "Saving Pages").
* Wiki page content interpreted as markdown and rendered.

### 2014-03-30

* Render pages via http.
* Create project on github and push (https://github.com/ScottMaclure/gowiki/)

### 2014-03-29

* Setup project. Getting started.
