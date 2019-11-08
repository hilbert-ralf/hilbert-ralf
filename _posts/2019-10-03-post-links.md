---
layout: post
title:  "Jekyll-Post-Links generieren"
date:   2019-10-03 01:00:00 +0200
category: ["tech"]
tags: [jekyll]
---

Dies ist ein kleines Script zum Anzeigen möglicher Post-Referenzen zum Einfügen in Markdown-Dateien. Mir spart es ein
 wenig händische Arbeit.

{% highlight bash linenos %}
{% include_relative utils/post_links.sh %}
{% endhighlight %}

Die Ausgabe sieht dann wie folgt aus.

{% highlight bash %}
{% raw %}
λ sh post_links.sh

just copy one of the following link jerkyll / markdown links and copy to your post:


1: ["Was tun mit dem Geld? Sparen!"]({% post_url 2018-08-11-gk1-sparen %})

2: ["Warum Sparen? Aka wie funktioniert unsere Rente?"]({% post_url 2018-08-18-gk2-rente %})

3: ["Grundlagen der passiven Geldanlage"]({% post_url 2018-08-29-gk3-grundlagen-der-passiven-geldanlage %})
{% endraw %}
{% endhighlight %}
