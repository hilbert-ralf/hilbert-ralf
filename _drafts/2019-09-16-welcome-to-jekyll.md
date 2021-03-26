---
layout: post
title:  "Welcome to struggle with Jekyll!"
date:   2019-09-16 18:24:19 +0200
category: ["tech"]
---

* follow https://jekyllrb.com/docs/installation/
* follow https://jekyllrb.com/docs/
* jekyll new .
* baseurl in _config.yml anpassen
* bundle install
* bundle exec jekyll build
* .gitignore editieren
* alles einchecken
* bundle exec jekyll serve

* github unterstützt viele plugins nicht, daher lokal bauen und **nur** docs ordner deployen:
  * bundle exec jekyll clean -d .\docs && bundle exec jekyll build -d .\docs