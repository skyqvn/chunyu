attrib -r .\docs\����.pdf
copy ..\����.pdf docs
attrib +r .\docs\����.pdf

attrib -r .\docs\ǧά�п�-�������.pdf
copy ..\ǧά�п�-�������.pdf docs
attrib +r .\docs\ǧά�п�-�������.pdf

attrib -r .\docs\�����ļ�.pdf
copy ..\�����ļ�.pdf docs
attrib +r .\docs\�����ļ�.pdf

attrib -r .\docs\����.docx
copy ..\����.docx docs
attrib +r .\docs\����.docx

attrib -r .\docs\ǧά�п�-�������.docx
copy ..\ǧά�п�-�������.docx docs
attrib +r .\docs\ǧά�п�-�������.docx

attrib -r .\docs\�����ļ�.docx
copy ..\�����ļ�.docx docs
attrib +r .\docs\�����ļ�.docx

set date=%date:~0,4%/%date:~5,2%/%date:~8,2%
echo %date% > .\docs\version.txt

git add .
git commit -m"..."
git push origin main
git push gitee main

pause
