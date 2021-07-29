# Articles README

This folder includes all articles, papers etc that have been created for Cosmos Cash. 

Each folder represents all the materials for a single articles, including bibliographies, images etc. The papers themselves are written in Markdown. These can then be published in various formats using [Pandoc](https://pandoc.org/MANUAL.html). 

For publishing to pdf we use this use a great latex template for a consistent look and feel. [here](https://github.com/Wandmalfarbe/pandoc-latex-template)

## Technical Set-up


### Pre-requisites

- GNU make
- [Pandoc](https://pandoc.org/installing.html)
- [Google's Inter font](https://fonts.google.com/specimen/Inter) is the Tendermint font. 

### Usage

- Create a folder with a self-explanatory name for the article
- Add your content to the folder as needed.
- To the top of the Markdown document add the following and customise as needed

```
---
title: "My title"
subtitle: "My sub title"
author: [Joe Smith]
date: "Publish date"
mainfont: Inter # Default Tendermint font
fontsize: 10pt
subject: "Cosmos Cash"
keywords: []
lang: "en"
book: true
titlepage: true
titlepage-background: ""
classoption: [oneside]
page-background: "../backgrounds/BackgroundTendermint.pdf"
footer-left: "footer"
header-left: "header"
footer-right: "footer"
header-right: "header"
---
```

- To create new report go just need to run `make buildmd` from top level directory NOTE: this won't work for the moment!!
- To clean outputs run `make clean`

