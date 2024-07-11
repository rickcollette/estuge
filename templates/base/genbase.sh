#!/bin/bash

tree > BASE_TEMPLATE.txt
echo "config.yaml" >> BASE_TEMPLATE.txt && cat config.yaml >> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
echo "footer.tmpl" >> BASE_TEMPLATE.txt && cat footer.tmpl >> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
echo "header.tmpl" >> BASE_TEMPLATE.txt && cat header.tmpl>> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
echo "main.tmpl" >> BASE_TEMPLATE.txt && cat main.tmpl>> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
echo "js/helper_funcs.js" >> BASE_TEMPLATE.txt && cat js/helper_funcs.js >> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
echo "style/style.css" >> BASE_TEMPLATE.txt && cat style/style.css >> BASE_TEMPLATE.txt && echo "===========" >> BASE_TEMPLATE.txt
