attrib -r .\docs\春语.pdf
copy ..\春语.pdf docs
attrib +r .\docs\春语.pdf

attrib -r .\docs\千维中考-假题分类.pdf
copy ..\千维中考-假题分类.pdf docs
attrib +r .\docs\千维中考-假题分类.pdf

attrib -r .\docs\二六文集.pdf
copy ..\二六文集.pdf docs
attrib +r .\docs\二六文集.pdf

attrib -r .\docs\春语.docx
copy ..\春语.docx docs
attrib +r .\docs\春语.docx

attrib -r .\docs\千维中考-假题分类.docx
copy ..\千维中考-假题分类.docx docs
attrib +r .\docs\千维中考-假题分类.docx

attrib -r .\docs\二六文集.docx
copy ..\二六文集.docx docs
attrib +r .\docs\二六文集.docx

set date=%date:~0,4%/%date:~5,2%/%date:~8,2%
echo %date% > .\docs\version.txt

git add .
git commit -m"..."
git push origin main
git push gitee main

pause
