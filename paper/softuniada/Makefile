.PHONY: latex pdf html

pdf:
	pandoc -s -o paper.pdf paper.md --toc

latex:
	pandoc -s -o paper.tex paper.md --toc 
	
pdf-redacted:
	pandoc -s -o paper.redacted.pdf paper.redacted.md --toc

latex-redacted:
	pandoc -s -o paper.redacted.tex paper.redacted.md --toc 
