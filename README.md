### How To: Publish your changes
After commiting and pushing your changes, create and push the next tag using the following as an example:
```
git tag v0.0.8
git push origin v0.0.8
```
Now, publish the module to https://pkg.go.dev/: with the latest tag
```
GOPROXY=proxy.golang.org go list -m github.com/semmet95/go_utils@v0.0.8
```