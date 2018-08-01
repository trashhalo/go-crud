# Go Crud

Proof of concept building an old school app with no javascript in golang. Forms. Flash. Templates. Oh My.

# Technology
1. Http - native
2. Router - http://www.gorillatoolkit.org/pkg/mux
3. Session - http://www.gorillatoolkit.org/pkg/sessions
4. DB - https://github.com/dgraph-io/badger
5. Templates - https://github.com/valyala/quicktemplate

# Building
`./build.sh`

# Why badger?
I wanted it to be self contained. If you dont want badger in your app swap out the db layer.
