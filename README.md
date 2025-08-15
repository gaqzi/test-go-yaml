Test for merging yaml lists in Go
=================================

I was trying to reuse a list of values with Drone CI and I saw [examples](https://smcleod.net/2022/11/yaml-anchors-and-aliases/) of [syntax](https://docs.gitlab.com/ci/yaml/yaml_optimization/) for making it work but I couldn't get any of them to work for me in Drone.

So I decided to figure out whether it was possible by using the underlying libraries used by Drone to parse the YAML.

I found the [go-task](https://github.com/drone/go-task/blob/main/go.mod) repo from Drone and saw they used github.com/ghodss/yaml which is a wrapper around gopkg.in/yaml.v3, so I whipped up the [main.go](./main.go) to show me what rendering my document in some versions look like.

The only version that works in both is:

```yaml
---
example: &example
  - "Hello"
  - "World!"

name: "hello"
commands:
  - << *example
  - "Oho!"
```

Which turns into:

```
0 -> main.Step{Name:"hello", Commands:[]string{"<< *example", "Oho!"}}
```

I.e. it's a string that contains all the valid symbols `<< *example`, everything else gives errors.

I finally found https://github.com/yaml/yaml/issues/35 which explains that there **is no spec** support for this feature at all, and there are a couple of implementations that support it but there's no consensus on how. So, it's not happening, and it took me quite a bit of looking before I found it. I literally ended up making this example because my searching at first didn't yield me the right result 😅
