# Generate PDFs from the Markdown source files
#
# In order to use this makefile, you need some tools:
# - GNU make
# - Pandoc

# Directory containing source (Markdown) files
source := ./rpts/CosmosCashRegs.md

# Directory containing pdf files
output := ./out/CosmosCashRegs.pdf

# bibliography
bibliography := ./bibs/cosmoscash.bib 

# template
template := eisvogel

clean:
	rm -f $(output)

buildmd:
	pandoc $(source) \
		--output $(output) \
		--from markdown \
		--template $(template) \
		--listings \
		--bibliography=$(bibliography) \
		--toc \
		--top-level-division=chapter \
		--number-sections
