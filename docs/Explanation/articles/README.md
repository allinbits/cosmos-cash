---
title: About Articles
---

# Articles README

This folder includes all articles, papers, and so on that have been created for Cosmos Cash. 

Each folder represents all of the materials and assets for a single article, including bibliographies, images, and so on. The papers themselves are written in Markdown. The Markdown files can then be published in various formats using [Pandoc](https://pandoc.org/MANUAL.html). 

For publishing to PDF format, we use this [Eisvogel LaTeX template](https://github.com/Wandmalfarbe/pandoc-latex-template) to ensure a consistent look and feel. 

## Technical Setup

Set up your environment to manage and publish articles.


### Prerequisites

- [GNU make](https://www.gnu.org/software/make/) utility
- [Pandoc](https://pandoc.org/installing.html)
- [Google Inter font](https://fonts.google.com/specimen/Inter) is the Tendermint font. 

### Usage

- Create a folder with a self-explanatory name for the article
- Add your content to the folder as needed.
- Add the following front matter to the top of the Markdown document and customize as needed:

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

- To create a new report, run the  `make buildmd` from the top-level directory NOTE: This build process is under development and won't work for the moment!!
- To clean outputs, run `make clean`.
