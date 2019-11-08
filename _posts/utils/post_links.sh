#{% raw %}
c=1

echo ""
echo "just copy one of the following link jerkyll / markdown links and copy to your post:"
echo ""

# set this variable with relative path to your posts
RELATIVE_PATH="../"

FILES="$(ls $RELATIVE_PATH -p | grep -v /)"

for n in $FILES
do
    # remove file ending
    REF="$(echo $n | cut -d'.' -f1)"

    # stripping down to the raw title
    TITLE="$(cat $RELATIVE_PATH$n | grep 'title:' | cut -d'"' -f2)"

    echo ""
    echo "$c: [\"$TITLE\"]({% post_url $REF %})"

    c=$((c+1))
done
#{% endraw %}