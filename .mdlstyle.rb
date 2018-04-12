# MDL style
# vim:set ft=ruby:

all

rule 'ul-style', style: :consistent  # MD004
rule 'ul-indent', indent: 4          # MD007
rule 'ol-prefix', style: 'ordered'   # MD029
rule 'hr-style', style: '---'        # MD035

exclude_rule 'first-header-h1'       # MD002
exclude_rule 'line-length'           # MD013
exclude_rule 'header-style'          # MD029
exclude_rule 'list-marker-space'     # MD030
exclude_rule 'no-inline-html'        # MD033
exclude_rule 'no-emphasis-as-header' # MD036
exclude_rule 'fenced-code-language'  # MD040
exclude_rule 'first-line-h1'         # MD041

# EOF
