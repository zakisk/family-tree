# Where should we write the GIF?
Output "tree.gif"

# Set up a 1200x600 terminal with 46px font.
Set FontSize 20
Set Width 1400
Set Height 1000
Set WindowBar Colorful
Set Margin 40
Set MarginFill "#688E26"

Type ./family-tree add person david
Sleep 700ms
Enter
Sleep 5s

Type ./family-tree add person jane
Sleep 700ms
Enter
Sleep 5s

Type ./family-tree add person john
Sleep 700ms
Enter
Sleep 5s

Type ./family-tree add person marry
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree connect jane --as wife --of david"
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree connect marry --as daughter --of david"
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree connect david --as father --of john"
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree count sons --of david"
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree count daughters --of david"
Sleep 700ms
Enter
Sleep 5s

Type "./family-tree count wives --of david"
Sleep 700ms
Enter
Sleep 5s

