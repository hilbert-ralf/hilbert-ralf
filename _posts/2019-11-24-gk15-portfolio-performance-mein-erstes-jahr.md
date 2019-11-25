---
layout: post
title:  "Portfolio Performance - mein erstes Jahr"
date:   2019-11-24 01:00:00 +0200
category: ["geldkram"]
tags: ["Portfolio Performance", Rebalancing, "Vermögensentwicklung"]
---


Mitte November 2018 habe ich das Tool [Portfolio Performance](https://www.portfolio-performance.info/portfolio/) für mich aufgesetzt. Jetzt, ein Jahr später kann ich sagen, dass ich die Fleißarbeit, die mit diesem Tool verbunden ist, nicht gescheut habe und es nicht bereue. Aber dazu später mehr, zunächst ...


# Was ist Portfolio Performance?

[Portfolio Performance](https://www.portfolio-performance.info/portfolio/) beschreibt sich auf der eigenen
 Homepage wie folgt.

> Ein Open Source Programm zur Berechnung der Performance eines Gesamtportfolios - über verschiedene Depots und
 Konten hinweg - anhand von True-Time Weighted Rate of Return und internem Zinsfuß.

Das umreißt es in groben Zügen.  Performance kann vieles sein, darüber hinaus finde ich Portfolio Performance aber auch nützlich für Strategiefestlegungen und Rebalancing, Erfassen von Dividenden(-änderungen) sowie dem Messen von Volatilität.


# Wie funktioniert Portfolio Performance?

Im Grunde stellt Portfolio Performance eine art Spiegelbild zum Depot dar. Man legt Konten und Depots an. Die Konten sind für Geld-Eingänge sowie -Ausgänge (Einzahlungen, Dividendenausschüttungen, Käufe und Verkäufe). Die Depots halten im Gegenzug die Wertpapiere (Aktien, Fonds und so weiter). Portfolio Performance ist in der Lage Kurse für Wertpapiere von verschiedenen Quellen aus dem Internet zu beziehen. Hierzu muss man dem Tool jedoch helfen, indem man einen entsprechenden Kurslieferanten angibt. Zur Auswahl stehen dabei beispielsweise [Yahoo Finance](https://de.finance.yahoo.com/) oder Tabellen auf Webseiten im Allgemeinen. Wenn das getan ist, ermittelt Portfolio Performance die Kurse beim start.


# Meine Erfahrungen mit Portfolio Performance

Nachdem ich das Tool nun ein Jahr lang benutzte kann ich sagen das ich im Großen und Ganzen zufrieden bin. Zahlen die ich vorher mit großem Aufwand erheben und berechnen hätte können, was ich aber nicht tat, habe ich nun auf einen Blick. Beispielsweise hatte ich in dem Jahr (22.11.2018 - 22.11.2019) eine "Rendite" von um die 14 % [^1] (14,63 % [Time Weighted Rate of Return](https://en.wikipedia.org/wiki/Time-weighted_return) beziehungsweise 14.82 % [Interner Zinsfuß](https://de.wikipedia.org/wiki/Interner_Zinsfu%C3%9F)). Spannend finde ich auch, das ich durch Schwankungen über eine Zeit von etwa 3 Monaten ein Wertverlust von 9,24 % meines Gesamtportfolios hinnehmen musste. Da spiegelt sich das viel beschrieene Risiko wieder, das viele mit dem Aktienmarkt verbinden. Auch kann ich nun schnell sehen, wie hoch meine Einlagen waren und wie viele Dividenden ich pro monat erhielt. Eine ganze Menge Diagramme die auch stark individualisierbar sind runden das Packet ab.


## Initiales Einpflegen der Daten

Leider bekommt man all dies nicht ganz "gratis". Der Preis ist die Arbeit, die man investieren muss. Gerade das initiale Pflegen der Daten ist aufwendig. In meinem Fall handelte es sich um mehrere Depots mit mehreren dutzend Wertpapieren. Der Einfachheit halber beschloss ich, die Historie (Kaufdatum sowie Dividendenzahlungen) nur begrenzt einzupflegen. Ich habe also nicht jeden Kauf in der Vergangenheit in Portfolio Performance eingetragen, sondern zu einem Stichtag meine Wertpapier- und Barbestände als "Einlieferung" eingetragen. Damit sehe ich nicht, wie die Entwicklung und die Performance in der Vergangenheit war. Aber die Arbeit, den aktuellen Stand einzutragen war mir bereits genug.


## Kontinuierliches Einpflegen

Da meine Depots natürlich nicht statisch sind, muss ich dann und wann noch Aktualisierungen in Portfolio Performance eintragen. Beispielsweise, wenn ich monatlich meinen festen Beitrag spare oder wenn ich von diesem Wertpapiere kaufe. Auch Dividendenzahlungen muss ich einpflegen. Interessant ist hierbei, das Portfolio Performance Bank-Dokumente in form von PDFs verarbeiten kann. Wenn ich also Dividenden bekomme, bekomme ich darüber Bescheid von meiner Bank in Form eines PDFs. Dieses kann Portfolio Performance verarbeiten und entsprechend die Werte übernehmen. Leider sind nur einige Banken unterstützt. Tatsächlich sind es [eine ganze Menge Banken](https://forum.portfolio-performance.info/t/buchungen-aus-pdf-dateien-importieren/38), von meinen 3 Banken sind es jedoch nur 2. Für die nicht unterstützte Bank muss ich diese Arbeit händisch durchführen. Mehr als ein paar Minuten im Monat dauert aber auch diese Aufgabe nicht.


## Rebalancing

Eine weitere wichtige Funktion von Portfolio Performance ist, das man seine Wertpapiere kategorisieren und mit einer gewollten Wertaufteilung versehen kann. Die aktuellen Werte in den Depots werden dann damit abgeglichen und man weiß, in welche Wertpapiere man investieren sollte, um die gewünschte Vermögensverteilung beizubehalten. Hier ein Beispiel wie so etwas aussehen kann anhand einer mit Portfolio Performance mitgelieferten Beispieldatei:

![Portfolio Performance Rebalancing]({% link assets/geldkram/gk15-portfolio-performance-rebalancing.JPG %})


## Open Scource

Zu guter Letzt ist ein weiterer Vorteil von Portfolio Performance, auf den ich zwar einen Blick geworfen, aber ihn nicht wahr genommen habe: Es ist open source. Das Tool ist in der Programmiersprache Java geschrieben und [auf GitHub verfügbar](https://github.com/buchen/portfolio). Änderungen an der Software, zum Beispiel für gewünschte Funktionalitäten oder Fehlerbehebungen kann man also vorschlagen oder gar selbst umsetzen, das richtige KnowHow und die nötige Zeit vorausgesetzt.


## Fazit und Ausblick

Auch wenn Portfolio Performance durchaus mit Arbeit verbunden ist, macht es mir Spass selber Dashboards zu erstellen und den Überblick zu behalten. Das mir jedoch auch beim Rebalancing unter die Arme gegriffen wird, ist vermutlich einer der wichtigsten Aspekte, die ich anders nicht abdecken könnte. Ich halte Portfolio Performance für ein tolles Tool, wenn man Interesse an der Entwicklung der eigenen Vermögenswerte hat und werde es auch weiterhin nutzen.

# Feedback

Feedback jeglicher Art kannst du mir gerne als Antwort unter folgendem Tweet hinterlassen.

{% twitter https://twitter.com/ralf_hilbert/status/1198996852739366913 %}



[^1]: Genau genommen waren es 14,63% [Time Weighted Rate of Return](https://en.wikipedia.org/wiki/Time-weighted_return) beziehungsweise 14.82% [Interner Zinsfuß](https://de.wikipedia.org/wiki/Interner_Zinsfu%C3%9F). Das klingt enorm, liegt aber vor allem an einem glücklichen Zeitfenster, das ich wählte. Der [MSCI World](https://www.onvista.de/index/MSCI-WORLD-Index-3193857) machte im gleichen Zeitraum etwa 14,39%, was einfach starken Schwankungen geschuldet ist. Langfristig kann man je nach Quelle von 5% - 10% p.a. ausgehen.
