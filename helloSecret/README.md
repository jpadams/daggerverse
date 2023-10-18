## HelloSecret

Use from the CLI with this PR: https://github.com/dagger/dagger/pull/5882 :smiley:

```
zenith ➤ dagger call --silent -m github.com/jpadams/daggerverse/helloSecret \
    with-name --name="Jeremy" with-greeting --greeting="hello" \
    shout
HELLO, JEREMY!!!!!!!

zenith ➤ dagger call --silent -m github.com/jpadams/daggerverse/helloSecret \
    with-name --name="Jeremy" with-greeting --greeting="hello" \
    message
hello, Jeremy!

zenith ➤ dagger call --silent -m github.com/jpadams/daggerverse/helloSecret \
    with-name --name="Jeremy" \
    message
Secret Hello, Jeremy!

zenith ➤ dagger call --silent -m github.com/jpadams/daggerverse/helloSecret \
    message
Secret Hello, World!
```
